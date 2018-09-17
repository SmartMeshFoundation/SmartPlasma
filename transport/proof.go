package transport

import (
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/pkg/errors"
)

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

// CreateUIDStateProof sends UID and checkpoint Hash to PlasmaCash RPC server.
// Returns merkle Proof for a UID.
func (c *Client) CreateUIDStateProof(
	uid *big.Int, checkpointHash common.Hash) ([]byte, *big.Int, error) {
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
			return nil, nil, replay.Error
		}
	case <-ctx.Done():
		return nil, nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, nil, errors.New(resp.Error)
	}

	return resp.Proof, resp.Nonce, nil
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
