package transport

import (
	"fmt"
	"math/big"
	"net/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

// Errors.
var (
	ErrTransactor = errors.New("transactor is missing")
)

// Smart Plasma RPC Methods.
const (
	AcceptTransactionMethod   = "SmartPlasma.AcceptTransaction"
	CreateProofMethod         = "SmartPlasma.CreateProof"
	AddCheckpointMethod       = "SmartPlasma.AddCheckpoint"
	CreateUIDStateProofMethod = "SmartPlasma.CreateUIDStateProof"

	PendingCodeAtMethod   = "SmartPlasma.PendingCodeAt"
	PendingNonceAtMethod  = "SmartPlasma.PendingNonceAt"
	SuggestGasPriceMethod = "SmartPlasma.SuggestGasPrice"
	EstimateGasMethod     = "SmartPlasma.EstimateGas"

	DepositMethod                         = "SmartPlasma.Deposit"
	WithdrawMethod                        = "SmartPlasma.Withdraw"
	StartExitMethod                       = "SmartPlasma.StartExit"
	ChallengeExitMethod                   = "SmartPlasma.ChallengeExit"
	ChallengeCheckpointMethod             = "SmartPlasma.ChallengeCheckpoint"
	RespondChallengeExitMethod            = "SmartPlasma.RespondChallengeExit"
	RespondCheckpointChallengeMethod      = "SmartPlasma.RespondCheckpointChallenge"
	RespondWithHistoricalCheckpointMethod = "SmartPlasma.RespondWithHistoricalCheckpoint"

	BuildBlockMethod      = "SmartPlasma.BuildBlock"
	SendBlockHashMethod   = "SmartPlasma.SendBlockHash"
	LastBlockNumberMethod = "SmartPlasma.LastBlockNumber"
	CurrentBlockMethod    = "SmartPlasma.CurrentBlock"
	SaveBlockToDBMethod   = "SmartPlasma.SaveBlockToDB"
	InitBlockMethod       = "SmartPlasma.InitBlock"
	VerifyTxProofMethod   = "SmartPlasma.VerifyTxProof"

	BuildCheckpointMethod       = "SmartPlasma.BuildCheckpoint"
	SendCheckpointHashMethod    = "SmartPlasma.SendCheckpointHash"
	CurrentCheckpointMethod     = "SmartPlasma.CurrentCheckpoint"
	SaveCheckpointToDBMethod    = "SmartPlasma.SaveCheckpointToDB"
	InitCheckpointMethod        = "SmartPlasma.InitCheckpoint"
	VerifyCheckpointProofMethod = "SmartPlasma.VerifyCheckpointProof"
)

// Client is RPC client for PlasmaCash.
type Client struct {
	connect          *rpc.Client
	sessionMediator  *mediator.MediatorSession
	sessionRootChain *rootchain.RootChainSession
	opts             *account.PlasmaTransactOpts
	timeout          uint64
	med              *build.Contract
	root             *build.Contract
}

// NewClient creates new PlasmaCash client.
// The Client must initialize RemoteEthereumClient or DirectEthereumClient.
func NewClient(timeout uint64, opts *account.PlasmaTransactOpts) *Client {
	return &Client{
		timeout: timeout,
		opts:    opts,
	}
}

// RemoteEthereumClient initializes work with remote ethereum client.
// Ethereum transactions are generated locally, signed locally,
// packaged and sent to the server. If this function is not called,
// then all transactions are sent directly to the Ethereum.
func (c *Client) RemoteEthereumClient(root, med *build.Contract) {
	c.med = med
	c.root = root
}

// DirectEthereumClient initializes work with direct ethereum client.
func (c *Client) DirectEthereumClient(sessionMediator *mediator.MediatorSession,
	sessionRootChain *rootchain.RootChainSession) {
	c.sessionRootChain = sessionRootChain
	c.sessionMediator = sessionMediator
}

// Connect tries to connect to a PlasmaCash RPC server.
func (c *Client) Connect(address string, port uint16) error {
	client, err := rpc.DialHTTP(tcpProtocol,
		fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return err
	}

	c.connect = client
	return nil
}

// ConnectString tries to connect to a PlasmaCash RPC server.
func (c *Client) ConnectString(str string) error {
	client, err := rpc.DialHTTP(tcpProtocol, str)
	if err != nil {
		return err
	}

	c.connect = client
	return nil
}

// Close closes connection to PlasmaCash RPC server.
func (c *Client) Close() error {
	return c.connect.Close()
}

// AcceptTransaction sends raw transaction to PlasmaCash RPC server.
func (c *Client) AcceptTransaction(rawTx []byte) (resp *AcceptTransactionResp,
	err error) {
	req := &AcceptTransactionReq{rawTx}

	if err = c.connect.Call(AcceptTransactionMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}

// CreateProof sends UID and Block number to PlasmaCash RPC server.
// Returns merkle Proof for a UID.
func (c *Client) CreateProof(uid *big.Int,
	block uint64) (resp *CreateProofResp, err error) {
	req := &CreateProofReq{UID: uid, Block: block}

	if err = c.connect.Call(CreateProofMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}

// AddCheckpoint sends UID and current transaction nonce
// for inclusion in a checkpoint.
func (c *Client) AddCheckpoint(uid,
	nonce *big.Int) (resp *AddCheckpointResp, err error) {
	req := &AddCheckpointReq{
		UID:   uid,
		Nonce: nonce,
	}

	if err = c.connect.Call(AddCheckpointMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}

// CreateUIDStateProof sends UID and checkpoint Hash to PlasmaCash RPC server.
// Returns merkle Proof for a UID.
func (c *Client) CreateUIDStateProof(uid *big.Int,
	checkpointHash common.Hash) (resp *CreateUIDStateProofResp, err error) {
	req := &CreateUIDStateProofReq{
		UID:            uid,
		CheckpointHash: checkpointHash,
	}

	if c.connect.Call(CreateUIDStateProofMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}

// Deposit transacts deposit function from Mediator contract.
func (c *Client) Deposit(currency common.Address,
	amount *big.Int) (tx *types.Transaction, err error) {
	if c.sessionMediator != nil {
		return c.sessionMediator.Deposit(currency, amount)
	}

	if c.med == nil {
		return nil, ErrTransactor
	}

	tx, err = c.med.Transaction(c.opts.TransactOpts,
		"deposit", currency, amount)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(DepositMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// Withdraw transacts withdraw function from Mediator contract.
func (c *Client) Withdraw(prevTx, prevTxProof []byte, prevTxBlkNum *big.Int,
	txRaw, txProof []byte, txBlkNum *big.Int) (*types.Transaction, error) {
	if c.sessionMediator != nil {
		return c.sessionMediator.Withdraw(prevTx, prevTxProof, prevTxBlkNum,
			txRaw, txProof, txBlkNum)
	}

	if c.med == nil {
		return nil, ErrTransactor
	}

	tx, err := c.med.Transaction(c.opts.TransactOpts,
		"withdraw", prevTx, prevTxProof, prevTxBlkNum, txRaw,
		txProof, txBlkNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(WithdrawMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// StartExit transacts startExit function from RootChain contract.
func (c *Client) StartExit(previousTx, previousTxProof []byte,
	previousTxBlockNum *big.Int, lastTx, lastTxProof []byte,
	lastTxBlockNum *big.Int) (*types.Transaction, error) {
	if c.sessionRootChain != nil {
		return c.sessionRootChain.StartExit(previousTx,
			previousTxProof, previousTxBlockNum, lastTx,
			lastTxProof, lastTxBlockNum)
	}

	if c.root == nil {
		return nil, ErrTransactor
	}

	tx, err := c.root.Transaction(c.opts.TransactOpts,
		"startExit", previousTx, previousTxProof, previousTxBlockNum,
		lastTx, lastTxProof, lastTxBlockNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(StartExitMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// ChallengeExit transacts challengeExit function from RootChain contract.
func (c *Client) ChallengeExit(uid *big.Int, challengeTx,
	proof []byte, challengeBlockNum *big.Int) (*types.Transaction, error) {
	if c.sessionRootChain != nil {
		return c.sessionRootChain.ChallengeExit(uid, challengeTx,
			proof, challengeBlockNum)
	}
	if c.root == nil {
		return nil, ErrTransactor
	}

	tx, err := c.root.Transaction(c.opts.TransactOpts,
		"challengeExit", uid, challengeTx, proof, challengeBlockNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(ChallengeExitMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// ChallengeCheckpoint transacts challengeCheckpoint function
// from RootChain contract.
func (c *Client) ChallengeCheckpoint(uid *big.Int, checkpointRoot [32]byte,
	checkpointProof []byte, wrongNonce *big.Int, lastTx,
	lastTxProof []byte, lastTxBlockNum *big.Int) (*types.Transaction, error) {
	if c.sessionRootChain != nil {
		return c.sessionRootChain.ChallengeCheckpoint(uid, checkpointRoot,
			checkpointProof, wrongNonce, lastTx, lastTxProof, lastTxBlockNum)
	}
	if c.root == nil {
		return nil, ErrTransactor
	}

	tx, err := c.root.Transaction(c.opts.TransactOpts,
		"challengeCheckpoint", uid, checkpointRoot, checkpointProof,
		wrongNonce, lastTx, lastTxProof, lastTxBlockNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(ChallengeCheckpointMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// RespondChallengeExit transacts respondChallengeExit function
// from RootChain contract.
func (c *Client) RespondChallengeExit(uid *big.Int, challengeTx, respondTx,
	proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	if c.sessionRootChain != nil {
		return c.sessionRootChain.RespondChallengeExit(uid, challengeTx,
			respondTx, proof, blockNum)
	}
	if c.root == nil {
		return nil, ErrTransactor
	}

	tx, err := c.root.Transaction(c.opts.TransactOpts,
		"respondChallengeExit", uid, challengeTx,
		respondTx, proof, blockNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(RespondChallengeExitMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// RespondCheckpointChallenge transacts respondCheckpointChallenge function
// from RootChain contract.
func (c *Client) RespondCheckpointChallenge(uid *big.Int,
	checkpointRoot [32]byte, challengeTx []byte, respondTx []byte,
	proof []byte, blockNum *big.Int) (*types.Transaction, error) {
	if c.sessionRootChain != nil {
		return c.sessionRootChain.RespondCheckpointChallenge(uid,
			checkpointRoot, challengeTx, respondTx, proof, blockNum)
	}
	if c.root == nil {
		return nil, ErrTransactor
	}

	tx, err := c.root.Transaction(c.opts.TransactOpts,
		"respondCheckpointChallenge", uid, checkpointRoot, challengeTx,
		respondTx, proof, blockNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(RespondCheckpointChallengeMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// RespondWithHistoricalCheckpoint transacts respondWithHistoricalCheckpoint
// function from RootChain contract.
func (c *Client) RespondWithHistoricalCheckpoint(uid *big.Int,
	checkpointRoot [32]byte, checkpointProof []byte,
	historicalCheckpointRoot [32]byte, historicalCheckpointProof []byte,
	challengeTx []byte, moreNonce *big.Int) (*types.Transaction, error) {
	if c.sessionRootChain != nil {
		return c.sessionRootChain.RespondWithHistoricalCheckpoint(uid,
			checkpointRoot, checkpointProof, historicalCheckpointRoot,
			historicalCheckpointProof, challengeTx, moreNonce)
	}
	if c.root == nil {
		return nil, ErrTransactor
	}

	tx, err := c.root.Transaction(c.opts.TransactOpts,
		"respondWithHistoricalCheckpoint", uid, checkpointRoot,
		checkpointProof, historicalCheckpointRoot, historicalCheckpointProof,
		challengeTx, moreNonce)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}

	var resp RawResp
	err = c.connect.Call(RespondWithHistoricalCheckpointMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// BuildBlock builds current transactions block on the server side.
func (c *Client) BuildBlock() (resp *BuildBlockResp,
	err error) {
	req := &BuildBlockReq{}
	err = c.connect.Call(BuildBlockMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// BuildCheckpoint  builds current checkpoint block on the server side.
func (c *Client) BuildCheckpoint() (resp *BuildCheckpointResp,
	err error) {
	req := &BuildCheckpointReq{}
	err = c.connect.Call(BuildCheckpointMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// SendBlockHash sends new transactions block hash to RootChain contract.
func (c *Client) SendBlockHash(hash common.Hash) (resp *SendBlockHashResp,
	err error) {
	req := &SendBlockHashReq{hash}
	err = c.connect.Call(SendBlockHashMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// SendCheckpointHash sends new checkpoints block hash to RootChain contract.
func (c *Client) SendCheckpointHash(hash common.Hash) (resp *SendBlockHashResp,
	err error) {

	req := &SendBlockHashReq{hash}
	err = c.connect.Call(SendCheckpointHashMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// LastBlockNumber returns last transactions block number
// from RootChain contract.
func (c *Client) LastBlockNumber() (number *big.Int, err error) {
	if c.sessionRootChain != nil {
		return c.sessionRootChain.BlockNumber()
	}
	req := &LastBlockNumberReq{}
	var resp LastBlockNumberResp

	err = c.connect.Call(LastBlockNumberMethod, req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errors.New(err.Error())
	}

	return resp.Number, err
}

// CurrentBlock returns raw current transactions block.
func (c *Client) CurrentBlock() (resp *CurrentBlockResp, err error) {
	req := &CurrentBlockReq{}
	err = c.connect.Call(CurrentBlockMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// CurrentCheckpoint returns raw current checkpoints block.
func (c *Client) CurrentCheckpoint() (resp *CurrentCheckpointResp, err error) {
	req := &CurrentCheckpointReq{}
	err = c.connect.Call(CurrentCheckpointMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// SaveBlockToDB saves raw transactions block in database on server side.
func (c *Client) SaveBlockToDB(number uint64,
	raw []byte) (resp *SaveBlockToDBResp, err error) {
	req := &SaveBlockToDBReq{
		Number: number,
		Block:  raw,
	}
	err = c.connect.Call(SaveBlockToDBMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// SaveCheckpointToDB saves raw checkpoints block in database on server side.
func (c *Client) SaveCheckpointToDB(
	raw []byte) (resp *SaveCheckpointToDBResp, err error) {
	req := &SaveCheckpointToDBReq{
		Block: raw,
	}
	err = c.connect.Call(SaveCheckpointToDBMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// InitBlock initializes new current transactions block on server side.
func (c *Client) InitBlock() (resp *InitBlockResp, err error) {
	req := &InitBlockReq{}
	err = c.connect.Call(InitBlockMethod, req, &resp)
	return resp, err
}

// InitCheckpoint initializes new current checkpoints block on server side.
func (c *Client) InitCheckpoint() (resp *InitCheckpointResp, err error) {
	req := &InitCheckpointReq{}
	err = c.connect.Call(InitCheckpointMethod, req, &resp)
	return resp, err
}

// VerifyTxProof checks whether the transaction is included
// in the transactions block.
func (c *Client) VerifyTxProof(uid *big.Int, hash common.Hash,
	block uint64, proof []byte) (resp *VerifyTxProofResp, err error) {
	req := &VerifyTxProofReq{
		UID:   uid,
		Hash:  hash,
		Block: block,
		Proof: proof,
	}
	err = c.connect.Call(VerifyTxProofMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// VerifyCheckpointProof checks whether the UID is included
// in the checkpoints block.
func (c *Client) VerifyCheckpointProof(uid *big.Int, number *big.Int,
	checkpoint common.Hash, proof []byte) (resp *VerifyCheckpointProofResp,
	err error) {
	req := &VerifyCheckpointProofReq{
		UID:        uid,
		Number:     number,
		Checkpoint: checkpoint,
		Proof:      proof,
	}
	err = c.connect.Call(VerifyCheckpointProofMethod, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
