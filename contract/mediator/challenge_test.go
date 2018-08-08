package mediator

import (
	"testing"

	"github.com/smartmeshfoundation/smartplasma/blockchan/block"
)

// test challenge #1
// normal flow
func TestChallengeNewOwner(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	txGood := testTx(t, zero, uid, one, zero, user1.From, user1)

	plasmaBlock1 := testBlock(t, txGood)

	ethTX, err := i.rootOwnerSession.NewBlock(plasmaBlock1.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX) {
		t.Fatal("failed to create new block")
	}

	proof1 := plasmaBlock1.CreateProof(uid)

	validateTX1 := block.CheckMembership(uid, txGood.Hash(),
		plasmaBlock1.Hash(), proof1)
	if !validateTX1 {
		t.Fatal("txGood invalid")
	}

	txBad := testTx(t, one, uid, one, one, user1.From, user1)

	plasmaBlock2 := testBlock(t, txBad)

	ethTX2, err := i.rootOwnerSession.NewBlock(plasmaBlock2.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX2) {
		t.Fatal("failed to create new block")
	}

	proof2 := plasmaBlock2.CreateProof(uid)

	validateTX2 := block.CheckMembership(uid, txBad.Hash(),
		plasmaBlock2.Hash(), proof2)
	if !validateTX2 {
		t.Fatal("txBad invalid")
	}

	rawGoodTx := txToBytes(t, txGood)
	rawBadTx := txToBytes(t, txBad)

	ethTX3, err := i.rootUser1Session.StartExit(rawGoodTx, proof1, one,
		rawBadTx, proof2, two)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTX3) {
		t.Fatal("failed to start exit")
	}

	exits1, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits1.State.Int64() != 2 {
		t.Fatal("the exit operation did not start")
	}

	_, err = i.rootUser2Session.ChallengeExit(uid, rawBadTx,
		proof2, two)
	if err == nil {
		t.Fatal("challenge exit was started")
	}

	exits2, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits2.State.Int64() != 2 {
		t.Fatal("the exit challenge operation was start")
	}
}

// test challenge #2
// nonce is wrong
func TestChallengeNonce(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	// tx1
	tx1 := testTx(t, zero, uid, one, zero, user2.From, user1)

	plasmaBlock1 := testBlock(t, tx1)

	_, err := i.rootOwnerSession.NewBlock(plasmaBlock1.Hash())
	if err != nil {
		t.Fatal(err)
	}

	proofTx1 := plasmaBlock1.CreateProof(uid)
	rawTx1 := txToBytes(t, tx1)

	// tx2
	tx2 := testTx(t, one, uid, one, one, user1.From, user2)

	plasmaBlock2 := testBlock(t, tx2)

	_, err = i.rootOwnerSession.NewBlock(plasmaBlock2.Hash())
	if err != nil {
		t.Fatal(err)
	}

	proofTx2 := plasmaBlock2.CreateProof(uid)
	rawTx2 := txToBytes(t, tx2)

	// tx3
	tx3 := testTx(t, two, uid, one, two, user2.From, user1)

	plasmaBlock3 := testBlock(t, tx3)

	_, err = i.rootOwnerSession.NewBlock(plasmaBlock3.Hash())
	if err != nil {
		t.Fatal(err)
	}

	proofTx3 := plasmaBlock3.CreateProof(uid)
	rawTx3 := txToBytes(t, tx3)

	_, err = i.rootUser1Session.StartExit(rawTx1, proofTx1, one,
		rawTx2, proofTx2, two)
	if err != nil {
		t.Fatal(err)
	}

	exits1, err := i.rootUser2Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits1.State.Int64() != 2 {
		t.Fatal("The exit operation did not start")
	}

	// challenge
	ethTx, err := i.rootUser2Session.ChallengeExit(uid, rawTx3,
		proofTx3, three)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx) {
		t.Fatal("failed to start challenge")
	}

	exits2, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits2.State.Int64() != 0 {
		t.Fatal("The exit challenge operation did not start")
	}
}

// test challenge #3
func TestChallengeDoubleSpending(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	txGood := testTx(t, zero, uid, one, zero, user1.From, user1)

	plasmaBlock1 := testBlock(t, txGood)

	ethTX, err := i.rootOwnerSession.NewBlock(plasmaBlock1.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX) {
		t.Fatal("failed to create new block")
	}

	proof1 := plasmaBlock1.CreateProof(uid)

	validateTX1 := block.CheckMembership(uid, txGood.Hash(),
		plasmaBlock1.Hash(), proof1)
	if !validateTX1 {
		t.Fatal("txGood invalid")
	}

	txGood2 := testTx(t, one, uid, one, one, user2.From, user1)

	plasmaBlock2 := testBlock(t, txGood2)

	ethTX2, err := i.rootOwnerSession.NewBlock(plasmaBlock2.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX2) {
		t.Fatal("failed to create new block")
	}

	proof2 := plasmaBlock2.CreateProof(uid)

	validateTX2 := block.CheckMembership(uid, txGood2.Hash(),
		plasmaBlock2.Hash(), proof2)
	if !validateTX2 {
		t.Fatal("txGood2 invalid")
	}

	txBad := testTx(t, one, uid, one, one, owner.From, user1)

	plasmaBlock3 := testBlock(t, txBad)

	ethTX3, err := i.rootOwnerSession.NewBlock(plasmaBlock3.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX3) {
		t.Fatal("failed to create new block")
	}

	proof3 := plasmaBlock3.CreateProof(uid)

	validateTX3 := block.CheckMembership(uid, txBad.Hash(),
		plasmaBlock3.Hash(), proof3)
	if !validateTX3 {
		t.Fatal("txBad invalid")
	}

	rawGood1Tx := txToBytes(t, txGood)
	rawGood2Tx := txToBytes(t, txGood2)
	rawBadTx := txToBytes(t, txBad)

	ethTX4, err := i.rootOwnerSession.StartExit(rawGood1Tx, proof1, one,
		rawBadTx, proof3, three)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTX4) {
		t.Fatal("failed to start exit")
	}

	exits1, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits1.State.Int64() != 2 {
		t.Fatal("The exit operation did not start")
	}

	ethTX5, err := i.rootUser2Session.ChallengeExit(uid, rawGood2Tx,
		proof2, two)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTX5) {
		t.Fatal("failed to challenge start exit")
	}

	exits2, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits2.State.Int64() != 0 {
		t.Fatal("The exit challenge operation did not start")
	}
}

// test challenge #4
// challenge the story
func TestEarlyChallengeDoubleSpending(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	// tx1
	tx1Good := testTx(t, zero, uid, one, zero, user2.From, user1)

	plasmaBlock1 := testBlock(t, tx1Good)

	_, err := i.rootOwnerSession.NewBlock(plasmaBlock1.Hash())
	if err != nil {
		t.Fatal(err)
	}

	// tx2
	tx2Good := testTx(t, one, uid, one, one, owner.From, user2)

	plasmaBlock2 := testBlock(t, tx2Good)

	_, err = i.rootOwnerSession.NewBlock(plasmaBlock2.Hash())
	if err != nil {
		t.Fatal(err)
	}

	proofTx2Good := plasmaBlock2.CreateProof(uid)
	rawTx2Good := txToBytes(t, tx2Good)

	// tx3 - bad transaction
	tx3Bad := testTx(t, two, uid, one, one, user1.From, user2)

	plasmaBlock3 := testBlock(t, tx3Bad)

	_, err = i.rootOwnerSession.NewBlock(plasmaBlock3.Hash())
	if err != nil {
		t.Fatal(err)
	}

	// tx4
	tx4 := testTx(t, three, uid, one, two, owner.From, user1)

	plasmaBlock4 := testBlock(t, tx4)

	_, err = i.rootOwnerSession.NewBlock(plasmaBlock4.Hash())
	if err != nil {
		t.Fatal(err)
	}

	proofTx4 := plasmaBlock4.CreateProof(uid)
	rawTx4 := txToBytes(t, tx4)

	// tx5
	tx5 := testTx(t, four, uid, one, three, user2.From, owner)

	plasmaBlock5 := testBlock(t, tx5)

	_, err = i.rootOwnerSession.NewBlock(plasmaBlock5.Hash())
	if err != nil {
		t.Fatal(err)
	}

	proofTx5 := plasmaBlock5.CreateProof(uid)
	rawTx5 := txToBytes(t, tx5)

	_, err = i.rootUser2Session.StartExit(rawTx4, proofTx4, four,
		rawTx5, proofTx5, five)
	if err != nil {
		t.Fatal(err)
	}

	exits1, err := i.rootUser2Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits1.State.Int64() != 2 {
		t.Fatal("The exit operation did not start")
	}

	// challenge
	ethTx, err := i.rootUser2Session.ChallengeExit(uid, rawTx2Good,
		proofTx2Good, two)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx) {
		t.Fatal("failed to start challenge")
	}

	exits2, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits2.State.Int64() != 1 {
		t.Fatal("The exit challenge operation did not start")
	}
}

// test respond to a challenge #1
func TestRespondToChallenge(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	// tx1
	tx1 := testTx(t, zero, uid, one, zero, user2.From, user1)

	plasmaBlock1 := testBlock(t, tx1)

	ethTx1, err := i.rootOwnerSession.NewBlock(plasmaBlock1.Hash())
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTx1) {
		t.Fatal("failed to mine tx")
	}

	// tx2
	tx2 := testTx(t, one, uid, one, one, owner.From, user2)

	plasmaBlock2 := testBlock(t, tx2)

	ethTx2, err := i.rootOwnerSession.NewBlock(plasmaBlock2.Hash())
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTx2) {
		t.Fatal("failed to mine tx")
	}

	proofTx2 := plasmaBlock2.CreateProof(uid)
	rawTx2 := txToBytes(t, tx2)

	// tx3
	tx3 := testTx(t, two, uid, one, two, user1.From, owner)

	plasmaBlock3 := testBlock(t, tx3)

	ethTx3, err := i.rootOwnerSession.NewBlock(plasmaBlock3.Hash())
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTx3) {
		t.Fatal("failed to mine tx")
	}

	proofTx3 := plasmaBlock3.CreateProof(uid)
	rawTx3 := txToBytes(t, tx3)

	// tx4
	tx4 := testTx(t, three, uid, one, three, user2.From, user1)

	plasmaBlock4 := testBlock(t, tx4)

	ethTx4, err := i.rootOwnerSession.NewBlock(plasmaBlock4.Hash())
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTx4) {
		t.Fatal("failed to mine tx")
	}

	proofTx4 := plasmaBlock4.CreateProof(uid)
	rawTx4 := txToBytes(t, tx4)

	// tx5
	tx5 := testTx(t, four, uid, one, four, user1.From, user2)

	plasmaBlock5 := testBlock(t, tx5)

	ethTx5, err := i.rootOwnerSession.NewBlock(plasmaBlock5.Hash())
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTx5) {
		t.Fatal("failed to mine tx")
	}

	proofTx5 := plasmaBlock5.CreateProof(uid)
	rawTx5 := txToBytes(t, tx5)

	ethTx6, err := i.rootUser1Session.StartExit(rawTx4, proofTx4, four,
		rawTx5, proofTx5, five)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTx6) {
		t.Fatal("failed to start exit")
	}

	exits1, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits1.State.Int64() != 2 {
		t.Fatal("The exit operation did not start")
	}

	ethTx7, err := i.rootOwnerSession.ChallengeExit(uid, rawTx2,
		proofTx2, two)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(ethTx7) {
		t.Fatal("failed to challenge start exit")
	}

	exits2, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits2.State.Int64() != 1 {
		t.Fatal("The exit challenge operation did not start")
	}

	cha, err := i.rootUser1Session.GetChallenge(uid, zero)
	if err != nil {
		t.Fatal(err)
	}

	ethTx8, err := i.rootUser1Session.RespondChallengeExit(uid,
		cha.ChallengeTx, rawTx3, proofTx3, three)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(ethTx8) {
		t.Fatal("failed to respond challenge")
	}

	exits3, err := i.rootUser1Session.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exits3.State.Int64() != 2 {
		t.Fatal("the exit operation did not continue")
	}
}
