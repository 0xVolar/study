// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Constrctor {
    address public owner;
    uint public x;

    constructor(uint _x) {
        owner = msg.sender;
        x = _x;
    }

    function returnAddress() public view returns(address) {
        return msg.sender;        
    }
}