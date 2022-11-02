// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract ValueTypes {
    bool public b = true;
    //uint默认为uint256，如果需要其他的数的话，须指明
    //uint为无符号正数
    uint public u = 123; 
    //定义负数
    int public i = -123;
    int public minInt = type(int).min;
    int public maxInt = type(int).max;

    //地址类型
    address public addr = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;

    //bytes类型
    // bytes32 public b32 = ;
}