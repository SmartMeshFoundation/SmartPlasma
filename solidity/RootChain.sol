pragma solidity ^0.4.23;

import "./libraries/datastructures/Challenge.sol";
import "./libraries/merkle.sol";
import "./libraries/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./libraries/datastructures/Transaction.sol";

contract RootChain is Ownable {
    using Challenge for Challenge.challenge[];
    using Merkle for bytes32;
    using Transaction for bytes;

    /*
     * Events
     */
    event Deposit(address depositor, uint256 amount, uint256 uid);

    /*
     * Storage
     */
    address public authority;
    uint public currentBlkNum;
    uint public depositCount;

    mapping(uint => bytes32) public childChain;
    mapping(bytes32 => uint) public wallet;
    mapping(uint => exit) public exits;
    mapping(uint => Challenge.challenge[]) public challenges;

    struct exit {
        bool hasValue;
        bool transferred;
        uint exitTime;
        uint exitTxBlkNum;
        bytes exitTx;
        uint txBeforeExitTxBlkNum;
        bytes txBeforeExitTx;
    }

    constructor () public {
        currentBlkNum = 0;
        depositCount = 0;
    }

    function deposit(
        address account,
        address currency,
        uint amount) public onlyOwner {

        bytes32 uid = keccak256(
            abi.encodePacked(currency, msg.sender, depositCount));
        wallet[uid] = amount;
        depositCount += 1;

        emit Deposit(account, amount, uint256(uid));
    }

    function startExit(
        bytes prevTx,
        bytes prevTxProof,
        uint prevTxBlkNum,
        bytes txRaw,
        bytes txProof,
        uint txBlkNum) public {

        Transaction.Tx memory prevTxObj = prevTx.createTx();
        Transaction.Tx memory txObj = txRaw.createTx();

        require(prevTxBlkNum == txObj.prevBlock);
        require(prevTxObj.uid == txObj.uid);
        require(prevTxObj.amount == txObj.amount);
        require(prevTxObj.newOwner == txObj.signer);
        require(msg.sender == txObj.newOwner);

        bytes32 prevMerkleHash = keccak256(prevTx);
        bytes32 prevRoot = childChain[prevTxBlkNum];
        bytes32 merkleHash = keccak256(txRaw);
        bytes32 root = childChain[txBlkNum];

        require(
            prevMerkleHash.checkMembership(
                prevTxObj.uid,
                prevRoot,
                prevTxProof
            )
        );
        require(
            merkleHash.checkMembership(
                txObj.uid,
                root,
                txProof
            )
        );

        // Record the exit tx.
        require(!exits[txObj.uid].hasValue);
        require(!exits[txObj.uid].transferred);

        exits[txObj.uid] = exit({
            hasValue: true,
            transferred: false,
            exitTime: now + 2 weeks,
            exitTxBlkNum: txBlkNum,
            exitTx: txRaw,
            txBeforeExitTxBlkNum: prevTxBlkNum,
            txBeforeExitTx: prevTx
            });
    }

    function finishExit(
        address account,
        bytes prevTx,
        bytes prevTxProof,
        uint prevTxBlkNum,
        bytes txRaw,
        bytes txProof,
        uint txBlkNum) public onlyOwner returns (bool) {

//        Transaction.Tx memory prevTxObj = prevTx.createTx();
//        Transaction.Tx memory txObj = txRaw.createTx();
//
//        require(prevTxBlkNum == txObj.prevBlock);
//        require(prevTxObj.uid == txObj.uid);
//        require(prevTxObj.amount == txObj.amount);
//        require(prevTxObj.newOwner == txObj.signer);
//        require(msg.sender == txObj.newOwner);
//
//        bytes32 prevMerkleHash = keccak256(prevTx);
//        bytes32 prevRoot = childChain[prevTxBlkNum];
//        bytes32 merkleHash = keccak256(txRaw);
//        bytes32 root = childChain[txBlkNum];
//
//        require(
//            prevMerkleHash.checkMembership(
//                prevTxObj.uid,
//                prevRoot,
//                prevTxProof
//            )
//        );
//        require(
//            merkleHash.checkMembership(
//                txObj.uid,
//                root,
//                txProof
//            )
//        );
//
//        require(exits[txObj.uid].hasValue);
//        require(exits[txObj.uid].exitTime > now);
//        require(!exits[txObj.uid].transferred);
//
//        exits[txObj.uid].transferred = true;
        return true;
    }

    function challengeExit(
        uint uid,
        bytes challengeTx,
        bytes proof,
        uint blkNum) public {
    }

    function respondChallengeExit(
        uint uid,
        bytes challengeTx,
        bytes respondTx,
        bytes proof,
        uint blkNum) public {
    }
}