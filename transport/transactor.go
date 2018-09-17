package transport

import (
	"context"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/pkg/errors"
)

// PendingCodeAt returns the code of the given Account
// in the pending state.
func (c *Client) PendingCodeAt(
	ctx context.Context, account common.Address) ([]byte, error) {
	req := &PendingCodeAtReq{
		Account: account,
	}
	var resp PendingCodeAtResp
	call := c.connect.Go(PendingCodeAtMethod, req, &resp, nil)

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

	return resp.Code, nil
}

// PendingNonceAt retrieves the current pending nonce
// associated with an Account.
func (c *Client) PendingNonceAt(
	ctx context.Context, account common.Address) (uint64, error) {
	req := &PendingNonceAtReq{
		Account: account,
	}
	var resp PendingNonceAtResp
	call := c.connect.Go(PendingNonceAtMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return 0, replay.Error
		}
	case <-ctx.Done():
		return 0, errors.New("timeout")
	}

	if resp.Error != "" {
		return 0, errors.New(resp.Error)
	}

	return resp.Nonce, nil
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	req := &SuggestGasPriceReq{}

	var resp SuggestGasPriceResp
	call := c.connect.Go(SuggestGasPriceMethod, req, &resp, nil)

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

	return resp.Price, nil
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.
func (c *Client) EstimateGas(
	ctx context.Context, call ethereum.CallMsg) (gas *big.Int, err error) {
	req := &EstimateGasReq{
		Call: call,
	}

	var resp EstimateGasResp
	call2 := c.connect.Go(EstimateGasMethod, req, &resp, nil)

	select {
	case replay := <-call2.Done:
		if replay.Error != nil {
			return nil, replay.Error
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Gas, nil
}

// WaitMined to wait mining.
func (c *Client) WaitMined(
	ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	if c.backend != nil {
		return c.backend.Mine(ctx, tx)
	}

	raw, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := &WaitMinedReq{
		Tx: raw,
	}

	var resp WaitMinedResp
	call := c.connect.Go(WaitMinedMethod, req, &resp, nil)

	select {
	case replay := <-call.Done:
		if replay.Error != nil {
			return nil, err
		}
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	tr := &types.Receipt{}
	err = tr.UnmarshalJSON(resp.Tr)
	if err != nil {
		return nil, err
	}

	return tr, nil
}
