package mediator

import (
	"testing"

	"github.com/smartmeshfoundation/smartplasma/blockchan/block/checkpoints"
)

func TestCheckpointChallenge(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	tx1 := testTx(t, zero, uid, one, zero, user1.From, user1)

	plasmaBlock1 := testBlock(t, tx1)

	ethTX, err := i.rootOwnerSession.NewBlock(plasmaBlock1.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX) {
		t.Fatal("failed to create new block")
	}

	tx2 := testTx(t, one, uid, one, one, user2.From, user1)

	plasmaBlock2 := testBlock(t, tx2)

	ethTx2, err := i.rootOwnerSession.NewBlock(plasmaBlock2.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx2) {
		t.Fatal("failed to create new block")
	}

	proof2 := plasmaBlock2.CreateProof(uid)
	rawTx2 := txToBytes(t, tx2)

	tx3 := testTx(t, two, uid, one, two, owner.From, user2)

	plasmaBlock3 := testBlock(t, tx3)

	ethTx3, err := i.rootOwnerSession.NewBlock(plasmaBlock3.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx3) {
		t.Fatal("failed to create new block")
	}

	proof3 := plasmaBlock3.CreateProof(uid)
	rawTx3 := txToBytes(t, tx3)

	chpt := checkpoints.NewBlock()

	err = chpt.AddCheckpoint(uid, three)
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
		chptProof, three, rawTx2, proof2, two)
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
		chptHash, rawTx2)
	if err != nil {
		t.Fatal(err)
	}

	if !suspiciousChpt {
		t.Fatal("checkpoint suspicious")
	}

	respondTx, err := i.rootOwnerSession.RespondCheckpointChallenge(uid,
		chptHash, rawTx2, rawTx3, proof3, three)
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
		chptHash, rawTx2)
	if err != nil {
		t.Fatal(err)
	}

	if suspiciousChpt {
		t.Fatal("checkpoint suspicious")
	}
}
