// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "Draco.sol";
import "Credit.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract DEX is Ownable {
    address _owner;
    Draco token;
    Credit credit;

    constructor () payable {}

    // decimal 숫자를 추가해주는 함수
    function decimal(uint256 num) private view returns (uint256) {
        return num * (10 ** uint256(token.decimals()));
    }

    function setDraco(address tokenAddress) public onlyOwner {
        require(tokenAddress != address(0x0));
        token = Draco(tokenAddress);
    }

    function setCredit(address creditAddress) public onlyOwner {
        require(creditAddress != address(0x0));
        credit = Credit(creditAddress);
    }


    // token -> credit으로 변경하는 함수 1(크레딧) : 5(토큰) 의 비율
    function TokenToCredit(address user, uint256 tokenAmount) public onlyOwner {
        // 일단 user의 tokenAmount가 요청한 만큼 있는지 확인한다.
        require(token.balanceOf(user) >= decimal(tokenAmount));
        // 그리고 DEX가 tokenAmount/5 만큼의 credit을 가지고 있는지 확인한다.
        require(credit.balanceOf(address(this)) >= decimal(tokenAmount) / 5);

        // 확인이 됐으면, user의 token을 DEX에게로 옮긴다.
        token.transfer(user, address(this), decimal(tokenAmount));
        // 이후 DEX의 credit을 user에게 옮긴다.
        credit.transfer(address(this), user, decimal(tokenAmount) / 5);
    }

    // credit -> token으로 변경하는 함수
    function CreditToToken(address user, uint256 creditAmount) public onlyOwner {
        // 일단 user의 creditAmount가 요청한 만큼 있는지 확인한다.
        require(credit.balanceOf(user) >= decimal(creditAmount));
        // 그리고 DEX가 creditAmount * 5 만큼의 token을 가지고 있는지 확인한다.
        require(token.balanceOf(address(this)) >= decimal(creditAmount) * 5);

        // 확인이 됐으면, user의 credit을 DEX에게로 옮긴다.
        credit.transfer(user, address(this), decimal(creditAmount));
        // 이후 DEX의 token을 user에게 옮긴다.
        token.transfer(address(this), user, decimal(creditAmount) * 5);
    }


    function sendTo(address recipient, uint256 amount) payable public {
        require(token.balanceOf(address(this)) >= amount * (10 ** uint256(token.decimals())));
        token.transfer(recipient, amount * (10 ** uint256(token.decimals())));
    }

}