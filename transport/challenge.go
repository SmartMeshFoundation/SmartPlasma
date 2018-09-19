package transport

import (
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/transport/handlers"
)

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

	req := &handlers.RawReq{
		RawTx: raw,
	}

	var resp handlers.RawResp
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

	req := &handlers.RawReq{
		RawTx: raw,
	}
	var resp handlers.RawResp
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

	req := &handlers.RawReq{
		RawTx: raw,
	}
	var resp handlers.RawResp
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

	req := &handlers.RawReq{
		RawTx: raw,
	}

	var resp handlers.RawResp
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

	req := &handlers.RawReq{
		RawTx: raw,
	}

	var resp handlers.RawResp
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
	req := &handlers.ChallengeExistsReq{
		UID:         uid,
		ChallengeTx: challengeTx,
	}
	var resp *handlers.ChallengeExistsResp
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
	req := &handlers.CheckpointIsChallengeReq{
		UID:         uid,
		Checkpoint:  checkpoint,
		ChallengeTx: challengeTx,
	}
	var resp *handlers.CheckpointIsChallengeResp
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
	req := &handlers.ChallengesLengthReq{
		UID: uid,
	}
	var resp *handlers.ChallengesLengthResp
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
	req := &handlers.CheckpointChallengesLengthReq{
		UID:        uid,
		Checkpoint: checkpoint,
	}
	var resp *handlers.CheckpointChallengesLengthResp
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
	uid, index *big.Int) (resp *handlers.GetChallengeResp, err error) {
	ctx, cancel := c.newContext()
	defer cancel()
	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		result, err := session.GetChallenge(uid, index)
		if err != nil {
			return nil, err
		}
		resp = &handlers.GetChallengeResp{
			ChallengeTx:    result.ChallengeTx,
			ChallengeBlock: result.ChallengeBlock,
		}
		return resp, err
	}
	req := &handlers.GetChallengeReq{
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
	index *big.Int) (resp *handlers.GetCheckpointChallengeResp, err error) {
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
		resp = &handlers.GetCheckpointChallengeResp{
			ChallengeTx:    result.ChallengeTx,
			ChallengeBlock: result.ChallengeBlock,
		}
		return resp, err
	}
	req := &handlers.GetCheckpointChallengeReq{
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
