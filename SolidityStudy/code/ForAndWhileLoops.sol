// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract ForAndWhileLoops {
    function loops() external pure {
        for (uint i = 0; i < 10; i++){
            if (i == 3) {
                continue;
            }
        }
    }
}