package build

import (
	"bytes"
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

func TestBuild(t *testing.T) {
	accounts := account.GenAccounts(3)
	owner := accounts[0]
	server := backend.NewSimulatedBackend(account.Addresses(accounts))

	address, _, err := rootchain.Deploy(owner.TransactOpts, server)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := abi.JSON(strings.NewReader(rootchain.RootChainABI))
	if err != nil {
		t.Fatal(err)
	}

	c, err := NewContract(address, parsed, server.Connect())
	if err != nil {
		t.Fatal(err)
	}

	blockHash := common.BigToHash(big.NewInt(123))

	tx1, err := c.Transaction(owner.TransactOpts,
		"newBlock", blockHash)
	if err != nil {
		t.Fatal(err)
	}

	rawTx, err := tx1.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	tx := &types.Transaction{}
	tx.UnmarshalJSON(rawTx)

	err = server.Connect().SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	tr, err := server.Mine(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	if tr.Status != 1 {
		t.Fatal("wrong transaction")
	}

	session, err := rootchain.NewRootChainSession(*owner.TransactOpts,
		address, server)
	if err != nil {
		t.Fatal(err)
	}

	hash, err := session.ChildChain(big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(blockHash.Bytes(), common.Hash(hash).Bytes()) {
		t.Fatal("hashes not equal")
	}
}
