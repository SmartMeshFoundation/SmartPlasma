package transport

import (
	"bytes"
	"io/ioutil"
	"math/big"
	"net/http/httptest"
	"net/rpc"
	"os"
	"path/filepath"
	"testing"

	"github.com/pborman/uuid"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
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
}

func testServer() *testService {
	rpcServer := rpc.NewServer()
	testAccounts := account.GenAccounts(1)
	owner := testAccounts[0]

	server := backend.NewSimulatedBackend(account.Addresses(testAccounts))

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

	smartPlasma := &SmartPlasma{
		currentChpt:  checkpoints.NewBlock(),
		currentBlock: transactions.NewTxBlock(),
		blockBase:    blockDB,
		chptBase:     chptDB,
		session:      session,
		backend:      server,
		timeout:      100,
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
	defer os.RemoveAll(s.dir)
	defer s.smartPlasma.blockBase.Close()
	defer s.smartPlasma.chptBase.Close()
	defer s.server.Close()
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
	s := testServer()
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
	s := testServer()
	defer s.Close()

	cli := testClient(t, s)
	defer cli.Close()

	newOwner := account.Account(account.GenKey())

	tx, err := transaction.NewTransaction(zero, one, two, three, newOwner.From)
	if err != nil {
		t.Fatal(err)
	}

	err = s.smartPlasma.currentBlock.AddTx(tx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.smartPlasma.currentBlock.Build()
	if err != nil {
		t.Fatal(err)
	}

	raw, err := s.smartPlasma.currentBlock.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.CreateProof(one, one.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error == "" {
		t.Fatal("error")
	}

	err = s.smartPlasma.blockBase.Set([]byte(one.String()), raw)
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
	s := testServer()
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
	s := testServer()
	defer s.Close()

	cli := testClient(t, s)
	defer cli.Close()

	err := s.smartPlasma.currentChpt.AddCheckpoint(one, two)
	if err != nil {
		t.Fatal(err)
	}

	hash := s.smartPlasma.currentChpt.Hash()

	raw, err := s.smartPlasma.currentChpt.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.CreateUIDStateProof(one, hash)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error == "" {
		t.Fatal("error")
	}

	err = s.smartPlasma.chptBase.Set(hash.Bytes(), raw)
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
