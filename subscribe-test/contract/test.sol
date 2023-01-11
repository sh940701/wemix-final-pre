// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.7;

contract test {
    event log(string, string);

    function set(string memory a, string memory b) public {
        emit log(a, b);
    }
}