package handlers

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
