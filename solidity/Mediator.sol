pragma solidity ^0.4.23;

import "./libraries/openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract Token {
    function totalSupply() public view returns (uint256);
    function balanceOf(address who) public view returns (uint256);
    function approve(
        address _spender,
        uint256 _value
    ) public returns (bool);
    function increaseApproval(
        address _spender,
        uint _addedValue
    ) public returns (bool);
    function transferFrom(
        address from,
        address to,
        uint256 value
    ) public returns (bool);
    function transfer(address to, uint256 value) public returns (bool);
}

contract RootChain {
    function deposit(address currency, uint amount) payable public;
    function startExit(
        bytes prevTx,
        bytes prevTxProof,
        uint prevTxBlkNum,
        bytes tx, bytes txProof,
        uint txBlkNum
    ) public;
    function challengeExit(
        uint uid,
        bytes challengeTx,
        bytes proof,
        uint blkNum
    ) public;
}

contract Mediator is Ownable {

    RootChain rootChain;

    constructor(address plasma) public {
        rootChain = RootChain(plasma);
    }

    function checkToken(address addr) view public returns(bool) {
        Token token = Token(addr);

        require(token.totalSupply() > 0); // checkToken test1
        require(token.balanceOf(msg.sender) > 0); // checkToken test2
        require(token.approve(this, 0)); // checkToken test3
        require(token.increaseApproval(this, 0)); // checkToken test4
        require(token.transferFrom(msg.sender, this, 0)); // checkToken test5

        return true;
    }

    function deposit(address currency, uint amount) payable public {
        Token token = Token(currency);

        token.transferFrom(msg.sender, this, amount);
    }
}
