pragma solidity ^0.4.18;

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
 * @title Mediator contract for SmartPlasma.
 *
 * Accumulates all deposits. Makes transfer tokens.
 */
contract Mediator is Ownable {
    RootChain public rootChain;

    /** @dev If it is true, than RootChain contact already joined. */
    bool joined;

    /** @dev Dictionary with deposits.
     *
     *  key - unique identifier of a deposit (uid).
     *  value - information about a deposit.
     */
    mapping(bytes32 => entry) cash;

    struct entry {
        /// Currency address.
        address currency;
        /// Amount amount of currency.
        uint amount;
    }

    function Mediator() public {
        joined = false;
    }

    /** @dev Joins new RootChain contact to Mediator contract.
     *  @param plasma Address of RootChain contact.
     */
    function joinPlasma(address plasma) public onlyOwner {
        require(!joined);

        rootChain = RootChain(plasma);
        joined = true;
    }

    /** @dev Checks token before deposit to Smart Plasma.
     *
     *  The user can call this function on his own to check the possibility
     *  of making a deposit to the Smart Plasma contract.
     *  This helper function can serve to verify a token contract in the future.
     *  @param addr Address of token contract.
     */
    function checkToken(address addr) view public returns(bool) {
        Token token = Token(addr);

        require(token.totalSupply() > 0); /// checkToken test1
        require(token.balanceOf(msg.sender) > 0); /// checkToken test2
        require(token.approve(this, 0)); /// checkToken test3
        require(token.increaseApproval(this, 0)); /// checkToken test4
        require(token.transferFrom(msg.sender, this, 0)); /// checkToken test5

        return true;
    }

    /** @dev Adds deposits on Smart Plasma.
     *  @param currency Currency address.
     *  @param amount Amount amount of currency.
     */
    function deposit(address currency, uint amount) public {
        require(amount > 0);

        Token token = Token(currency);
        token.transferFrom(msg.sender, this, amount); /// deposit test1

        bytes32 uid = rootChain.deposit(msg.sender, currency, amount); /// deposit test2
        cash[uid] = entry({
            currency: currency,
            amount: amount
        });
    }

    /** @dev withdraws deposit from Smart Plasma.
     *  @param prevTx Penultimate deposit transaction.
     *  @param prevTxProof Proof of inclusion of a penultimate transaction in a Smart Plasma block.
     *  @param prevTxBlkNum The number of the block in which the penultimate transaction is included.
     *  @param txRaw lastTx Last deposit transaction.
     *  @param txProof Proof of inclusion of a last transaction in a Smart Plasma block.
     *  @param txBlkNum The number of the block in which the last transaction is included.
     */
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
        token.transfer(msg.sender, invoice.amount); /// withdraw test 1

        delete(cash[uid]); /// withdraw test 2
    }
}
