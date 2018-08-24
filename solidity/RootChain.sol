pragma solidity ^0.4.23;

import "./libraries/PlasmaLib.sol";
import "./libraries/datastructures/Transaction.sol";
import "./libraries/MerkleProof.sol";
import "./libraries/ownership/Ownable.sol";
import "./libraries/math/SafeMath.sol";

/** @title RootChain contract for SmartPlasma.
 *
 *  SmartPlasma is based on Plasma Cash.
 */
contract RootChain is Ownable {
    using Merkle for bytes32;
    using Transaction for bytes;
    using SafeMath for uint256;

    event Deposit(address depositor, uint256 amount, uint256 uid);
    event NewBlock(bytes32 hash);
    event NewCheckpoint(bytes32 hash);
    event StartExit(uint256 uid, uint256 previousBlock,uint256 lastBlock);
    event FinishExit(uint256 uid);
    event ChallengeExit(uint256 uid);
    event ChallengeCheckpoint(uint256 uid, bytes32 checkpoint);
    event RespondChallengeExit(uint256 uid);
    event RespondCheckpointChallenge(uint256 uid, bytes32 checkpoint);
    event RespondWithHistoricalCheckpoint(uint256 uid, bytes32 checkpoint, bytes32 historicalCheckpoint);

    /** @dev Counter of deposits. */
    uint256 public depositCount;

    /** @dev Current block number of ChildChain. */
    uint256 public blockNumber;

    /** @dev The period for challenging. */
    uint256 challengePeriod;

    /** @dev Plasma Cash operator address. */
    address operator;

    /** @dev Dictionary of child chain blocks.
     *
     *  key = block number.
     *  value = block hash.
     */
    mapping(uint256 => bytes32) public childChain;

    /** @dev Dictionary of incomplete exits from SmartPlasma.
     *
     *  key = unique identifier of a deposit (uid).
     *  value = a exit information.
     */
    mapping(uint256 => exit) public exits;

    /** @dev Dictionary of deposits.
     *
     *  key = unique identifier of a deposit (uid).
     *  value = the amount of currency corresponding to this uid.
     */
    mapping(bytes32 => uint256) public wallet;

    /** @dev Dictionary of current disputes.
     *
     *  key = unique identifier of a deposit (uid).
     *  value = a dispute information.
     */
    mapping(uint256 => dispute) disputes;

    /** @dev Dictionary of checkpoints.
     *
     *  key = checkpoint hash - checkpoint merkle root.
     *  value = unix timestamp - checkpoint create time.
     */
    mapping(bytes32 => uint256) public checkpoints;

    /** @dev Dictionary of current checkpoint disputes.
     *
     *  key = unique identifier of a deposit (uid).
     *  value = dictionary of disputes.
     *  key2 (dictionary of disputes) = checkpoint hash.
     *  value2 (dictionary of disputes) = a dispute information.
     */
    mapping(uint256 => mapping(bytes32 => dispute)) checkpointDisputes;

    /** @dev Exit information.
     *
     *  state - state of exit.
     *  exitTime - unix timestamp of start exit.
     *  exitTxBlkNum - block number of last transaction.
     *  exitTx - decoded last Smart Plasma transaction.
     *  txBeforeExitTxBlkNum - block number of penultimate transaction.
     *  txBeforeExitTx - decoded penultimate Smart Plasma transaction.
     */
    struct exit {
        /** @dev exit states:
         *
         *  0 - did not request to exit,
         *  1 - in challenge proceeding, it blocks a exit,
         *  2 - in anticipation of exit,
         *  3 - a exit was made.
         */
        uint256 state;
        uint256 exitTime;
        uint256 exitTxBlkNum;
        bytes exitTx;
        uint256 txBeforeExitTxBlkNum;
        bytes txBeforeExitTx;
    }

    /** @dev Challenge information.
     *
     *  exists - if is true then challenge is exists.
     *  challengeTx - it transaction has caused challenge.
     *  blockNumber - block number of challenge transaction.
     */
    struct challenge {
        bool exists;
        bytes challengeTx;
        uint256 blockNumber;
    }

    /** @dev Dispute information.
     *
     *  len - number of outstanding disputes.
     *  challenges - dictionary of current challenges.
     *  key(dictionary of current challenges) - index of challenge.
     *  value(dictionary of current challenges) - a challenge information.
     *  indexes - dictionary of challenge indexes.
     *  key(dictionary of challenge indexes) - decoded challenge transaction.
     *  value(dictionary of challenge indexes) - index of challenge.
     */
    struct dispute {
        uint256 len;
        mapping(uint256 => challenge) challenges;
        mapping(bytes => uint256) indexes;
    }

    /** @dev Constructor of RootChain contract.
     *  @param _operator Address of Plasma Cash operator.
     */
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

    /** @dev Creates deposit. Can only call the owner.
     *       Usually the owner is the mediator contract.
     *  @param account Depositor address.
     *  @param currency Currency address.
     *  @param amount Currency amount.
     */
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

    /** @dev .
     */
    function newBlock(bytes32 hash) public onlyOperator {
        blockNumber = blockNumber.add(uint256(1));
        childChain[blockNumber] = hash;

        emit NewBlock(hash);
    }

    /** @dev .
     */
    function newCheckpoint(bytes32 hash) public onlyOperator {
        require(checkpoints[hash] == 0);

        checkpoints[hash] = now;

        emit NewCheckpoint(hash);
    }

    /** @dev .
     */
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
        require(decodedTx.nonce == prevDecodedTx.nonce.add(uint256(1)));
        require(msg.sender == decodedTx.newOwner);
        require(wallet[bytes32(decodedTx.uid)] != 0);

        bytes32 prevTxHash = prevDecodedTx.hash;
        bytes32 prevBlockRoot = childChain[previousTxBlockNum];
        bytes32 txHash = decodedTx.hash;
        bytes32 blockRoot = childChain[lastTxBlockNum];

        require(
            prevTxHash.verifyProof(
                prevDecodedTx.uid,
                prevBlockRoot,
                previousTxProof
            )
        );
        require(
            txHash.verifyProof(
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

        emit StartExit(prevDecodedTx.uid, previousTxBlockNum, lastTxBlockNum);
    }

    /** @dev .
     */
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
            prevTxHash.verifyProof(
                prevDecodedTx.uid,
                prevBlockRoot,
                previousTxProof
            )
        );

        require(
            txHash.verifyProof(
                decodedTx.uid,
                blockRoot,
                lastTxProof
            )
        );

        require(exits[decodedTx.uid].exitTime < now);
        require(exits[decodedTx.uid].state == 2);
        require(challengesLength(decodedTx.uid) == 0);

        exits[decodedTx.uid].state = 3;

        delete(wallet[bytes32(decodedTx.uid)]);

        emit FinishExit(decodedTx.uid);

        return bytes32(decodedTx.uid);
    }

    /** @dev .
     */
    function challengeExit(
        uint256 uid,
        bytes challengeTx,
        bytes proof,
        uint256 challengeBlockNum
    )
        public
    {
        require(exits[uid].state == 2);

        Transaction.Tx memory exitDecodedTx = (exits[uid].exitTx).createTx();
        Transaction.Tx memory beforeExitDecodedTx = (exits[uid].txBeforeExitTx).createTx();
        Transaction.Tx memory challengeDecodedTx = challengeTx.createTx();

        require(exitDecodedTx.uid == challengeDecodedTx.uid);
        require(exitDecodedTx.amount == challengeDecodedTx.amount);

        bytes32 txHash = challengeDecodedTx.hash;
        bytes32 blockRoot = childChain[challengeBlockNum];

        require(txHash.verifyProof(uid, blockRoot, proof));

        // test challenge #1 & test challenge #2
        if (exitDecodedTx.newOwner == challengeDecodedTx.signer &&
        exitDecodedTx.nonce < challengeDecodedTx.nonce) {
            delete exits[uid];
            return;
        }

        // test challenge #3
        if (challengeBlockNum < exits[uid].exitTxBlkNum &&
            (beforeExitDecodedTx.newOwner == challengeDecodedTx.signer &&
            challengeDecodedTx.nonce > beforeExitDecodedTx.nonce)) {
            delete exits[uid];
            return;
        }

        // test challenge #4
        if (challengeBlockNum < exits[uid].txBeforeExitTxBlkNum ) {
            exits[uid].state = 1;
            addChallenge(uid, challengeTx, challengeBlockNum);
        }

        require(exits[uid].state == 1);

        emit ChallengeExit(uid);
    }

    /** @dev .
     */
    function challengeCheckpoint(
        uint256 uid,
        bytes32 checkpointRoot,
        bytes checkpointProof,
        uint256 wrongNonce,
        bytes lastTx,
        bytes lastTxProof,
        uint lastTxBlockNum
    )
        public
    {
        require(
            checkpoints[checkpointRoot] != 0 &&
            checkpoints[checkpointRoot].add(challengePeriod) > now
        );
        require(!checkpointIsChallenge(uid, checkpointRoot, lastTx));

        Transaction.Tx memory lastTxDecoded = lastTx.createTx();

        bytes32 txHash = lastTxDecoded.hash;
        bytes32 blockRoot = childChain[lastTxBlockNum];
        bytes32 wrongNonceHash = bytes32(wrongNonce);

        require(
            txHash.verifyProof(
                uid,
                blockRoot,
                lastTxProof
            )
        );
        require(
            wrongNonceHash.verifyProof(
                uid,
                checkpointRoot,
                checkpointProof
            )
        );

        if (wrongNonce > lastTxDecoded.nonce) {
            addCheckpointChallenge(
                uid,
                checkpointRoot,
                lastTx,
                lastTxBlockNum
            );
        }

        emit ChallengeCheckpoint(uid, checkpointRoot);
    }

    /** @dev .
     */
    function respondChallengeExit(
        uint256 uid,
        bytes challengeTx,
        bytes respondTx,
        bytes proof,
        uint blockNum
    )
        public
    {
        require(challengeExists(uid, challengeTx));
        require(exits[uid].state == 1);

        Transaction.Tx memory challengeDecodedTx = challengeTx.createTx();
        Transaction.Tx memory respondDecodedTx = respondTx.createTx();

        require(challengeDecodedTx.uid == respondDecodedTx.uid);
        require(challengeDecodedTx.amount == respondDecodedTx.amount);
        require(challengeDecodedTx.newOwner == respondDecodedTx.signer);
        require(challengeDecodedTx.nonce.add(uint256(1)) == respondDecodedTx.nonce);
        require(blockNum < exits[uid].txBeforeExitTxBlkNum);

        bytes32 txHash = respondDecodedTx.hash;
        bytes32 blockRoot = childChain[blockNum];

        require(txHash.verifyProof(uid, blockRoot, proof));

        removeChallenge(uid, challengeTx);

        if (challengesLength(uid) == 0) {
            exits[uid].state = 2;
        }

        emit RespondChallengeExit(uid);
    }

    /** @dev .
     */
    function respondCheckpointChallenge(
        uint256 uid,
        bytes32 checkpointRoot,
        bytes challengeTx,
        bytes respondTx,
        bytes proof,
        uint blockNum
    )
        public
    {
        require(checkpointIsChallenge(uid, checkpointRoot, challengeTx));

        Transaction.Tx memory challengeDecodedTx = challengeTx.createTx();
        Transaction.Tx memory respondDecodedTx = respondTx.createTx();

        require(challengeDecodedTx.uid == respondDecodedTx.uid);
        require(challengeDecodedTx.amount == respondDecodedTx.amount);
        require(challengeDecodedTx.newOwner == respondDecodedTx.signer);
        require(challengeDecodedTx.nonce.add(uint256(1)) == respondDecodedTx.nonce);

        bytes32 txHash = respondDecodedTx.hash;
        bytes32 blockRoot = childChain[blockNum];

        require(txHash.verifyProof(uid, blockRoot, proof));

        removeCheckpointChallenge(uid, checkpointRoot, challengeTx);

        emit RespondCheckpointChallenge(uid, checkpointRoot);
    }

    /** @dev .
     */
    function respondWithHistoricalCheckpoint(
        uint256 uid,
        bytes32 checkpointRoot,
        bytes checkpointProof,
        bytes32 historicalCheckpointRoot,
        bytes historicalCheckpointProof,
        bytes challengeTx,
        uint256 moreNonce
    )
        public
    {
        require(checkpointIsChallenge(uid, checkpointRoot, challengeTx));

        Transaction.Tx memory challengeDecodedTx = challengeTx.createTx();

        bytes32 moreNonceBytes = bytes32(moreNonce);

        require(moreNonce > challengeDecodedTx.nonce);
        require(
            moreNonceBytes.verifyProof(
                uid,
                historicalCheckpointRoot,
                historicalCheckpointProof
            )
        );
        require(checkpoints[historicalCheckpointRoot].add(challengePeriod) < now);
        require(checkpoints[historicalCheckpointRoot] < checkpoints[checkpointRoot]);

        removeCheckpointChallenge(uid, checkpointRoot, challengeTx);

        emit RespondWithHistoricalCheckpoint(uid, checkpointRoot, historicalCheckpointRoot);
    }

    /** @dev .
     */
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

    /** @dev .
     */
    function checkpointIsChallenge(
        uint256 uid,
        bytes32 checkpoint,
        bytes challengeTx
    )
        public
        view
        returns(bool)
    {
        uint256 index = checkpointDisputes[uid][checkpoint].indexes[challengeTx];
        if (index == 0) {
            return false;
        }
        return checkpointDisputes[uid][checkpoint].challenges[index].exists;
    }

    /** @dev .
     */
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

    /** @dev .
     */
    function checkpointChallengesLength(
        uint256 uid,
        bytes32 checkpoint
    )
        public
        view
        returns(uint256)
    {
        uint256 origLen = checkpointDisputes[uid][checkpoint].len;

        if (origLen == 0) {
            return uint256(0);
        }
        return(origLen.sub(uint256(1)));
    }

    /** @dev .
     */
    function getChallenge(
        uint256 uid,
        uint256 index
    )
        public
        view
        returns(bytes challengeTx, uint256 challengeBlock)
    {
        challenge storage che = disputes[uid].challenges[index.add(uint256(1))];

        return(che.challengeTx, che.blockNumber);
    }

    /** @dev .
     */
    function getCheckpointChallenge(
        uint256 uid,
        bytes32 checkpoint,
        uint256 index
    )
        public
        view
        returns(bytes challengeTx, uint256 challengeBlock)
    {
        challenge storage che = checkpointDisputes[uid][checkpoint].challenges[index.add(uint256(1))];

        return(che.challengeTx, che.blockNumber);
    }

    /** @dev .
     */
    function addCheckpointChallenge(
        uint256 uid,
        bytes32 checkpoint,
        bytes challengeTx,
        uint challengeBlockNumber
    )
        private
    {
        uint256 indexTx = checkpointDisputes[uid][checkpoint].indexes[challengeTx];

        require(indexTx == 0);

        challenge memory cha = challenge({
            exists: true,
            challengeTx: challengeTx,
            blockNumber: challengeBlockNumber
            });

        // index 1 is magic number
        if (checkpointDisputes[uid][checkpoint].len == 0) {
            checkpointDisputes[uid][checkpoint].len = 1;
        }

        uint256 currentLen = checkpointDisputes[uid][checkpoint].len;

        checkpointDisputes[uid][checkpoint].challenges[currentLen] = cha;
        checkpointDisputes[uid][checkpoint].indexes[challengeTx] = currentLen;
        checkpointDisputes[uid][checkpoint].len = currentLen.add(uint256(1));
    }

    /** @dev .
     */
    function addChallenge(
        uint256 uid,
        bytes challengeTx,
        uint challengeBlockNumber
    )
        private
    {
        uint256 indexTx = disputes[uid].indexes[challengeTx];

        require(indexTx == 0);

        challenge memory cha = challenge({
            exists: true,
            challengeTx: challengeTx,
            blockNumber: challengeBlockNumber
        });

        // index 1 is magic number
        if (disputes[uid].len == 0) {
            disputes[uid].len = 1;
        }

        disputes[uid].challenges[disputes[uid].len] = cha;
        disputes[uid].indexes[challengeTx] = disputes[uid].len;
        disputes[uid].len = disputes[uid].len.add(uint256(1));
    }

    /** @dev .
     */
    function removeCheckpointChallenge(
        uint256 uid,
        bytes32 checkpoint,
        bytes challengeTx
    )
        private
    {
        uint256 indexTx = checkpointDisputes[uid][checkpoint].indexes[challengeTx];

        require(indexTx != 0);

        delete(checkpointDisputes[uid][checkpoint].challenges[indexTx]);
        delete(checkpointDisputes[uid][checkpoint].indexes[challengeTx]);

        uint256 lastIndex = checkpointDisputes[uid][checkpoint].len.sub(uint256(1));

        if (indexTx != lastIndex) {
            challenge storage lastChe = checkpointDisputes[uid][checkpoint].challenges[lastIndex];
            checkpointDisputes[uid][checkpoint].challenges[indexTx] = lastChe;
            checkpointDisputes[uid][checkpoint].indexes[lastChe.challengeTx] = indexTx;
            delete(checkpointDisputes[uid][checkpoint].challenges[lastIndex]);
        }

        // index 1 is magic number
        if (lastIndex == 1) {
            checkpointDisputes[uid][checkpoint].len = 0;
            return;
        }

        checkpointDisputes[uid][checkpoint].len = lastIndex;
    }

    /** @dev .
     */
    function removeChallenge(
        uint256 uid,
        bytes challengeTx
    )
        private
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