package transport

import (
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

// BuildCheckpoint  builds current checkpoint block on the server side.
func (c *Client) BuildCheckpoint() (hash common.Hash, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &BuildCheckpointReq{}
	var resp *BuildCheckpointResp
	call := c.connect.Go(BuildCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return common.Hash{}, replay.Error
		}
	case <-ctx.Done():
		return common.Hash{}, errors.New("timeout")
	}

	if resp.Error != "" {
		return common.Hash{}, errors.New(resp.Error)
	}
	return resp.Hash, err
}

// SendCheckpointHash sends new checkpoints block hash to RootChain contract.
func (c *Client) SendCheckpointHash(hash common.Hash) (tx *types.Transaction,
	err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.NewCheckpoint(hash)
	}

	req := &SendCheckpointHashReq{Hash: hash}
	var resp *SendCheckpointHashResp
	call := c.connect.Go(SendCheckpointHashMethod, req, &resp, nil)

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

	tx = &types.Transaction{}
	err = tx.UnmarshalJSON(resp.Tx)
	if err != nil {
		return nil, err
	}

	return tx, err
}

// CurrentCheckpoint returns raw current checkpoints block.
func (c *Client) CurrentCheckpoint() (checkpoint []byte, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &CurrentCheckpointReq{}
	var resp *CurrentCheckpointResp
	call := c.connect.Go(CurrentCheckpointMethod, req, &resp, nil)

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

	return resp.Checkpoint, err
}

// SaveCheckpointToDB saves raw checkpoints block in database on server side.
func (c *Client) SaveCheckpointToDB(raw []byte) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &SaveCheckpointToDBReq{
		Block: raw,
	}
	var resp *SaveCheckpointToDBResp
	call := c.connect.Go(SaveCheckpointToDBMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

// InitCheckpoint initializes new current checkpoints block on server side.
func (c *Client) InitCheckpoint() error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &InitCheckpointReq{}
	var resp *InitCheckpointResp
	call := c.connect.Go(InitCheckpointMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}

	return nil
}

// SaveCurrentCheckpointBlock saves current checkpoints block
// in database on server side.
func (c *Client) SaveCurrentCheckpointBlock() error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &SaveCurrentCheckpointBlockReq{}

	var resp *SaveCurrentCheckpointBlockResp
	call := c.connect.Go(SaveCurrentCheckpointBlockMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return replay.Error
		}
	case <-ctx.Done():
		return errors.New("timeout")
	}

	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

// GetCheckpointsBlock gets and builds checkpoints block.
func (c *Client) GetCheckpointsBlock(
	hash common.Hash) (checkpoints.CheckpointBlock, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &GetCheckpointsBlockReq{
		Hash: hash,
	}

	var resp *GetCheckpointsBlockResp
	call := c.connect.Go(GetCheckpointsBlockMethod, req, &resp, nil)

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

	bl := checkpoints.NewBlock()
	err := bl.Unmarshal(resp.Block)
	if err != nil {
		return nil, err
	}
	_, err = bl.Build()
	if err != nil {
		return nil, err
	}

	return bl, nil
}
