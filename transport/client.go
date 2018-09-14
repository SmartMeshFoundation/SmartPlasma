package transport

import (
	"context"
	"fmt"
	"math/big"
	"net/rpc"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
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
	WaitMinedMethod       = "SmartPlasma.WaitMined"

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

	DepositCountMethod               = "SmartPlasma.DepositCount"
	ChallengePeriodMethod            = "SmartPlasma.ChallengePeriod"
	OperatorMethod                   = "SmartPlasma.Operator"
	ChildChainMethod                 = "SmartPlasma.ChildChain"
	ExitsMethod                      = "SmartPlasma.Exits"
	WalletMethod                     = "SmartPlasma.Wallet"
	ChallengeExistsMethod            = "SmartPlasma.ChallengeExists"
	CheckpointIsChallengeMethod      = "SmartPlasma.CheckpointIsChallenge"
	ChallengesLengthMethod           = "SmartPlasma.ChallengesLength"
	CheckpointChallengesLengthMethod = "SmartPlasma.CheckpointChallengesLength"
	GetChallengeMethod               = "SmartPlasma.GetChallenge"
	GetCheckpointChallengeMethod     = "SmartPlasma.GetCheckpointChallenge"

	// additional methods
	SaveCurrentBlockMethod           = "SmartPlasma.SaveCurrentBlock"
	SaveCurrentCheckpointBlockMethod = "SmartPlasma.SaveCurrentCheckpointBlock"
	GetTransactionsBlockMethod       = "SmartPlasma.GetTransactionsBlock"
	GetCheckpointsBlockMethod        = "SmartPlasma.GetCheckpointsBlock"
)

// Client is RPC client for PlasmaCash.
type Client struct {
	connect          *rpc.Client
	backend          backend.Backend
	sessionMediator  *mediator.MediatorSession
	sessionRootChain *rootchain.RootChainSession
	opts             *account.PlasmaTransactOpts
	timeout          uint64 // in seconds
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
func (c *Client) DirectEthereumClient(opts bind.TransactOpts,
	mediatorAddress, rootChainAddress common.Address, backend backend.Backend) {
	mSession, err := mediator.NewMediatorSession(
		opts, mediatorAddress, backend)
	if err != nil {
		panic(err)
	}

	rSession, err := rootchain.NewRootChainSession(
		opts, rootChainAddress, backend)
	if err != nil {
		panic(err)
	}

	c.sessionRootChain = rSession
	c.sessionMediator = mSession
	c.backend = backend
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

func (c *Client) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(), time.Duration(c.timeout)*time.Second)
}

// AcceptTransaction sends raw transaction to PlasmaCash RPC server.
func (c *Client) AcceptTransaction(rawTx []byte) (err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &AcceptTransactionReq{rawTx}
	var resp *AcceptTransactionResp
	call := c.connect.Go(AcceptTransactionMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
}

// CreateProof sends UID and Block number to PlasmaCash RPC server.
// Returns merkle Proof for a UID.
func (c *Client) CreateProof(uid *big.Int, block uint64) ([]byte, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &CreateProofReq{UID: uid, Block: block}
	var resp *CreateProofResp
	call := c.connect.Go(CreateProofMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Proof, nil
}

// AddCheckpoint sends UID and current transaction nonce
// for inclusion in a checkpoint.
func (c *Client) AddCheckpoint(uid, nonce *big.Int) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &AddCheckpointReq{
		UID:   uid,
		Nonce: nonce,
	}
	var resp *AddCheckpointResp
	call := c.connect.Go(AddCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
}

// CreateUIDStateProof sends UID and checkpoint Hash to PlasmaCash RPC server.
// Returns merkle Proof for a UID.
func (c *Client) CreateUIDStateProof(
	uid *big.Int, checkpointHash common.Hash) ([]byte, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &CreateUIDStateProofReq{
		UID:            uid,
		CheckpointHash: checkpointHash,
	}
	var resp *CreateUIDStateProofResp
	call := c.connect.Go(CreateUIDStateProofMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Proof, nil
}

// Deposit transacts deposit function from Mediator contract.
func (c *Client) Deposit(currency common.Address,
	amount *big.Int) (tx *types.Transaction, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionMediator != nil {
		session := mediator.CopySession(c.sessionMediator)
		session.TransactOpts.Context = ctx
		return session.Deposit(currency, amount)
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
	call := c.connect.Go(DepositMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// Withdraw transacts withdraw function from Mediator contract.
func (c *Client) Withdraw(prevTx, prevTxProof []byte, prevTxBlkNum *big.Int,
	txRaw, txProof []byte, txBlkNum *big.Int) (*types.Transaction, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionMediator != nil {
		session := mediator.CopySession(c.sessionMediator)
		session.TransactOpts.Context = ctx
		return session.Withdraw(prevTx, prevTxProof, prevTxBlkNum,
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
	call := c.connect.Go(WithdrawMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
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
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.StartExit(previousTx,
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
	call := c.connect.Go(StartExitMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// ChallengeExit transacts challengeExit function from RootChain contract.
func (c *Client) ChallengeExit(uid *big.Int, challengeTx,
	proof []byte, challengeBlockNum *big.Int) (*types.Transaction, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.ChallengeExit(uid, challengeTx,
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
	call := c.connect.Go(ChallengeExitMethod, req, &resp, nil)
	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
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
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.ChallengeCheckpoint(uid, checkpointRoot,
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
	call := c.connect.Go(ChallengeCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
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
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.RespondChallengeExit(uid, challengeTx,
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
	call := c.connect.Go(RespondChallengeExitMethod, req, &resp, nil)
	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
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
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.RespondCheckpointChallenge(uid,
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
	call := c.connect.Go(RespondCheckpointChallengeMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
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
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.RespondWithHistoricalCheckpoint(uid,
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
	call := c.connect.Go(
		RespondWithHistoricalCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return tx, err
}

// BuildBlock builds current transactions block on the server side.
func (c *Client) BuildBlock() (hash common.Hash, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &BuildBlockReq{}
	var resp *BuildBlockResp
	call := c.connect.Go(BuildBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return common.Hash{}, replay.Error
		}
	case <-ctx.Done():
		return common.Hash{}, errors.New("timeout")
	}

	if resp.Error != "" {
		return common.Hash{}, errors.New(resp.Error)
	}

	return resp.Hash, err
}

// BuildCheckpoint  builds current checkpoint block on the server side.
func (c *Client) BuildCheckpoint() (hash common.Hash, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &BuildCheckpointReq{}
	var resp *BuildCheckpointResp
	call := c.connect.Go(BuildCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return common.Hash{}, replay.Error
		}
	case <-ctx.Done():
		return common.Hash{}, errors.New("timeout")
	}

	if resp.Error != "" {
		return common.Hash{}, errors.New(resp.Error)
	}
	return resp.Hash, err
}

// SendBlockHash sends new transactions block hash to RootChain contract.
func (c *Client) SendBlockHash(hash common.Hash) (*types.Transaction, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.NewBlock(hash)
	}

	req := &SendBlockHashReq{hash}
	var resp *SendBlockHashResp
	call := c.connect.Go(SendBlockHashMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	tx := &types.Transaction{}
	err := tx.UnmarshalJSON(resp.Tx)
	if err != nil {
		return nil, err
	}

	return tx, err
}

// SendCheckpointHash sends new checkpoints block hash to RootChain contract.
func (c *Client) SendCheckpointHash(hash common.Hash) (tx *types.Transaction,
	err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.NewCheckpoint(hash)
	}

	req := &SendCheckpointHashReq{hash}
	var resp *SendCheckpointHashResp
	call := c.connect.Go(SendCheckpointHashMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	tx = &types.Transaction{}
	err = tx.UnmarshalJSON(resp.Tx)
	if err != nil {
		return nil, err
	}

	return tx, err
}

// LastBlockNumber returns last transactions block number
// from RootChain contract.
func (c *Client) LastBlockNumber() (number *big.Int, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.BlockNumber()
	}
	req := &LastBlockNumberReq{}
	var resp LastBlockNumberResp
	call := c.connect.Go(LastBlockNumberMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Number, err
}

// CurrentBlock returns raw current transactions block.
func (c *Client) CurrentBlock() (block []byte, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &CurrentBlockReq{}
	var resp *CurrentBlockResp
	call := c.connect.Go(CurrentBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Block, err
}

// CurrentCheckpoint returns raw current checkpoints block.
func (c *Client) CurrentCheckpoint() (checkpoint []byte, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &CurrentCheckpointReq{}
	var resp *CurrentCheckpointResp
	call := c.connect.Go(CurrentCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Checkpoint, err
}

// SaveBlockToDB saves raw transactions block in database on server side.
func (c *Client) SaveBlockToDB(number uint64, raw []byte) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &SaveBlockToDBReq{
		Number: number,
		Block:  raw,
	}
	var resp *SaveBlockToDBResp
	call := c.connect.Go(SaveBlockToDBMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

// SaveCheckpointToDB saves raw checkpoints block in database on server side.
func (c *Client) SaveCheckpointToDB(raw []byte) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &SaveCheckpointToDBReq{
		Block: raw,
	}
	var resp *SaveCheckpointToDBResp
	call := c.connect.Go(SaveCheckpointToDBMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

// InitBlock initializes new current transactions block on server side.
func (c *Client) InitBlock() error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &InitBlockReq{}
	var resp *InitBlockResp
	call := c.connect.Go(InitBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
}

// InitCheckpoint initializes new current checkpoints block on server side.
func (c *Client) InitCheckpoint() error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &InitCheckpointReq{}
	var resp *InitCheckpointResp
	call := c.connect.Go(InitCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
}

// VerifyTxProof checks whether the transaction is included
// in the transactions block.
func (c *Client) VerifyTxProof(uid *big.Int, hash common.Hash,
	block uint64, proof []byte) (exists bool, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &VerifyTxProofReq{
		UID:   uid,
		Hash:  hash,
		Block: block,
		Proof: proof,
	}
	var resp *VerifyTxProofResp
	call := c.connect.Go(VerifyTxProofMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return false, replay.Error
		}
	case <-ctx.Done():
		return false, errors.New("timeout")
	}

	if resp.Error != "" {
		return false, errors.New(resp.Error)
	}

	return resp.Exists, err
}

// VerifyCheckpointProof checks whether the UID is included
// in the checkpoints block.
func (c *Client) VerifyCheckpointProof(uid *big.Int, number *big.Int,
	checkpoint common.Hash, proof []byte) (exists bool, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &VerifyCheckpointProofReq{
		UID:        uid,
		Number:     number,
		Checkpoint: checkpoint,
		Proof:      proof,
	}
	var resp *VerifyCheckpointProofResp
	call := c.connect.Go(VerifyCheckpointProofMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return false, replay.Error
		}
	case <-ctx.Done():
		return false, errors.New("timeout")
	}

	if resp.Error != "" {
		return false, errors.New(resp.Error)
	}

	return resp.Exists, err
}

// WaitMined to wait mining.
func (c *Client) WaitMined(
	ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if c.backend != nil {
		return c.backend.Mine(ctx, tx)
	}

	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &WaitMinedReq{
		Tx: raw,
	}

	var resp WaitMinedResp
	call := c.connect.Go(WaitMinedMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, err
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	tr := &types.Receipt{}
	err = tr.UnmarshalJSON(resp.Tr)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// DepositCount returns a deposit counter.
func (c *Client) DepositCount() (count *big.Int, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.DepositCount()
	}
	req := &DepositCountReq{}
	var resp *DepositCountResp
	call := c.connect.Go(DepositCountMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Count, err
}

// ChallengePeriod returns a period for challenging in seconds.
func (c *Client) ChallengePeriod() (count *big.Int, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.ChallengePeriod()
	}

	req := &ChallengePeriodReq{}
	var resp *ChallengePeriodResp
	call := c.connect.Go(ChallengePeriodMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.ChallengePeriod, err
}

// Operator returns a Plasma Cash operator address.
func (c *Client) Operator() (address common.Address, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.Operator()
	}
	req := &OperatorReq{}
	var resp *OperatorResp
	call := c.connect.Go(OperatorMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return common.Address{}, replay.Error
		}
	case <-ctx.Done():
		return common.Address{}, errors.New("timeout")
	}

	if resp.Error != "" {
		return common.Address{}, errors.New(resp.Error)
	}

	return resp.Operator, err
}

// ChildChain returns a block hash by a block number.
func (c *Client) ChildChain(
	blockNumber *big.Int) (hash common.Hash, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.ChildChain(blockNumber)
	}
	req := &ChildChainReq{
		BlockNumber: blockNumber,
	}
	var resp *ChildChainResp
	call := c.connect.Go(ChildChainMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return common.Hash{}, replay.Error
		}
	case <-ctx.Done():
		return common.Hash{}, errors.New("timeout")
	}

	if resp.Error != "" {
		return common.Hash{}, errors.New(resp.Error)
	}

	return resp.BlockHash, err
}

// Exits returns a incomplete exit by UID.
func (c *Client) Exits(uid *big.Int) (resp *ExitsResp, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		result, err := session.Exits(uid)
		if err != nil {
			return nil, err
		}
		resp = &ExitsResp{
			State:                result.State,
			ExitTime:             result.ExitTime,
			ExitTxBlkNum:         result.ExitTxBlkNum,
			ExitTx:               result.ExitTx,
			TxBeforeExitTxBlkNum: result.TxBeforeExitTxBlkNum,
			TxBeforeExitTx:       result.TxBeforeExitTx,
		}
		return resp, err
	}
	req := &ExitsReq{
		UID: uid,
	}

	call := c.connect.Go(ExitsMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp, err
}

// Wallet returns a deposit amount.
func (c *Client) Wallet(uid *big.Int) (amount *big.Int, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.Wallet(common.BigToHash(uid))
	}
	req := &WalletReq{
		UID: uid,
	}
	var resp *WalletResp
	call := c.connect.Go(WalletMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Amount, err
}

// ChallengeExists if this is true,
// that a exit is blocked by a transaction of challenge.
func (c *Client) ChallengeExists(
	uid *big.Int, challengeTx []byte) (exists bool, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.ChallengeExists(uid, challengeTx)
	}
	req := &ChallengeExistsReq{
		UID:         uid,
		ChallengeTx: challengeTx,
	}
	var resp *ChallengeExistsResp
	call := c.connect.Go(ChallengeExistsMethod, req, &resp, nil)
	if err != nil {
		return false, err
	}

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return false, replay.Error
		}
	case <-ctx.Done():
		return false, errors.New("timeout")
	}

	if resp.Error != "" {
		return false, errors.New(resp.Error)
	}

	return resp.Exists, err
}

// CheckpointIsChallenge if this is true,
// that a checkpoint is blocked by a transaction of challenge.
func (c *Client) CheckpointIsChallenge(
	uid *big.Int, checkpoint common.Hash,
	challengeTx []byte) (exists bool, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.CheckpointIsChallenge(
			uid, checkpoint, challengeTx)
	}
	req := &CheckpointIsChallengeReq{
		UID:         uid,
		Checkpoint:  checkpoint,
		ChallengeTx: challengeTx,
	}
	var resp *CheckpointIsChallengeResp
	call := c.connect.Go(CheckpointIsChallengeMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return false, replay.Error
		}
	case <-ctx.Done():
		return false, errors.New("timeout")
	}

	if resp.Error != "" {
		return false, errors.New(resp.Error)
	}

	return resp.Exists, err
}

// ChallengesLength returns number of disputes on withdrawal of uid.
func (c *Client) ChallengesLength(uid *big.Int) (length *big.Int, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.ChallengesLength(uid)
	}
	req := &ChallengesLengthReq{
		UID: uid,
	}
	var resp *ChallengesLengthResp
	call := c.connect.Go(ChallengesLengthMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Length, err
}

// CheckpointChallengesLength returns number of disputes
// for checkpoint by a uid.
func (c *Client) CheckpointChallengesLength(
	uid *big.Int, checkpoint common.Hash) (length *big.Int, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.CheckpointChallengesLength(uid, checkpoint)
	}
	req := &CheckpointChallengesLengthReq{
		UID:        uid,
		Checkpoint: checkpoint,
	}
	var resp *CheckpointChallengesLengthResp
	call := c.connect.Go(CheckpointChallengesLengthMethod, req, &resp, nil)
	if err != nil {
		return nil, err
	}

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Length, err
}

// GetChallenge returns exit challenge transaction by uid and index.
func (c *Client) GetChallenge(
	uid, index *big.Int) (resp *GetChallengeResp, err error) {
	ctx, cancel := c.newContext()
	defer cancel()
	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		result, err := session.GetChallenge(uid, index)
		if err != nil {
			return nil, err
		}
		resp = &GetChallengeResp{
			ChallengeTx:    result.ChallengeTx,
			ChallengeBlock: result.ChallengeBlock,
		}
		return resp, err
	}
	req := &GetChallengeReq{
		UID:   uid,
		Index: index,
	}
	call := c.connect.Go(GetChallengeMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp, err
}

// GetCheckpointChallenge Returns checkpoint challenge transaction
// by checkpoint merkle root, uid and index.
func (c *Client) GetCheckpointChallenge(uid *big.Int, checkpoint common.Hash,
	index *big.Int) (resp *GetCheckpointChallengeResp, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		result, err := session.GetCheckpointChallenge(
			uid, checkpoint, index)
		if err != nil {
			return nil, err
		}
		resp = &GetCheckpointChallengeResp{
			ChallengeTx:    result.ChallengeTx,
			ChallengeBlock: result.ChallengeBlock,
		}
		return resp, err
	}
	req := &GetCheckpointChallengeReq{
		UID:        uid,
		Checkpoint: checkpoint,
		Index:      index,
	}
	call := c.connect.Go(GetCheckpointChallengeMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp, err
}

// SaveCurrentBlock saves current transactions block in database on server side.
func (c *Client) SaveCurrentBlock(number uint64) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &SaveCurrentBlockReq{
		Number: number,
	}

	var resp *SaveCurrentBlockResp
	call := c.connect.Go(SaveCurrentBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

// SaveCurrentCheckpointBlock saves current checkpoints block in database on server side.
func (c *Client) SaveCurrentCheckpointBlock() error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &SaveCurrentCheckpointBlockReq{}

	var resp *SaveCurrentCheckpointBlockResp
	call := c.connect.Go(SaveCurrentCheckpointBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

// GetTransactionsBlock gets and builds transactions block.
func (c *Client) GetTransactionsBlock(
	number uint64) (transactions.TxBlock, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &GetTransactionsBlockReq{
		Number: number,
	}

	var resp *GetTransactionsBlockResp
	call := c.connect.Go(GetTransactionsBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	bl := transactions.NewTxBlock()
	err := bl.Unmarshal(resp.Block)
	if err != nil {
		return nil, err
	}
	_, err = bl.Build()
	if err != nil {
		return nil, err
	}

	return bl, nil
}

// GetCheckpointsBlock gets and builds checkpoints block.
func (c *Client) GetCheckpointsBlock(
	hash common.Hash) (checkpoints.CheckpointBlock, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &GetCheckpointsBlockReq{
		Hash: hash,
	}

	var resp *GetCheckpointsBlockResp
	call := c.connect.Go(GetCheckpointsBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	bl := checkpoints.NewBlock()
	err := bl.Unmarshal(resp.Block)
	if err != nil {
		return nil, err
	}
	_, err = bl.Build()
	if err != nil {
		return nil, err
	}

	return bl, nil
}
