// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Array {
    uint[] public nums = [1, 2, 3];
    uint[3] public numsFixed = [4, 5, 6];

    function examples() external {
        nums.push(4);
        uint x = nums[1];
        nums[2] = 777;
        delete nums[1];

    }

    function returnArray() external view returns (uint[] memory) { 
        return nums;
    }
}