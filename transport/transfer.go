package transport

import (
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

// Deposit transacts deposit function from Mediator contract.
func (c *Client) Deposit(currency common.Address,
	amount *big.Int) (tx *types.Transaction, err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionMediator != nil {
		session := mediator.CopySession(c.sessionMediator)
		session.TransactOpts.Context = ctx
		return session.Deposit(currency, amount)
	}

	if c.med == nil {
		return nil, ErrTransactor
	}

	tx, err = c.med.Transaction(c.opts.TransactOpts,
		"deposit", currency, amount)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}
	var resp RawResp
	call := c.connect.Go(DepositMethod, req, &resp, nil)

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

// Withdraw transacts withdraw function from Mediator contract.
func (c *Client) Withdraw(prevTx, prevTxProof []byte, prevTxBlkNum *big.Int,
	txRaw, txProof []byte, txBlkNum *big.Int) (*types.Transaction, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionMediator != nil {
		session := mediator.CopySession(c.sessionMediator)
		session.TransactOpts.Context = ctx
		return session.Withdraw(prevTx, prevTxProof, prevTxBlkNum,
			txRaw, txProof, txBlkNum)
	}

	if c.med == nil {
		return nil, ErrTransactor
	}

	tx, err := c.med.Transaction(c.opts.TransactOpts,
		"withdraw", prevTx, prevTxProof, prevTxBlkNum, txRaw,
		txProof, txBlkNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}
	var resp RawResp
	call := c.connect.Go(WithdrawMethod, req, &resp, nil)

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

// StartExit transacts startExit function from RootChain contract.
func (c *Client) StartExit(previousTx, previousTxProof []byte,
	previousTxBlockNum *big.Int, lastTx, lastTxProof []byte,
	lastTxBlockNum *big.Int) (*types.Transaction, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	if c.sessionRootChain != nil {
		session := rootchain.CopySession(c.sessionRootChain)
		session.TransactOpts.Context = ctx
		return session.StartExit(previousTx,
			previousTxProof, previousTxBlockNum, lastTx,
			lastTxProof, lastTxBlockNum)
	}

	if c.root == nil {
		return nil, ErrTransactor
	}

	tx, err := c.root.Transaction(c.opts.TransactOpts,
		"startExit", previousTx, previousTxProof, previousTxBlockNum,
		lastTx, lastTxProof, lastTxBlockNum)
	if err != nil {
		return nil, err
	}
	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &RawReq{
		RawTx: raw,
	}
	var resp RawResp
	call := c.connect.Go(StartExitMethod, req, &resp, nil)

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

// AcceptTransaction sends raw transaction to PlasmaCash RPC server.
func (c *Client) AcceptTransaction(rawTx []byte) (err error) {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &AcceptTransactionReq{Tx: rawTx}
	var resp *AcceptTransactionResp
	call := c.connect.Go(AcceptTransactionMethod, req, &resp, nil)

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

// AddCheckpoint sends UID and current transaction nonce
// for inclusion in a checkpoint.
func (c *Client) AddCheckpoint(uid, nonce *big.Int) error {
	ctx, cancel := c.newContext()
	defer cancel()

	req := &AddCheckpointReq{
		UID:   uid,
		Nonce: nonce,
	}
	var resp *AddCheckpointResp
	call := c.connect.Go(AddCheckpointMethod, req, &resp, nil)

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
