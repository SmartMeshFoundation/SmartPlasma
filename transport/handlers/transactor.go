package handlers

import "github.com/SmartMeshFoundation/Spectrum/core/types"

// PendingCodeAt returns the code of the given Account in the pending state.
func (api *SmartPlasma) PendingCodeAt(req *PendingCodeAtReq,
	resp *PendingCodeAtResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	code, err := api.service.PendingCodeAt(ctx, req.Account)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Code = code
	return nil
}

// PendingNonceAt retrieves the current pending nonce
// associated with an Account.
func (api *SmartPlasma) PendingNonceAt(req *PendingNonceAtReq,
	resp *PendingNonceAtResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	nonce, err := api.service.PendingNonceAt(ctx, req.Account)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Nonce = nonce
	return nil
}

// SuggestGasPrice retrieves the currently suggested gas price
// to allow a timely execution of a transaction.
func (api *SmartPlasma) SuggestGasPrice(req *SuggestGasPriceReq,
	resp *SuggestGasPriceResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	price, err := api.service.SuggestGasPrice(ctx)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Price = price
	return nil
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.
func (api *SmartPlasma) EstimateGas(req *EstimateGasReq,
	resp *EstimateGasResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	gas, err := api.service.EstimateGas(ctx, req.Call)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}
	resp.Gas = gas
	return nil
}

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func (api *SmartPlasma) WaitMined(req *WaitMinedReq,
	resp *WaitMinedResp) error {
	tx := &types.Transaction{}
	err := tx.UnmarshalJSON(req.Tx)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	ctx, cancel := api.newContext()
	defer cancel()

	tr, err := api.service.Mine(ctx, tx)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	raw, err := tr.MarshalJSON()
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	resp.Tr = raw
	return nil
}
