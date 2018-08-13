package rootchain

import (
	"testing"

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
