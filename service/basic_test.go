package service

import (
	"testing"

	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

func TestMediatorTransaction(t *testing.T) {
	i := newInstance(t)
	tokenAddr, _ := deployToken(t, owner.TransactOpts)
	tokOwnerSession := tokenSession(t, owner.TransactOpts, tokenAddr)
	tokUserSession := tokenSession(t, user1.TransactOpts, tokenAddr)

	mint(t, tokOwnerSession, user1.From, one)
	increaseApproval(t, tokUserSession, i.mediatorAddress, one)

	tx, err := i.service.mediatorContractWrapper.Transaction(
		user1.TransactOpts, "deposit", tokenAddr, one)
	if err != nil {
		t.Fatal(err)
	}

	rawTx, err := tx.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	err = i.service.MediatorTransaction(rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(tx) {
		t.Fatal("failed to deposit tokens")
	}

	logs, err := rootchain.LogsDeposit(i.rootOwnerSession.Contract)
	if err != nil {
		t.Fatalf("failed to parse deposit logs %s", err)
	}

	if len(logs) != 1 {
		t.Fatal("wrong number of logs")
	}

	if logs[0].Depositor.String() != user1.From.String() {
		t.Fatal("wrong depositor")
	}
}

func TestServiceRootChainTransaction(t *testing.T) {
	// test exists contract/build/TestBuild
}
