pragma solidity ^0.4.23;

import "./libraries/datastructures/PlasmaLib.sol";
import "./libraries/datastructures/Transaction.sol";
import "./libraries/merkle.sol";
import "./libraries/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./libraries/openzeppelin-solidity/contracts/math/SafeMath.sol";

contract RootChain is Ownable {
    using Merkle for bytes32;
    using Transaction for bytes;
    using SafeMath for uint256;

    event Deposit(address depositor, uint256 amount, uint256 uid);

    uint public depositCount;
    uint public blockNumber;
    uint challengePeriod;

    address operator;

    mapping(uint => bytes32) public childChain;
    mapping(uint => exit) public exits;
    mapping(bytes32 => uint) public wallet;

    struct exit {
        // 0 - did not request to exit,
        // 1 - in challenge proceeding,
        // 2 - in anticipation of exit,
        // 3 - the exit was made.
        uint state;
        uint exitTime;
        uint exitTxBlkNum;
        bytes exitTx;
        uint txBeforeExitTxBlkNum;
        bytes txBeforeExitTx;
    }

    constructor (address _operator) public {
        blockNumber = 0;
        challengePeriod = 2 weeks;
        depositCount = 0;
        operator =_operator;
    }

    modifier onlyOperator() {
        require(msg.sender == operator);
        _;
    }

    function challengeExit(
        uint uid,
        bytes challengeTx,
        bytes proof,
        uint blkNum) public {
    }

    // TODO: not payable
    function deposit(
        address account,
        address currency,
        uint amount) public onlyOwner returns (bytes32){
        bytes32 uid = PlasmaLib.generateUID(
            account,
            currency,
            depositCount
        );
        wallet[uid] = amount;
        depositCount = depositCount.add(uint256(1));

        emit Deposit(account, amount, uint256(uid));

        return uid;
    }

    function newBlock(bytes32 hash) public onlyOperator {
        blockNumber = blockNumber.add(uint256(1));
        childChain[blockNumber] = hash;
    }

    function startExit(
        bytes previousTx,
        bytes previousTxProof,
        uint previousTxBlockNum,
        bytes lastTx,
        bytes lastTxProof,
        uint lastTxBlockNum
    ) public {
        Transaction.Tx memory prevDecodedTx = previousTx.createTx();
        Transaction.Tx memory decodedTx = lastTx.createTx();

        require(previousTxBlockNum == decodedTx.prevBlock);
        require(prevDecodedTx.uid == decodedTx.uid);
        require(prevDecodedTx.amount == decodedTx.amount);
        require(prevDecodedTx.newOwner == decodedTx.signer);
        require(msg.sender == decodedTx.newOwner);

        bytes32 prevTxHash = prevDecodedTx.hash;
        bytes32 prevBlockRoot = childChain[previousTxBlockNum];
        bytes32 txHash = decodedTx.hash;
        bytes32 blockRoot = childChain[lastTxBlockNum];

        require(
            prevTxHash.checkMembership(
                prevDecodedTx.uid,
                prevBlockRoot,
                previousTxProof
            )
        );
        require(
            txHash.checkMembership(
                decodedTx.uid,
                blockRoot,
                lastTxProof
            )
        );

        // Record the exit tx.
        require(exits[decodedTx.uid].state == 0);

        exits[decodedTx.uid] = exit({
            state: 2,
            exitTime: now.add(challengePeriod),
            exitTxBlkNum: lastTxBlockNum,
            exitTx: lastTx,
            txBeforeExitTxBlkNum: previousTxBlockNum,
            txBeforeExitTx: previousTx
        });
    }

    function finishExit(
        address account,
        bytes prevTx,
        bytes prevTxProof,
        uint prevTxBlkNum,
        bytes txRaw,
        bytes txProof,
        uint txBlkNum
    ) public onlyOwner returns (bytes32) {
        Transaction.Tx memory prevTxObj = prevTx.createTx();
        Transaction.Tx memory txObj = txRaw.createTx();

        require(prevTxBlkNum == txObj.prevBlock);
        require(prevTxObj.uid == txObj.uid);
        require(prevTxObj.amount == txObj.amount);
        require(prevTxObj.newOwner == txObj.signer);
        require(account == txObj.newOwner);

        bytes32 prevMerkleHash = prevTxObj.hash;
        bytes32 prevRoot = childChain[prevTxBlkNum];
        bytes32 merkleHash = txObj.hash;
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

        require(exits[txObj.uid].exitTime < now);
        require(exits[txObj.uid].state == 2);

        exits[txObj.uid].state = 3;

        return bytes32(txObj.uid);
    }
}