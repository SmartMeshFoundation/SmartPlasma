package transport

import (
	"bytes"
	"context"
	"math/big"
	"time"

	"github.com/SmartMeshFoundation/Spectrum"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/service"
)

// SmartPlasma implements PlasmaCash methods to RPC server.
type SmartPlasma struct {
	timeout int
	service *service.Service
}

// AcceptTransactionReq is request for send Plasma transaction to PRC server.
type AcceptTransactionReq struct {
	Tx []byte
}

// AcceptTransactionResp is response for send Plasma transaction to PRC server.
type AcceptTransactionResp struct {
	Error string
}

// CreateProofReq is request for CreateProof method.
type CreateProofReq struct {
	UID   *big.Int
	Block uint64
}

// CreateProofResp is response for CreateProof method.
type CreateProofResp struct {
	Proof []byte
	Error string
}

// AddCheckpointReq is request for AddCheckpoint method.
type AddCheckpointReq struct {
	UID   *big.Int
	Nonce *big.Int
}

// AddCheckpointResp is response for AddCheckpoint method.
type AddCheckpointResp struct {
	Error string
}

// CreateUIDStateProofReq is request for CreateUIDStateProof method.
type CreateUIDStateProofReq struct {
	UID            *big.Int
	CheckpointHash common.Hash
}

// CreateUIDStateProofResp is response for CreateUIDStateProof method.
type CreateUIDStateProofResp struct {
	Nonce *big.Int
	Proof []byte
	Error string
}

// PendingCodeAtReq is request for PendingCodeAt method.
type PendingCodeAtReq struct {
	Account common.Address
}

// PendingCodeAtResp is response for PendingCodeAt method.
type PendingCodeAtResp struct {
	Code  []byte
	Error string
}

// PendingNonceAtReq is request for PendingNonceAt method.
type PendingNonceAtReq struct {
	Account common.Address
}

// PendingNonceAtResp is response for PendingNonceAt method.
type PendingNonceAtResp struct {
	Nonce uint64
	Error string
}

// SuggestGasPriceReq is request for SuggestGasPrice method.
type SuggestGasPriceReq struct {
}

// SuggestGasPriceResp is response for SuggestGasPrice method.
type SuggestGasPriceResp struct {
	Price *big.Int
	Error string
}

// EstimateGasReq is request for EstimateGas method.
type EstimateGasReq struct {
	Call ethereum.CallMsg
}

// EstimateGasResp is response for EstimateGas method.
type EstimateGasResp struct {
	Gas   *big.Int
	Error string
}

// WaitMinedReq is request for WaitMined method.
type WaitMinedReq struct {
	Tx []byte
}

// WaitMinedResp is response for WaitMined method.
type WaitMinedResp struct {
	Tr    []byte
	Error string
}

// RawReq is request for methods that works raw transactions.
type RawReq struct {
	RawTx []byte
}

// RawResp is response for methods that works raw transactions.
type RawResp struct {
	Error string
}

// BuildBlockReq is request for BuildBlock method.
type BuildBlockReq struct{}

// BuildBlockResp is response for BuildBlock method.
type BuildBlockResp struct {
	Hash  common.Hash
	Error string
}

// BuildCheckpointReq is request for BuildCheckpoint method.
type BuildCheckpointReq struct{}

// BuildCheckpointResp is response for BuildCheckpoint method.
type BuildCheckpointResp struct {
	Hash  common.Hash
	Error string
}

// SendBlockHashReq is request for SendBlockHash method.
type SendBlockHashReq struct {
	Hash common.Hash
}

// SendBlockHashResp is response for SendBlockHash method.
type SendBlockHashResp struct {
	Tx    []byte
	Error string
}

// SendCheckpointHashReq is request for SendCheckpointHash method.
type SendCheckpointHashReq struct {
	Hash common.Hash
}

// SendCheckpointHashResp is response for SendCheckpointHash method.
type SendCheckpointHashResp struct {
	Tx    []byte
	Error string
}

// LastBlockNumberReq is request for LastBlockNumber method.
type LastBlockNumberReq struct{}

// LastBlockNumberResp is response for LastBlockNumber method.
type LastBlockNumberResp struct {
	Number *big.Int
	Error  string
}

// CurrentBlockReq is request for CurrentBlock method.
type CurrentBlockReq struct{}

// CurrentBlockResp is response for CurrentBlock method.
type CurrentBlockResp struct {
	Block []byte
	Error string
}

// CurrentCheckpointReq is request for CurrentCheckpoint method.
type CurrentCheckpointReq struct{}

// CurrentCheckpointResp is response for CurrentCheckpoint method.
type CurrentCheckpointResp struct {
	Checkpoint []byte
	Error      string
}

// SaveBlockToDBReq is request for SaveBlockToDB method.
type SaveBlockToDBReq struct {
	Number uint64
	Block  []byte
}

// SaveBlockToDBResp is response for SaveBlockToDB method.
type SaveBlockToDBResp struct {
	Error string
}

// SaveCheckpointToDBReq is request for SaveCheckpointToDB method.
type SaveCheckpointToDBReq struct {
	Block []byte
}

// SaveCheckpointToDBResp is response for SaveCheckpointToDB method.
type SaveCheckpointToDBResp struct {
	Error string
}

// InitBlockReq is request for InitBlock method.
type InitBlockReq struct {
}

// InitBlockResp is response for InitBlock method.
type InitBlockResp struct {
	Error string
}

// InitCheckpointReq is request for InitCheckpoint method.
type InitCheckpointReq struct {
}

// InitCheckpointResp is response for InitCheckpoint method.
type InitCheckpointResp struct {
	Error string
}

// VerifyTxProofReq is request for VerifyTxProof method.
type VerifyTxProofReq struct {
	UID   *big.Int
	Hash  common.Hash
	Block uint64
	Proof []byte
}

// VerifyTxProofResp is response for VerifyTxProof method.
type VerifyTxProofResp struct {
	Exists bool
	Error  string
}

// VerifyCheckpointProofReq is request for VerifyCheckpointProof method.
type VerifyCheckpointProofReq struct {
	UID        *big.Int
	Number     *big.Int
	Checkpoint common.Hash
	Proof      []byte
}

// VerifyCheckpointProofResp is response for VerifyCheckpointProof method.
type VerifyCheckpointProofResp struct {
	Exists bool
	Error  string
}

// DepositCountReq is request for DepositCount method.
type DepositCountReq struct{}

// DepositCountResp is response for DepositCount method.
type DepositCountResp struct {
	Count *big.Int
	Error string
}

// ChallengePeriodReq is request for ChallengePeriod method.
type ChallengePeriodReq struct{}

// ChallengePeriodResp is response for ChallengePeriod method.
type ChallengePeriodResp struct {
	ChallengePeriod *big.Int
	Error           string
}

// OperatorReq is request for Operator method.
type OperatorReq struct{}

// OperatorResp is response for Operator method.
type OperatorResp struct {
	Operator common.Address
	Error    string
}

// ChildChainReq is request for ChildChain method.
type ChildChainReq struct {
	BlockNumber *big.Int
}

// ChildChainResp is response for ChildChain method.
type ChildChainResp struct {
	BlockHash common.Hash
	Error     string
}

// ExitsReq is request for Exits method.
type ExitsReq struct {
	UID *big.Int
}

// ExitsResp is response for Exits method.
type ExitsResp struct {
	State                *big.Int
	ExitTime             *big.Int
	ExitTxBlkNum         *big.Int
	ExitTx               []byte
	TxBeforeExitTxBlkNum *big.Int
	TxBeforeExitTx       []byte
	Error                string
}

// WalletReq is request for Wallet method.
type WalletReq struct {
	UID *big.Int
}

// WalletResp is response for Wallet method.
type WalletResp struct {
	Amount *big.Int
	Error  string
}

// ChallengeExistsReq is request for ChallengeExists method.
type ChallengeExistsReq struct {
	UID         *big.Int
	ChallengeTx []byte
}

// ChallengeExistsResp is response for ChallengeExists method.
type ChallengeExistsResp struct {
	Exists bool
	Error  string
}

// CheckpointIsChallengeReq is request for CheckpointIsChallenge method.
type CheckpointIsChallengeReq struct {
	UID         *big.Int
	Checkpoint  common.Hash
	ChallengeTx []byte
}

// CheckpointIsChallengeResp is response for CheckpointIsChallenge method.
type CheckpointIsChallengeResp struct {
	Exists bool
	Error  string
}

// ChallengesLengthReq is request for ChallengesLength method.
type ChallengesLengthReq struct {
	UID *big.Int
}

// ChallengesLengthResp is response for ChallengesLength method.
type ChallengesLengthResp struct {
	Length *big.Int
	Error  string
}

// CheckpointChallengesLengthReq is request
// for CheckpointChallengesLength method.
type CheckpointChallengesLengthReq struct {
	UID        *big.Int
	Checkpoint common.Hash
}

// CheckpointChallengesLengthResp is response
// for CheckpointChallengesLength method.
type CheckpointChallengesLengthResp struct {
	Length *big.Int
	Error  string
}

// GetChallengeReq is request for GetChallenge method.
type GetChallengeReq struct {
	UID   *big.Int
	Index *big.Int
}

// GetChallengeResp is response for GetChallenge method.
type GetChallengeResp struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
	Error          string
}

// GetCheckpointChallengeReq is request for GetCheckpointChallenge method.
type GetCheckpointChallengeReq struct {
	UID        *big.Int
	Checkpoint common.Hash
	Index      *big.Int
}

// GetCheckpointChallengeResp is response for GetCheckpointChallenge method.
type GetCheckpointChallengeResp struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
	Error          string
}

// SaveCurrentBlockReq is request for SaveCurrentBlock method.
type SaveCurrentBlockReq struct {
	Number uint64
}

// SaveCurrentBlockResp is response for SaveCurrentBlock method.
type SaveCurrentBlockResp struct {
	Error string
}

// SaveCurrentCheckpointBlockReq is request for SaveCurrentBlock method.
type SaveCurrentCheckpointBlockReq struct {
}

// SaveCurrentCheckpointBlockResp is response for SaveCurrentBlock method.
type SaveCurrentCheckpointBlockResp struct {
	Error string
}

// GetTransactionsBlockReq is request for GetTransactionsBlock method.
type GetTransactionsBlockReq struct {
	Number uint64
}

// GetTransactionsBlockResp is response for GetTransactionsBlock method.
type GetTransactionsBlockResp struct {
	Block []byte
	Error string
}

// GetCheckpointsBlockReq is request for GetCheckpointsBlock method.
type GetCheckpointsBlockReq struct {
	Hash common.Hash
}

// GetCheckpointsBlockResp is response for GetCheckpointsBlock method.
type GetCheckpointsBlockResp struct {
	Block []byte
	Error string
}

// AcceptTransaction accepts a raw transaction and returns a response.
func (api *SmartPlasma) AcceptTransaction(req *AcceptTransactionReq,
	resp *AcceptTransactionResp) error {
	tx := &transaction.Transaction{}

	if err := transaction.DecodeRLP(bytes.NewBuffer(req.Tx), tx); err != nil {
		resp.Error = err.Error()
		return nil
	}
	if err := api.service.AcceptTransaction(tx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// CreateProof creates merkle Proof for particular UID.
func (api *SmartPlasma) CreateProof(req *CreateProofReq,
	resp *CreateProofResp) error {

	proof, err := api.service.CreateProof(req.UID, req.Block)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Proof = proof
	return nil
}

// AddCheckpoint accepts UID with transaction number for current checkpoint.
func (api *SmartPlasma) AddCheckpoint(req *AddCheckpointReq,
	resp *AddCheckpointResp) error {
	if err := api.service.AcceptUIDState(req.UID, req.Nonce); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// CreateUIDStateProof creates merkle Proof for particular UID.
func (api *SmartPlasma) CreateUIDStateProof(req *CreateUIDStateProofReq,
	resp *CreateUIDStateProofResp) error {
	proof, nonce, err := api.service.CreateUIDStateProof(
		req.UID, req.CheckpointHash)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Proof = proof
	resp.Nonce = nonce
	return nil
}

// PendingCodeAt returns the code of the given Account in the pending state.
func (api *SmartPlasma) PendingCodeAt(req *PendingCodeAtReq,
	resp *PendingCodeAtResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	code, err := api.service.PendingCodeAt(ctx, req.Account)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Code = code
	return nil
}

// PendingNonceAt retrieves the current pending nonce
// associated with an Account.
func (api *SmartPlasma) PendingNonceAt(req *PendingNonceAtReq,
	resp *PendingNonceAtResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	nonce, err := api.service.PendingNonceAt(ctx, req.Account)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Nonce = nonce
	return nil
}

// SuggestGasPrice retrieves the currently suggested gas price
// to allow a timely execution of a transaction.
func (api *SmartPlasma) SuggestGasPrice(req *SuggestGasPriceReq,
	resp *SuggestGasPriceResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	price, err := api.service.SuggestGasPrice(ctx)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Price = price
	return nil
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.
func (api *SmartPlasma) EstimateGas(req *EstimateGasReq,
	resp *EstimateGasResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	gas, err := api.service.EstimateGas(ctx, req.Call)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Gas = gas
	return nil
}

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func (api *SmartPlasma) WaitMined(req *WaitMinedReq,
	resp *WaitMinedResp) error {
	tx := &types.Transaction{}
	err := tx.UnmarshalJSON(req.Tx)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	ctx, cancel := api.newContext()
	defer cancel()

	tr, err := api.service.Mine(ctx, tx)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	raw, err := tr.MarshalJSON()
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	resp.Tr = raw
	return nil
}

// Deposit invokes deposit method on Mediator contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) Deposit(req *RawReq, resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.MediatorTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// Withdraw invokes withdraw method
// on Mediator contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) Withdraw(req *RawReq, resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.MediatorTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// StartExit invokes startExit method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) StartExit(req *RawReq, resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.RootChainTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// ChallengeExit invokes challengeExit method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) ChallengeExit(req *RawReq, resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.RootChainTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// ChallengeCheckpoint invokes challengeCheckpoint method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) ChallengeCheckpoint(req *RawReq,
	resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.RootChainTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// RespondChallengeExit invokes respondChallengeExit method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) RespondChallengeExit(req *RawReq,
	resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.RootChainTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// RespondCheckpointChallenge invokes respondCheckpointChallenge method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) RespondCheckpointChallenge(req *RawReq,
	resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.RootChainTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// RespondWithHistoricalCheckpoint invokes respondWithHistoricalCheckpoint
// method on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) RespondWithHistoricalCheckpoint(req *RawReq,
	resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.RootChainTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// BuildBlock builds current transactions block.
func (api *SmartPlasma) BuildBlock(req *BuildBlockReq,
	resp *BuildBlockResp) error {
	hash, err := api.service.BuildBlock()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Hash = hash
	return nil
}

// BuildCheckpoint builds current checkpoints block.
func (api *SmartPlasma) BuildCheckpoint(req *BuildCheckpointReq,
	resp *BuildCheckpointResp) error {
	hash, err := api.service.BuildCheckpoint()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Hash = hash
	return nil
}

// SendBlockHash sends hash of transactions block to RootChain contract.
func (api *SmartPlasma) SendBlockHash(req *SendBlockHashReq,
	resp *SendBlockHashResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	tx, err := api.service.SendBlockHash(ctx, req.Hash)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	rawTx, err := tx.MarshalJSON()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Tx = rawTx
	return nil
}

// SendCheckpointHash sends hash of checkpoints block to RootChain contract.
func (api *SmartPlasma) SendCheckpointHash(req *SendCheckpointHashReq,
	resp *SendCheckpointHashResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	tx, err := api.service.SendChptHash(ctx, req.Hash)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	rawTx, err := tx.MarshalJSON()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Tx = rawTx
	return nil
}

// LastBlockNumber gets number by transactions block from RootChain contract.
func (api *SmartPlasma) LastBlockNumber(req *LastBlockNumberReq,
	resp *LastBlockNumberResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	number, err := api.service.LastBlockNumber(ctx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Number = number
	return nil
}

// CurrentBlock returns raw current transactions block.
func (api *SmartPlasma) CurrentBlock(req *CurrentBlockReq,
	resp *CurrentBlockResp) error {
	block := api.service.CurrentBlock()
	raw, err := block.Marshal()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Block = raw
	return nil
}

// CurrentCheckpoint returns raw current checkpoints block.
func (api *SmartPlasma) CurrentCheckpoint(req *CurrentCheckpointReq,
	resp *CurrentCheckpointResp) error {
	block := api.service.CurrentCheckpoint()
	raw, err := block.Marshal()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Checkpoint = raw
	return nil
}

// SaveBlockToDB saves transactions block in database.
func (api *SmartPlasma) SaveBlockToDB(req *SaveBlockToDBReq,
	resp *SaveBlockToDBResp) error {
	blk := transactions.NewTxBlock()
	err := blk.Unmarshal(req.Block)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	err = api.service.SaveBlockToDB(req.Number, blk)
	if err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// SaveCheckpointToDB saves checkpoints block in database.
func (api *SmartPlasma) SaveCheckpointToDB(req *SaveCheckpointToDBReq,
	resp *SaveCheckpointToDBResp) error {
	blk := checkpoints.NewBlock()
	err := blk.Unmarshal(req.Block)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	_, err = blk.Build()
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	err = api.service.SaveCheckpointToDB(blk)
	if err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// InitBlock initializes the new current transactions block.
func (api *SmartPlasma) InitBlock(req *InitBlockReq,
	resp *InitBlockResp) error {
	api.service.InitBlock()
	return nil
}

// InitCheckpoint initializes the new current checkpoints block.
func (api *SmartPlasma) InitCheckpoint(req *InitCheckpointReq,
	resp *InitCheckpointResp) error {
	api.service.InitCheckpoint()
	return nil
}

// VerifyTxProof checks whether the Plasma Cash
// transaction is included in the block.
func (api *SmartPlasma) VerifyTxProof(req *VerifyTxProofReq,
	resp *VerifyTxProofResp) error {
	exists, err := api.service.VerifyTxProof(req.UID, req.Hash,
		req.Block, req.Proof)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Exists = exists
	return nil
}

// VerifyCheckpointProof checks whether the UID is included in the block.
func (api *SmartPlasma) VerifyCheckpointProof(req *VerifyCheckpointProofReq,
	resp *VerifyCheckpointProofResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	exists, err := api.service.IsValidCheckpoint(
		ctx, req.UID, req.Number,
		req.Checkpoint, req.Proof)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Exists = exists
	return nil
}

// DepositCount returns a deposit counter.
func (api *SmartPlasma) DepositCount(
	req *DepositCountReq, resp *DepositCountResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	count, err := api.service.DepositCount(ctx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Count = count
	return nil
}

// ChallengePeriod returns a period for challenging in seconds.
func (api *SmartPlasma) ChallengePeriod(
	req *ChallengePeriodReq, resp *ChallengePeriodResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	secs, err := api.service.ChallengePeriod(ctx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.ChallengePeriod = secs
	return nil
}

// Operator returns a Plasma Cash operator address.
func (api *SmartPlasma) Operator(req *OperatorReq, resp *OperatorResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	operator, err := api.service.Operator(ctx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Operator = operator
	return nil
}

// ChildChain returns a block hash by a block number.
func (api *SmartPlasma) ChildChain(
	req *ChildChainReq, resp *ChildChainResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	hash, err := api.service.ChildChain(ctx, req.BlockNumber)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.BlockHash = hash
	return nil
}

// Exits returns a incomplete exit by UID.
func (api *SmartPlasma) Exits(req *ExitsReq, resp *ExitsResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	result, err := api.service.Exits(ctx, req.UID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.State = result.State
	resp.ExitTime = result.ExitTime
	resp.ExitTxBlkNum = result.ExitTxBlkNum
	resp.ExitTx = result.ExitTx
	resp.TxBeforeExitTxBlkNum = result.TxBeforeExitTxBlkNum
	resp.TxBeforeExitTx = result.TxBeforeExitTx
	return nil
}

// Wallet returns a deposit amount.
func (api *SmartPlasma) Wallet(
	req *WalletReq, resp *WalletResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	amount, err := api.service.Wallet(ctx, req.UID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Amount = amount
	return nil
}

// ChallengeExists if this is true,
// that a exit is blocked by a transaction of challenge.
func (api *SmartPlasma) ChallengeExists(
	req *ChallengeExistsReq, resp *ChallengeExistsResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	exists, err := api.service.ChallengeExists(
		ctx, req.UID, req.ChallengeTx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Exists = exists
	return nil
}

// CheckpointIsChallenge if this is true,
// that a checkpoint is blocked by a transaction of challenge.
func (api *SmartPlasma) CheckpointIsChallenge(
	req *CheckpointIsChallengeReq, resp *CheckpointIsChallengeResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	exists, err := api.service.CheckpointIsChallenge(ctx,
		req.UID, req.Checkpoint, req.ChallengeTx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Exists = exists
	return nil
}

// ChallengesLength returns number of disputes on withdrawal of uid.
func (api *SmartPlasma) ChallengesLength(
	req *ChallengesLengthReq, resp *ChallengesLengthResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	length, err := api.service.ChallengesLength(ctx, req.UID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Length = length
	return nil
}

// CheckpointChallengesLength returns number of disputes
// for checkpoint by a uid.
func (api *SmartPlasma) CheckpointChallengesLength(
	req *CheckpointChallengesLengthReq,
	resp *CheckpointChallengesLengthResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	length, err := api.service.CheckpointChallengesLength(
		ctx, req.UID, req.Checkpoint)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Length = length
	return nil
}

// GetChallenge returns exit challenge transaction by uid and index.
func (api *SmartPlasma) GetChallenge(
	req *GetChallengeReq, resp *GetChallengeResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	result, err := api.service.GetChallenge(
		ctx, req.UID, req.Index)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.ChallengeTx = result.ChallengeTx
	resp.ChallengeBlock = result.ChallengeBlock
	return nil
}

// GetCheckpointChallenge Returns checkpoint challenge transaction
// by checkpoint merkle root, uid and index.
func (api *SmartPlasma) GetCheckpointChallenge(
	req *GetCheckpointChallengeReq, resp *GetCheckpointChallengeResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	result, err := api.service.GetCheckpointChallenge(
		ctx, req.UID, req.Checkpoint, req.Index)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.ChallengeTx = result.ChallengeTx
	resp.ChallengeBlock = result.ChallengeBlock
	return nil
}

func (api *SmartPlasma) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(), time.Duration(api.timeout)*time.Second)
}

// SaveCurrentBlock saves current block to database.
func (api *SmartPlasma) SaveCurrentBlock(req *SaveCurrentBlockReq,
	resp *SaveCurrentBlockResp) error {
	err := api.service.SaveBlockToDB(req.Number, api.service.CurrentBlock())
	if err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// SaveCurrentCheckpointBlock saves current checkpoints block to database.
func (api *SmartPlasma) SaveCurrentCheckpointBlock(
	req *SaveCurrentCheckpointBlockReq,
	resp *SaveCurrentCheckpointBlockResp) error {
	err := api.service.SaveCheckpointToDB(api.service.CurrentCheckpoint())
	if err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// GetTransactionsBlock returns transactions block by number.
func (api *SmartPlasma) GetTransactionsBlock(req *GetTransactionsBlockReq,
	resp *GetTransactionsBlockResp) error {
	raw, err := api.service.RawBlockFromDB(req.Number)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Block = raw
	return nil
}

// GetCheckpointsBlock returns checkpoints block by number.
func (api *SmartPlasma) GetCheckpointsBlock(req *GetCheckpointsBlockReq,
	resp *GetCheckpointsBlockResp) error {
	raw, err := api.service.RawCheckpointFromDB(req.Hash)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Block = raw
	return nil
}
