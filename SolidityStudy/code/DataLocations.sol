// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract DataLocations {
    struct MyStruct {
        uint foo;
        string text;
    }

    mapping(address => MyStruct) public myStructs;

    function example(uint[] calldata y, string calldata s) external returns (uint[] memory) {
        myStructs[msg.sender] = MyStruct({foo: 123, text: "bar"});

        MyStruct storage myStruct = myStructs[msg.sender];
        myStruct.text = "foo";

        uint[] memory memArr = new uint[](3);
        memArr[0] = 234;
        return memArr;
    }
}