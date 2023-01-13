// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract DEX is Ownable {
    address _owner;
    ERC20 token;

    constructor () payable {}

    function setToken(address tokenAddress) public onlyOwner returns (bool) {
    require(tokenAddress != address(0x0));
    token = ERC20(tokenAddress);
    return true;
    }

    function sendTo(address recipient, uint256 amount) payable public {
        require(token.balanceOf(address(this)) > amount * (10 ** uint256(token.decimals())));
        token.transfer(recipient, amount * (10 ** uint256(token.decimals())));
    }

}