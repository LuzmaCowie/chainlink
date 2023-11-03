// SPDX-License-Identifier: MIT

pragma solidity 0.8.6;

import {ConfirmedOwner} from "../../shared/access/ConfirmedOwner.sol";
import {IKeeperRegistryMaster} from "../interfaces/v2_1/IKeeperRegistryMaster.sol";
import {IAutomationForwarder} from "../interfaces/IAutomationForwarder.sol";
import {IAutomationRegistryConsumer} from "../interfaces/IAutomationRegistryConsumer.sol";
import {LinkTokenInterface} from "../../shared/interfaces/LinkTokenInterface.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";

/// @title The UpkeepBalanceMonitor contract.
/// @notice A keeper-compatible contract that monitors and funds Chainlink Automation upkeeps.
contract UpkeepBalanceMonitor is ConfirmedOwner, Pausable {
  LinkTokenInterface public immutable LINK_TOKEN;

  event FundsWithdrawn(uint256 amountWithdrawn, address payee);
  event ConfigSet(uint96 minPercentage, uint96 targetPercentage);
  event TopUpFailed(uint256 indexed upkeepId);
  event TopUpSucceeded(uint256 indexed upkeepId, uint96 amount);

  error DuplicateSubcriptionId(uint256 duplicate);
  error InvalidConfig();
  error LengthMismatch();
  error OnlyKeeperRegistry();

  struct Config {
    uint8 maxBatchSize;
    uint24 minPercentage;
    uint24 targetPercentage; // max target is 160K times the min balance
    uint96 maxTopUpAmount;
  }

  IKeeperRegistryMaster private s_registry;
  uint256[] private s_watchList; // the watchlist on which subscriptions are stored
  Config private s_config;

  /// @param linkTokenAddress the Link token address
  /// @param keeperRegistryAddress the address of the keeper registry contract
  /// @param maxBatchSize the maximum number of upkeeps to fund in a single transaction
  /// @param minPercentage the percentage of the min balance at which to trigger top ups
  /// @param targetPercentage the percentage of the min balance to target during top ups
  /// @param maxTopUpAmount the maximum amount to top up an upkeep
  constructor(
    address linkTokenAddress,
    IKeeperRegistryMaster keeperRegistryAddress,
    uint8 maxBatchSize,
    uint24 minPercentage,
    uint24 targetPercentage,
    uint96 maxTopUpAmount
  ) ConfirmedOwner(msg.sender) {
    require(linkTokenAddress != address(0));
    LINK_TOKEN = LinkTokenInterface(linkTokenAddress);
    setConfig(maxBatchSize, minPercentage, targetPercentage, maxTopUpAmount);
    LinkTokenInterface(linkTokenAddress).approve(address(keeperRegistryAddress), type(uint256).max);
  }

  // ================================================================
  // |                    AUTOMATION COMPATIBLE                     |
  // ================================================================

  /// @notice Gets list of upkeeps ids that are underfunded and returns a keeper-compatible payload.
  /// @return upkeepNeeded signals if upkeep is needed, performData is an abi encoded list of subscription ids that need funds
  function checkUpkeep(
    bytes calldata
  ) external view whenNotPaused returns (bool upkeepNeeded, bytes memory performData) {
    (uint256[] memory needsFunding, uint256[] memory topUpAmounts) = getUnderfundedUpkeeps();
    upkeepNeeded = needsFunding.length > 0;
    performData = abi.encode(needsFunding, topUpAmounts);
    return (upkeepNeeded, performData);
  }

  /// @notice Called by the keeper to send funds to underfunded addresses.
  /// @param performData the abi encoded list of addresses to fund
  function performUpkeep(bytes calldata performData) external whenNotPaused {
    // if (msg.sender != address(s_registry)) revert OnlyKeeperRegistry();
    // TODO - forwarder contract
    (uint256[] memory needsFunding, uint96[] memory topUpAmounts) = abi.decode(performData, (uint256[], uint96[]));
    if (needsFunding.length != topUpAmounts.length) revert LengthMismatch();
    IAutomationForwarder forwarder = IAutomationForwarder(msg.sender);
    IAutomationRegistryConsumer registry = forwarder.getRegistry();
    uint256 contractBalance = LINK_TOKEN.balanceOf(address(this));
    for (uint256 i = 0; i < needsFunding.length; i++) {
      try registry.addFunds(needsFunding[i], topUpAmounts[i]) {
        emit TopUpSucceeded(needsFunding[i], topUpAmounts[i]);
      } catch {
        emit TopUpFailed(needsFunding[i]);
      }
    }
  }

  // ================================================================
  // |                            ADMIN                             |
  // ================================================================

  /// @notice Withdraws the contract balance in LINK.
  /// @param amount the amount of LINK (in juels) to withdraw
  /// @param payee the address to pay
  function withdraw(uint256 amount, address payee) external onlyOwner {
    require(payee != address(0));
    LINK_TOKEN.transfer(payee, amount);
    emit FundsWithdrawn(amount, payee);
  }

  /// @notice Pause the contract, which prevents executing performUpkeep.
  function pause() external onlyOwner {
    _pause();
  }

  /// @notice Unpause the contract.
  function unpause() external onlyOwner {
    _unpause();
  }

  // ================================================================
  // |                           SETTERS                            |
  // ================================================================

  /// @notice Sets the list of upkeeps to watch and their funding parameters.
  /// @param watchlist the list of subscription ids to watch
  function setWatchList(uint256[] calldata watchlist) external onlyOwner {
    s_watchList = watchlist;
  }

  /// @notice Sets the contract config
  /// @param maxBatchSize the maximum number of upkeeps to fund in a single transaction
  /// @param minPercentage the percentage of the min balance at which to trigger top ups
  /// @param targetPercentage the percentage of the min balance to target during top ups
  /// @param maxTopUpAmount the maximum amount to top up an upkeep
  function setConfig(
    uint8 maxBatchSize,
    uint24 minPercentage,
    uint24 targetPercentage,
    uint96 maxTopUpAmount
  ) public onlyOwner {
    if (maxBatchSize == 0 || minPercentage < 100 || targetPercentage <= minPercentage || maxTopUpAmount == 0)
      revert InvalidConfig();
    s_config = Config({
      maxBatchSize: maxBatchSize,
      minPercentage: minPercentage,
      targetPercentage: targetPercentage,
      maxTopUpAmount: maxTopUpAmount
    });
    emit ConfigSet(minPercentage, targetPercentage);
  }

  // ================================================================
  // |                           GETTERS                            |
  // ================================================================

  /// @notice Gets a list of upkeeps that are underfunded.
  /// @return needsFunding list of underfunded upkeepIDs
  /// @return topUpAmounts amount to top up each upkeep
  function getUnderfundedUpkeeps() public view returns (uint256[] memory, uint256[] memory) {
    uint256 numUpkeeps = s_watchList.length;
    uint256[] memory needsFunding = new uint256[](numUpkeeps);
    uint256[] memory topUpAmounts = new uint256[](numUpkeeps);
    Config memory config = s_config;
    uint256 availableFunds = LINK_TOKEN.balanceOf(address(this));
    uint256 count;
    uint256 upkeepID;
    for (uint256 i = 0; i < numUpkeeps; i++) {
      upkeepID = s_watchList[i];
      uint96 upkeepBalance = s_registry.getBalance(upkeepID);
      uint256 minBalance = uint256(s_registry.getMinBalance(upkeepID));
      uint256 topUpThreshold = (minBalance * config.minPercentage) / 100;
      uint256 topUpAmount = (minBalance * config.targetPercentage) / 100;
      if (topUpAmount > config.maxTopUpAmount) {
        topUpAmount = config.maxTopUpAmount;
      }
      if (upkeepBalance <= topUpThreshold && availableFunds >= topUpAmount) {
        needsFunding[count] = upkeepID;
        topUpAmounts[count] = topUpAmount;
        count++;
        availableFunds -= topUpAmount;
      }
    }
    if (count < numUpkeeps) {
      assembly {
        mstore(needsFunding, count)
        mstore(topUpAmounts, count)
      }
    }
    return (needsFunding, topUpAmounts);
  }

  /// @notice Gets the list of upkeeps ids being watched.
  function getWatchList() external view returns (uint256[] memory) {
    return s_watchList;
  }

  /// @notice Gets the contract config
  function getConfig() external view returns (Config memory) {
    return s_config;
  }
}