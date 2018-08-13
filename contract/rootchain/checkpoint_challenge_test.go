package rootchain

import (
	"testing"
	"time"

	"github.com/smartmeshfoundation/smartplasma/blockchan/block/checkpoints"
)

func TestCheckpointChallenge(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	newPlasmaTestTx(t, i, zero, uid, one, zero, user1.From, user1)
	tx2 := newPlasmaTestTx(t, i, one, uid, one, one, user2.From, user1)
	tx3 := newPlasmaTestTx(t, i, two, uid, one, two, owner.From, user2)

	chpt := checkpoints.NewBlock()

	err := chpt.AddCheckpoint(uid, three)
	if err != nil {
		t.Fatal(err)
	}

	chptHash, err := chpt.Build()
	if err != nil {
		t.Fatal(err)
	}

	chptProof := chpt.CreateProof(uid)

	ethTx4, err := i.rootOwnerSession.NewCheckpoint(chptHash)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx4) {
		t.Fatal("failed to create new checkpoint")
	}

	ethTx5, err := i.rootOwnerSession.ChallengeCheckpoint(uid, chptHash,
		chptProof, three, tx2.rawTx, tx2.proof, two)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx5) {
		t.Fatal("failed to create new checkpoint challenge")
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid, chptHash)
	if err != nil {
		t.Fatal(err)
	}

	if length.Int64() != 1 {
		t.Fatal("length is wrong")
	}

	suspiciousChpt, err := i.rootOwnerSession.CheckpointIsChallenge(uid,
		chptHash, tx2.rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if !suspiciousChpt {
		t.Fatal("checkpoint not suspicious")
	}

	respondTx, err := i.rootOwnerSession.RespondCheckpointChallenge(uid,
		chptHash, tx2.rawTx, tx3.rawTx, tx3.proof, three)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(respondTx) {
		t.Fatal("failed to respond checkpoint challenge")
	}

	length, err = i.rootOwnerSession.CheckpointChallengesLength(uid, chptHash)
	if err != nil {
		t.Fatal(err)
	}

	if length.Int64() != 0 {
		t.Fatal("length is wrong")
	}

	suspiciousChpt, err = i.rootOwnerSession.CheckpointIsChallenge(uid,
		chptHash, tx2.rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if suspiciousChpt {
		t.Fatal("checkpoint suspicious")
	}
}

func TestRespondWithHistoricalCheckpoint(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	newPlasmaTestTx(t, i, zero, uid, one, zero, user1.From, user1)
	tx2 := newPlasmaTestTx(t, i, one, uid, one, one, user2.From, user1)

	chpt1 := checkpoints.NewBlock()

	err := chpt1.AddCheckpoint(uid, two)
	if err != nil {
		t.Fatal(err)
	}

	chpt1Hash, err := chpt1.Build()
	if err != nil {
		t.Fatal(err)
	}

	chpt1Proof := chpt1.CreateProof(uid)

	ethTx1, err := i.rootOwnerSession.NewCheckpoint(chpt1Hash)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx1) {
		t.Fatal("failed to create new checkpoint")
	}

	// + 3 week
	timeMachine(t, time.Duration(504*time.Hour))

	chpt2 := checkpoints.NewBlock()

	err = chpt2.AddCheckpoint(uid, three)
	if err != nil {
		t.Fatal(err)
	}

	chpt2Hash, err := chpt2.Build()
	if err != nil {
		t.Fatal(err)
	}

	chpt2Proof := chpt1.CreateProof(uid)

	ethTx2, err := i.rootOwnerSession.NewCheckpoint(chpt2Hash)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx2) {
		t.Fatal("failed to create new checkpoint")
	}

	ethTx3, err := i.rootOwnerSession.ChallengeCheckpoint(uid, chpt2Hash,
		chpt2Proof, three, tx2.rawTx, tx2.proof, two)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx3) {
		t.Fatal("failed to create new checkpoint challenge")
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid, chpt2Hash)
	if err != nil {
		t.Fatal(err)
	}

	if length.Int64() != 1 {
		t.Fatal("length is wrong")
	}

	suspiciousChpt, err := i.rootOwnerSession.CheckpointIsChallenge(uid,
		chpt2Hash, tx2.rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if !suspiciousChpt {
		t.Fatal("checkpoint not suspicious")
	}

	respondTx, err := i.rootOwnerSession.RespondWithHistoricalCheckpoint(uid,
		chpt2Hash, chpt2Proof, chpt1Hash, chpt1Proof, tx2.rawTx, two)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(respondTx) {
		t.Fatal("failed to respond checkpoint challenge")
	}

	length, err = i.rootOwnerSession.CheckpointChallengesLength(uid, chpt1Hash)
	if err != nil {
		t.Fatal(err)
	}

	if length.Int64() != 0 {
		t.Fatal("length is wrong")
	}

	suspiciousChpt, err = i.rootOwnerSession.CheckpointIsChallenge(uid,
		chpt2Hash, tx2.rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if suspiciousChpt {
		t.Fatal("checkpoint suspicious")
	}
}
