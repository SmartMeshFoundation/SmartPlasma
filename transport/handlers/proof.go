package handlers

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
