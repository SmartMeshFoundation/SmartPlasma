package rootchain

import (
	"testing"
)

// test challenge #1
// normal flow
func TestChallengeNewOwner(t *testing.T) {
	i := newInstance(t)

	uid := testDeposit(t, i)

	txGood := newPlasmaTestTx(t, i, zero, uid, one, zero, user1.From, user1)
	txBad := newPlasmaTestTx(t, i, one, uid, one, one, user1.From, user1)

	ethTX3, err := i.rootUser1Session.StartExit(txGood.rawTx, txGood.proof, one,
		txBad.rawTx, txBad.proof, two)
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

	_, err = i.rootUser2Session.ChallengeExit(uid, txBad.rawTx,
		txBad.proof, two)
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

	tx1 := newPlasmaTestTx(t, i, zero, uid, one, zero, user2.From, user1)
	tx2 := newPlasmaTestTx(t, i, one, uid, one, one, user1.From, user2)
	tx3 := newPlasmaTestTx(t, i, two, uid, one, two, user2.From, user1)

	_, err := i.rootUser1Session.StartExit(tx1.rawTx, tx1.proof, one,
		tx2.rawTx, tx2.proof, two)
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
	ethTx, err := i.rootUser2Session.ChallengeExit(uid, tx3.rawTx,
		tx3.proof, three)
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

	txGood := newPlasmaTestTx(t, i, zero, uid, one, zero, user1.From, user1)
	txGood2 := newPlasmaTestTx(t, i, one, uid, one, one, user2.From, user1)
	txBad := newPlasmaTestTx(t, i, one, uid, one, one, owner.From, user1)

	ethTX4, err := i.rootOwnerSession.StartExit(txGood.rawTx, txGood.proof,
		one, txBad.rawTx, txBad.proof, three)
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

	ethTX5, err := i.rootUser2Session.ChallengeExit(uid, txGood2.rawTx,
		txGood2.proof, two)
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

	newPlasmaTestTx(t, i, zero, uid, one, zero, user2.From, user1)
	tx2Good := newPlasmaTestTx(t, i, one, uid, one, one, owner.From, user2)
	newPlasmaTestTx(t, i, one, uid, one, one, user1.From, user2)
	tx4 := newPlasmaTestTx(t, i, three, uid, one, two, owner.From, user1)
	tx5 := newPlasmaTestTx(t, i, four, uid, one, three, user2.From, owner)

	_, err := i.rootUser2Session.StartExit(tx4.rawTx, tx4.proof, four,
		tx5.rawTx, tx5.proof, five)
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
	ethTx, err := i.rootUser2Session.ChallengeExit(uid, tx2Good.rawTx,
		tx2Good.proof, two)
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

	newPlasmaTestTx(t, i, zero, uid, one, zero, user2.From, user1)
	tx2 := newPlasmaTestTx(t, i, one, uid, one, one, owner.From, user2)
	tx3 := newPlasmaTestTx(t, i, two, uid, one, two, user1.From, owner)
	tx4 := newPlasmaTestTx(t, i, three, uid, one, three, user2.From, user1)
	tx5 := newPlasmaTestTx(t, i, four, uid, one, four, user1.From, user2)

	ethTx6, err := i.rootUser1Session.StartExit(tx4.rawTx, tx4.proof,
		four, tx5.rawTx, tx5.proof, five)
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

	ethTx7, err := i.rootOwnerSession.ChallengeExit(uid, tx2.rawTx,
		tx2.proof, two)
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
		cha.ChallengeTx, tx3.rawTx, tx3.proof, three)
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
