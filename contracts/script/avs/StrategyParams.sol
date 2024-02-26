// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { EigenM2GoerliDeployments } from "test/avs/eigen/deploy/EigenM2GoerliDeployments.sol";

/**
 * @title StrategyParams
 * @dev Defines OmniAVS strategy params for different chains
 */
library StrategyParams {
    /// @notice standar strategy multiplier, matches OmniAVS.WEIGHTING_DIVISOR
    uint96 public constant STD_MULTIPLIER = 1e18;

    /// @notice EigenLayer's canonical, virtual beacon chain ETH strategy
    address public constant BEACON_CHAIN_ETH_STRATEGY = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

    /// @notice Goerli strategy params
    function goerli() external pure returns (IOmniAVS.StrategyParams[] memory params) {
        params = new IOmniAVS.StrategyParams[](3);

        params[0] = IOmniAVS.StrategyParams({
            strategy: IStrategy(EigenM2GoerliDeployments.stETHStrategy),
            multiplier: STD_MULTIPLIER
        });

        params[1] = IOmniAVS.StrategyParams({
            strategy: IStrategy(EigenM2GoerliDeployments.rETHStrategy),
            multiplier: STD_MULTIPLIER
        });

        params[2] =
            IOmniAVS.StrategyParams({ strategy: IStrategy(BEACON_CHAIN_ETH_STRATEGY), multiplier: STD_MULTIPLIER });
    }

    /// @notice Mainnet strategy params
    function mainnet() external pure returns (IOmniAVS.StrategyParams[] memory) {
        revert("Not implemented");
    }
}