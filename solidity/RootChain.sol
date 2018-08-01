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

    uint256 public depositCount;
    uint256 public blockNumber;
    uint256 challengePeriod;
    address operator;

    mapping(uint256 => bytes32) public childChain;
    mapping(uint256 => exit) public exits;
    mapping(bytes32 => uint256) public wallet;
    mapping(uint256 => dispute) public disputes;

    struct exit {
        // 0 - did not request to exit,
        // 1 - in challenge proceeding,
        // 2 - in anticipation of exit,
        // 3 - the exit was made.
        uint256 state;
        uint256 exitTime;
        uint256 exitTxBlkNum;
        bytes exitTx;
        uint256 txBeforeExitTxBlkNum;
        bytes txBeforeExitTx;
    }

    struct challenge {
        bool exists;
        bytes challengeTx;
        uint256 blockNumber;
    }

    struct dispute {
        uint256 len;
        mapping(uint256 => challenge) challenges;
        mapping(bytes => uint256) indexes;
    }

    constructor (address _operator) public {
        blockNumber = 0;
        challengePeriod = 2 weeks;
        depositCount = 0;
        operator = _operator;
    }

    modifier onlyOperator() {
        require(msg.sender == operator);
        _;
    }

    function deposit(
        address account,
        address currency,
        uint256 amount
    )
        public
        onlyOwner
        returns (bytes32)
    {
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
        uint256 previousTxBlockNum,
        bytes lastTx,
        bytes lastTxProof,
        uint256 lastTxBlockNum
    )
        public
    {
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
        require(challengesLength(decodedTx.uid) == 0);

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
        bytes previousTx,
        bytes previousTxProof,
        uint256 previousTxBlockNum,
        bytes lastTx,
        bytes lastTxProof,
        uint256 lastTxBlockNum
    )
        public
        onlyOwner
        returns (bytes32)
    {
        Transaction.Tx memory prevDecodedTx = previousTx.createTx();
        Transaction.Tx memory decodedTx = lastTx.createTx();

        require(previousTxBlockNum == decodedTx.prevBlock);
        require(prevDecodedTx.uid == decodedTx.uid);
        require(prevDecodedTx.amount == decodedTx.amount);
        require(prevDecodedTx.newOwner == decodedTx.signer);
        require(account == decodedTx.newOwner);

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

        require(exits[decodedTx.uid].exitTime < now);
        require(exits[decodedTx.uid].state == 2);
        require(challengesLength(decodedTx.uid) == 0);

        exits[decodedTx.uid].state = 3;

        return bytes32(decodedTx.uid);
    }

    function startDispute() public {

    }

    function respondToDispute() public {

    }

    function challengeExists(
        uint256 uid,
        bytes challengeTx
    )
        public
        view
        returns(bool)
    {
        uint256 index = disputes[uid].indexes[challengeTx];
        if (index == 0) {
            return false;
        }
        return disputes[uid].challenges[index].exists;
    }

    function challengesLength(
        uint256 uid
    )
        public
        view
        returns(uint256)
    {
        uint256 origLen = disputes[uid].len;

        if (origLen == 0) {
            return uint256(0);
        }
        return(origLen.sub(uint256(1)));
    }

    function getChallenge(
        uint256 uid,
        uint256 index
    )
        public
        view
        returns(bytes challengeTx, uint256 blockNumber)
    {
        challenge storage che = disputes[uid].challenges[index.add(uint256(1))];

        return(che.challengeTx, che.blockNumber);
    }

    function addChallenge(
        uint256 uid,
        bytes challengeTx,
        uint blockNumber
    )
        public
    {
        uint256 indexTx = disputes[uid].indexes[challengeTx];

        require(indexTx == 0);

        challenge memory cha = challenge({
            exists: true,
            challengeTx: challengeTx,
            blockNumber: blockNumber
        });

        // index 1 is magic number
        if (disputes[uid].len == 0) {
            disputes[uid].len = 1;
        }

        disputes[uid].challenges[disputes[uid].len] = cha;
        disputes[uid].indexes[challengeTx] = disputes[uid].len;
        disputes[uid].len = disputes[uid].len.add(uint256(1));
    }

    function removeChallenge(
        uint256 uid,
        bytes challengeTx
    )
        public
    {
        uint256 indexTx = disputes[uid].indexes[challengeTx];

        require(indexTx != 0);

        delete(disputes[uid].challenges[indexTx]);
        delete(disputes[uid].indexes[challengeTx]);

        uint256 lastIndex = disputes[uid].len.sub(uint256(1));

        if (indexTx != lastIndex) {
            challenge storage lastChe = disputes[uid].challenges[lastIndex];
            disputes[uid].challenges[indexTx] = lastChe;
            disputes[uid].indexes[lastChe.challengeTx] = indexTx;
            delete(disputes[uid].challenges[lastIndex]);
        }

        // index 1 is magic number
        if (lastIndex == 1) {
            disputes[uid].len = 0;
            return;
        }

        disputes[uid].len = lastIndex;
    }
}