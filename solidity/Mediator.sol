pragma solidity ^0.4.23;

import "./libraries/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./RootChain.sol";

contract Token {
    function totalSupply() public view returns (uint256);
    function balanceOf(address who) public view returns (uint256);
    function approve(
        address _spender,
        uint256 _value) public returns (bool);
    function increaseApproval(
        address _spender,
        uint _addedValue) public returns (bool);
    function transferFrom(
        address from,
        address to,
        uint256 value) public returns (bool);
    function transfer(address to, uint256 value) public returns (bool);
}

contract Mediator is Ownable {

    RootChain public rootChain;

    mapping(bytes32 => entry) cash;

    struct entry {
        address currency;
        uint amount;
    }

    constructor() public {
        rootChain = new RootChain();
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

    function deposit(address currency, uint amount) public {
        require(amount > 0);

        Token token = Token(currency);
        token.transferFrom(msg.sender, this, amount); // deposit test1

        bytes32 uid = rootChain.deposit(msg.sender, currency, amount); // deposit test2
        cash[uid] = entry({
            currency: currency,
            amount: amount
        });
    }

    function withdraw(
        bytes prevTx,
        bytes prevTxProof,
        uint prevTxBlkNum,
        bytes txRaw,
        bytes txProof,
        uint txBlkNum
    ) public {
        bytes32 uid = rootChain.finishExit(
            msg.sender,
            prevTx,
            prevTxProof,
            prevTxBlkNum,
            txRaw,
            txProof,
            txBlkNum
        );

        entry deposit = cash[uid];

        Token token = Token(deposit.currency);
        token.transfer(msg.sender, deposit.amount); // withdraw test 1
        delete(cash[uid]); // withdraw test 2
    }
}
