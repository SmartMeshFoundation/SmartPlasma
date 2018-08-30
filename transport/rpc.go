package transport

import (
	"bytes"
	"context"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum"
	"github.com/SmartMeshFoundation/Spectrum/common"

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
	Error string
}

// SendCheckpointHashReq is request for SendCheckpointHash method.
type SendCheckpointHashReq struct {
	Hash common.Hash
}

// SendCheckpointHashResp is response for SendCheckpointHash method.
type SendCheckpointHashResp struct {
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
	proof, err := api.service.CreateUIDStateProof(req.UID, req.CheckpointHash)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Proof = proof
	return nil
}

// PendingCodeAt returns the code of the given Account in the pending state.
func (api *SmartPlasma) PendingCodeAt(req *PendingCodeAtReq,
	resp *PendingCodeAtResp) error {
	code, err := api.service.PendingCodeAt(context.Background(), req.Account)
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
	nonce, err := api.service.PendingNonceAt(context.Background(), req.Account)
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
	price, err := api.service.SuggestGasPrice(context.Background())
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
	gas, err := api.service.EstimateGas(context.Background(), req.Call)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Gas = gas
	return nil
}

// Deposit invokes deposit method on Mediator contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) Deposit(req *RawReq, resp *RawResp) error {
	if err := api.service.MediatorTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// Withdraw invokes withdraw method
// on Mediator contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) Withdraw(req *RawReq, resp *RawResp) error {
	if err := api.service.MediatorTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// StartExit invokes startExit method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) StartExit(req *RawReq, resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// ChallengeExit invokes challengeExit method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) ChallengeExit(req *RawReq, resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// ChallengeCheckpoint invokes challengeCheckpoint method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) ChallengeCheckpoint(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// RespondChallengeExit invokes respondChallengeExit method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) RespondChallengeExit(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// RespondCheckpointChallenge invokes respondCheckpointChallenge method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) RespondCheckpointChallenge(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// RespondWithHistoricalCheckpoint invokes respondWithHistoricalCheckpoint
// method on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) RespondWithHistoricalCheckpoint(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
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
	err := api.service.SendBlockHash(context.Background(), req.Hash)
	if err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// SendCheckpointHash sends hash of checkpoints block to RootChain contract.
func (api *SmartPlasma) SendCheckpointHash(req *SendCheckpointHashReq,
	resp *SendBlockHashResp) error {
	err := api.service.SendChptHash(context.Background(), req.Hash)
	if err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// LastBlockNumber gets number by transactions block from RootChain contract.
func (api *SmartPlasma) LastBlockNumber(req *LastBlockNumberReq,
	resp *LastBlockNumberResp) error {
	number, err := api.service.LastBlockNumber()
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
	exists, err := api.service.IsValidCheckpoint(req.UID, req.Number,
		req.Checkpoint, req.Proof)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Exists = exists
	return nil
}
