package service

import (
	"bytes"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pborman/uuid"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
)

var (
	server backend.Backend

	owner *account.PlasmaTransactOpts

	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
	five  = big.NewInt(5)
	zero  = big.NewInt(0)
)

type instance struct {
	service       *Service
	rootChainAddr common.Address
	owner         bind.TransactOpts
}

func newInstance(t *testing.T) *instance {
	rootChainAddr, _, err := rootchain.Deploy(
		owner.TransactOpts, server)
	if err != nil {
		panic(err)
	}

	session, err := rootchain.NewRootChainSession(
		*owner.TransactOpts, rootChainAddr, server)
	if err != nil {
		panic(err)
	}

	dir, err := ioutil.TempDir("", uuid.NewUUID().String())
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	blockDB, err := bolt.NewDB(filepath.Join(dir, bolt.BlocksBucket),
		bolt.BlocksBucket, nil)
	if err != nil {
		t.Fatal(err)
	}

	chptDB, err := bolt.NewDB(filepath.Join(dir, bolt.CheckpointsBucket),
		bolt.CheckpointsBucket, nil)
	if err != nil {
		t.Fatal(err)
	}

	service, err := NewService(session, server, blockDB, chptDB)
	if err != nil {
		t.Fatal(err)
	}

	return &instance{
		service:       service,
		rootChainAddr: rootChainAddr,
	}
}

func testTx(t *testing.T, prevBlock, uid,
	amount *big.Int, nonce *big.Int, newOwner common.Address,
	signer *account.PlasmaTransactOpts) *transaction.Transaction {
	unsignedTx, err := transaction.NewTransaction(
		prevBlock, uid, amount, nonce, newOwner)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := signer.PlasmaSigner(signer.From, unsignedTx)
	if err != nil {
		t.Fatalf("failed to sign transaction %s", err)
	}

	addr, err := transaction.Sender(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(signer.From.Bytes(), addr.Bytes()) {
		t.Fatal("addresses not equal")
	}
	return tx
}

func TestNewService(t *testing.T) {
	newInstance(t)
}

func TestMain(m *testing.M) {
	testAccounts := account.GenAccounts(1)
	server = backend.NewSimulatedBackend(account.Addresses(testAccounts))
	owner = testAccounts[0]

	os.Exit(m.Run())
}
