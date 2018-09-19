package main

import (
	"bytes"
	"context"
	"fmt"
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
	"github.com/SmartMeshFoundation/Spectrum/core/types"
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

// Total 1000 users.
// Each user owns 20 deposits.
// Each cycle, each user transfers all his deposits to another user.
// Then the operator collects the block.
// He stores it in the database and sends the block hash to the blockchain.
// Assumptions:
// 1) Clients do not save transactions on their side.
// 2) Transactions is not checked.
var (
	one         = big.NewInt(1)
	deposits    = 20
	supply      = new(big.Int).Mul(one, big.NewInt(int64(deposits)))
	usersNumber = 1000
	cycles      = 3
)

type user struct {
	acc      *account.PlasmaTransactOpts
	uids     map[string]uint64
	incoming map[string]uint64
	cli      *transport.Client
}

type environment struct {
	dir              string
	accounts         []*user
	mediatorAddress  common.Address
	rootChainAddress common.Address
	tokenAddress     common.Address
	mediatorABI      abi.ABI
	rootChainABI     abi.ABI
	backend          backend.Backend
	server           *transport.Server
	port             uint16
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
		rootChainContract, mediatorContract, false)

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

	testAccounts := account.GenAccounts(usersNumber)
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

	parsed, err := abi.JSON(strings.NewReader(rootchain.RootChainABI))
	if err != nil {
		panic(err)
	}

	mParsed, err := abi.JSON(strings.NewReader(mediator.MediatorABI))
	if err != nil {
		panic(err)
	}

	tAddr, _ := deployToken(owner.TransactOpts, server)

	tOwnerSession := tokenSession(owner.TransactOpts,
		tAddr, server)

	var users []*user

	// mint and approval to Mediator contract.
	for _, acc := range testAccounts {
		mint(tOwnerSession, acc.From, supply, server)

		tSession := tokenSession(acc.TransactOpts,
			tAddr, server)
		increaseApproval(tSession, mAddr, supply, server)

		users = append(users, &user{acc: acc})
	}

	// get free port
	port := getPort()

	// creates new Plasma server.
	srv := newServer(
		*testAccounts[0], rAddr, mAddr, parsed,
		mParsed, server, dir, port)
	go srv.ListenAndServe()

	time.Sleep(time.Microsecond * 200)

	return &environment{
		dir:              dir,
		accounts:         users,
		mediatorAddress:  mAddr,
		rootChainAddress: rAddr,
		tokenAddress:     tAddr,
		mediatorABI:      mParsed,
		rootChainABI:     parsed,
		backend:          server,
		server:           srv,
		port:             port,
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

func (u *user) deposit(env *environment) {
	cli1 := newDirectClient(u.acc, env.rootChainAddress,
		env.mediatorAddress, env.backend)
	err := cli1.Connect("", env.port)
	if err != nil {
		panic(err)
	}
	u.cli = cli1
	u.uids = make(map[string]uint64)
	u.incoming = make(map[string]uint64)

	for i := 0; i < deposits; i++ {
		tx, err := cli1.Deposit(env.tokenAddress, one)
		if err != nil {
			panic(err)
		}
		tr, err := cli1.WaitMined(context.Background(), tx)
		if err != nil {
			panic(err)
		}

		if tr.Status == types.ReceiptStatusFailed {
			panic("receipt status is failed")
		}

		uid := new(big.Int).SetBytes(tr.Logs[1].Data[64:96])
		u.uids[uid.String()] = 0
	}
}

func (u *user) transact(newOwner *user) {
	for uidStr, nonceUint64 := range u.uids {
		// create Plasma Cash transaction
		uid, _ := new(big.Int).SetString(uidStr, 10)
		prevBlock := new(big.Int).SetUint64(nonceUint64)
		nonce := new(big.Int).SetUint64(nonceUint64)

		tx, err := transaction.NewTransaction(
			prevBlock, uid, one, nonce, newOwner.acc.From)
		if err != nil {
			panic(err)
		}
		// sign Plasma Cash transaction
		tx1, err := u.acc.PlasmaSigner(u.acc.From, tx)
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
		err = u.cli.AcceptTransaction(buf.Bytes())
		if err != nil {
			panic(err)
		}
		delete(u.uids, uidStr)
		newOwner.incoming[uidStr] = nonceUint64 + 1
	}

	for k, v := range u.incoming {
		u.uids[k] = v
	}

	u.incoming = make(map[string]uint64)
}

func (env *environment) cycle(iteration int) {
	cliOwner := env.accounts[0].cli

	for i := 0; i < iteration; i++ {
		fmt.Println("start transaction cycle")
		start := time.Now().UnixNano()
		for k := 0; k < len(env.accounts); k++ {
			u := env.accounts[k]

			var newOwnerIndex int

			if len(env.accounts) == k+1 {
				newOwnerIndex = 0
			} else {
				newOwnerIndex = k + 1
			}

			newOwner := env.accounts[newOwnerIndex]

			u.transact(newOwner)
		}
		stop := time.Now().UnixNano()
		fmt.Printf("stop transaction cycle. time: %d ms\n", (stop-start)/1000000)

		fmt.Println("start building block ")
		start = time.Now().UnixNano()
		hash, err := cliOwner.BuildBlock()
		if err != nil {
			panic(err)
		}
		stop = time.Now().UnixNano()
		fmt.Printf("finish building block, time: %d ms\n", (stop-start)/1000000)

		lastBlock, err := cliOwner.LastBlockNumber()
		if err != nil {
			panic(err)
		}

		fmt.Println("start saving block ", lastBlock.Uint64()+1)
		start = time.Now().UnixNano()

		err = cliOwner.SaveCurrentBlock(lastBlock.Uint64() + 1)
		if err != nil {
			panic(err)
		}

		// Operator publishes hash for block
		sendBlock1HashTx, err := cliOwner.SendBlockHash(hash)
		if err != nil {
			panic(err)
		}

		sendBlock1Tr, err := cliOwner.WaitMined(
			context.Background(), sendBlock1HashTx)
		if err != nil {
			panic(err)
		}

		if sendBlock1Tr.Status != 1 {
			panic("wrong tx status")
		}
		stop = time.Now().UnixNano()
		fmt.Printf("finish saving block %d time: %d ms\n",
			lastBlock.Uint64()+1, (stop-start)/1000000)

		err = cliOwner.InitBlock()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	env := newTestEnvironment()
	defer os.RemoveAll(env.dir)
	defer env.server.Close()

	fmt.Println("started the process of creating deposits")
	for _, u := range env.accounts {
		u.deposit(env)
	}
	fmt.Println("finished the process of creating deposits")

	env.cycle(cycles)
}
