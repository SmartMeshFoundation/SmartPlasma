package transport

import (
	"bytes"
	"io/ioutil"
	"math/big"
	"net/http/httptest"
	"net/rpc"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pborman/uuid"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
	"github.com/SmartMeshFoundation/SmartPlasma/service"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
	five  = big.NewInt(5)
	zero  = big.NewInt(0)
)

type testService struct {
	dir         string
	server      *httptest.Server
	owner       *account.PlasmaTransactOpts
	smartPlasma *SmartPlasma
	blockBase   database.Database
	chptBase    database.Database
}

func testServer(t *testing.T) *testService {
	rpcServer := rpc.NewServer()
	testAccounts := account.GenAccounts(1)
	owner := testAccounts[0]

	server := backend.NewSimulatedBackend(account.Addresses(testAccounts))

	mediatorAddr, _, err := mediator.Deploy(owner.TransactOpts, server)
	if err != nil {
		t.Fatal(err)
	}

	mSession, err := mediator.NewMediatorSession(*owner.TransactOpts,
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

	chptDB, err := bolt.NewDB(filepath.Join(dir, bolt.CheckpointsBucket),
		bolt.CheckpointsBucket, nil)
	if err != nil {
		panic(err)
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

	s := service.NewService(session, server, blockDB, chptDB, rchc, mc)

	smartPlasma := &SmartPlasma{
		timeout: 100,
		service: s,
	}

	rpcServer.RegisterName("SmartPlasma", smartPlasma)

	httpServer := httptest.NewServer(rpcServer)

	return &testService{
		dir:         dir,
		server:      httpServer,
		owner:       owner,
		smartPlasma: smartPlasma,
	}
}

func (s *testService) Close() {
	os.RemoveAll(s.dir)
	s.smartPlasma.service.Close()
	s.server.Close()
}

func testClient(t *testing.T, srv *testService) *Client {
	cli := NewClient(100)
	err := cli.ConnectString(srv.server.URL[7:])
	if err != nil {
		t.Fatal(err)
	}
	return cli
}

func TestAcceptTransaction(t *testing.T) {
	s := testServer(t)
	defer s.Close()

	cli := testClient(t, s)
	defer cli.Close()

	tx, err := transaction.NewTransaction(zero, one, two, three, s.owner.From)
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBuffer([]byte{})

	err = tx.EncodeRLP(buf)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.AcceptTransaction(buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error != "" {
		t.Fatal("error")
	}
}

func TestCreateProof(t *testing.T) {
	s := testServer(t)
	defer s.Close()

	cli := testClient(t, s)
	defer cli.Close()

	newOwner := account.Account(account.GenKey())

	tx, err := transaction.NewTransaction(zero, one, two, three, newOwner.From)
	if err != nil {
		t.Fatal(err)
	}

	err = s.smartPlasma.service.AcceptTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.smartPlasma.service.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}

	curBlock := s.smartPlasma.service.CurrentBlock()

	resp, err := cli.CreateProof(one, one.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error == "" {
		t.Fatal("error")
	}

	err = s.smartPlasma.service.SaveBlockToDB(one.Uint64(), curBlock)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = cli.CreateProof(one, one.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error != "" || len(resp.Proof) == 0 {
		t.Fatal("error")
	}
}

func TestAddCheckpoint(t *testing.T) {
	s := testServer(t)
	defer s.Close()

	cli := testClient(t, s)
	defer cli.Close()

	resp, err := cli.AddCheckpoint(one, two)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error != "" {
		t.Fatal("error")
	}
}

func TestCreateUIDStateProof(t *testing.T) {
	s := testServer(t)
	defer s.Close()

	cli := testClient(t, s)
	defer cli.Close()

	err := s.smartPlasma.service.AcceptUIDState(one, two)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.smartPlasma.service.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	curChpt := s.smartPlasma.service.CurrentCheckpoint()

	hash := s.smartPlasma.service.CurrentCheckpoint().Hash()

	resp, err := cli.CreateUIDStateProof(one, hash)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error == "" {
		t.Fatal("error")
	}

	err = s.smartPlasma.service.SaveCheckpointToDB(curChpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = cli.CreateUIDStateProof(one, hash)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error != "" || len(resp.Proof) == 0 {
		t.Fatal("error")
	}
}
