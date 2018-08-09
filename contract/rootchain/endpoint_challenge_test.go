package rootchain

// Change addCheckpointChallenge and removeCheckpointChallenge from private to public.
/*import (
	"math/big"
	"os"
	"testing"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartmeshfoundation/smartplasma/blockchan/account"
	"github.com/smartmeshfoundation/smartplasma/blockchan/backend"
	"github.com/smartmeshfoundation/smartplasma/contract/mediator"
	"bytes"
)

const (
	rawTx1data = "0x1"
	rawTx2data = "0x2"
	rawTx3data = "0x3"
	rawTx4data = "0x4"
	rawTx5data = "0x5"
)

var (
	uid = big.NewInt(123)

	server backend.Backend

	owner *account.PlasmaTransactOpts

	challengeData = []*plasmaChallengeData{
		{
			challengeTx: []byte(rawTx1data),
			blockNumber: big.NewInt(101),
		},
		{
			challengeTx: []byte(rawTx2data),
			blockNumber: big.NewInt(102),
		},
		{
			challengeTx: []byte(rawTx3data),
			blockNumber: big.NewInt(103),
		},
		{
			challengeTx: []byte(rawTx4data),
			blockNumber: big.NewInt(104),
		},
		{
			challengeTx: []byte(rawTx5data),
			blockNumber: big.NewInt(105),
		},
	}
)

type plasmaChallengeData struct {
	challengeTx []byte
	blockNumber *big.Int
}

type instance struct {
	rootOwnerSession *RootChainSession
	rootChainAddr    common.Address
}

func challengeGenerator(number int) map[int]*plasmaChallengeData {
	result := make(map[int]*plasmaChallengeData)
	for i := 0; i < number; i++ {
		key, _ := crypto.GenerateKey()
		item := &plasmaChallengeData{
			challengeTx: key.Y.Bytes(),
			blockNumber: big.NewInt(1),
		}
		result[i] = item
	}

	return result
}

func hashGenerator() common.Hash{
	key, _ := crypto.GenerateKey()
	return common.BigToHash(key.D)
}

func newInstance(t *testing.T) *instance {
	i := &instance{}

	_, med, err := deployMediator(owner.TransactOpts)
	if err != nil {
		t.Fatal(err)
	}

	i.rootChainAddr, err = med.RootChain(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}

	i.rootOwnerSession = rootChainSession(t, owner.TransactOpts,
		i.rootChainAddr)

	return i
}

func rootChainSession(t *testing.T, account *bind.TransactOpts,
	contact common.Address) (session *RootChainSession) {
	session, err := NewRootChainSession(*account,
		contact, server)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func deployMediator(account *bind.TransactOpts) (address common.Address,
	contract *mediator.Mediator, err error) {
	address, contract, err = mediator.Deploy(account, server)
	return
}

func TestMain(m *testing.M) {
	accounts := account.GenAccounts(3)
	owner = accounts[0]

	server = backend.NewSimulatedBackend(account.Addresses(accounts))

	os.Exit(m.Run())
}

func TestAddCheckpointChallenge(t *testing.T) {
	i := newInstance(t)

	checkpointHash := hashGenerator()

	for index, data := range challengeData {
		tx, err := i.rootOwnerSession.AddCheckpointChallenge(uid,
			checkpointHash, data.challengeTx, data.blockNumber)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to add challenge")
		}

		str, err := i.rootOwnerSession.GetCheckpointChallenge(uid,
			checkpointHash, big.NewInt(int64(index)))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(str.ChallengeTx, data.challengeTx) ||
			!bytes.Equal(str.ChallengeBlock.Bytes(),
				data.blockNumber.Bytes()) {
			t.Fatal("challenge is wrong")
		}

		exists, err := i.rootOwnerSession.CheckpointChallengeExists(uid,
			checkpointHash, data.challengeTx)
		if err != nil {
			t.Fatal(err)
		}

		if !exists {
			t.Fatal("challenge not found")
		}
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid,
		checkpointHash)
	if err != nil {
		t.Fatal(err)
	}
	if length.Int64() != int64(len(challengeData)) {
		t.Fatal("disputes length is wrong")
	}
}

func TestRemoveCheckpointChallengeFromMiddle(t *testing.T) {
	i := newInstance(t)

	checkpointHash := hashGenerator()

	var index int64 = 2

	for index, data := range challengeData {
		tx, err := i.rootOwnerSession.AddCheckpointChallenge(uid,
			checkpointHash, data.challengeTx, data.blockNumber)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to add challenge")
		}

		str, err := i.rootOwnerSession.GetCheckpointChallenge(uid,
			checkpointHash, big.NewInt(int64(index)))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(str.ChallengeTx, data.challengeTx) ||
			!bytes.Equal(str.ChallengeBlock.Bytes(),
				data.blockNumber.Bytes()) {
			t.Fatal("challenge is wrong")
		}
	}

	tx, err := i.rootOwnerSession.RemoveCheckpointChallenge(uid,
		checkpointHash, challengeData[index].challengeTx)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to remove challenge")
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid,
		checkpointHash)
	if err != nil {
		t.Fatal(err)
	}
	if length.Int64() != int64(len(challengeData)-1) {
		t.Fatal("disputes length is wrong")
	}

	str, err := i.rootOwnerSession.GetCheckpointChallenge(uid,
		checkpointHash, big.NewInt(index))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(str.ChallengeTx,
		challengeData[len(challengeData)-1].challengeTx) {
		t.Fatal("incorrect removal from the middle")
	}
}

func TestRemoveChallengeFromStart(t *testing.T) {
	i := newInstance(t)

	checkpointHash := hashGenerator()

	var index int64 = 0

	for index, data := range challengeData {
		tx, err := i.rootOwnerSession.AddCheckpointChallenge(uid,
			checkpointHash, data.challengeTx, data.blockNumber)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to add challenge")
		}

		str, err := i.rootOwnerSession.GetCheckpointChallenge(uid,
			checkpointHash, big.NewInt(int64(index)))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(str.ChallengeTx, data.challengeTx) ||
			!bytes.Equal(str.ChallengeBlock.Bytes(),
				data.blockNumber.Bytes()) {
			t.Fatal("challenge is wrong")
		}
	}

	tx, err := i.rootOwnerSession.RemoveCheckpointChallenge(uid,
		checkpointHash, challengeData[index].challengeTx)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to remove challenge")
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid,
		checkpointHash)
	if err != nil {
		t.Fatal(err)
	}
	if length.Int64() != int64(len(challengeData)-1) {
		t.Fatal("disputes length is wrong")
	}

	str, err := i.rootOwnerSession.GetCheckpointChallenge(uid,
		checkpointHash, big.NewInt(index))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(str.ChallengeTx,
		challengeData[len(challengeData)-1].challengeTx) {
		t.Fatal("incorrect removal from the start")
	}
}

func TestRemoveChallengeFromFinish(t *testing.T) {
	i := newInstance(t)

	checkpointHash := hashGenerator()

	var index int64 = 4

	for index, data := range challengeData {
		tx, err := i.rootOwnerSession.AddCheckpointChallenge(uid,
			checkpointHash, data.challengeTx, data.blockNumber)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to add challenge")
		}

		str, err := i.rootOwnerSession.GetCheckpointChallenge(uid,
			checkpointHash, big.NewInt(int64(index)))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(str.ChallengeTx, data.challengeTx) ||
			!bytes.Equal(str.ChallengeBlock.Bytes(),
				data.blockNumber.Bytes()) {
			t.Fatal("challenge is wrong")
		}
	}

	tx, err := i.rootOwnerSession.RemoveCheckpointChallenge(uid,
		checkpointHash, challengeData[index].challengeTx)
	if err != nil {
		t.Fatal(err)
	}
	if !server.GoodTransaction(tx) {
		t.Fatal("failed to remove challenge")
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid,
		checkpointHash)
	if err != nil {
		t.Fatal(err)
	}
	if length.Int64() != int64(len(challengeData)-1) {
		t.Fatal("disputes length is wrong")
	}

	str, err := i.rootOwnerSession.GetCheckpointChallenge(uid, checkpointHash,
		big.NewInt(index))
	if err != nil {
		t.Fatal(err)
	}

	if len(str.ChallengeTx) != 0 {
		t.Fatal("incorrect removal from the finish")
	}
}

func TestRemoveAllChallenge(t *testing.T) {
	i := newInstance(t)

	checkpointHash := hashGenerator()

	for index, data := range challengeData {
		tx, err := i.rootOwnerSession.AddCheckpointChallenge(uid,
			checkpointHash, data.challengeTx, data.blockNumber)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to add challenge")
		}

		str, err := i.rootOwnerSession.GetCheckpointChallenge(uid,
			checkpointHash, big.NewInt(int64(index)))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(str.ChallengeTx, data.challengeTx) ||
			!bytes.Equal(str.ChallengeBlock.Bytes(),
				data.blockNumber.Bytes()) {
			t.Fatal("challenge is wrong")
		}
	}

	for _, data := range challengeData {
		tx, err := i.rootOwnerSession.RemoveCheckpointChallenge(uid,
			checkpointHash, data.challengeTx)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to remove challenge")
		}
	}

	_, err := i.rootOwnerSession.RemoveCheckpointChallenge(uid,
		checkpointHash, challengeData[0].challengeTx)
	if err == nil {
		t.Fatal("failed to remove challenge")
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid,
		checkpointHash)
	if err != nil {
		t.Fatal(err)
	}
	if length.Int64() != 0 {
		t.Fatal("disputes length is wrong")
	}
}

func TestMegaChallenge(t *testing.T) {
	i := newInstance(t)

	var number int64 = 1000

	chs := challengeGenerator(int(number))

	checkpointHash := hashGenerator()

	for _, data := range chs {
		tx, err := i.rootOwnerSession.AddCheckpointChallenge(uid,
			checkpointHash, data.challengeTx, data.blockNumber)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to add challenge")
		}
	}

	length, err := i.rootOwnerSession.CheckpointChallengesLength(uid,
		checkpointHash)
	if err != nil {
		t.Fatal(err)
	}

	if length.Int64() != number {
		t.Fatal("disputes length is wrong")
	}

	for _, data := range chs {
		tx, err := i.rootOwnerSession.RemoveCheckpointChallenge(uid,
			checkpointHash, data.challengeTx)
		if err != nil {
			t.Fatal(err)
		}
		if !server.GoodTransaction(tx) {
			t.Fatal("failed to remove challenge")
		}
	}

	length, err = i.rootOwnerSession.ChallengesLength(uid)
	if err != nil {
		t.Fatal(err)
	}

	if length.Int64() != 0 {
		t.Fatal("disputes length is wrong")
	}
}*/
