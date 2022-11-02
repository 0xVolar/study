// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Mapping {
    mapping(address => uint) public balances;
    mapping(address => mapping(address => bool)) public isFirend;

    function examples() external {
        balances[msg.sender] = 123;
        uint bal = balances[msg.sender];

        delete balances[msg.sender];

        isFirend[msg.sender][address(this)] = true;
    }
}