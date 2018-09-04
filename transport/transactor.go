package transport

import (
	"context"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum"
	"github.com/SmartMeshFoundation/Spectrum/common"
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
	if err := c.connect.Call(PendingCodeAtMethod, req, &resp); err != nil {
		return nil, err
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
	if err := c.connect.Call(PendingNonceAtMethod, req, &resp); err != nil {
		return 0, err
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
	if err := c.connect.Call(SuggestGasPriceMethod, req, &resp); err != nil {
		return nil, err
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
	if err := c.connect.Call(EstimateGasMethod, req, &resp); err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Gas, nil
}
