package transport

import (
	"bytes"
	"context"
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
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/erc20token"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database"
	"github.com/SmartMeshFoundation/SmartPlasma/database/bolt"
	"github.com/SmartMeshFoundation/SmartPlasma/service"
	"github.com/SmartMeshFoundation/SmartPlasma/transport/handlers"
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
	smartPlasma      *handlers.SmartPlasma
	blockBase        database.Database
	chptBase         database.Database
	backend          backend.Backend
	rootChainAddress common.Address
	mediatorAddress  common.Address
	service          *service.Service
}

type txData struct {
	rawTx []byte
	proof []byte
	block uint64
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

	s := service.NewService(session, server, blockDB, chptDB, rchc, mc, false)

	smartPlasma := handlers.NewSmartPlasma(100, s)

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
		service:          s,
	}
}

func (s *testService) Close() {
	os.RemoveAll(s.dir)
	s.service.Close()
	s.server.Close()
}

func testClient(t *testing.T, srv *testService, direct bool,
	user *account.PlasmaTransactOpts) *Client {
	cli := NewClient(100, user)
	err := cli.ConnectString(srv.server.URL[7:])
	if err != nil {
		t.Fatal(err)
	}

	if direct {
		cli.DirectEthereumClient(
			*user.TransactOpts, srv.mediatorAddress,
			srv.rootChainAddress, srv.backend)
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

	_, err = cli.ChallengePeriod()
	if err != nil {
		t.Fatal(err)
	}

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

	tr, err := cli.WaitMined(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	if tr.Status != 1 {
		t.Fatal("transaction is failed")
	}

	amount2, err := cli.Wallet(uid)
	if err != nil {
		t.Fatal(err)
	}

	if amount2.Uint64() != amount.Uint64() {
		t.Fatal("amount is wrong")
	}

	count, err := cli.DepositCount()
	if err != nil {
		t.Fatal(err)
	}

	if count.Uint64() != 1 {
		t.Fatal("wrong deposit count")

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
	goodTxs, badTxs []*transaction.Transaction, cli *Client, validate bool) (map[string]*txData, common.Hash) {
	result := make(map[string]*txData)

	for _, tx := range goodTxs {
		buf := bytes.NewBuffer([]byte{})

		err := tx.EncodeRLP(buf)
		if err != nil {
			t.Fatal(err)
		}

		err = cli.AcceptTransaction(buf.Bytes())
		if err != nil {
			t.Fatal(err)
		}

		data := &txData{
			rawTx: buf.Bytes(),
		}

		result[tx.UID().String()] = data
	}

	if validate {
		for _, tx := range badTxs {
			buf := bytes.NewBuffer([]byte{})

			err := tx.EncodeRLP(buf)
			if err != nil {
				t.Fatal(err)
			}

			err = cli.AcceptTransaction(buf.Bytes())
			if err != nil {
				t.Fatal(err)
			}
		}

		err := cli.ValidateBlock()
		if err != nil {
			t.Fatal(err)
		}

		curBlock, err := cli.CurrentBlock()
		if err != nil {
			t.Fatal(err)
		}

		testBlock := transactions.NewBlock()
		err = testBlock.Unmarshal(curBlock)
		if err != nil {
			t.Fatal(err)
		}

		if testBlock.NumberOfTX() == 0 && len(goodTxs) == 0 {
			return nil, common.Hash{}
		}
	}

	buildResp, err := cli.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}

	sendBlock1Tx, err := cli.SendBlockHash(buildResp)
	if err != nil {
		t.Fatal(err)
	}

	sendBlock1Tr, err := cli.WaitMined(context.Background(), sendBlock1Tx)
	if err != nil {
		t.Fatal(err)
	}

	if sendBlock1Tr.Status != 1 {
		t.Fatal("wrong tx status")
	}

	lastBlock, err := cli.LastBlockNumber()
	if err != nil {
		t.Fatal(err)
	}

	err = cli.SaveCurrentBlock(lastBlock.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	bl, err := cli.GetTransactionsBlock(lastBlock.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if (bl.Hash() == common.Hash{}) && len(goodTxs) != 0 {
		t.Fatal("hash is empty")
	}

	err = cli.InitBlock()
	if err != nil {
		t.Fatal(err)
	}

	for _, tx := range goodTxs {
		proof, err := cli.CreateProof(tx.UID(), lastBlock.Uint64())
		if err != nil {
			t.Fatal(err)
		}

		respProof, err := cli.VerifyTxProof(tx.UID(), tx.Hash(),
			lastBlock.Uint64(), proof)
		if err != nil {
			t.Fatal(err)
		}

		if !respProof {
			t.Fatal("not exists")
		}

		result[tx.UID().String()].proof = proof
		result[tx.UID().String()].block = lastBlock.Uint64()
	}

	for _, tx := range badTxs {
		proof, err := cli.CreateProof(tx.UID(), lastBlock.Uint64())
		if err != nil {
			t.Fatal(err)
		}

		respProof, err := cli.VerifyTxProof(tx.UID(), tx.Hash(),
			lastBlock.Uint64(), proof)
		if err != nil {
			t.Fatal(err)
		}

		if respProof {
			t.Fatal("bad transaction exists in the block")
		}
	}

	return result, buildResp
}

func testAcceptTransaction(t *testing.T, direct bool) {
	s := newTestService(t, 1)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
	defer cli.Close()

	operator, err := cli.Operator()
	if err != nil {
		t.Fatal(err)
	}

	if operator.String() != s.accounts[0].From.String() {
		t.Fatal("wrong operator")
	}

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

	err = cli.AcceptTransaction(buf.Bytes())
	if err != nil {
		t.Fatal(err)
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

	err = s.service.AcceptTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.service.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}

	curBlock := s.service.CurrentBlock()

	err = s.service.SaveBlockToDB(one.Uint64(), curBlock)
	if err != nil {
		t.Fatal(err)
	}

	proof, err := cli.CreateProof(one, one.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	if len(proof) == 0 {
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

	tx1 := testTx(t, zero, one, one, four, s.accounts[0].From, s.accounts[0])
	objects1, _ := addTx(t,
		one, []*transaction.Transaction{tx1}, nil, cli, false)
	tx1Obj := objects1[tx1.UID().String()]

	err := cli.AddCheckpoint(one, four, tx1Obj.block)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddCheckpoint(t *testing.T) {
	testAddCheckpoint(t, true)
	testAddCheckpoint(t, false)
}

func testCreateUIDStateProof(t *testing.T, direct bool) {
	s := newTestService(t, 2)
	defer s.Close()

	cli := testClient(t, s, direct, s.accounts[0])
	defer cli.Close()

	uid := one
	badUID := new(big.Int).Add(uid, one)
	nonce := four
	badNonce := new(big.Int).Add(nonce, one)
	blockNum := three
	badBlockNumber := new(big.Int).Add(blockNum, one)

	tx := testTx(t, one, uid, one, nonce, s.accounts[1].From, s.accounts[0])

	blk := transactions.NewBlock()

	err := blk.AddTx(tx)
	if err != nil {
		t.Fatal(err)
	}

	rawBlock, err := blk.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	err = cli.SaveBlockToDB(blockNum.Uint64(), rawBlock)
	if err != nil {
		t.Fatal(err)
	}

	err = s.service.AcceptUIDState(badUID, nonce, blockNum.Uint64())
	if err == nil {
		t.Fatal("checkpoint should not be accepted")
	}

	err = s.service.AcceptUIDState(uid, badNonce, blockNum.Uint64())
	if err == nil {
		t.Fatal("checkpoint should not be accepted")
	}

	err = s.service.AcceptUIDState(uid, nonce, badBlockNumber.Uint64())
	if err == nil {
		t.Fatal("checkpoint should not be accepted")
	}

	err = s.service.AcceptUIDState(uid, nonce, blockNum.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.service.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	curChpt := s.service.CurrentCheckpoint()

	hash := s.service.CurrentCheckpoint().Hash()

	err = s.service.SaveCheckpointToDB(curChpt)
	if err != nil {
		t.Fatal(err)
	}

	proof, expected, err := cli.CreateUIDStateProof(one, hash)
	if err != nil {
		t.Fatal(err)
	}

	if expected.Uint64() != nonce.Uint64() {
		t.Fatal("wrong nonce")
	}

	if len(proof) == 0 {
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

	objects0, block0Hash := addTx(t, uid, []*transaction.Transaction{tx1}, nil, cli, false)
	objects1, _ := addTx(t, uid, []*transaction.Transaction{tx2}, nil, cli, false)

	tx1Obj := objects0[tx1.UID().String()]
	tx2Obj := objects1[tx2.UID().String()]

	startExitTx, err := cli2.StartExit(tx1Obj.rawTx, tx1Obj.proof,
		new(big.Int).SetUint64(tx1Obj.block), tx2Obj.rawTx,
		tx2Obj.proof, new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(startExitTx) {
		t.Fatal("failed to start exit")
	}

	exitsResp, err := cli.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if exitsResp.State.Uint64() != 2 {
		t.Fatal("wrong exit state")
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

	block1Hash, err := cli.ChildChain(big.NewInt(int64(tx1Obj.block)))
	if err != nil {
		t.Fatal(err)
	}

	if block1Hash.String() != block0Hash.String() {
		t.Fatal("wrong child chain block")
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

	objects1, _ := addTx(t, uid, []*transaction.Transaction{tx1}, nil, cli, false)
	objects2, _ := addTx(t, uid, []*transaction.Transaction{tx2}, nil, cli, false)
	objects3, _ := addTx(t, uid, []*transaction.Transaction{tx3}, nil, cli, false)

	tx1Obj := objects1[tx1.UID().String()]
	tx2Obj := objects2[tx2.UID().String()]
	tx3Obj := objects3[tx3.UID().String()]

	startExitTx, err := cli2.StartExit(tx1Obj.rawTx, tx1Obj.proof,
		new(big.Int).SetUint64(tx1Obj.block), tx2Obj.rawTx,
		tx2Obj.proof, new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli3.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if resp.State.Uint64() != 2 {
		t.Fatal("exit not exist")
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

	resp, err = cli3.Exits(uid)
	if err != nil {
		t.Fatal(err)
	}

	if resp.State.Uint64() != 0 {
		t.Fatal("exit is exist")
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

	objects1, _ := addTx(t, uid, []*transaction.Transaction{tx1}, nil, cli, false)
	objects2, _ := addTx(t, uid, []*transaction.Transaction{tx2}, nil, cli, false)
	objects3, _ := addTx(t, uid, []*transaction.Transaction{tx3}, nil, cli, false)

	tx1Obj := objects1[tx1.UID().String()]
	tx2Obj := objects2[tx2.UID().String()]
	tx3Obj := objects3[tx3.UID().String()]

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

	addTx(t, uid, []*transaction.Transaction{tx1}, nil, cli, false)
	objects2, _ := addTx(t, uid, []*transaction.Transaction{tx2}, nil, cli, false)
	addTx(t, uid, []*transaction.Transaction{tx3}, nil, cli, false)
	objects4, _ := addTx(t, uid, []*transaction.Transaction{tx4}, nil, cli, false)
	objects5, _ := addTx(t, uid, []*transaction.Transaction{tx5}, nil, cli, false)

	tx2Obj := objects2[tx2.UID().String()]
	tx4Obj := objects4[tx4.UID().String()]
	tx5Obj := objects5[tx5.UID().String()]

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

	exist, err := cli.ChallengeExists(uid, tx2Obj.rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if !exist {
		t.Fatal("challenge not exists")
	}

	length, err := cli.ChallengesLength(uid)
	if err != nil {
		t.Fatal(err)
	}

	if length.Uint64() != 1 {
		t.Fatal("challenges length is null")
	}

	challenge1, err := cli.GetChallenge(uid, zero)
	if err != nil {
		t.Fatal(err)
	}

	if challenge1.ChallengeBlock.Uint64() != tx2Obj.block {
		t.Fatal("wrong challenge block")
	}

	if !bytes.Equal(challenge1.ChallengeTx, tx2Obj.rawTx) {
		t.Fatal("wrong challenge tx")
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

	addTx(t, uid, []*transaction.Transaction{tx1}, nil, cli, false)
	objects2, _ := addTx(t, uid, []*transaction.Transaction{tx2}, nil, cli, false)
	objects3, _ := addTx(t, uid, []*transaction.Transaction{tx3}, nil, cli, false)
	objects4, _ := addTx(t, uid, []*transaction.Transaction{tx4}, nil, cli, false)
	objects5, _ := addTx(t, uid, []*transaction.Transaction{tx5}, nil, cli, false)

	tx2Obj := objects2[tx2.UID().String()]
	tx3Obj := objects3[tx3.UID().String()]
	tx4Obj := objects4[tx4.UID().String()]
	tx5Obj := objects5[tx5.UID().String()]

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

	exist, err := cli.ChallengeExists(uid, tx2Obj.rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if exist {
		t.Fatal("challenge is exists")
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

	addTx(t, uid, []*transaction.Transaction{tx1}, nil, cli, false)
	objects2, _ := addTx(t, uid, []*transaction.Transaction{tx2}, nil, cli, false)
	objects3, _ := addTx(t, uid, []*transaction.Transaction{tx3}, nil, cli, false)

	tx2Obj := objects2[tx2.UID().String()]
	tx3Obj := objects3[tx3.UID().String()]

	badNonce := three
	badBlockNum := four

	badTx := testTx(t, one, uid, one, badNonce, s.accounts[1].From, s.accounts[0])

	blk := transactions.NewBlock()

	err := blk.AddTx(badTx)
	if err != nil {
		t.Fatal(err)
	}

	rawBlock, err := blk.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	err = cli.SaveBlockToDB(badBlockNum.Uint64(), rawBlock)
	if err != nil {
		t.Fatal(err)
	}

	err = cli.AddCheckpoint(uid, badNonce, badBlockNum.Uint64())
	if err != nil {
		t.Fatal(err)
	}

	buildCheckpointResp, err := cli.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	sendCheckpointHashTx, err := cli.SendCheckpointHash(
		buildCheckpointResp)
	if err != nil {
		t.Fatal(err)
	}

	sendCheckpointHashTr, err := cli.WaitMined(
		context.Background(), sendCheckpointHashTx)
	if err != nil {
		t.Fatal(err)
	}

	if sendCheckpointHashTr.Status != 1 {
		t.Fatal("wrong transaction status")
	}

	err = cli.SaveCurrentCheckpointBlock()
	if err != nil {
		t.Fatal(err)
	}

	createUIDStateProof, nonce, err := cli.CreateUIDStateProof(
		uid, buildCheckpointResp)
	if err != nil {
		t.Fatal(err)
	}

	if nonce.Uint64() != three.Uint64() {
		t.Fatal("wrong nonce")
	}

	resp, err := cli.VerifyCheckpointProof(
		uid, three, buildCheckpointResp, createUIDStateProof)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !resp {
		t.Fatal("uid not exist in a checkpoint")
	}

	challengeCheckpointTx, err := cli.ChallengeCheckpoint(
		uid, buildCheckpointResp, createUIDStateProof, three,
		tx2Obj.rawTx, tx2Obj.proof, new(big.Int).SetUint64(tx2Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeCheckpointTx) {
		t.Fatal("failed to start exit")
	}

	exist, err := cli.CheckpointIsChallenge(
		uid, buildCheckpointResp, tx2Obj.rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if !exist {
		t.Fatal("checkpoint not challenges")
	}

	length, err := cli.CheckpointChallengesLength(
		uid, buildCheckpointResp)
	if err != nil {
		t.Fatal(err)
	}

	if length.Uint64() != 1 {
		t.Fatal("challenges length not equal 1")
	}

	challenge1, err := cli.GetCheckpointChallenge(
		uid, buildCheckpointResp, zero)
	if err != nil {
		t.Fatal(err)
	}

	if challenge1.ChallengeBlock.Uint64() != tx2Obj.block {
		t.Fatal("wrong challenge block")
	}

	if !bytes.Equal(challenge1.ChallengeTx, tx2Obj.rawTx) {
		t.Fatal("wrong challenge tx")
	}

	respTx, err := cli.RespondCheckpointChallenge(uid,
		buildCheckpointResp, tx2Obj.rawTx,
		tx3Obj.rawTx, tx3Obj.proof,
		new(big.Int).SetUint64(tx3Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(respTx) {
		t.Fatal("failed to start exit")
	}

	err = cli.InitCheckpoint()
	if err != nil {
		t.Fatal(err)
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

	tx1 := testTx(t, zero, uid, one, one, u1.From, owner)
	tx2 := testTx(t, one, uid, one, two, u2.From, owner)
	tx3 := testTx(t, two, uid, one, three, u2.From, owner)

	objects1, _ := addTx(t,
		uid, []*transaction.Transaction{tx1}, nil, cli, false)
	objects2, _ := addTx(t,
		uid, []*transaction.Transaction{tx2}, nil, cli, false)
	objects3, _ := addTx(t,
		uid, []*transaction.Transaction{tx3}, nil, cli, false)

	tx1Obj := objects1[tx1.UID().String()]
	tx2Obj := objects2[tx2.UID().String()]
	tx3Obj := objects3[tx3.UID().String()]

	err := cli.AddCheckpoint(uid, tx2.Nonce(), tx2Obj.block)
	if err != nil {
		t.Fatal(err)
	}

	buildCheckpointResp, err := cli.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	sendCheckpointHashTx, err := cli.SendCheckpointHash(
		buildCheckpointResp)
	if err != nil {
		t.Fatal(err)
	}

	sendCheckpointHashTr, err := cli.WaitMined(
		context.Background(), sendCheckpointHashTx)
	if err != nil {
		t.Fatal(err)
	}

	if sendCheckpointHashTr.Status != 1 {
		t.Fatal("wrong transaction status")
	}

	currentCheckpointResp, err := cli.CurrentCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	err = cli.SaveCheckpointToDB(currentCheckpointResp)
	if err != nil {
		t.Fatal(err)
	}

	chptbl, err := cli.GetCheckpointsBlock(buildCheckpointResp)
	if err != nil {
		t.Fatal(err)
	}

	if chptbl.Hash().String() != buildCheckpointResp.String() {
		t.Fatal("wrong checkpoint hash")
	}

	createUIDStateProof, nonce, err := cli.CreateUIDStateProof(
		uid, buildCheckpointResp)
	if err != nil {
		t.Fatal(err)
	}

	if nonce.Uint64() != two.Uint64() {
		t.Fatal("wrong nonce")
	}

	verifyCheckpointProofResp, err := cli.VerifyCheckpointProof(
		uid, two, buildCheckpointResp, createUIDStateProof)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !verifyCheckpointProofResp {
		t.Fatal("uid not exist in a checkpoint")
	}

	// + 3 weeks
	timeMachine(t, time.Duration(504*time.Hour), s.backend)

	err = cli.InitCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	err = cli.AddCheckpoint(uid, tx3.Nonce(), tx3Obj.block)
	if err != nil {
		t.Fatal(err)
	}

	buildCheckpointResp2, err := cli.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	sendCheckpointHashTx2, err := cli.SendCheckpointHash(
		buildCheckpointResp2)
	if err != nil {
		t.Fatal(err)
	}

	sendCheckpointHashTr2, err := cli.WaitMined(
		context.Background(), sendCheckpointHashTx2)
	if err != nil {
		t.Fatal(err)
	}

	if sendCheckpointHashTr2.Status != 1 {
		t.Fatal("wrong transaction status")
	}

	currentCheckpointResp2, err := cli.CurrentCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	err = cli.SaveCheckpointToDB(currentCheckpointResp2)
	if err != nil {
		t.Fatal(err)
	}

	createUIDStateProof2, nonce, err := cli.CreateUIDStateProof(
		uid, buildCheckpointResp2)
	if err != nil {
		t.Fatal(err)
	}

	if nonce.Uint64() != three.Uint64() {
		t.Fatal("wrong nonce")
	}

	challengeCheckpointTx, err := cli.ChallengeCheckpoint(
		uid, buildCheckpointResp2, createUIDStateProof2, three,
		tx1Obj.rawTx, tx1Obj.proof, new(big.Int).SetUint64(tx1Obj.block))
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(challengeCheckpointTx) {
		t.Fatal("failed to start exit")
	}

	respTx, err := cli.RespondWithHistoricalCheckpoint(
		uid, buildCheckpointResp2, createUIDStateProof2,
		buildCheckpointResp, createUIDStateProof, tx1Obj.rawTx, tx2.Nonce())
	if err != nil {
		t.Fatal(err)
	}

	if !s.backend.GoodTransaction(respTx) {
		t.Fatal("failed to start exit")
	}
}

func TestValidateBlock(t *testing.T) {
	s := newTestService(t, 3)
	defer s.Close()

	cli0 := testClient(t, s, true, s.accounts[0])
	defer cli0.Close()

	cli1 := testClient(t, s, true, s.accounts[0])
	defer cli1.Close()

	uid := deposit(t, s, cli0, one)
	badUID := big.NewInt(1000)

	validTx1 := testTx(t, zero, uid, one, zero, s.accounts[0].From, s.accounts[0])
	validTx2 := testTx(t, one, uid, one, one, s.accounts[1].From, s.accounts[0])
	validTx3 := testTx(t, two, uid, one, two, s.accounts[2].From, s.accounts[1])

	//bad new owner
	badTx0 := testTx(t, zero, uid, one, zero, s.accounts[1].From, s.accounts[0])

	// no deposit
	badTx1 := testTx(t, zero, badUID, one, zero, s.accounts[0].From, s.accounts[0])

	// bad nonce
	badTx2 := testTx(t, zero, uid, one, zero, s.accounts[0].From, s.accounts[0])

	// bad nonce
	badTx3 := testTx(t, one, uid, one, two, s.accounts[1].From, s.accounts[0])

	// bad amount
	badTx4 := testTx(t, one, uid, two, one, s.accounts[1].From, s.accounts[0])

	// bad previous block
	badTx5 := testTx(t, two, uid, one, one, s.accounts[1].From, s.accounts[0])

	// bad new owner
	badTx6 := testTx(t, one, uid, one, one, s.accounts[0].From, s.accounts[0])

	addTx(t, uid, nil, []*transaction.Transaction{badTx0}, cli0, true)

	addTx(t, uid, []*transaction.Transaction{validTx1},
		[]*transaction.Transaction{badTx1}, cli0, true)

	addTx(t, uid, nil, []*transaction.Transaction{badTx2}, cli0, true)
	addTx(t, uid, nil, []*transaction.Transaction{badTx3}, cli0, true)
	addTx(t, uid, nil, []*transaction.Transaction{badTx4}, cli0, true)
	addTx(t, uid, nil, []*transaction.Transaction{badTx5}, cli0, true)
	addTx(t, uid, nil, []*transaction.Transaction{badTx6}, cli0, true)

	addTx(t, uid, []*transaction.Transaction{validTx2}, nil, cli0, true)
	addTx(t, uid, []*transaction.Transaction{validTx3}, nil, cli1, true)
}
