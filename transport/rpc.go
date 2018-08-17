package transport

import (
	"bytes"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database"
)

// SmartPlasma implements PlasmaCash methods to RPC server.
type SmartPlasma struct {
	timeout      int
	currentBlock transactions.TxBlock
	currentChpt  checkpoints.CheckpointBlock
	blockBase    database.Database
	chptBase     database.Database
	session      *rootchain.RootChainSession
	backend      backend.Backend
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

// AcceptTransaction accepts a raw transaction and returns a response.
func (api *SmartPlasma) AcceptTransaction(req *AcceptTransactionReq,
	resp *AcceptTransactionResp) error {
	tx := &transaction.Transaction{}

	if err := transaction.DecodeRLP(bytes.NewBuffer(req.Tx), tx); err != nil {
		resp.Error = err.Error()
		return nil
	}

	if err := api.currentBlock.AddTx(tx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// CreateProof creates merkle proof for particular uid.
func (api *SmartPlasma) CreateProof(req *CreateProofReq,
	resp *CreateProofResp) error {
	raw, err := api.rawBlockFromDB(req.Block)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	blk := transactions.NewTxBlock()
	err = buildBlockFromBytes(blk, raw)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	resp.Proof = blk.CreateProof(req.UID)
	return nil
}

// AddCheckpoint accepts uid with transaction number for current checkpoint.
func (api *SmartPlasma) AddCheckpoint(req *AddCheckpointReq,
	resp *AddCheckpointResp) error {
	if err := api.currentChpt.AddCheckpoint(req.UID, req.Nonce); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// CreateUIDStateProof creates merkle proof for particular uid.
func (api *SmartPlasma) CreateUIDStateProof(req *CreateUIDStateProofReq,
	resp *CreateUIDStateProofResp) error {
	raw, err := api.rawCheckpointFromDB(req.CheckpointHash)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	blk := checkpoints.NewBlock()
	err = buildBlockFromBytes(blk, raw)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	resp.Proof = blk.CreateProof(req.UID)
	return nil
}

func (api *SmartPlasma) rawBlockFromDB(number uint64) ([]byte, error) {
	return api.blockBase.Get(strconv.AppendUint(nil, number, 10))
}

func (api *SmartPlasma) rawCheckpointFromDB(hash common.Hash) ([]byte, error) {
	return api.chptBase.Get(hash.Bytes())
}

func buildBlockFromBytes(blk block.Block, raw []byte) error {
	err := blk.Unmarshal(raw)
	if err != nil {
		return err
	}

	_, err = blk.Build()
	return err
}
