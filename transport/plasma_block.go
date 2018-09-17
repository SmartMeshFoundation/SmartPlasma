package transport

import (
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/transport/handlers"
)

// BuildBlock builds current transactions block on the server side.
func (c *Client) BuildBlock() (hash common.Hash, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &handlers.BuildBlockReq{}
	var resp *handlers.BuildBlockResp
	call := c.connect.Go(BuildBlockMethod, req, &resp, nil)

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

// SendBlockHash sends new transactions block hash to RootChain contract.
func (c *Client) SendBlockHash(hash common.Hash) (*types.Transaction, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.NewBlock(hash)
	}

	req := &handlers.SendBlockHashReq{Hash: hash}
	var resp *handlers.SendBlockHashResp
	call := c.connect.Go(SendBlockHashMethod, req, &resp, nil)

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

	tx := &types.Transaction{}
	err := tx.UnmarshalJSON(resp.Tx)
	if err != nil {
		return nil, err
	}

	return tx, err
}

// LastBlockNumber returns last transactions block number
// from RootChain contract.
func (c *Client) LastBlockNumber() (number *big.Int, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.BlockNumber()
	}
	req := &handlers.LastBlockNumberReq{}
	var resp handlers.LastBlockNumberResp
	call := c.connect.Go(LastBlockNumberMethod, req, &resp, nil)

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

	return resp.Number, err
}

// CurrentBlock returns raw current transactions block.
func (c *Client) CurrentBlock() (block []byte, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &handlers.CurrentBlockReq{}
	var resp *handlers.CurrentBlockResp
	call := c.connect.Go(CurrentBlockMethod, req, &resp, nil)

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

	return resp.Block, err
}

// SaveBlockToDB saves raw transactions block in database on server side.
func (c *Client) SaveBlockToDB(number uint64, raw []byte) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &handlers.SaveBlockToDBReq{
		Number: number,
		Block:  raw,
	}
	var resp *handlers.SaveBlockToDBResp
	call := c.connect.Go(SaveBlockToDBMethod, req, &resp, nil)

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

// InitBlock initializes new current transactions block on server side.
func (c *Client) InitBlock() error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &handlers.InitBlockReq{}
	var resp *handlers.InitBlockResp
	call := c.connect.Go(InitBlockMethod, req, &resp, nil)

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

// SaveCurrentBlock saves current transactions block in database on server side.
func (c *Client) SaveCurrentBlock(number uint64) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &handlers.SaveCurrentBlockReq{
		Number: number,
	}

	var resp *handlers.SaveCurrentBlockResp
	call := c.connect.Go(SaveCurrentBlockMethod, req, &resp, nil)

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

// GetTransactionsBlock gets and builds transactions block.
func (c *Client) GetTransactionsBlock(
	number uint64) (transactions.TxBlock, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &handlers.GetTransactionsBlockReq{
		Number: number,
	}

	var resp *handlers.GetTransactionsBlockResp
	call := c.connect.Go(GetTransactionsBlockMethod, req, &resp, nil)

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

	bl := transactions.NewBlock()
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
