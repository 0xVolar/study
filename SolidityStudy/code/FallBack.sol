// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Fallback {
    event Log(string func, address sender, uint value, bytes data);

    fallback() external payable {
        emit Log("dallback", msg.sender, 5, msg.data);
    }

    // receive() external payable {
    //     emit Log("receive", msg.sender, msg.value, "");
    // }

    function getBalance() external returns (uint) {
        return address(this).balance;
    }
}