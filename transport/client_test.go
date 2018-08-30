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
	"github.com/SmartMeshFoundation/SmartPlasma/database"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
	"github.com/SmartMeshFoundation/SmartPlasma/service"
)

var (
	zero  = big.NewInt(0)
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
)

type testService struct {
	dir              string
	server           *httptest.Server
	accounts         []*account.PlasmaTransactOpts
	smartPlasma      *SmartPlasma
	blockBase        database.Database
	chptBase         database.Database
	backend          backend.Backend
	rootChainAddress common.Address
	mediatorAddress  common.Address
}

type txData struct {
	rawTx     []byte
	proof     []byte
	block     uint64
	blockHash common.Hash
}

func newTestService(t *testing.T, numberAcc int) *testService {
	rpcServer := rpc.NewServer()
	testAccounts := account.GenAccounts(numberAcc)
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
		dir:              dir,
		server:           httpServer,
		backend:          server,
		accounts:         testAccounts,
		smartPlasma:      smartPlasma,
		mediatorAddress:  mediatorAddr,
		rootChainAddress: rootChainAddr,
	}
}

func (s *testService) Close() {
	os.RemoveAll(s.dir)
	s.smartPlasma.service.Close()
	s.server.Close()
}

func testClient(t *testing.T, srv *testService, direct bool, user *account.PlasmaTransactOpts) *Client {
	cli := NewClient(100, user)
	err := cli.ConnectString(srv.server.URL[7:])
	if err != nil {
		t.Fatal(err)
	}

	if direct {
		mSession, err := mediator.NewMediatorSession(*user.TransactOpts,
			srv.mediatorAddress, srv.backend)
		if err != nil {
			t.Fatal(err)
		}

		rSession, err := rootchain.NewRootChainSession(*user.TransactOpts,
			srv.rootChainAddress, srv.backend)
		if err != nil {
			t.Fatal(err)
		}
		cli.DirectEthereumClient(mSession, rSession)
	}

	parsed, err := abi.JSON(strings.NewReader(rootchain.RootChainABI))
	if err != nil {
		t.Fatal(err)
	}

	rc, err := build.NewContract(srv.rootChainAddress, parsed, cli)
	if err != nil {
		t.Fatal(err)
	}

	parsed2, err := abi.JSON(strings.NewReader(mediator.MediatorABI))
	if err != nil {
		t.Fatal(err)
	}

	mc, err := build.NewContract(srv.mediatorAddress, parsed2, cli)
	if err != nil {
		t.Fatal(err)
	}
	cli.RemoteEthereumClient(rc, mc)

	return cli
}

func deployToken(t *testing.T, account *bind.TransactOpts,
	backend backend.Backend) (address common.Address,
	contract *erc20token.ExampleToken) {
	address, contract, err := erc20token.Deploy(account, backend)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func mint(t *testing.T, session *erc20token.ExampleTokenSession,
	acc common.Address, val *big.Int, backend backend.Backend) {
	tx, err := session.Mint(acc, val)
	if err != nil {
		t.Fatal(err)
	}
	if !backend.GoodTransaction(tx) {
		t.Fatal("failed to mint tokens")
	}
}

func increaseApproval(t *testing.T, session *erc20token.ExampleTokenSession,
	spender common.Address, addedValue *big.Int, backend backend.Backend) {
	tx, err := session.IncreaseApproval(spender, addedValue)
	if err != nil {
		t.Fatal(err)
	}
	if !backend.GoodTransaction(tx) {
		t.Fatal("failed to approval tokens")
	}
}

func tokenSession(t *testing.T, account *bind.TransactOpts,
	contact common.Address,
	backend backend.Backend) (session *erc20token.ExampleTokenSession) {
	session, err := erc20token.NewExampleTokenSession(*account,
		contact, backend)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func deposit(t *testing.T, s *testService, cli *Client,
	amount *big.Int) *big.Int {
	tokenAddr, _ := deployToken(t, s.accounts[0].TransactOpts, s.backend)
	tokOwnerSession := tokenSession(t, s.accounts[0].TransactOpts,
		tokenAddr, s.backend)
	mint(t, tokOwnerSession, s.accounts[0].From, amount, s.backend)
	increaseApproval(t, tokOwnerSession, s.mediatorAddress, amount, s.backend)

	// for test
	rSession, err := rootchain.NewRootChainSession(*s.accounts[0].TransactOpts,
		s.rootChainAddress, s.backend)
	if err != nil {
		t.Fatal(err)
	}

	// for test
	uid, err := rootchain.GenerateNextUID(rSession,
		cli.opts.From, tokenAddr)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := cli.Deposit(tokenAddr, amount)
	if err != nil {
		t.Fatal(err)
	}
	if (tx.Hash() == common.Hash{}) {
		t.Fatal("Hash is null")
	}

	if !s.backend.GoodTransaction(tx) {
		t.Fatal("transaction is failed")
	}

	amount2, err := rSession.Wallet(common.BigToHash(uid))
	if err != nil {
		t.Fatal(err)
	}

	if amount2.Uint64() != amount.Uint64() {
		t.Fatal("amount is wrong")
	}

	return uid
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

func addTx(t *testing.T, uid *big.Int,
	tx *transaction.Transaction, cli *Client) *txData {
	buf := bytes.NewBuffer([]byte{})

	err := tx.EncodeRLP(buf)
	if err != nil {
		t.Fatal(err)
	}

	acceptResp, err := cli.AcceptTransaction(buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	if acceptResp.Error != "" {
		t.Fatal(acceptResp.Error)
	}

	buildResp, err := cli.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}

	if buildResp.Error != "" {
		t.Fatal(buildResp.Error)
	}

	sendBlockHashResp, err := cli.SendBlockHash(buildResp.Hash)
	if err != nil {
		t.Fatal(err)
	}

	if sendBlockHashResp.Error != "" {
		t.Fatal(sendBlockHashResp.Error)
	}

	lastBlock, err := cli.LastBlockNumber()
	if err != nil {
		t.Fatal(err)
	}

	currentBlockResp, err := cli.CurrentBlock()
	if err != nil {
		t.Fatal(err)
	}

	saveBlockResp, err := cli.SaveBlockToDB(lastBlock.Uint64(),
		currentBlockResp.Block)
	if err != nil {
		t.Fatal(err)
	}

	if saveBlockResp.Error != "" {
		t.Fatal(saveBlockResp.Error)
	}

	profResp, err := cli.CreateProof(uid, lastBlock.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if profResp.Error != "" {
		t.Fatal(profResp.Error)
	}

	initResp, err := cli.InitBlock()
	if err != nil {
		t.Fatal(err)
	}

	if initResp.Error != "" {
		t.Fatal(initResp.Error)
	}

	respProof, err := cli.VerifyTxProof(uid, tx.Hash(),
		lastBlock.Uint64(), profResp.Proof)
	if err != nil {
		t.Fatal(err)
	}

	if respProof.Error != "" {
		t.Fatal(respProof.Error)
	}

	if !respProof.Exists {
		t.Fatal("not exists")
	}

	return &txData{
		rawTx:     buf.Bytes(),
		proof:     profResp.Proof,
		block:     lastBlock.Uint64(),
		blockHash: buildResp.Hash,
	}
}

func testAcceptTransaction(t *testing.T, direct bool) {
	s := newTestService(t, 1)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
	defer cli.Close()

	tx, err := transaction.NewTransaction(zero, one, two,
		three, s.accounts[0].From)
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

func TestAcceptTransaction(t *testing.T) {
	testAcceptTransaction(t, true)
	testAcceptTransaction(t, false)
}

func testCreateProof(t *testing.T, direct bool) {
	s := newTestService(t, 1)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
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

func TestCreateProof(t *testing.T) {
	testCreateProof(t, true)
	testCreateProof(t, false)
}

func testAddCheckpoint(t *testing.T, direct bool) {
	s := newTestService(t, 1)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
	defer cli.Close()

	resp, err := cli.AddCheckpoint(one, two)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error != "" {
		t.Fatal("error")
	}
}

func TestAddCheckpoint(t *testing.T) {
	testAddCheckpoint(t, true)
	testAddCheckpoint(t, false)
}

func testCreateUIDStateProof(t *testing.T, direct bool) {
	s := newTestService(t, 1)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
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

func TestCreateUIDStateProof(t *testing.T) {
	testCreateUIDStateProof(t, true)
	testCreateUIDStateProof(t, false)
}

func testDeposit(t *testing.T, direct bool) {
	s := newTestService(t, 1)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
	defer cli.Close()

	deposit(t, s, cli, one)
}

func TestDeposit(t *testing.T) {
	testDeposit(t, true)
	testDeposit(t, false)
}

func timeMachine(t *testing.T, adjustment time.Duration,
	server backend.Backend) {
	if sim, ok := server.(backend.Simulator); ok {
		if err := sim.AdjustTime(adjustment); err != nil {
			t.Fatal(err)
		}
	}
}

func testWithdraw(t *testing.T, direct bool) {
	s := newTestService(t, 2)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
	defer cli.Close()

	uid := deposit(t, s, cli, one)

	cli2 := testClient(t, s, direct, s.accounts[1])
	defer cli2.Close()

	tx1 := testTx(t, zero, uid, one, zero, s.accounts[0].From, s.accounts[0])
	tx2 := testTx(t, one, uid, one, one, s.accounts[1].From, s.accounts[0])

	tx1Obj := addTx(t, uid, tx1, cli)
	tx2Obj := addTx(t, uid, tx2, cli2)

	startExitTx, err := cli2.StartExit(tx1Obj.rawTx, tx1Obj.proof,
		new(big.Int).SetUint64(tx1Obj.block), tx2Obj.rawTx,
		tx2Obj.proof, new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(startExitTx) {
		t.Fatal("failed to start exit")
	}
	timeMachine(t, time.Duration(504*time.Hour), s.backend)

	withdrawTx, err := cli2.Withdraw(tx1Obj.rawTx, tx1Obj.proof,
		new(big.Int).SetUint64(tx1Obj.block), tx2Obj.rawTx,
		tx2Obj.proof, new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(withdrawTx) {
		t.Fatal("failed to start exit")
	}
}

func TestWithdraw(t *testing.T) {
	testWithdraw(t, true)
	testWithdraw(t, false)
}

func testChallengeNonce(t *testing.T, direct bool) {
	s := newTestService(t, 4)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
	defer cli.Close()

	cli2 := testClient(t, s, direct, s.accounts[2])
	defer cli.Close()

	cli3 := testClient(t, s, direct, s.accounts[3])
	defer cli.Close()

	uid := deposit(t, s, cli, one)

	tx1 := testTx(t, zero, uid, one, zero, s.accounts[1].From, s.accounts[0])
	tx2 := testTx(t, one, uid, one, one, s.accounts[2].From, s.accounts[1])
	tx3 := testTx(t, two, uid, one, two, s.accounts[3].From, s.accounts[2])

	tx1Obj := addTx(t, uid, tx1, cli)
	tx2Obj := addTx(t, uid, tx2, cli)
	tx3Obj := addTx(t, uid, tx3, cli)

	startExitTx, err := cli2.StartExit(tx1Obj.rawTx, tx1Obj.proof,
		new(big.Int).SetUint64(tx1Obj.block), tx2Obj.rawTx,
		tx2Obj.proof, new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(startExitTx) {
		t.Fatal("failed to start exit")
	}

	challengeExitTx, err := cli3.ChallengeExit(uid, tx3Obj.rawTx, tx3Obj.proof,
		new(big.Int).SetUint64(tx3Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeExitTx) {
		t.Fatal("failed to start exit")
	}
}

func TestChallengeNonce(t *testing.T) {
	testChallengeNonce(t, true)
	testChallengeNonce(t, false)
}

func testChallengeDoubleSpending(t *testing.T, direct bool) {
	s := newTestService(t, 3)
	defer s.Close()

	owner := s.accounts[0]
	u1 := s.accounts[1]
	u2 := s.accounts[2]

	cli := testClient(t, s, direct, owner)
	defer cli.Close()

	cli1 := testClient(t, s, direct, u1)
	defer cli.Close()

	cli2 := testClient(t, s, direct, u2)
	defer cli.Close()

	uid := deposit(t, s, cli, one)

	tx1 := testTx(t, zero, uid, one, zero, owner.From, owner)
	tx2 := testTx(t, one, uid, one, one, u1.From, owner)
	tx3 := testTx(t, one, uid, one, one, u2.From, owner)

	tx1Obj := addTx(t, uid, tx1, cli)
	tx2Obj := addTx(t, uid, tx2, cli)
	tx3Obj := addTx(t, uid, tx3, cli)

	startExitTx, err := cli2.StartExit(tx1Obj.rawTx, tx1Obj.proof,
		new(big.Int).SetUint64(tx1Obj.block), tx3Obj.rawTx,
		tx3Obj.proof, new(big.Int).SetUint64(tx3Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(startExitTx) {
		t.Fatal("failed to start exit")
	}

	challengeExitTx, err := cli1.ChallengeExit(uid, tx2Obj.rawTx, tx2Obj.proof,
		new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeExitTx) {
		t.Fatal("failed to start exit")
	}
}

func TestChallengeDoubleSpending(t *testing.T) {
	testChallengeDoubleSpending(t, true)
	testChallengeDoubleSpending(t, false)
}

func testEarlyChallengeDoubleSpending(t *testing.T, direct bool) {
	s := newTestService(t, 3)
	defer s.Close()

	owner := s.accounts[0]
	u1 := s.accounts[1]
	u2 := s.accounts[2]

	cli := testClient(t, s, direct, owner)
	defer cli.Close()

	cli1 := testClient(t, s, direct, u1)
	defer cli.Close()

	cli2 := testClient(t, s, direct, u2)
	defer cli.Close()

	uid := deposit(t, s, cli, one)

	tx1 := testTx(t, zero, uid, one, zero, owner.From, owner)
	tx2 := testTx(t, one, uid, one, one, u1.From, owner)
	tx3 := testTx(t, one, uid, one, one, u2.From, owner)
	tx4 := testTx(t, three, uid, one, two, owner.From, u2)
	tx5 := testTx(t, four, uid, one, three, u2.From, owner)

	addTx(t, uid, tx1, cli)
	tx2Obj := addTx(t, uid, tx2, cli)
	addTx(t, uid, tx3, cli)
	tx4Obj := addTx(t, uid, tx4, cli)
	tx5Obj := addTx(t, uid, tx5, cli)

	startExitTx, err := cli2.StartExit(tx4Obj.rawTx, tx4Obj.proof,
		new(big.Int).SetUint64(tx4Obj.block), tx5Obj.rawTx,
		tx5Obj.proof, new(big.Int).SetUint64(tx5Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(startExitTx) {
		t.Fatal("failed to start exit")
	}

	challengeExitTx, err := cli1.ChallengeExit(uid, tx2Obj.rawTx, tx2Obj.proof,
		new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeExitTx) {
		t.Fatal("failed to start exit")
	}
}

func TestEarlyChallengeDoubleSpending(t *testing.T) {
	testEarlyChallengeDoubleSpending(t, true)
	testEarlyChallengeDoubleSpending(t, false)
}

func testRespondToChallenge(t *testing.T, direct bool) {
	s := newTestService(t, 3)
	defer s.Close()

	owner := s.accounts[0]
	u1 := s.accounts[1]
	u2 := s.accounts[2]

	cli := testClient(t, s, direct, owner)
	defer cli.Close()

	cli1 := testClient(t, s, direct, u1)
	defer cli.Close()

	cli2 := testClient(t, s, direct, u2)
	defer cli.Close()

	uid := deposit(t, s, cli, one)

	tx1 := testTx(t, zero, uid, one, zero, owner.From, owner)
	tx2 := testTx(t, one, uid, one, one, u1.From, owner)
	tx3 := testTx(t, two, uid, one, two, u2.From, u1)
	tx4 := testTx(t, three, uid, one, three, owner.From, u2)
	tx5 := testTx(t, four, uid, one, four, u2.From, owner)

	addTx(t, uid, tx1, cli)
	tx2Obj := addTx(t, uid, tx2, cli)
	tx3Obj := addTx(t, uid, tx3, cli)
	tx4Obj := addTx(t, uid, tx4, cli)
	tx5Obj := addTx(t, uid, tx5, cli)

	startExitTx, err := cli2.StartExit(tx4Obj.rawTx, tx4Obj.proof,
		new(big.Int).SetUint64(tx4Obj.block), tx5Obj.rawTx,
		tx5Obj.proof, new(big.Int).SetUint64(tx5Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(startExitTx) {
		t.Fatal("failed to start exit")
	}

	challengeExitTx, err := cli1.ChallengeExit(uid, tx2Obj.rawTx, tx2Obj.proof,
		new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeExitTx) {
		t.Fatal("failed to start exit")
	}

	respondChallengeExitTx, err := cli2.RespondChallengeExit(uid, tx2Obj.rawTx,
		tx3Obj.rawTx, tx3Obj.proof, new(big.Int).SetUint64(tx3Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(respondChallengeExitTx) {
		t.Fatal("failed to start exit")
	}
}

func TestRespondToChallenge(t *testing.T) {
	testRespondToChallenge(t, true)
	testRespondToChallenge(t, false)
}

func TestCheckpointChallenge(t *testing.T) {
	s := newTestService(t, 3)
	defer s.Close()

	owner := s.accounts[0]
	u1 := s.accounts[1]
	u2 := s.accounts[2]

	cli := testClient(t, s, true, owner)
	defer cli.Close()

	uid := deposit(t, s, cli, one)

	tx1 := testTx(t, zero, uid, one, zero, owner.From, owner)
	tx2 := testTx(t, one, uid, one, one, u1.From, owner)
	tx3 := testTx(t, two, uid, one, two, u2.From, u1)

	addTx(t, uid, tx1, cli)
	tx2Obj := addTx(t, uid, tx2, cli)
	tx3Obj := addTx(t, uid, tx3, cli)

	addCheckpointResp, err := cli.AddCheckpoint(uid, three)
	if err != nil {
		t.Fatal(err)
	}

	if addCheckpointResp.Error != "" {
		t.Fatal(addCheckpointResp.Error)
	}

	buildCheckpointResp, err := cli.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	if buildCheckpointResp.Error != "" {
		t.Fatal(buildCheckpointResp.Error)
	}

	sendCheckpointHashResp, err := cli.SendCheckpointHash(
		buildCheckpointResp.Hash)
	if err != nil {
		t.Fatal(err)
	}

	if sendCheckpointHashResp.Error != "" {
		t.Fatal(sendCheckpointHashResp.Error)
	}

	currentCheckpointResp, err := cli.CurrentCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	if currentCheckpointResp.Error != "" {
		t.Fatal(currentCheckpointResp.Error)
	}

	saveCheckpointToDBResp, err := cli.SaveCheckpointToDB(
		currentCheckpointResp.Checkpoint)
	if err != nil {
		t.Fatal(err)
	}

	if saveCheckpointToDBResp.Error != "" {
		t.Fatal(saveCheckpointToDBResp.Error)
	}

	createUIDStateProofResp, err := cli.CreateUIDStateProof(
		uid, buildCheckpointResp.Hash)
	if err != nil {
		t.Fatal(err)
	}

	if createUIDStateProofResp.Error != "" {
		t.Fatal(createUIDStateProofResp.Error)
	}

	resp, err := cli.VerifyCheckpointProof(uid, three,
		buildCheckpointResp.Hash,
		createUIDStateProofResp.Proof)
	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.Error != "" {
		t.Fatal(resp.Error)
	}

	if !resp.Exists {
		t.Fatal("uid not exist in a checkpoint")
	}

	challengeCheckpointTx, err := cli.ChallengeCheckpoint(uid,
		buildCheckpointResp.Hash, createUIDStateProofResp.Proof,
		three, tx2Obj.rawTx, tx2Obj.proof,
		new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeCheckpointTx) {
		t.Fatal("failed to start exit")
	}

	respTx, err := cli.RespondCheckpointChallenge(uid,
		buildCheckpointResp.Hash, tx2Obj.rawTx,
		tx3Obj.rawTx, tx3Obj.proof,
		new(big.Int).SetUint64(tx3Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(respTx) {
		t.Fatal("failed to start exit")
	}

	initCheckpointResp, err := cli.InitCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	if initCheckpointResp.Error != "" {
		t.Fatal(initCheckpointResp.Error)
	}
}

func TestRespondWithHistoricalCheckpoint(t *testing.T) {
	s := newTestService(t, 3)
	defer s.Close()

	owner := s.accounts[0]
	u1 := s.accounts[1]
	u2 := s.accounts[2]

	cli := testClient(t, s, true, owner)
	defer cli.Close()

	uid := deposit(t, s, cli, one)

	tx1 := testTx(t, zero, uid, one, zero, u1.From, owner)
	tx2 := testTx(t, one, uid, one, one, u2.From, owner)

	addTx(t, uid, tx1, cli)
	tx2Obj := addTx(t, uid, tx2, cli)

	addCheckpointResp, err := cli.AddCheckpoint(uid, two)
	if err != nil {
		t.Fatal(err)
	}

	if addCheckpointResp.Error != "" {
		t.Fatal(addCheckpointResp.Error)
	}

	buildCheckpointResp, err := cli.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	if buildCheckpointResp.Error != "" {
		t.Fatal(buildCheckpointResp.Error)
	}

	sendCheckpointHashResp, err := cli.SendCheckpointHash(
		buildCheckpointResp.Hash)
	if err != nil {
		t.Fatal(err)
	}

	if sendCheckpointHashResp.Error != "" {
		t.Fatal(sendCheckpointHashResp.Error)
	}

	currentCheckpointResp, err := cli.CurrentCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	if currentCheckpointResp.Error != "" {
		t.Fatal(currentCheckpointResp.Error)
	}

	saveCheckpointToDBResp, err := cli.SaveCheckpointToDB(
		currentCheckpointResp.Checkpoint)
	if err != nil {
		t.Fatal(err)
	}

	if saveCheckpointToDBResp.Error != "" {
		t.Fatal(saveCheckpointToDBResp.Error)
	}

	createUIDStateProofResp, err := cli.CreateUIDStateProof(
		uid, buildCheckpointResp.Hash)
	if err != nil {
		t.Fatal(err)
	}

	if createUIDStateProofResp.Error != "" {
		t.Fatal(createUIDStateProofResp.Error)
	}

	verifyCheckpointProofResp, err := cli.VerifyCheckpointProof(uid,
		two, buildCheckpointResp.Hash,
		createUIDStateProofResp.Proof)
	if err != nil {
		t.Fatal(err.Error())
	}

	if verifyCheckpointProofResp.Error != "" {
		t.Fatal(verifyCheckpointProofResp.Error)
	}

	if !verifyCheckpointProofResp.Exists {
		t.Fatal("uid not exist in a checkpoint")
	}

	// + 3 weeks
	timeMachine(t, time.Duration(504*time.Hour), s.backend)

	cli.InitCheckpoint()

	cli.AddCheckpoint(uid, three)

	buildCheckpointResp2, err := cli.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	if buildCheckpointResp2.Error != "" {
		t.Fatal(buildCheckpointResp2.Error)
	}

	sendCheckpointHashResp, err = cli.SendCheckpointHash(
		buildCheckpointResp2.Hash)
	if err != nil {
		t.Fatal(err)
	}

	if sendCheckpointHashResp.Error != "" {
		t.Fatal(sendCheckpointHashResp.Error)
	}

	currentCheckpointResp2, err := cli.CurrentCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	if currentCheckpointResp2.Error != "" {
		t.Fatal(currentCheckpointResp2.Error)
	}

	saveCheckpointToDBResp, err = cli.SaveCheckpointToDB(
		currentCheckpointResp2.Checkpoint)
	if err != nil {
		t.Fatal(err)
	}

	if saveCheckpointToDBResp.Error != "" {
		t.Fatal(saveCheckpointToDBResp.Error)
	}

	createUIDStateProofResp2, err := cli.CreateUIDStateProof(
		uid, buildCheckpointResp2.Hash)
	if err != nil {
		t.Fatal(err)
	}

	if createUIDStateProofResp2.Error != "" {
		t.Fatal(createUIDStateProofResp2.Error)
	}

	challengeCheckpointTx, err := cli.ChallengeCheckpoint(uid,
		buildCheckpointResp2.Hash, createUIDStateProofResp2.Proof,
		three, tx2Obj.rawTx, tx2Obj.proof,
		new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeCheckpointTx) {
		t.Fatal("failed to start exit")
	}

	respTx, err := cli.RespondWithHistoricalCheckpoint(uid,
		buildCheckpointResp2.Hash,
		createUIDStateProofResp2.Proof, buildCheckpointResp.Hash,
		createUIDStateProofResp.Proof, tx2Obj.rawTx,
		new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(respTx) {
		t.Fatal("failed to start exit")
	}
}
