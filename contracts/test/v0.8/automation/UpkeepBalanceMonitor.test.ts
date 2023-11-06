import { ethers } from 'hardhat'
import chai, { assert, expect } from 'chai'
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { randomAddress } from '../../test-helpers/helpers'
import { loadFixture } from '@nomicfoundation/hardhat-network-helpers'
import { IKeeperRegistryMaster__factory as RegistryFactory } from '../../../typechain/factories/IKeeperRegistryMaster__factory'
import { IAutomationForwarder__factory as ForwarderFactory } from '../../../typechain/factories/IAutomationForwarder__factory'
import { UpkeepBalanceMonitor } from '../../../typechain/UpkeepBalanceMonitor'
import { LinkToken } from '../../../typechain/LinkToken'
import { BigNumber } from 'ethers'
import {
  deployMockContract,
  MockContract,
} from '@ethereum-waffle/mock-contract'

let owner: SignerWithAddress
let stranger: SignerWithAddress
let registry: MockContract
let forwarder: MockContract
let linkToken: LinkToken
let upkeepBalanceMonitor: UpkeepBalanceMonitor

const setup = async () => {
  const accounts = await ethers.getSigners()
  owner = accounts[0]
  stranger = accounts[1]

  const ltFactory = await ethers.getContractFactory(
    'src/v0.4/LinkToken.sol:LinkToken',
    owner,
  )
  linkToken = (await ltFactory.deploy()) as LinkToken
  const bmFactory = await ethers.getContractFactory(
    'UpkeepBalanceMonitor',
    owner,
  )
  upkeepBalanceMonitor = await bmFactory.deploy(linkToken.address, {
    maxBatchSize: 10,
    minPercentage: 120,
    targetPercentage: 300,
    maxTopUpAmount: ethers.utils.parseEther('100'),
  })
  registry = await deployMockContract(owner, RegistryFactory.abi)
  forwarder = await deployMockContract(owner, ForwarderFactory.abi)
  await forwarder.mock.getRegistry.returns(registry.address)
  await upkeepBalanceMonitor.setForwarder(forwarder.address)
  await linkToken
    .connect(owner)
    .transfer(upkeepBalanceMonitor.address, ethers.utils.parseEther('10000'))
  await upkeepBalanceMonitor
    .connect(owner)
    .setWatchList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12])
  for (let i = 1; i < 13; i++) {
    await registry.mock.getMinBalance.withArgs(i).returns(100)
    await registry.mock.getBalance.withArgs(i).returns(121) // all upkeeps are sufficiently funded
  }
}

describe('UpkeepBalanceMonitor', () => {
  beforeEach(async () => {
    await loadFixture(setup)
  })

  describe('constructor', () => {
    it('should set the initial values correctly', async () => {
      const config = await upkeepBalanceMonitor.getConfig()
      expect(config.maxBatchSize).to.equal(10)
      expect(config.minPercentage).to.equal(120)
      expect(config.targetPercentage).to.equal(300)
      expect(config.maxTopUpAmount).to.equal(ethers.utils.parseEther('100'))
    })
  })

  describe('setConfig', () => {
    it('should set config correctly', async () => {
      const newConfig = {
        maxBatchSize: 100,
        minPercentage: 150,
        targetPercentage: 500,
        maxTopUpAmount: 1,
      }
      await upkeepBalanceMonitor.connect(owner).setConfig(newConfig)
      const config = await upkeepBalanceMonitor.getConfig()
      expect(config.maxBatchSize).to.equal(newConfig.maxBatchSize)
      expect(config.minPercentage).to.equal(newConfig.minPercentage)
      expect(config.targetPercentage).to.equal(newConfig.targetPercentage)
      expect(config.maxTopUpAmount).to.equal(newConfig.maxTopUpAmount)
    })
  })

  describe('setForwarder', () => {
    it('should set the forwarder correctly', async () => {
      const expected = randomAddress()
      await upkeepBalanceMonitor.connect(owner).setForwarder(expected)
      const forwarderAddress = await upkeepBalanceMonitor.getForwarder()
      expect(forwarderAddress).to.equal(expected)
    })
  })

  describe('setWatchList', () => {
    it('should add addresses to the watchlist', async () => {
      const expected = [
        BigNumber.from(1),
        BigNumber.from(2),
        BigNumber.from(10),
      ]
      await upkeepBalanceMonitor.connect(owner).setWatchList(expected)
      const watchList = await upkeepBalanceMonitor.getWatchList()
      expect(watchList).to.deep.equal(expected)
    })
  })

  describe('withdraw', () => {
    it('should withdraw funds to a payee', async () => {
      const payee = randomAddress()
      const initialBalance = await linkToken.balanceOf(
        upkeepBalanceMonitor.address,
      )
      const withdrawAmount = 100
      await upkeepBalanceMonitor.connect(owner).withdraw(withdrawAmount, payee)
      const finalBalance = await linkToken.balanceOf(
        upkeepBalanceMonitor.address,
      )
      const payeeBalance = await linkToken.balanceOf(payee)
      expect(finalBalance).to.equal(initialBalance.sub(withdrawAmount))
      expect(payeeBalance).to.equal(withdrawAmount)
    })
  })

  describe('pause and unpause', () => {
    it('should pause and unpause the contract', async () => {
      await upkeepBalanceMonitor.connect(owner).pause()
      expect(await upkeepBalanceMonitor.paused()).to.be.true
      await upkeepBalanceMonitor.connect(owner).unpause()
    })
  })

  describe('getUnderfundedUpkeeps', () => {
    it('should find the underfunded upkeeps', async () => {
      let [upkeepIDs, topUpAmounts] =
        await upkeepBalanceMonitor.getUnderfundedUpkeeps()
      expect(upkeepIDs.length).to.equal(0)
      expect(topUpAmounts.length).to.equal(0)
      // update the balance for some upkeeps
      await registry.mock.getBalance.withArgs(2).returns(120)
      await registry.mock.getBalance.withArgs(4).returns(15)
      await registry.mock.getBalance.withArgs(5).returns(0)
      ;[upkeepIDs, topUpAmounts] =
        await upkeepBalanceMonitor.getUnderfundedUpkeeps()
      expect(upkeepIDs).to.deep.equal([2, 4, 5].map(BigNumber.from))
      expect(topUpAmounts).to.deep.equal([180, 285, 300].map(BigNumber.from))
      // update all to need funding
      for (let i = 1; i < 13; i++) {
        await registry.mock.getBalance.withArgs(i).returns(0)
      }
      // test that only up to max batch size are included in the list
      ;[upkeepIDs, topUpAmounts] =
        await upkeepBalanceMonitor.getUnderfundedUpkeeps()
      expect(upkeepIDs.length).to.equal(10)
      expect(topUpAmounts.length).to.equal(10)
    })
  })
})