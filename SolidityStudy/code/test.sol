// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract test {
    uint[] private a;

    function find() external returns (uint) {
        a.push(5);
        a.push(8);
        a.push(0);

        return a.length;
    }
    revi

    function test1() external pure returns (bytes32) {
        string memory s = "654894s51ad65";
        string memory x = "sadas";
        bool y = true;
        return keccak256(abi.encode(s, x, y));
    }
}