// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract HashFunc{
    function hash(string memory text, uint num, address add) external pure returns(bytes32) {
        return keccak256(abi.encodePacked(text, num, add));
    } 
}