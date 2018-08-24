pragma solidity ^0.4.23;

import "./libraries/ownership/Ownable.sol";

contract Token {
    function approve(address _spender, uint256 _value) public returns (bool);

    function balanceOf(address who) public view returns (uint256);

    function increaseApproval(address _spender, uint _addedValue) public returns (bool);

    function totalSupply() public view returns (uint256);

    function transfer(address to, uint256 value) public returns (bool);

    function transferFrom(address from, address to, uint256 value) public returns (bool);
}

contract RootChain {
    function deposit(
        address account,
        address currency,
        uint256 amount
    )
        public returns (bytes32);

    function finishExit(
        address account,
        bytes previousTx,
        bytes previousTxProof,
        uint256 previousTxBlockNum,
        bytes lastTx,
        bytes lastTxProof,
        uint256 lastTxBlockNum
    )
        public returns (bytes32);
}

/**
 * @title Mediator
 */
contract Mediator is Ownable {
    RootChain public rootChain;

    bool joined;

    mapping(bytes32 => entry) cash;

    struct entry {
        address currency;
        uint amount;
    }

    constructor() public {
        joined = false;
    }

    function joinPlasma(address plasma) public onlyOwner {
        require(!joined);

        rootChain = RootChain(plasma);
        joined = true;
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
    )
        public
    {
        bytes32 uid = rootChain.finishExit(
            msg.sender,
            prevTx,
            prevTxProof,
            prevTxBlkNum,
            txRaw,
            txProof,
            txBlkNum
        );

        entry invoice = cash[uid];

        Token token = Token(invoice.currency);
        token.transfer(msg.sender, invoice.amount); // withdraw test 1

        delete(cash[uid]); // withdraw test 2
    }
}
