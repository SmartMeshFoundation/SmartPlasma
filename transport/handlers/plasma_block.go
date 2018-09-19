package handlers

import (
	"context"
	"time"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
)

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

// SaveBlockToDB saves transactions block in database.
func (api *SmartPlasma) SaveBlockToDB(req *SaveBlockToDBReq,
	resp *SaveBlockToDBResp) error {
	blk := transactions.NewBlock()
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

// InitBlock initializes the new current transactions block.
func (api *SmartPlasma) InitBlock(req *InitBlockReq,
	resp *InitBlockResp) error {
	api.service.InitBlock()
	return nil
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

// ValidateBlock check current block and remove bad transactions.
func (api *SmartPlasma) ValidateBlock(req *ValidateBlockReq,
	resp *ValidateBlockResp) error {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second*600) // TODO: timeout hardcoded
	defer cancel()

	if err := api.service.ValidateBlock(ctx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}
