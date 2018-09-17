package transport

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/pborman/uuid"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
	"github.com/SmartMeshFoundation/SmartPlasma/service"
)

const (
	tcp = "tcp"
)

var (
	clients  []*Client
	numUsers = 1000
	rpcPort  uint16
	env      *Env
)

type Env struct {
	dir             string
	server          *Server
	accounts        []*account.PlasmaTransactOpts
	mediatorAddress common.Address
	rootChainAddr   common.Address
	mediatorABI     abi.ABI
	rootChainABI    abi.ABI
	backend         backend.Backend
}

func getPort() (port int, err error) {
	addr, err := net.ResolveTCPAddr(tcp, "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP(tcp, addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

func newEnv() *Env {
	testAccounts := account.GenAccounts(numUsers)
	owner := testAccounts[0]

	server := backend.NewSimulatedBackend(account.Addresses(testAccounts))

	mediatorAddr, _, err := mediator.Deploy(owner.TransactOpts, server)
	if err != nil {
		panic(err)
	}

	mSession, err := mediator.NewMediatorSession(*owner.TransactOpts,
		mediatorAddr, server)
	if err != nil {
		panic(err)
	}

	rootChainAddr, err := mSession.RootChain()
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
		panic(err)
	}

	blockDB, err := bolt.NewDB(filepath.Join(dir, bolt.BlocksBucket),
		bolt.BlocksBucket, nil)
	if err != nil {
		panic(err)
	}

	checkpointDB, err := bolt.NewDB(filepath.Join(dir, bolt.CheckpointsBucket),
		bolt.CheckpointsBucket, nil)
	if err != nil {
		panic(err)
	}

	parsed, err := abi.JSON(strings.NewReader(rootchain.RootChainABI))
	if err != nil {
		panic(err)
	}

	rchc, err := build.NewContract(rootChainAddr, parsed, server.Connect())
	if err != nil {
		panic(err)
	}

	mParsed, err := abi.JSON(strings.NewReader(mediator.MediatorABI))
	if err != nil {
		panic(err)
	}

	mc, err := build.NewContract(mediatorAddr, mParsed, server.Connect())
	if err != nil {
		panic(err)
	}

	s := service.NewService(session, server, blockDB, checkpointDB, rchc, mc)

	srv := NewServer(100, rpcPort, s)

	fatal := make(chan error)

	go func() {
		fatal <- srv.ListenAndServe()
	}()

	select {
	case err := <-fatal:
		if err != nil {
			panic(err)
		}
	case <-time.After(time.Microsecond * 100):
	}

	return &Env{
		dir:             dir,
		server:          srv,
		accounts:        testAccounts,
		mediatorAddress: mediatorAddr,
		rootChainAddr:   rootChainAddr,
		backend:         server,
		rootChainABI:    parsed,
		mediatorABI:     mParsed,
	}
}

func (e *Env) Close() error {
	defer os.RemoveAll(e.dir)
	return e.server.Close()
}

func newClient(user *account.PlasmaTransactOpts,
	rootChainAddr, mediatorAddress common.Address, rootChainABI,
	mediatorABI abi.ABI) *Client {
	cli := NewClient(100, user)

	rc, err := build.NewContract(rootChainAddr, rootChainABI, cli)
	if err != nil {
		panic(err)
	}

	mc, err := build.NewContract(mediatorAddress, mediatorABI, cli)
	if err != nil {
		panic(err)
	}
	cli.RemoteEthereumClient(rc, mc)

	return cli
}

func initEnv() {
	port, err := getPort()
	if err != nil {
		panic(err)
	}

	rpcPort = uint16(port)

	env = newEnv()
	defer env.Close()

	for _, acc := range env.accounts {
		cli := newClient(acc, env.rootChainAddr,
			env.mediatorAddress, env.rootChainABI, env.mediatorABI)
		err := cli.Connect("localhost", rpcPort)
		if err != nil {
			panic(err)
		}
		clients = append(clients, cli)
	}
}

// BenchmarkCurrentBlock checks the speed of receiving the block.
func BenchmarkCurrentBlock(b *testing.B) {
	initEnv()
	defer env.Close()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := clients[0].CurrentBlock()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkAcceptTransaction checks the speed
// of write transaction to current block.
func BenchmarkAcceptTransaction(b *testing.B) {
	initEnv()
	defer env.Close()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		tx, err := transaction.NewTransaction(zero, one, two,
			three, clients[1].Opts().From)
		if err != nil {
			b.Fatal(err)
		}

		signedTx, err := clients[0].Opts().PlasmaSigner(
			clients[0].Opts().From, tx)
		if err != nil {
			b.Fatal(err)
		}

		buf := bytes.NewBuffer([]byte{})

		err = signedTx.EncodeRLP(buf)
		if err != nil {
			b.Fatal(err)
		}

		err = clients[0].AcceptTransaction(buf.Bytes())
		if err != nil {
			b.Fatal(err)
		}
	}
}
