package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/pborman/uuid"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/erc20token"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
	"github.com/SmartMeshFoundation/SmartPlasma/service"
	"github.com/SmartMeshFoundation/SmartPlasma/transport"
)

type environment struct {
	dir              string
	accounts         []*account.PlasmaTransactOpts
	mediatorAddress  common.Address
	rootChainAddress common.Address
	tokenAddress     common.Address
	mediatorABI      abi.ABI
	rootChainABI     abi.ABI
	backend          backend.Backend
	rootChainSession *rootchain.RootChainSession
}

// newServer creates new PlasmaCash server.
func newServer(owner account.PlasmaTransactOpts, rootChainAddr,
	mediatorAddr common.Address, rootChainABI, mediatorABI abi.ABI,
	backend backend.Backend, databaseDir string,
	port uint16) *transport.Server {

	// session for communication RootChain smart contract.
	rSession, err := rootchain.NewRootChainSession(*owner.TransactOpts,
		rootChainAddr, backend)
	if err != nil {
		panic(err)
	}

	// object for handling remote transactions
	// from clients to RootChain contract.
	rootChainContract, err := build.NewContract(rootChainAddr,
		rootChainABI, backend.Connect())
	if err != nil {
		panic(err)
	}

	// object for handling remote transactions
	// from clients to Mediator contract.
	mediatorContract, err := build.NewContract(mediatorAddr,
		mediatorABI, backend.Connect())
	if err != nil {
		panic(err)
	}

	// object for communication local database for Plasma blocks.
	blockDB, err := bolt.NewDB(filepath.Join(databaseDir, bolt.BlocksBucket),
		bolt.BlocksBucket, nil)
	if err != nil {
		panic(err)
	}

	// object for communication local database for checkpoint blocks.
	checkpointDB, err := bolt.NewDB(filepath.Join(databaseDir, bolt.CheckpointsBucket),
		bolt.CheckpointsBucket, nil)
	if err != nil {
		panic(err)
	}

	// create new PlasmaCash service.
	s := service.NewService(rSession, backend, blockDB, checkpointDB,
		rootChainContract, mediatorContract)

	// create new RPC server for communication with PlasmaCash clients.
	return transport.NewServer(100, port, s)
}

// newDirectClient the client that interacts with the blockchain directly.
// He creates the transactions himself and sends them to the Spectrum.
func newDirectClient(user *account.PlasmaTransactOpts, rootChainAddr,
	mediatorAddr common.Address, backend backend.Backend) *transport.Client {
	cli := transport.NewClient(100, user)
	cli.DirectEthereumClient(
		*user.TransactOpts, mediatorAddr, rootChainAddr, backend)

	return cli
}

// newRemoveClient the client does not interact with the blockchain directly.
// It creates transactions locally. He signs them with his private key,
// packs it and sends it to the server. The server has access to the Spectrum,
// Server translates them to the network without changing the contents.
func newRemoveClient(user *account.PlasmaTransactOpts, rootChainAddr,
	mediatorAddr common.Address, rootChainABI,
	mediatorABI abi.ABI) *transport.Client {
	cli := transport.NewClient(100, user)

	rc, err := build.NewContract(rootChainAddr, rootChainABI, cli)
	if err != nil {
		panic(err)
	}

	mc, err := build.NewContract(mediatorAddr, mediatorABI, cli)
	if err != nil {
		panic(err)
	}
	cli.RemoteEthereumClient(rc, mc)

	return cli
}

// getPort gets free network port on host.
func getPort() (port uint16) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return uint16(l.Addr().(*net.TCPAddr).Port)
}

// newTestEnvironment is environment for test.
func newTestEnvironment() *environment {
	dir, err := ioutil.TempDir("", uuid.NewUUID().String())
	if err != nil {
		panic(err)
	}

	testAccounts := account.GenAccounts(3)
	owner := testAccounts[0]

	server := backend.NewSimulatedBackend(account.Addresses(testAccounts))

	mAddr, _, err := mediator.Deploy(owner.TransactOpts, server)
	if err != nil {
		panic(err)
	}

	mSession, err := mediator.NewMediatorSession(*owner.TransactOpts,
		mAddr, server)
	if err != nil {
		panic(err)
	}

	rAddr, err := mSession.RootChain()
	if err != nil {
		panic(err)
	}

	rSession, err := rootchain.NewRootChainSession(
		*owner.TransactOpts, rAddr, server)
	if err != nil {
		panic(err)
	}

	parsed, err := abi.JSON(strings.NewReader(rootchain.RootChainABI))
	if err != nil {
		panic(err)
	}

	mParsed, err := abi.JSON(strings.NewReader(mediator.MediatorABI))
	if err != nil {
		panic(err)
	}

	tAddr, _ := deployToken(owner.TransactOpts, server)

	one := big.NewInt(1)

	tOwnerSession := tokenSession(owner.TransactOpts,
		tAddr, server)

	// mint and approval to Mediator contract.
	for _, acc := range testAccounts {
		mint(tOwnerSession, acc.From, one, server)

		tSession := tokenSession(acc.TransactOpts,
			tAddr, server)
		increaseApproval(tSession, mAddr, one, server)
	}

	return &environment{
		dir:              dir,
		accounts:         testAccounts,
		mediatorAddress:  mAddr,
		rootChainAddress: rAddr,
		tokenAddress:     tAddr,
		mediatorABI:      mParsed,
		rootChainABI:     parsed,
		backend:          server,
		rootChainSession: rSession,
	}
}

func tokenSession(account *bind.TransactOpts, contact common.Address,
	backend backend.Backend) (session *erc20token.ExampleTokenSession) {
	session, err := erc20token.NewExampleTokenSession(*account,
		contact, backend)
	if err != nil {
		panic(err)
	}
	return session
}

func deployToken(account *bind.TransactOpts,
	backend backend.Backend) (address common.Address,
	contract *erc20token.ExampleToken) {
	address, contract, err := erc20token.Deploy(account, backend)
	if err != nil {
		panic(err)
	}
	return address, contract
}

func mint(session *erc20token.ExampleTokenSession,
	acc common.Address, val *big.Int, backend backend.Backend) {
	tx, err := session.Mint(acc, val)
	if err != nil {
		panic(err)
	}
	if !backend.GoodTransaction(tx) {
		panic("failed to mint tokens")
	}
}

func increaseApproval(session *erc20token.ExampleTokenSession,
	spender common.Address, addedValue *big.Int, backend backend.Backend) {
	tx, err := session.IncreaseApproval(spender, addedValue)
	if err != nil {
		panic(err)
	}
	if !backend.GoodTransaction(tx) {
		panic("failed to approval tokens")
	}
}

func timeMachine(adjustment time.Duration,
	server backend.Backend) {
	if sim, ok := server.(backend.Simulator); ok {
		if err := sim.AdjustTime(adjustment); err != nil {
			panic(err)
		}
	}
}

func main() {
	env := newTestEnvironment()
	defer os.RemoveAll(env.dir)

	// get free port
	port := getPort()

	// creates new Plasma server.
	srv := newServer(*env.accounts[0], env.rootChainAddress,
		env.mediatorAddress, env.rootChainABI,
		env.mediatorABI, env.backend, env.dir, port)

	go srv.ListenAndServe()
	defer srv.Close()

	time.Sleep(time.Microsecond * 200)

	owner := env.accounts[0]
	user1 := env.accounts[1]
	user2 := env.accounts[2]

	cli0 := newDirectClient(owner, env.rootChainAddress,
		env.mediatorAddress, env.backend)

	err := cli0.Connect("", port)
	if err != nil {
		panic(err)
	}

	cli1 := newDirectClient(user1, env.rootChainAddress,
		env.mediatorAddress, env.backend)

	err = cli1.Connect("", port)
	if err != nil {
		panic(err)
	}

	cli2 := newRemoveClient(user2, env.rootChainAddress,
		env.mediatorAddress, env.rootChainABI, env.mediatorABI)

	err = cli2.Connect("", port)
	if err != nil {
		panic(err)
	}

	// new deposit
	txDepositU1, err := cli1.Deposit(env.tokenAddress, big.NewInt(1))
	if err != nil {
		panic(err)
	}

	// waiting for a transaction to be written to a Spectrum block.
	tr, err := cli1.WaitMined(context.Background(), txDepositU1)
	if err != nil {
		panic(err)
	}

	if tr.Status != 1 {
		panic("no deposit")
	}

	// uid from Spectrum log - Deposit(account, amount, uint256(uid));
	uid := new(big.Int).SetBytes(tr.Logs[1].Data[64:96])

	// create Plasma Cash transaction
	unTx1, err := transaction.NewTransaction(big.NewInt(0), uid,
		big.NewInt(1), big.NewInt(0), user1.From)

	// sign Plasma Cash transaction
	tx1, err := user1.PlasmaSigner(user1.From, unTx1)
	if err != nil {
		panic(err)
	}

	// encode Plasma Cash transaction
	buf := bytes.NewBuffer([]byte{})
	err = tx1.EncodeRLP(buf)
	if err != nil {
		panic(err)
	}

	// send Plasma Cash transaction to server
	err = cli1.AcceptTransaction(buf.Bytes())
	if err != nil {
		panic(err)
	}

	// build current Plasma Cash block
	respBuildBlock1, err := cli1.BuildBlock()
	if err != nil {
		panic(err)
	}

	// get last Plasma Cash block number from root chain
	lastBlock, err := cli1.LastBlockNumber()
	if err != nil {
		panic(err)
	}

	// get current block
	rawBlock1, err := cli1.CurrentBlock()
	if err != nil {
		panic(err)
	}

	// save current block to database
	err = cli1.SaveBlockToDB(lastBlock.Uint64()+1, rawBlock1)
	if err != nil {
		panic(err)
	}

	// Operator publishes hash for block 1
	sendBlock1HashTx, err := cli0.SendBlockHash(respBuildBlock1)
	if err != nil {
		panic(err)
	}

	sendBlock1Tr, err := cli1.WaitMined(context.Background(), sendBlock1HashTx)
	if err != nil {
		panic(err)
	}

	if sendBlock1Tr.Status != 1 {
		panic("wrong tx status")
	}

	// create proof for transaction #1 on Plasma Cash block #1.
	proof1, err := cli1.CreateProof(uid, 1)
	if err != nil {
		panic(err)
	}

	// verify proof for transaction #1 on Plasma Cash block #1
	verifyTx1Resp, err := cli1.VerifyTxProof(
		uid, tx1.Hash(), 1, proof1)

	if !verifyTx1Resp {
		panic("transaction #1 not verified")
	}

	err = cli1.InitBlock()
	if err != nil {
		panic(err)
	}

	// create Plasma Cash transaction #2
	unTx2, err := transaction.NewTransaction(big.NewInt(1), uid,
		big.NewInt(1), big.NewInt(1), user2.From)

	// sign Plasma Cash transaction #2
	tx2, err := user1.PlasmaSigner(user1.From, unTx2)
	if err != nil {
		panic(err)
	}

	// encode Plasma Cash transaction
	buf = bytes.NewBuffer([]byte{})
	err = tx2.EncodeRLP(buf)
	if err != nil {
		panic(err)
	}

	// send Plasma Cash transaction #2 to server
	err = cli1.AcceptTransaction(buf.Bytes())
	if err != nil {
		panic(err)
	}

	// build current Plasma Cash block
	respBuildBlock2, err := cli1.BuildBlock()
	if err != nil {
		panic(err)
	}

	// get last Plasma Cash block number from root chain
	lastBlock, err = cli1.LastBlockNumber()
	if err != nil {
		panic(err)
	}

	// get current block
	rawBlock2, err := cli1.CurrentBlock()
	if err != nil {
		panic(err)
	}

	// save current block to database
	err = cli1.SaveBlockToDB(lastBlock.Uint64()+1, rawBlock2)
	if err != nil {
		panic(err)
	}

	// Operator publishes hash for block 2
	sendBlock2HashTx, err := cli0.SendBlockHash(respBuildBlock2)
	if err != nil {
		panic(err)
	}

	sendBlock2Tr, err := cli1.WaitMined(context.Background(), sendBlock2HashTx)
	if err != nil {
		panic(err)
	}

	if sendBlock2Tr.Status != 1 {
		panic("wrong tx status")
	}

	// create proof for transaction #1 on Plasma Cash block #1.
	proof2, err := cli1.CreateProof(uid, 2)
	if err != nil {
		panic(err)
	}

	// verify proof for transaction #1 on Plasma Cash block #1
	verifyTx2Resp, err := cli1.VerifyTxProof(
		uid, tx2.Hash(), 2, proof2)

	if !verifyTx2Resp {
		panic("transaction #1 not verified")
	}

	bufTx1 := bytes.NewBuffer([]byte{})

	if err := tx1.EncodeRLP(bufTx1); err != nil {
		panic(err)
	}

	bufTx2 := bytes.NewBuffer([]byte{})

	if err := tx2.EncodeRLP(bufTx2); err != nil {
		panic(err)
	}

	// start exit
	exitTx, err := cli2.StartExit(
		bufTx1.Bytes(), proof1, big.NewInt(1),
		bufTx2.Bytes(), proof2, big.NewInt(2))
	if err != nil {
		panic(err)
	}
	trExit, err := cli2.WaitMined(context.Background(), exitTx)
	if err != nil {
		panic(err)
	}

	if trExit.Status != 1 {
		panic("no deposit")
	}

	// attempt to withdraw
	_, err = cli2.Withdraw(
		bufTx1.Bytes(), proof1, big.NewInt(1),
		bufTx2.Bytes(), proof2, big.NewInt(2))
	if err == nil {
		panic("the period of challenge has not yet been completed")
	}

	// time.Now + 3 weeks
	timeMachine(time.Duration(504*time.Hour), env.backend)

	tokSession := tokenSession(
		user2.TransactOpts, env.tokenAddress, env.backend)

	balance, err := tokSession.BalanceOf(user2.From)
	if err != nil {
		panic(err)
	}

	// balance = 1 token
	if balance.Uint64() != 1 {
		panic("wrong balance")
	}

	// withdraw
	withdrawTx, err := cli2.Withdraw(
		bufTx1.Bytes(), proof1, big.NewInt(1),
		bufTx2.Bytes(), proof2, big.NewInt(2))
	if err != nil {
		panic(err)
	}

	trWithdraw, err := cli2.WaitMined(context.Background(), withdrawTx)
	if err != nil {
		panic(err)
	}

	if trWithdraw.Status != 1 {
		panic("no withdraw")
	}

	// attempt to withdraw
	_, err = cli2.Withdraw(
		bufTx1.Bytes(), proof1, big.NewInt(1),
		bufTx2.Bytes(), proof2, big.NewInt(2))
	if err == nil {
		panic("can not withdraw several times")
	}

	balance, err = tokSession.BalanceOf(user2.From)
	if err != nil {
		panic(err)
	}

	// balance = 2 tokens
	if balance.Uint64() != 2 {
		panic("wrong balance")
	}
}
