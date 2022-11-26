// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract AbiDecode {
    struct Mystruct {
        string mame;
        uint[2] nums;
    }

    fallback() external payable{}
    receive() external payable{}

    constructor() payable {
    }

    function encode (uint x, address addr, uint[] calldata arr, Mystruct calldata myStruct) external pure returns (bytes memory) {
        return abi.encode(x, addr, arr, myStruct);
    }

    function decode(bytes calldata data) external pure returns (uint x, address addr, uint[] memory arr, Mystruct memory myStruct) {
        (x, addr, arr, myStruct) = abi.decode(data, (uint, address, uint[], Mystruct));
    }

    function testTransfer() external {
        payable(msg.sender).transfer(1 ether);
    }

}