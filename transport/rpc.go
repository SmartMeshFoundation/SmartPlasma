package transport

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

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

type PendingCodeAtReq struct {
	Account common.Address
}

type PendingCodeAtResp struct {
	Code  []byte
	Error string
}

type PendingNonceAtReq struct {
	Account common.Address
}

type PendingNonceAtResp struct {
	Nonce uint64
	Error string
}

type SuggestGasPriceReq struct {
}

type SuggestGasPriceResp struct {
	Price *big.Int
	Error string
}

type EstimateGasReq struct {
	Call ethereum.CallMsg
}

type EstimateGasResp struct {
	Gas   uint64
	Error string
}

type RawReq struct {
	RawTx []byte
}

type RawResp struct {
	Error string
}

type BuildBlockReq struct{}

type BuildBlockResp struct {
	Hash  common.Hash
	Error string
}

type SendBlockHashReq struct {
	Hash common.Hash
}

type SendBlockHashResp struct {
	Error string
}

type LastBlockNumberReq struct{}

type LastBlockNumberResp struct {
	Number *big.Int
	Error  string
}

type CurrentBlockReq struct{}

type CurrentBlockResp struct {
	Block []byte
	Error string
}

type SaveBlockToDBReq struct {
	Number uint64
	Block  []byte
}

type SaveBlockToDBResp struct {
	Error string
}

type InitBlockReq struct {
}

type InitBlockResp struct {
	Error string
}

type VerifyTxProofReq struct {
	Uid   *big.Int
	Hash  common.Hash
	Block uint64
	Proof []byte
}

type VerifyTxProofResp struct {
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

// CreateProof creates merkle Proof for particular Uid.
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

// AddCheckpoint accepts Uid with transaction number for current checkpoint.
func (api *SmartPlasma) AddCheckpoint(req *AddCheckpointReq,
	resp *AddCheckpointResp) error {
	if err := api.service.AcceptUIDState(req.UID, req.Nonce); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// CreateUIDStateProof creates merkle Proof for particular Uid.
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

func (api *SmartPlasma) Deposit(req *RawReq, resp *RawResp) error {
	if err := api.service.MediatorTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) Withdraw(req *RawReq, resp *RawResp) error {
	if err := api.service.MediatorTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) StartExit(req *RawReq, resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) ChallengeExit(req *RawReq, resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) ChallengeCheckpoint(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) RespondChallengeExit(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) RespondCheckpointChallenge(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) RespondWithHistoricalCheckpoint(req *RawReq,
	resp *RawResp) error {
	if err := api.service.RootChainTransaction(req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) BuildBlock(req *BuildBlockReq,
	resp *BuildBlockResp) error {
	hash, err := api.service.BuildBlock()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Hash = hash
	return nil
}

func (api *SmartPlasma) SendBlockHash(req *SendBlockHashReq,
	resp *SendBlockHashResp) error {
	err := api.service.SendBlockHash(context.Background(), req.Hash)
	if err != nil {
		resp.Error = err.Error()
	}
	return nil
}

func (api *SmartPlasma) LastBlockNumber(req *LastBlockNumberReq,
	resp *LastBlockNumberResp) error {
	number, err := api.service.LastBlockNumber()
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Number = number
	return nil
}

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

func (api *SmartPlasma) InitBlock(req *InitBlockReq,
	resp *InitBlockResp) error {
	api.service.InitBlock()
	return nil
}

func (api *SmartPlasma) VerifyTxProof(req *VerifyTxProofReq,
	resp *VerifyTxProofResp) error {
	exists, err := api.service.VerifyTxProof(req.Uid, req.Hash,
		req.Block, req.Proof)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Exists = exists
	return nil
}
