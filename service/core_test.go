package service

import (
	"bytes"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pborman/uuid"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/erc20token"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	server backend.Backend

	owner *account.PlasmaTransactOpts
	user1 *account.PlasmaTransactOpts

	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	zero  = big.NewInt(0)
)

type instance struct {
	service              *Service
	rootChainAddr        common.Address
	mediatorAddress      common.Address
	mediatorUser1Session *mediator.MediatorSession
	mediatorOwnerSession *mediator.MediatorSession
	rootOwnerSession     *rootchain.RootChainSession
	rootUser1Session     *rootchain.RootChainSession
}

func newInstance(t *testing.T) *instance {
	mediatorAddr, _, err := mediator.Deploy(owner.TransactOpts, server)
	if err != nil {
		t.Fatal(err)
	}

	mSession, err := mediator.NewMediatorSession(*owner.TransactOpts,
		mediatorAddr, server)
	if err != nil {
		t.Fatal(err)
	}

	mUserSession, err := mediator.NewMediatorSession(*user1.TransactOpts,
		mediatorAddr, server)
	if err != nil {
		t.Fatal(err)
	}

	rootChainAddr, err := mSession.RootChain()
	if err != nil {
		t.Fatal(err)
	}

	session, err := rootchain.NewRootChainSession(
		*owner.TransactOpts, rootChainAddr, server)
	if err != nil {
		t.Fatal(err)
	}

	rSession, err := rootchain.NewRootChainSession(
		*user1.TransactOpts, rootChainAddr, server)
	if err != nil {
		t.Fatal(err)
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

	parsed, err := abi.JSON(strings.NewReader(rootchain.RootChainABI))
	if err != nil {
		t.Fatal(err)
	}

	rchc, err := build.NewContract(rootChainAddr, parsed, server.Connect())
	if err != nil {
		t.Fatal(err)
	}

	mParsed, err := abi.JSON(strings.NewReader(mediator.MediatorABI))
	if err != nil {
		t.Fatal(err)
	}

	mc, err := build.NewContract(mediatorAddr, mParsed, server.Connect())
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(session, server, blockDB, chptDB, rchc, mc)

	return &instance{
		service:              service,
		rootChainAddr:        rootChainAddr,
		mediatorAddress:      mediatorAddr,
		mediatorOwnerSession: mSession,
		mediatorUser1Session: mUserSession,
		rootOwnerSession:     session,
		rootUser1Session:     rSession,
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

func deployToken(t *testing.T,
	account *bind.TransactOpts) (address common.Address,
	contract *erc20token.ExampleToken) {
	address, contract, err := erc20token.Deploy(account, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func mint(t *testing.T, session *erc20token.ExampleTokenSession,
	acc common.Address, val *big.Int) {
	tx, err := session.Mint(acc, val)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}
}

func increaseApproval(t *testing.T, session *erc20token.ExampleTokenSession,
	spender common.Address, addedValue *big.Int) {
	tx, err := session.IncreaseApproval(spender, addedValue)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to approval tokens")
	}
}

func deposit(t *testing.T, session *mediator.MediatorSession,
	token common.Address, amount *big.Int, executable bool) {
	tx, err := session.Deposit(token, amount)
	if !executable {
		if err == nil {
			t.Fatal("It should not be executed")
		}
		return
	}
	if err != nil {
		t.Fatal(err)
	}

	if !server.GoodTransaction(tx) {
		t.Fatal("failed to deposit tokens")
	}
}

func tokenSession(t *testing.T, account *bind.TransactOpts,
	contact common.Address) (session *erc20token.ExampleTokenSession) {
	session, err := erc20token.NewExampleTokenSession(*account,
		contact, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

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

func TestNewService(t *testing.T) {
	newInstance(t)
}

func TestMain(m *testing.M) {
	testAccounts := account.GenAccounts(2)
	server = backend.NewSimulatedBackend(account.Addresses(testAccounts))
	owner = testAccounts[0]
	user1 = testAccounts[1]
	os.Exit(m.Run())
}
