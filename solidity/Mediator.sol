pragma solidity ^0.4.23;

import "./libraries/openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract Token {
    function totalSupply() public view returns (uint256);
    function balanceOf(address who) public view returns (uint256);
    function approve(address _spender, uint256 _value) public returns (bool);
    function transferFrom(address from, address to, uint256 value) public returns (bool);
}

contract RootChain {
    function deposit(address currency, uint amount) payable public;
}

contract Mediator is Ownable {

    constructor() public {

    }

    function checkToken(address addr) view public returns(bool) {
        Token token = Token(addr);

        require(token.totalSupply() > 0);
        require(token.balanceOf(msg.sender) > 0);
        require(token.approve(this, 0));
        require(token.transferFrom(msg.sender, this, 0));
        return true;
    }
}
