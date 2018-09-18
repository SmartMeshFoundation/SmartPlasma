package handlers

import "github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"

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

// InitCheckpoint initializes the new current checkpoints block.
func (api *SmartPlasma) InitCheckpoint(req *InitCheckpointReq,
	resp *InitCheckpointResp) error {
	api.service.InitCheckpoint()
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
