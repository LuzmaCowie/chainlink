// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type CommonAddressAndWeight struct {
	Addr   common.Address
	Weight *big.Int
}

var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierProxyAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadVerification\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"CannotDeactivateLatestConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DigestEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"DigestInactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"DigestNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSigners\",\"type\":\"uint256\"}],\"name\":\"ExcessSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FaultToleranceMustBePositive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedIdEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"InactiveFeed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedNumSigners\",\"type\":\"uint256\"}],\"name\":\"IncorrectSignatureCount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSigners\",\"type\":\"uint256\"}],\"name\":\"InsufficientSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"InvalidFeed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ssLength\",\"type\":\"uint256\"}],\"name\":\"MismatchedSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"ConfigActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"ConfigDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"FeedActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"FeedDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"ReportVerified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"activateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"activateFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"deactivateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"deactivateFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.AddressAndWeight[]\",\"name\":\"recipientAddressesAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isVerifier\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signedReport\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002150380380620021508339810160408190526200003491620001a6565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000fb565b5050506001600160a01b038116620000e95760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0316608052620001d8565b336001600160a01b03821603620001555760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001b957600080fd5b81516001600160a01b0381168114620001d157600080fd5b9392505050565b608051611f55620001fb6000396000818161031d015261091e0152611f556000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806379ba50971161008c578063b70d929d11610066578063b70d929d14610221578063ded6307c14610280578063e84f128e14610293578063f2fde38b146102f057600080fd5b806379ba5097146101de5780638da5cb5b146101e657806394d959801461020e57600080fd5b80633dd86430116100bd5780633dd86430146101a357806352ba27d6146101b8578063564a0a7a146101cb57600080fd5b806301ffc9a7146100e4578063181f5a771461014e5780633d3ac1b514610190575b600080fd5b6101396100f2366004611487565b7fffffffff00000000000000000000000000000000000000000000000000000000167f3d3ac1b5000000000000000000000000000000000000000000000000000000001490565b60405190151581526020015b60405180910390f35b60408051808201909152600e81527f566572696669657220312e312e3000000000000000000000000000000000000060208201525b6040516101459190611534565b61018361019e366004611570565b610303565b6101b66101b13660046115f1565b61049d565b005b6101b66101c6366004611902565b61054f565b6101b66101d93660046115f1565b610a3f565b6101b6610b00565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610145565b6101b661021c3660046119ff565b610bfd565b61025d61022f3660046115f1565b6000908152600260205260408120600181015490549192909168010000000000000000900463ffffffff1690565b604080519315158452602084019290925263ffffffff1690820152606001610145565b6101b661028e3660046119ff565b610d5e565b6102cd6102a13660046115f1565b6000908152600260205260409020805460019091015463ffffffff808316936401000000009093041691565b6040805163ffffffff948516815293909216602084015290820152606001610145565b6101b66102fe366004611a21565b610e6f565b60603373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610374576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080808080610386888a018a611a3c565b9450945094509450945060008461039c90611b17565b60008181526002602052604090208054919250906c01000000000000000000000000900460ff1615610402576040517f36dbe748000000000000000000000000000000000000000000000000000000008152600481018390526024015b60405180910390fd5b8651600081815260028301602052604090206104218483898985610e83565b61042b8984610f7f565b87516020890120610440818b8a8a8a87610fe7565b60405173ffffffffffffffffffffffffffffffffffffffff8d16815285907f58ca9502e98a536e06e72d680fcc251e5d10b72291a281665a2c2dc0ac30fcc59060200160405180910390a250969c9b505050505050505050505050565b6104a5611263565b60008181526002602052604081208054909163ffffffff90911690036104fa576040517fa25b0b96000000000000000000000000000000000000000000000000000000008152600481018390526024016103f9565b80547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff16815560405182907ff438564f793525caa89c6e3a26d41e16aa39d1e589747595751e3f3df75cb2b490600090a25050565b86518560ff168060000361058f576040517f0743bae600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601f8211156105d4576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101839052601f60248201526044016103f9565b6105df816003611bba565b821161063757816105f1826003611bba565b6105fc906001611bf7565b6040517f9dd9e6d8000000000000000000000000000000000000000000000000000000008152600481019290925260248201526044016103f9565b61063f611263565b60008a81526002602052604081208054909163ffffffff90911690829061066583611c10565b82546101009290920a63ffffffff81810219909316918316021790915582546000925061069a918e91168d8d8d8d8d8d6112e6565b6000818152600284016020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8d16176101001790559091505b8b518160ff1610156108db5760008c8260ff168151811061070057610700611b5c565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610770576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806000858152600287016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452600190810190925290912054610100900460ff16908111156107c3576107c3611c33565b14801591506107fe576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff8416815260208101600190526000858152600287016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684526001908101835292208351815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00821681178355928501519193919284927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090921617906101009084908111156108c0576108c0611c33565b02179055509050505050806108d490611c62565b90506106dd565b5060018201546040517f589ede2800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163589ede2891610956919085908a90600401611c81565b600060405180830381600087803b15801561097057600080fd5b505af1158015610984573d6000803e3d6000fd5b505050508b7fa23a88453230b183877098801ff5a8f771a120e2573eea559ce6c4c2e305a4da8360000160049054906101000a900463ffffffff16838560000160009054906101000a900463ffffffff168f8f8f8f8f8f6040516109f099989796959493929190611d78565b60405180910390a281547fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff1664010000000063ffffffff43160217825560019091015550505050505050505050565b610a47611263565b60008181526002602052604081208054909163ffffffff9091169003610a9c576040517fa25b0b96000000000000000000000000000000000000000000000000000000008152600481018390526024016103f9565b80547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff166c0100000000000000000000000017815560405182907ffc4f79b8c65b6be1773063461984c0974400d1e99654c79477a092ace83fd06190600090a25050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610b81576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103f9565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610c05611263565b600082815260026020526040902081610c4a576040517fe332262700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260028201602052604081205460ff169003610ca0576040517f8bca631100000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044016103f9565b80600101548203610ce7576040517fa403c01600000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044016103f9565b60008281526002820160205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555183907f0e173bea63a8c59ec70bf87043f2a729693790183f16a1a54b705de9e989cc4c90610d519085815260200190565b60405180910390a2505050565b610d66611263565b600082815260026020526040902081610dab576040517fe332262700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260028201602052604081205460ff169003610e01576040517f8bca631100000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044016103f9565b60008281526002820160205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555183907f54f8872b9b94ebea6577f33576d55847bd8ea22641ccc886b965f6e50bfe774690610d519085815260200190565b610e77611263565b610e8081611392565b50565b8054600090610e969060ff166001611e0e565b8254909150610100900460ff16610ee3576040517ffc10a28300000000000000000000000000000000000000000000000000000000815260048101879052602481018690526044016103f9565b8060ff16845114610f2f5783516040517f5348a282000000000000000000000000000000000000000000000000000000008152600481019190915260ff821660248201526044016103f9565b8251845114610f7757835183516040517ff0d31408000000000000000000000000000000000000000000000000000000008152600481019290925260248201526044016103f9565b505050505050565b6020820151815463ffffffff600883901c81169168010000000000000000900416811115610fe15782547fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff166801000000000000000063ffffffff8316021783555b50505050565b60008686604051602001610ffc929190611e27565b6040516020818303038152906040528051906020012090506000611030604080518082019091526000808252602082015290565b8651600090815b818110156111fb5760018689836020811061105457611054611b5c565b61106191901a601b611e0e565b8c848151811061107357611073611b5c565b60200260200101518c858151811061108d5761108d611b5c565b6020026020010151604051600081526020016040526040516110cb949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156110ed573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526001808d01602090815291859020848601909552845460ff80821686529399509395509085019261010090049091169081111561117257611172611c33565b600181111561118357611183611c33565b90525093506001846020015160018111156111a0576111a0611c33565b146111d7576040517f4df18f0700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836000015160080260ff166001901b85019450806111f490611e63565b9050611037565b50837e01010101010101010101010101010101010101010101010101010101010101851614611256576040517f4df18f0700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146112e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103f9565b565b6000808946308b8b8b8b8b8b8b60405160200161130c9a99989796959493929190611e9b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e060000000000000000000000000000000000000000000000000000000000001791505098975050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611411576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103f9565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561149957600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146114c957600080fd5b9392505050565b6000815180845260005b818110156114f6576020818501810151868301820152016114da565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006114c960208301846114d0565b803573ffffffffffffffffffffffffffffffffffffffff8116811461156b57600080fd5b919050565b60008060006040848603121561158557600080fd5b833567ffffffffffffffff8082111561159d57600080fd5b818601915086601f8301126115b157600080fd5b8135818111156115c057600080fd5b8760208285010111156115d257600080fd5b6020928301955093506115e89186019050611547565b90509250925092565b60006020828403121561160357600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561165c5761165c61160a565b60405290565b6040516060810167ffffffffffffffff8111828210171561165c5761165c61160a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156116cc576116cc61160a565b604052919050565b600067ffffffffffffffff8211156116ee576116ee61160a565b5060051b60200190565b600082601f83011261170957600080fd5b8135602061171e611719836116d4565b611685565b82815260059290921b8401810191818101908684111561173d57600080fd5b8286015b8481101561175f5761175281611547565b8352918301918301611741565b509695505050505050565b600082601f83011261177b57600080fd5b8135602061178b611719836116d4565b82815260059290921b840181019181810190868411156117aa57600080fd5b8286015b8481101561175f57803583529183019183016117ae565b803560ff8116811461156b57600080fd5b600082601f8301126117e757600080fd5b813567ffffffffffffffff8111156118015761180161160a565b61183260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611685565b81815284602083860101111561184757600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff8116811461156b57600080fd5b600082601f83011261188d57600080fd5b8135602061189d611719836116d4565b82815260069290921b840181019181810190868411156118bc57600080fd5b8286015b8481101561175f57604081890312156118d95760008081fd5b6118e1611639565b6118ea82611547565b815281850135858201528352918301916040016118c0565b600080600080600080600080610100898b03121561191f57600080fd5b88359750602089013567ffffffffffffffff8082111561193e57600080fd5b61194a8c838d016116f8565b985060408b013591508082111561196057600080fd5b61196c8c838d0161176a565b975061197a60608c016117c5565b965060808b013591508082111561199057600080fd5b61199c8c838d016117d6565b95506119aa60a08c01611864565b945060c08b01359150808211156119c057600080fd5b6119cc8c838d016117d6565b935060e08b01359150808211156119e257600080fd5b506119ef8b828c0161187c565b9150509295985092959890939650565b60008060408385031215611a1257600080fd5b50508035926020909101359150565b600060208284031215611a3357600080fd5b6114c982611547565b600080600080600060e08688031215611a5457600080fd5b86601f870112611a6357600080fd5b611a6b611662565b806060880189811115611a7d57600080fd5b885b81811015611a97578035845260209384019301611a7f565b5090965035905067ffffffffffffffff80821115611ab457600080fd5b611ac089838a016117d6565b95506080880135915080821115611ad657600080fd5b611ae289838a0161176a565b945060a0880135915080821115611af857600080fd5b50611b058882890161176a565b9598949750929560c001359392505050565b80516020808301519190811015611b56577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8160200360031b1b821691505b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611bf257611bf2611b8b565b500290565b80820180821115611c0a57611c0a611b8b565b92915050565b600063ffffffff808316818103611c2957611c29611b8b565b6001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600060ff821660ff8103611c7857611c78611b8b565b60010192915050565b600060608201858352602085818501526040606081860152828651808552608087019150838801945060005b81811015611ce8578551805173ffffffffffffffffffffffffffffffffffffffff168452850151858401529484019491830191600101611cad565b50909998505050505050505050565b600081518084526020808501945080840160005b83811015611d3d57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611d0b565b509495945050505050565b600081518084526020808501945080840160005b83811015611d3d57815187529582019590820190600101611d5c565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152611da88184018a611cf7565b90508281036080840152611dbc8189611d48565b905060ff871660a084015282810360c0840152611dd981876114d0565b905067ffffffffffffffff851660e0840152828103610100840152611dfe81856114d0565b9c9b505050505050505050505050565b60ff8181168382160190811115611c0a57611c0a611b8b565b828152600060208083018460005b6003811015611e5257815183529183019190830190600101611e35565b505050506080820190509392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e9457611e94611b8b565b5060010190565b60006101408c83528b602084015273ffffffffffffffffffffffffffffffffffffffff8b16604084015267ffffffffffffffff808b166060850152816080850152611ee88285018b611cf7565b915083820360a0850152611efc828a611d48565b915060ff881660c085015283820360e0850152611f1982886114d0565b9086166101008501528381036101208501529050611f3781856114d0565b9d9c5050505050505050505050505056fea164736f6c6343000810000a",
}

var VerifierABI = VerifierMetaData.ABI

var VerifierBin = VerifierMetaData.Bin

func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, verifierProxyAddr common.Address) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierBin), backend, verifierProxyAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

type Verifier struct {
	address common.Address
	abi     abi.ABI
	VerifierCaller
	VerifierTransactor
	VerifierFilterer
}

type VerifierCaller struct {
	contract *bind.BoundContract
}

type VerifierTransactor struct {
	contract *bind.BoundContract
}

type VerifierFilterer struct {
	contract *bind.BoundContract
}

type VerifierSession struct {
	Contract     *Verifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierCallerSession struct {
	Contract *VerifierCaller
	CallOpts bind.CallOpts
}

type VerifierTransactorSession struct {
	Contract     *VerifierTransactor
	TransactOpts bind.TransactOpts
}

type VerifierRaw struct {
	Contract *Verifier
}

type VerifierCallerRaw struct {
	Contract *VerifierCaller
}

type VerifierTransactorRaw struct {
	Contract *VerifierTransactor
}

func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{address: address, abi: abi, VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

func (_Verifier *VerifierCaller) LatestConfigDetails(opts *bind.CallOpts, feedId [32]byte) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "latestConfigDetails", feedId)

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_Verifier *VerifierSession) LatestConfigDetails(feedId [32]byte) (LatestConfigDetails,

	error) {
	return _Verifier.Contract.LatestConfigDetails(&_Verifier.CallOpts, feedId)
}

func (_Verifier *VerifierCallerSession) LatestConfigDetails(feedId [32]byte) (LatestConfigDetails,

	error) {
	return _Verifier.Contract.LatestConfigDetails(&_Verifier.CallOpts, feedId)
}

func (_Verifier *VerifierCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts, feedId [32]byte) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "latestConfigDigestAndEpoch", feedId)

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_Verifier *VerifierSession) LatestConfigDigestAndEpoch(feedId [32]byte) (LatestConfigDigestAndEpoch,

	error) {
	return _Verifier.Contract.LatestConfigDigestAndEpoch(&_Verifier.CallOpts, feedId)
}

func (_Verifier *VerifierCallerSession) LatestConfigDigestAndEpoch(feedId [32]byte) (LatestConfigDigestAndEpoch,

	error) {
	return _Verifier.Contract.LatestConfigDigestAndEpoch(&_Verifier.CallOpts, feedId)
}

func (_Verifier *VerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Verifier *VerifierSession) Owner() (common.Address, error) {
	return _Verifier.Contract.Owner(&_Verifier.CallOpts)
}

func (_Verifier *VerifierCallerSession) Owner() (common.Address, error) {
	return _Verifier.Contract.Owner(&_Verifier.CallOpts)
}

func (_Verifier *VerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Verifier *VerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Verifier.Contract.SupportsInterface(&_Verifier.CallOpts, interfaceId)
}

func (_Verifier *VerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Verifier.Contract.SupportsInterface(&_Verifier.CallOpts, interfaceId)
}

func (_Verifier *VerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Verifier *VerifierSession) TypeAndVersion() (string, error) {
	return _Verifier.Contract.TypeAndVersion(&_Verifier.CallOpts)
}

func (_Verifier *VerifierCallerSession) TypeAndVersion() (string, error) {
	return _Verifier.Contract.TypeAndVersion(&_Verifier.CallOpts)
}

func (_Verifier *VerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "acceptOwnership")
}

func (_Verifier *VerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _Verifier.Contract.AcceptOwnership(&_Verifier.TransactOpts)
}

func (_Verifier *VerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Verifier.Contract.AcceptOwnership(&_Verifier.TransactOpts)
}

func (_Verifier *VerifierTransactor) ActivateConfig(opts *bind.TransactOpts, feedId [32]byte, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "activateConfig", feedId, configDigest)
}

func (_Verifier *VerifierSession) ActivateConfig(feedId [32]byte, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.ActivateConfig(&_Verifier.TransactOpts, feedId, configDigest)
}

func (_Verifier *VerifierTransactorSession) ActivateConfig(feedId [32]byte, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.ActivateConfig(&_Verifier.TransactOpts, feedId, configDigest)
}

func (_Verifier *VerifierTransactor) ActivateFeed(opts *bind.TransactOpts, feedId [32]byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "activateFeed", feedId)
}

func (_Verifier *VerifierSession) ActivateFeed(feedId [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.ActivateFeed(&_Verifier.TransactOpts, feedId)
}

func (_Verifier *VerifierTransactorSession) ActivateFeed(feedId [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.ActivateFeed(&_Verifier.TransactOpts, feedId)
}

func (_Verifier *VerifierTransactor) DeactivateConfig(opts *bind.TransactOpts, feedId [32]byte, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "deactivateConfig", feedId, configDigest)
}

func (_Verifier *VerifierSession) DeactivateConfig(feedId [32]byte, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.DeactivateConfig(&_Verifier.TransactOpts, feedId, configDigest)
}

func (_Verifier *VerifierTransactorSession) DeactivateConfig(feedId [32]byte, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.DeactivateConfig(&_Verifier.TransactOpts, feedId, configDigest)
}

func (_Verifier *VerifierTransactor) DeactivateFeed(opts *bind.TransactOpts, feedId [32]byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "deactivateFeed", feedId)
}

func (_Verifier *VerifierSession) DeactivateFeed(feedId [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.DeactivateFeed(&_Verifier.TransactOpts, feedId)
}

func (_Verifier *VerifierTransactorSession) DeactivateFeed(feedId [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.DeactivateFeed(&_Verifier.TransactOpts, feedId)
}

func (_Verifier *VerifierTransactor) SetConfig(opts *bind.TransactOpts, feedId [32]byte, signers []common.Address, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "setConfig", feedId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, recipientAddressesAndWeights)
}

func (_Verifier *VerifierSession) SetConfig(feedId [32]byte, signers []common.Address, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _Verifier.Contract.SetConfig(&_Verifier.TransactOpts, feedId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, recipientAddressesAndWeights)
}

func (_Verifier *VerifierTransactorSession) SetConfig(feedId [32]byte, signers []common.Address, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _Verifier.Contract.SetConfig(&_Verifier.TransactOpts, feedId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, recipientAddressesAndWeights)
}

func (_Verifier *VerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "transferOwnership", to)
}

func (_Verifier *VerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.TransferOwnership(&_Verifier.TransactOpts, to)
}

func (_Verifier *VerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.TransferOwnership(&_Verifier.TransactOpts, to)
}

func (_Verifier *VerifierTransactor) Verify(opts *bind.TransactOpts, signedReport []byte, sender common.Address) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "verify", signedReport, sender)
}

func (_Verifier *VerifierSession) Verify(signedReport []byte, sender common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.Verify(&_Verifier.TransactOpts, signedReport, sender)
}

func (_Verifier *VerifierTransactorSession) Verify(signedReport []byte, sender common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.Verify(&_Verifier.TransactOpts, signedReport, sender)
}

type VerifierConfigActivatedIterator struct {
	Event *VerifierConfigActivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierConfigActivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierConfigActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierConfigActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierConfigActivatedIterator) Error() error {
	return it.fail
}

func (it *VerifierConfigActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierConfigActivated struct {
	FeedId       [32]byte
	ConfigDigest [32]byte
	Raw          types.Log
}

func (_Verifier *VerifierFilterer) FilterConfigActivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierConfigActivatedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ConfigActivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierConfigActivatedIterator{contract: _Verifier.contract, event: "ConfigActivated", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchConfigActivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigActivated, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ConfigActivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierConfigActivated)
				if err := _Verifier.contract.UnpackLog(event, "ConfigActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseConfigActivated(log types.Log) (*VerifierConfigActivated, error) {
	event := new(VerifierConfigActivated)
	if err := _Verifier.contract.UnpackLog(event, "ConfigActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierConfigDeactivatedIterator struct {
	Event *VerifierConfigDeactivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierConfigDeactivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierConfigDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierConfigDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierConfigDeactivatedIterator) Error() error {
	return it.fail
}

func (it *VerifierConfigDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierConfigDeactivated struct {
	FeedId       [32]byte
	ConfigDigest [32]byte
	Raw          types.Log
}

func (_Verifier *VerifierFilterer) FilterConfigDeactivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierConfigDeactivatedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ConfigDeactivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierConfigDeactivatedIterator{contract: _Verifier.contract, event: "ConfigDeactivated", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchConfigDeactivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigDeactivated, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ConfigDeactivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierConfigDeactivated)
				if err := _Verifier.contract.UnpackLog(event, "ConfigDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseConfigDeactivated(log types.Log) (*VerifierConfigDeactivated, error) {
	event := new(VerifierConfigDeactivated)
	if err := _Verifier.contract.UnpackLog(event, "ConfigDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierConfigSetIterator struct {
	Event *VerifierConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierConfigSet struct {
	FeedId                    [32]byte
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	OffchainTransmitters      [][32]byte
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_Verifier *VerifierFilterer) FilterConfigSet(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierConfigSetIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ConfigSet", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierConfigSetIterator{contract: _Verifier.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierConfigSet, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ConfigSet", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierConfigSet)
				if err := _Verifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseConfigSet(log types.Log) (*VerifierConfigSet, error) {
	event := new(VerifierConfigSet)
	if err := _Verifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierFeedActivatedIterator struct {
	Event *VerifierFeedActivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierFeedActivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierFeedActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierFeedActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierFeedActivatedIterator) Error() error {
	return it.fail
}

func (it *VerifierFeedActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierFeedActivated struct {
	FeedId [32]byte
	Raw    types.Log
}

func (_Verifier *VerifierFilterer) FilterFeedActivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierFeedActivatedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "FeedActivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierFeedActivatedIterator{contract: _Verifier.contract, event: "FeedActivated", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchFeedActivated(opts *bind.WatchOpts, sink chan<- *VerifierFeedActivated, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "FeedActivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierFeedActivated)
				if err := _Verifier.contract.UnpackLog(event, "FeedActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseFeedActivated(log types.Log) (*VerifierFeedActivated, error) {
	event := new(VerifierFeedActivated)
	if err := _Verifier.contract.UnpackLog(event, "FeedActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierFeedDeactivatedIterator struct {
	Event *VerifierFeedDeactivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierFeedDeactivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierFeedDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierFeedDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierFeedDeactivatedIterator) Error() error {
	return it.fail
}

func (it *VerifierFeedDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierFeedDeactivated struct {
	FeedId [32]byte
	Raw    types.Log
}

func (_Verifier *VerifierFilterer) FilterFeedDeactivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierFeedDeactivatedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "FeedDeactivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierFeedDeactivatedIterator{contract: _Verifier.contract, event: "FeedDeactivated", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchFeedDeactivated(opts *bind.WatchOpts, sink chan<- *VerifierFeedDeactivated, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "FeedDeactivated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierFeedDeactivated)
				if err := _Verifier.contract.UnpackLog(event, "FeedDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseFeedDeactivated(log types.Log) (*VerifierFeedDeactivated, error) {
	event := new(VerifierFeedDeactivated)
	if err := _Verifier.contract.UnpackLog(event, "FeedDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierOwnershipTransferRequestedIterator struct {
	Event *VerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Verifier *VerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnershipTransferRequestedIterator{contract: _Verifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierOwnershipTransferRequested)
				if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierOwnershipTransferRequested, error) {
	event := new(VerifierOwnershipTransferRequested)
	if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierOwnershipTransferredIterator struct {
	Event *VerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Verifier *VerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnershipTransferredIterator{contract: _Verifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierOwnershipTransferred)
				if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierOwnershipTransferred, error) {
	event := new(VerifierOwnershipTransferred)
	if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierReportVerifiedIterator struct {
	Event *VerifierReportVerified

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierReportVerifiedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierReportVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VerifierReportVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VerifierReportVerifiedIterator) Error() error {
	return it.fail
}

func (it *VerifierReportVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierReportVerified struct {
	FeedId    [32]byte
	Requester common.Address
	Raw       types.Log
}

func (_Verifier *VerifierFilterer) FilterReportVerified(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierReportVerifiedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ReportVerified", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierReportVerifiedIterator{contract: _Verifier.contract, event: "ReportVerified", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchReportVerified(opts *bind.WatchOpts, sink chan<- *VerifierReportVerified, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ReportVerified", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierReportVerified)
				if err := _Verifier.contract.UnpackLog(event, "ReportVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Verifier *VerifierFilterer) ParseReportVerified(log types.Log) (*VerifierReportVerified, error) {
	event := new(VerifierReportVerified)
	if err := _Verifier.contract.UnpackLog(event, "ReportVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_Verifier *Verifier) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Verifier.abi.Events["ConfigActivated"].ID:
		return _Verifier.ParseConfigActivated(log)
	case _Verifier.abi.Events["ConfigDeactivated"].ID:
		return _Verifier.ParseConfigDeactivated(log)
	case _Verifier.abi.Events["ConfigSet"].ID:
		return _Verifier.ParseConfigSet(log)
	case _Verifier.abi.Events["FeedActivated"].ID:
		return _Verifier.ParseFeedActivated(log)
	case _Verifier.abi.Events["FeedDeactivated"].ID:
		return _Verifier.ParseFeedDeactivated(log)
	case _Verifier.abi.Events["OwnershipTransferRequested"].ID:
		return _Verifier.ParseOwnershipTransferRequested(log)
	case _Verifier.abi.Events["OwnershipTransferred"].ID:
		return _Verifier.ParseOwnershipTransferred(log)
	case _Verifier.abi.Events["ReportVerified"].ID:
		return _Verifier.ParseReportVerified(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierConfigActivated) Topic() common.Hash {
	return common.HexToHash("0x54f8872b9b94ebea6577f33576d55847bd8ea22641ccc886b965f6e50bfe7746")
}

func (VerifierConfigDeactivated) Topic() common.Hash {
	return common.HexToHash("0x0e173bea63a8c59ec70bf87043f2a729693790183f16a1a54b705de9e989cc4c")
}

func (VerifierConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa23a88453230b183877098801ff5a8f771a120e2573eea559ce6c4c2e305a4da")
}

func (VerifierFeedActivated) Topic() common.Hash {
	return common.HexToHash("0xf438564f793525caa89c6e3a26d41e16aa39d1e589747595751e3f3df75cb2b4")
}

func (VerifierFeedDeactivated) Topic() common.Hash {
	return common.HexToHash("0xfc4f79b8c65b6be1773063461984c0974400d1e99654c79477a092ace83fd061")
}

func (VerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VerifierReportVerified) Topic() common.Hash {
	return common.HexToHash("0x58ca9502e98a536e06e72d680fcc251e5d10b72291a281665a2c2dc0ac30fcc5")
}

func (_Verifier *Verifier) Address() common.Address {
	return _Verifier.address
}

type VerifierInterface interface {
	LatestConfigDetails(opts *bind.CallOpts, feedId [32]byte) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts, feedId [32]byte) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateConfig(opts *bind.TransactOpts, feedId [32]byte, configDigest [32]byte) (*types.Transaction, error)

	ActivateFeed(opts *bind.TransactOpts, feedId [32]byte) (*types.Transaction, error)

	DeactivateConfig(opts *bind.TransactOpts, feedId [32]byte, configDigest [32]byte) (*types.Transaction, error)

	DeactivateFeed(opts *bind.TransactOpts, feedId [32]byte) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, feedId [32]byte, signers []common.Address, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Verify(opts *bind.TransactOpts, signedReport []byte, sender common.Address) (*types.Transaction, error)

	FilterConfigActivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierConfigActivatedIterator, error)

	WatchConfigActivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigActivated, feedId [][32]byte) (event.Subscription, error)

	ParseConfigActivated(log types.Log) (*VerifierConfigActivated, error)

	FilterConfigDeactivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierConfigDeactivatedIterator, error)

	WatchConfigDeactivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigDeactivated, feedId [][32]byte) (event.Subscription, error)

	ParseConfigDeactivated(log types.Log) (*VerifierConfigDeactivated, error)

	FilterConfigSet(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierConfigSet, feedId [][32]byte) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VerifierConfigSet, error)

	FilterFeedActivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierFeedActivatedIterator, error)

	WatchFeedActivated(opts *bind.WatchOpts, sink chan<- *VerifierFeedActivated, feedId [][32]byte) (event.Subscription, error)

	ParseFeedActivated(log types.Log) (*VerifierFeedActivated, error)

	FilterFeedDeactivated(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierFeedDeactivatedIterator, error)

	WatchFeedDeactivated(opts *bind.WatchOpts, sink chan<- *VerifierFeedDeactivated, feedId [][32]byte) (event.Subscription, error)

	ParseFeedDeactivated(log types.Log) (*VerifierFeedDeactivated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierOwnershipTransferred, error)

	FilterReportVerified(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierReportVerifiedIterator, error)

	WatchReportVerified(opts *bind.WatchOpts, sink chan<- *VerifierReportVerified, feedId [][32]byte) (event.Subscription, error)

	ParseReportVerified(log types.Log) (*VerifierReportVerified, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
