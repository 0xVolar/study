// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract GlobalVariables {
    function globalVars() external view returns (address, uint, uint) {
        address sender = msg.sender;
        //区块时间戳
        uint timestamp = block.timestamp;
        //
        uint blockNumber = block.number;

        return (sender, timestamp, blockNumber);
    }
}