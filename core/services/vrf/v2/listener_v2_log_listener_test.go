package v2

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	evmmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/mocks"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/log_emitter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_coordinator_v2"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_log_emitter"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/vrf/vrfcommon"
	"github.com/smartcontractkit/chainlink/v2/core/testdata/testspecs"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var (
	EmitterABI, _    = abi.JSON(strings.NewReader(log_emitter.LogEmitterABI))
	VRFEmitterABI, _ = abi.JSON(strings.NewReader(vrf_log_emitter.VRFLogEmitterABI))
)

type TestHarness struct {
	Lggr              logger.Logger
	ChainID           *big.Int
	ORM               *logpoller.DbORM
	LogPoller         logpoller.LogPollerTest
	Client            *backends.SimulatedBackend
	Emitter           *log_emitter.LogEmitter
	EmitterAddress    common.Address
	VRFLogEmitter     *vrf_log_emitter.VRFLogEmitter
	VRFEmitterAddress common.Address
	Owner             *bind.TransactOpts
	EthDB             ethdb.Database
	Db                *sqlx.DB
	Listener          *listenerV2
	Ctx               context.Context
}

func SetupTH(t *testing.T,
	useFinalityTag bool,
	finalityDepth, backfillBatchSize,
	rpcBatchSize, keepFinalizedBlocksDepth int64,
	mockChainUpdateFn func(*evmmocks.Chain, *TestHarness)) *TestHarness {

	lggr := logger.TestLogger(t)
	chainID := testutils.NewRandomEVMChainID()
	db := pgtest.NewSqlxDB(t)

	o := logpoller.NewORM(chainID, db, lggr, pgtest.NewQConfig(true))
	owner := testutils.MustNewSimTransactor(t)
	ethDB := rawdb.NewMemoryDatabase()
	ec := backends.NewSimulatedBackendWithDatabase(ethDB, map[common.Address]core.GenesisAccount{
		owner.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10e6)
	// VRF Listener relies on block timestamps, but SimulatedBackend uses by default clock starting from 1970-01-01
	// This trick is used to move the clock closer to the current time. We set first block to be X hours ago.
	// FirstBlockAge is used to compute first block's timestamp in SimulatedBackend (time.Now() - FirstBlockAge)
	const FirstBlockAge = 24 * time.Hour
	blockTime := time.UnixMilli(int64(ec.Blockchain().CurrentHeader().Time))
	err := ec.AdjustTime(time.Since(blockTime) - FirstBlockAge)
	require.NoError(t, err)
	ec.Commit()

	// Poll period doesn't matter, we intend to call poll and save logs directly in the test.
	// Set it to some insanely high value to not interfere with any tests.
	esc := client.NewSimulatedBackendClient(t, ec, chainID)
	// Mark genesis block as finalized to avoid any nulls in the tests
	head := esc.Backend().Blockchain().CurrentHeader()
	esc.Backend().Blockchain().SetFinalized(head)

	lp := logpoller.NewLogPoller(o, esc, lggr, 1*time.Hour, useFinalityTag, finalityDepth, backfillBatchSize, rpcBatchSize, keepFinalizedBlocksDepth)

	emitterAddress1, _, emitter1, err := log_emitter.DeployLogEmitter(owner, ec)
	require.NoError(t, err)
	vrfLogEmitterAddress, _, vrfLogEmitter, err := vrf_log_emitter.DeployVRFLogEmitter(owner, ec)
	require.NoError(t, err)
	ec.Commit()

	// Log Poller Listener
	cfg := pgtest.NewQConfig(false)
	ks := keystore.NewInMemory(db, utils.FastScryptParams, lggr, cfg)
	require.NoError(t, ks.Unlock("blah"))
	j, err := vrfcommon.ValidatedVRFSpec(testspecs.GenerateVRFSpec(testspecs.VRFSpecParams{
		RequestedConfsDelay: 10,
		EVMChainID:          chainID.String(),
	}).Toml())
	require.NoError(t, err)

	coordinatorV2, err := vrf_coordinator_v2.NewVRFCoordinatorV2(vrfLogEmitter.Address(), ec)
	require.Nil(t, err)
	coordinator := NewCoordinatorV2(coordinatorV2)

	chain := evmmocks.NewChain(t)
	listener := &listenerV2{
		respCount:   map[string]uint64{},
		job:         j,
		chain:       chain,
		l:           logger.Sugared(lggr),
		coordinator: coordinator,
	}
	ctx := testutils.Context(t)

	// Filter registration is idempotent, so we can just call it every time
	// and retry on errors using the ticker.
	err = lp.RegisterFilter(logpoller.Filter{
		Name: fmt.Sprintf("vrf_%s_keyhash_%s_job_%d", "v2", listener.job.VRFSpec.PublicKey.MustHash().String(), listener.job.ID),
		EventSigs: evmtypes.HashArray{
			vrf_log_emitter.VRFLogEmitterRandomWordsRequested{}.Topic(),
			vrf_log_emitter.VRFLogEmitterRandomWordsFulfilled{}.Topic(),
		},
		// EventSigs: []common.Hash{VRFEmitterABI.Events["RandomWordsRequested"].ID},
		Addresses: evmtypes.AddressArray{
			vrfLogEmitter.Address(),
			// listener.job.VRFSpec.CoordinatorAddress.Address(),
		},
	})
	require.Nil(t, err)
	require.NoError(t, lp.RegisterFilter(logpoller.Filter{
		Name:      "Integration test",
		EventSigs: []common.Hash{EmitterABI.Events["Log1"].ID},
		Addresses: []common.Address{emitterAddress1},
		Retention: 0}))
	require.Nil(t, err)
	require.Len(t, lp.Filter(nil, nil, nil).Addresses, 2)
	require.Len(t, lp.Filter(nil, nil, nil).Topics, 1)

	th := &TestHarness{
		Lggr:              lggr,
		ChainID:           chainID,
		ORM:               o,
		LogPoller:         lp,
		Emitter:           emitter1,
		EmitterAddress:    emitterAddress1,
		VRFLogEmitter:     vrfLogEmitter,
		VRFEmitterAddress: vrfLogEmitterAddress,
		Client:            ec,
		Owner:             owner,
		EthDB:             ethDB,
		Db:                db,
		Listener:          listener,
		Ctx:               ctx,
	}
	mockChainUpdateFn(chain, th)
	return th
}

/* Tests for initializeLastProcessedBlock: BEGIN
 * TestInitProcessedBlock_NoVRFReqs
 * TestInitProcessedBlock_NoUnfulfilledVRFReqs
 * TestInitProcessedBlock_OneUnfulfilledVRFReq
 * TestInitProcessedBlock_SomeUnfulfilledVRFReqs
 * TestInitProcessedBlock_UnfulfilledNFulfilledVRFReqs
 */

func TestInitProcessedBlock_NoVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, th *TestHarness) {
		// txstore := txmgr.NewTxStore(db, lggr, cfg)
		// txm := makeTestTxm(t, txstore, ks)
		// chain.On("TxManager").Return(txm)
		mockChain.On("ID").Return(th.ChainID)
		mockChain.On("LogPoller").Return(th.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Emit some logs in blocks finalityDepth+1 to finalityDepth+4.
	numBlocks := 5
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 5 (EmitLog blocks) = 9

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(testutils.Context(t)))

	// The poller starts on a new chain at latest-finality (finalityDepth + 5 in this case),
	// Replaying from block 4 should guarantee we have block 4 immediately.  (We will also get
	// block 3 once the backup poller runs, since it always starts 100 blocks behind.)
	require.NoError(t, th.LogPoller.Replay(testutils.Context(t), 4))

	// We should immediately have at least logs 4-7
	logs, err := th.LogPoller.Logs(4, 7, EmitterABI.Events["Log1"].ID, th.EmitterAddress,
		pg.WithParentCtx(testutils.Context(t)))
	require.NoError(t, err)
	require.Equal(t, 3, len(logs))

	lastProcessedBlock, err := th.Listener.initializeLastProcessedBlock(th.Ctx)
	require.Nil(t, err)
	require.Equal(t, int64(6), lastProcessedBlock)
}

func TestInitProcessedBlock_NoUnfulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("ID").Return(curTH.ChainID)
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Create VRF request block and a fulfillment block
	_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
		[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
		big.NewInt(1), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
	require.NoError(t, err2)
	th.Client.Commit()
	_, err2 = th.VRFLogEmitter.EmitRandomWordsFulfilled(th.Owner, big.NewInt(1), big.NewInt(105), big.NewInt(10), true)
	require.NoError(t, err2)
	th.Client.Commit()

	// Emit some logs in blocks to make the VRF req and fulfillment older than finalityDepth from latestBlock
	numBlocks := 5
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 2 (VRF req/resp block) + 5 (EmitLog blocks) = 11

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	latestBlock := finalityDepth + int64(numBlocks) + 1
	require.NoError(t, th.LogPoller.Replay(th.Ctx, latestBlock))

	// initializeLastProcessedBlock must return the finalizedBlockNumber (8) instead of
	// VRF request block number (5), since all VRF requests are fulfilled
	lastProcessedBlock, err := th.Listener.initializeLastProcessedBlock(th.Ctx)
	require.Nil(t, err)
	require.Equal(t, int64(8), lastProcessedBlock)
}

func TestInitProcessedBlock_OneUnfulfilledVRFReq(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("ID").Return(curTH.ChainID)
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Make a VRF request without fulfilling it
	_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
		[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
		big.NewInt(1), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
	require.NoError(t, err2)
	th.Client.Commit()

	// Emit some logs in blocks to make the VRF req and fulfillment older than finalityDepth from latestBlock
	numBlocks := 5
	th.Client.Commit()
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 1 (VRF req block) + 5 (EmitLog blocks) = 10

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	latestBlock := finalityDepth + int64(numBlocks) + 1
	require.NoError(t, th.LogPoller.Replay(th.Ctx, latestBlock))

	// initializeLastProcessedBlock must return the unfulfilled VRF
	// request block number (5) instead of finalizedBlockNumber (8)
	lastProcessedBlock, err := th.Listener.initializeLastProcessedBlock(th.Ctx)
	require.Nil(t, err)
	require.Equal(t, int64(5), lastProcessedBlock)
}

func TestInitProcessedBlock_SomeUnfulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("ID").Return(curTH.ChainID)
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Emit some logs in blocks with VRF reqs interspersed
	// No fulfillment for any VRF requests
	numBlocks := 5
	th.Client.Commit()
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()

		// Create 2 blocks with VRF requests in each iteration
		_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i)), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		th.Client.Commit()
		_, err2 = th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i+1)), big.NewInt(106), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 3*5 (EmitLog + VRF req/resp blocks) = 19

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	latestBlock := finalityDepth + int64(numBlocks) + 1
	require.NoError(t, th.LogPoller.Replay(th.Ctx, latestBlock))

	// initializeLastProcessedBlock must return the earliest unfulfilled VRF request block
	// number instead of finalizedBlockNumber
	lastProcessedBlock, err := th.Listener.initializeLastProcessedBlock(th.Ctx)
	require.Nil(t, err)
	require.Equal(t, int64(7), lastProcessedBlock)
}

func TestInitProcessedBlock_UnfulfilledNFulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("ID").Return(curTH.ChainID)
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Emit some logs in blocks with VRF reqs interspersed
	// One VRF request in each iteration is fulfilled to imitate mixed workload
	numBlocks := 5
	th.Client.Commit()
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()

		// Create 2 blocks with VRF requests in each iteration and fulfill one
		// of them. This creates a mixed workload of fulfilled and unfulfilled
		// VRF requests for testing the VRF listener
		_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i)), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		th.Client.Commit()
		_, err2 = th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i+1)), big.NewInt(106), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		_, err2 = th.VRFLogEmitter.EmitRandomWordsFulfilled(th.Owner, big.NewInt(int64(2*i)), big.NewInt(105), big.NewInt(10), true)
		require.NoError(t, err2)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 3*5 (EmitLog + VRF req/resp blocks) = 19

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	latestBlock := finalityDepth + int64(numBlocks) + 1
	require.NoError(t, th.LogPoller.Replay(th.Ctx, latestBlock))

	// initializeLastProcessedBlock must return the earliest unfulfilled VRF request block
	// number instead of finalizedBlockNumber
	lastProcessedBlock, err := th.Listener.initializeLastProcessedBlock(th.Ctx)
	require.Nil(t, err)
	require.Equal(t, int64(8), lastProcessedBlock)
}

/* Tests for initializeLastProcessedBlock: END */

/* Tests for updateLastProcessedBlock: BEGIN
 * TestUpdateLastProcessedBlock_NoVRFReqs
 * TestUpdateLastProcessedBlock_NoUnfulfilledVRFReqs
 * TestUpdateLastProcessedBlock_OneUnfulfilledVRFReq
 * TestUpdateLastProcessedBlock_SomeUnfulfilledVRFReqs
 * TestUpdateLastProcessedBlock_UnfulfilledNFulfilledVRFReqs
 */

func TestUpdateLastProcessedBlock_NoVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Create VRF request logs
	_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
		[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
		big.NewInt(int64(1)), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
	require.NoError(t, err2)
	th.Client.Commit()
	_, err2 = th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
		[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
		big.NewInt(int64(2)), big.NewInt(106), 1, 10, 10000, 2, th.Owner.From)
	require.NoError(t, err2)
	th.Client.Commit()

	// Emit some logs in blocks to make the VRF req and fulfillment older than finalityDepth from latestBlock
	numBlocks := 5
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 2 (VRF req blocks) + 5 (EmitLog blocks) = 11

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	// We've to replay from before VRF request log, since updateLastProcessedBlock
	// does not internally call LogPoller.Replay
	require.NoError(t, th.LogPoller.Replay(th.Ctx, 4))

	// updateLastProcessedBlock must return the finalizedBlockNumber as there are
	// no VRF requests, after currLastProcessedBlock (block 6). The VRF requests
	// made above are before the currLastProcessedBlock (7) passed in below
	lastProcessedBlock, err := th.Listener.updateLastProcessedBlock(th.Ctx, 7)
	require.Nil(t, err)
	require.Equal(t, int64(8), lastProcessedBlock)
}

func TestUpdateLastProcessedBlock_NoUnfulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Create VRF request log block with a fulfillment log block
	_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
		[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
		big.NewInt(int64(1)), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
	require.NoError(t, err2)
	th.Client.Commit()
	_, err2 = th.VRFLogEmitter.EmitRandomWordsFulfilled(th.Owner, big.NewInt(1), big.NewInt(105), big.NewInt(10), true)
	require.NoError(t, err2)
	th.Client.Commit()

	// Emit some logs in blocks to make the VRF req and fulfillment older than finalityDepth from latestBlock
	numBlocks := 5
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 2 (VRF req/resp blocks) + 5 (EmitLog blocks) = 11

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	// We've to replay from before VRF request log, since updateLastProcessedBlock
	// does not internally call LogPoller.Replay
	require.NoError(t, th.LogPoller.Replay(th.Ctx, 4))

	// updateLastProcessedBlock must return the finalizedBlockNumber (8) though we have
	// a VRF req at block (5) after currLastProcessedBlock (4) passed below, because
	// the VRF request is fulfilled
	lastProcessedBlock, err := th.Listener.updateLastProcessedBlock(th.Ctx, 4)
	require.Nil(t, err)
	require.Equal(t, int64(8), lastProcessedBlock)
}

func TestUpdateLastProcessedBlock_OneUnfulfilledVRFReq(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Create VRF request logs without a fulfillment log block
	_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
		[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
		big.NewInt(int64(1)), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
	require.NoError(t, err2)
	th.Client.Commit()

	// Emit some logs in blocks to make the VRF req and fulfillment older than finalityDepth from latestBlock
	numBlocks := 5
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 1 (VRF req block) + 5 (EmitLog blocks) = 10

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	// We've to replay from before VRF request log, since updateLastProcessedBlock
	// does not internally call LogPoller.Replay
	require.NoError(t, th.LogPoller.Replay(th.Ctx, 4))

	// updateLastProcessedBlock must return the VRF req at block (5) instead of
	// finalizedBlockNumber (8) after currLastProcessedBlock (4) passed below,
	// because the VRF request is unfulfilled
	lastProcessedBlock, err := th.Listener.updateLastProcessedBlock(th.Ctx, 4)
	require.Nil(t, err)
	require.Equal(t, int64(5), lastProcessedBlock)
}

func TestUpdateLastProcessedBlock_SomeUnfulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Emit some logs in blocks to make the VRF req and fulfillment older than finalityDepth from latestBlock
	numBlocks := 5
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()

		// Create 2 blocks with VRF requests in each iteration
		_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i)), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		th.Client.Commit()
		_, err2 = th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i+1)), big.NewInt(106), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 3*5 (EmitLog + VRF req blocks) = 19

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	// We've to replay from before VRF request log, since updateLastProcessedBlock
	// does not internally call LogPoller.Replay
	require.NoError(t, th.LogPoller.Replay(th.Ctx, 4))

	// updateLastProcessedBlock must return the VRF req at block (6) instead of
	// finalizedBlockNumber (16) after currLastProcessedBlock (4) passed below,
	// as block 6 contains the earliest unfulfilled VRF request
	lastProcessedBlock, err := th.Listener.updateLastProcessedBlock(th.Ctx, 4)
	require.Nil(t, err)
	require.Equal(t, int64(6), lastProcessedBlock)
}

func TestUpdateLastProcessedBlock_UnfulfilledNFulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	finalityDepth := int64(3)
	th := SetupTH(t, false, finalityDepth, 3, 2, 1000, func(mockChain *evmmocks.Chain, curTH *TestHarness) {
		mockChain.On("LogPoller").Return(curTH.LogPoller)
	})

	// Block 3 to finalityDepth. Ensure we have finality number of blocks
	for i := 1; i < int(finalityDepth); i++ {
		th.Client.Commit()
	}

	// Emit some logs in blocks to make the VRF req and fulfillment older than finalityDepth from latestBlock
	numBlocks := 5
	for i := 0; i < numBlocks; i++ {
		_, err1 := th.Emitter.EmitLog1(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		_, err1 = th.Emitter.EmitLog2(th.Owner, []*big.Int{big.NewInt(int64(i))})
		require.NoError(t, err1)
		th.Client.Commit()

		// Create 2 blocks with VRF requests in each iteration and fulfill one
		// of them. This creates a mixed workload of fulfilled and unfulfilled
		// VRF requests for testing the VRF listener
		_, err2 := th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i)), big.NewInt(105), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		th.Client.Commit()
		_, err2 = th.VRFLogEmitter.EmitRandomWordsRequested(th.Owner,
			[32]byte(th.Listener.job.VRFSpec.PublicKey.MustHash().Bytes()),
			big.NewInt(int64(2*i+1)), big.NewInt(106), 1, 10, 10000, 2, th.Owner.From)
		require.NoError(t, err2)
		_, err2 = th.VRFLogEmitter.EmitRandomWordsFulfilled(th.Owner, big.NewInt(int64(2*i)), big.NewInt(105), big.NewInt(10), true)
		require.NoError(t, err2)
		th.Client.Commit()
	}

	// Blocks till now: 2 (in SetupTH) + 2 (empty blocks) + 3*5 (EmitLog + VRF req blocks) = 19

	// Calling Start() after RegisterFilter() simulates a node restart after job creation, should reload Filter from db.
	require.NoError(t, th.LogPoller.Start(th.Ctx))

	// We've to replay from before VRF request log, since updateLastProcessedBlock
	// does not internally call LogPoller.Replay
	require.NoError(t, th.LogPoller.Replay(th.Ctx, 4))

	lh, _ := th.LogPoller.LatestLogByEventSigWithConfs(VRFEmitterABI.Events["RandomWordsFulfilled"].ID,
		th.Listener.coordinator.Address(),
		logpoller.Finalized,
	)
	for i, topic := range lh.Topics {
		fmt.Printf("LogTopic[%d]: %v\n", i, common.Bytes2Hex(topic))
	}
	fmt.Printf("LogData: %v\n", common.Bytes2Hex(lh.Data))

	// updateLastProcessedBlock must return the VRF req at block (7) instead of
	// finalizedBlockNumber (16) after currLastProcessedBlock (4) passed below,
	// as block 7 contains the earliest unfulfilled VRF request. VRF request
	// in block 6 has been fulfilled in block 7.
	lastProcessedBlock, err := th.Listener.updateLastProcessedBlock(th.Ctx, 4)
	require.Nil(t, err)
	require.Equal(t, int64(7), lastProcessedBlock)
}

/* Tests for updateLastProcessedBlock: END */

/* Tests for getUnfulfilled: BEGIN
 * TestGetUnfulfilled_NoVRFReqs
 * TestGetUnfulfilled_NoUnfulfilledVRFReqs
 * TestGetUnfulfilled_OneUnfulfilledVRFReq
 * TestGetUnfulfilled_SomeUnfulfilledVRFReqs
 * TestGetUnfulfilled_UnfulfilledNFulfilledVRFReqs
 */

func SetupGetUnfulfilledTH(t *testing.T) (*listenerV2, *utils.Big) {
	chainID := utils.NewBig(big.NewInt(12345))
	lggr := logger.TestLogger(t)
	j, err := vrfcommon.ValidatedVRFSpec(testspecs.GenerateVRFSpec(testspecs.VRFSpecParams{
		RequestedConfsDelay: 10,
		EVMChainID:          chainID.String(),
	}).Toml())
	require.NoError(t, err)
	chain := evmmocks.NewChain(t)

	// Construct CoordinatorV2_X object for VRF listener
	owner := testutils.MustNewSimTransactor(t)
	ethDB := rawdb.NewMemoryDatabase()
	ec := backends.NewSimulatedBackendWithDatabase(ethDB, map[common.Address]core.GenesisAccount{
		owner.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10e6)
	_, _, vrfLogEmitter, err := vrf_log_emitter.DeployVRFLogEmitter(owner, ec)
	require.NoError(t, err)
	ec.Commit()
	coordinatorV2, err := vrf_coordinator_v2.NewVRFCoordinatorV2(vrfLogEmitter.Address(), ec)
	require.Nil(t, err)
	coordinator := NewCoordinatorV2(coordinatorV2)

	listener := &listenerV2{
		respCount:   map[string]uint64{},
		job:         j,
		chain:       chain,
		l:           logger.Sugared(lggr),
		coordinator: coordinator,
	}
	return listener, chainID
}

func TestGetUnfulfilled_NoVRFReqs(t *testing.T) {
	t.Parallel()

	listener, chainID := SetupGetUnfulfilledTH(t)

	logs := []logpoller.Log{}
	for i := 0; i < 10; i++ {
		logs = append(logs, logpoller.Log{
			EvmChainId:     chainID,
			LogIndex:       0,
			BlockHash:      common.BigToHash(big.NewInt(int64(i))),
			BlockNumber:    int64(i),
			BlockTimestamp: time.Now(),
			Topics: [][]byte{
				[]byte("0x7299721c72d5436afdfb18da500ab7e08ac04c812d7204e1bddd322f67d8be4c"),
			},
			EventSig:  EmitterABI.Events["Log1"].ID,
			Address:   common.Address{},
			TxHash:    common.BigToHash(big.NewInt(int64(i))),
			Data:      nil,
			CreatedAt: time.Now(),
		})
	}

	unfulfilled, _, fulfilled := listener.getUnfulfilled(logs, listener.l)
	require.Empty(t, unfulfilled)
	require.Empty(t, fulfilled)
}

func TestGetUnfulfilled_NoUnfulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	listener, chainID := SetupGetUnfulfilledTH(t)

	logs := []logpoller.Log{}
	for i := 0; i < 10; i++ {
		eventSig := EmitterABI.Events["Log1"].ID
		topics := [][]byte{
			common.FromHex("0x46692c0e59ca9cd1ad8f984a9d11715ec83424398b7eed4e05c8ce84662415a8"),
		}
		if i%2 == 0 {
			eventSig = VRFEmitterABI.Events["RandomWordsRequested"].ID
			topics = [][]byte{
				common.FromHex("0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772"),
				common.FromHex("0xc0a6c424ac7157ae408398df7e5f4552091a69125d5dfcb7b8c2659029395bdf"),
				common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001"),
				common.FromHex("0x0000000000000000000000005ee3b50502b5c4c9184dcb281471a0614d4b2ef9"),
			}
		}
		logs = append(logs, logpoller.Log{
			EvmChainId:     chainID,
			LogIndex:       0,
			BlockHash:      common.BigToHash(big.NewInt(int64(2 * i))),
			BlockNumber:    int64(2 * i),
			BlockTimestamp: time.Now(),
			Topics:         topics,
			EventSig:       eventSig,
			Address:        common.Address{},
			TxHash:         common.BigToHash(big.NewInt(int64(2 * i))),
			Data:           common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000" + fmt.Sprintf("%d", i) + "000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000027100000000000000000000000000000000000000000000000000000000000000002"),
			CreatedAt:      time.Now(),
		})
		if i%2 == 0 {
			logs = append(logs, logpoller.Log{
				EvmChainId:     chainID,
				LogIndex:       0,
				BlockHash:      common.BigToHash(big.NewInt(int64(2*i + 1))),
				BlockNumber:    int64(2*i + 1),
				BlockTimestamp: time.Now(),
				Topics: [][]byte{
					common.FromHex("0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4"),
					common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000" + fmt.Sprintf("%d", i)),
				},
				EventSig:  VRFEmitterABI.Events["RandomWordsFulfilled"].ID,
				Address:   common.Address{},
				TxHash:    common.BigToHash(big.NewInt(int64(2*i + 1))),
				Data:      common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000069000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000001"),
				CreatedAt: time.Now(),
			})
		}
	}

	unfulfilled, _, fulfilled := listener.getUnfulfilled(logs, listener.l)
	require.Empty(t, unfulfilled)
	require.Len(t, fulfilled, 5)
}

func TestGetUnfulfilled_OneUnfulfilledVRFReq(t *testing.T) {
	t.Parallel()

	listener, chainID := SetupGetUnfulfilledTH(t)

	logs := []logpoller.Log{}
	for i := 0; i < 10; i++ {
		eventSig := EmitterABI.Events["Log1"].ID
		topics := [][]byte{
			common.FromHex("0x46692c0e59ca9cd1ad8f984a9d11715ec83424398b7eed4e05c8ce84662415a8"),
		}
		if i == 4 {
			eventSig = VRFEmitterABI.Events["RandomWordsRequested"].ID
			topics = [][]byte{
				common.FromHex("0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772"),
				common.FromHex("0xc0a6c424ac7157ae408398df7e5f4552091a69125d5dfcb7b8c2659029395bdf"),
				common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001"),
				common.FromHex("0x0000000000000000000000005ee3b50502b5c4c9184dcb281471a0614d4b2ef9"),
			}
		}
		logs = append(logs, logpoller.Log{
			EvmChainId:     chainID,
			LogIndex:       0,
			BlockHash:      common.BigToHash(big.NewInt(int64(2 * i))),
			BlockNumber:    int64(2 * i),
			BlockTimestamp: time.Now(),
			Topics:         topics,
			EventSig:       eventSig,
			Address:        common.Address{},
			TxHash:         common.BigToHash(big.NewInt(int64(2 * i))),
			Data:           common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000" + fmt.Sprintf("%d", i) + "000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000027100000000000000000000000000000000000000000000000000000000000000002"),
			CreatedAt:      time.Now(),
		})
	}

	unfulfilled, _, fulfilled := listener.getUnfulfilled(logs, listener.l)
	require.Equal(t, unfulfilled[0].RequestID().Int64(), big.NewInt(4).Int64())
	require.Len(t, unfulfilled, 1)
	require.Empty(t, fulfilled)
}

func TestGetUnfulfilled_SomeUnfulfilledVRFReq(t *testing.T) {
	t.Parallel()

	listener, chainID := SetupGetUnfulfilledTH(t)

	logs := []logpoller.Log{}
	for i := 0; i < 10; i++ {
		eventSig := EmitterABI.Events["Log1"].ID
		topics := [][]byte{
			common.FromHex("0x46692c0e59ca9cd1ad8f984a9d11715ec83424398b7eed4e05c8ce84662415a8"),
		}
		if i%2 == 0 {
			eventSig = VRFEmitterABI.Events["RandomWordsRequested"].ID
			topics = [][]byte{
				common.FromHex("0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772"),
				common.FromHex("0xc0a6c424ac7157ae408398df7e5f4552091a69125d5dfcb7b8c2659029395bdf"),
				common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001"),
				common.FromHex("0x0000000000000000000000005ee3b50502b5c4c9184dcb281471a0614d4b2ef9"),
			}
		}
		logs = append(logs, logpoller.Log{
			EvmChainId:     chainID,
			LogIndex:       0,
			BlockHash:      common.BigToHash(big.NewInt(int64(2 * i))),
			BlockNumber:    int64(2 * i),
			BlockTimestamp: time.Now(),
			Topics:         topics,
			EventSig:       eventSig,
			Address:        common.Address{},
			TxHash:         common.BigToHash(big.NewInt(int64(2 * i))),
			Data:           common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000" + fmt.Sprintf("%d", i) + "000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000027100000000000000000000000000000000000000000000000000000000000000002"),
			CreatedAt:      time.Now(),
		})
	}

	unfulfilled, _, fulfilled := listener.getUnfulfilled(logs, listener.l)
	require.Len(t, unfulfilled, 5)
	require.Len(t, fulfilled, 0)
	require.Equal(t, unfulfilled[0].RequestID().Int64(), big.NewInt(0).Int64())
	require.Equal(t, unfulfilled[1].RequestID().Int64(), big.NewInt(2).Int64())
	require.Equal(t, unfulfilled[2].RequestID().Int64(), big.NewInt(4).Int64())
	require.Equal(t, unfulfilled[3].RequestID().Int64(), big.NewInt(6).Int64())
	require.Equal(t, unfulfilled[4].RequestID().Int64(), big.NewInt(8).Int64())
}

func TestGetUnfulfilled_UnfulfilledNFulfilledVRFReqs(t *testing.T) {
	t.Parallel()

	listener, chainID := SetupGetUnfulfilledTH(t)

	logs := []logpoller.Log{}
	for i := 0; i < 10; i++ {
		eventSig := EmitterABI.Events["Log1"].ID
		topics := [][]byte{
			common.FromHex("0x46692c0e59ca9cd1ad8f984a9d11715ec83424398b7eed4e05c8ce84662415a8"),
		}
		if i%2 == 0 {
			eventSig = VRFEmitterABI.Events["RandomWordsRequested"].ID
			topics = [][]byte{
				common.FromHex("0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772"),
				common.FromHex("0xc0a6c424ac7157ae408398df7e5f4552091a69125d5dfcb7b8c2659029395bdf"),
				common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001"),
				common.FromHex("0x0000000000000000000000005ee3b50502b5c4c9184dcb281471a0614d4b2ef9"),
			}
		}
		logs = append(logs, logpoller.Log{
			EvmChainId:     chainID,
			LogIndex:       0,
			BlockHash:      common.BigToHash(big.NewInt(int64(2 * i))),
			BlockNumber:    int64(2 * i),
			BlockTimestamp: time.Now(),
			Topics:         topics,
			EventSig:       eventSig,
			Address:        common.Address{},
			TxHash:         common.BigToHash(big.NewInt(int64(2 * i))),
			Data:           common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000" + fmt.Sprintf("%d", i) + "000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000027100000000000000000000000000000000000000000000000000000000000000002"),
			CreatedAt:      time.Now(),
		})
		if i%2 == 0 && i < 6 {
			logs = append(logs, logpoller.Log{
				EvmChainId:     chainID,
				LogIndex:       0,
				BlockHash:      common.BigToHash(big.NewInt(int64(2*i + 1))),
				BlockNumber:    int64(2*i + 1),
				BlockTimestamp: time.Now(),
				Topics: [][]byte{
					common.FromHex("0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4"),
					common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000" + fmt.Sprintf("%d", i)),
				},
				EventSig:  VRFEmitterABI.Events["RandomWordsFulfilled"].ID,
				Address:   common.Address{},
				TxHash:    common.BigToHash(big.NewInt(int64(2*i + 1))),
				Data:      common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000069000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000001"),
				CreatedAt: time.Now(),
			})
		}
	}

	unfulfilled, _, fulfilled := listener.getUnfulfilled(logs, listener.l)
	require.Len(t, unfulfilled, 2)
	require.Len(t, fulfilled, 3)
	require.Equal(t, unfulfilled[0].RequestID().Int64(), big.NewInt(6).Int64())
	require.Equal(t, unfulfilled[1].RequestID().Int64(), big.NewInt(8).Int64())
}

/* Tests for getUnfulfilled: END */
