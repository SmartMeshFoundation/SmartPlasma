package handlers

import (
	"bytes"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
)

// Deposit invokes deposit method on Mediator contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) Deposit(req *RawReq, resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.MediatorTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// Withdraw invokes withdraw method
// on Mediator contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) Withdraw(req *RawReq, resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.MediatorTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// StartExit invokes startExit method
// on RootChain contract from a specific account.
// Function received raw signed Ethereum transaction.
func (api *SmartPlasma) StartExit(req *RawReq, resp *RawResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	if err := api.service.RootChainTransaction(
		ctx, req.RawTx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// AcceptTransaction accepts a raw transaction and returns a response.
func (api *SmartPlasma) AcceptTransaction(req *AcceptTransactionReq,
	resp *AcceptTransactionResp) error {
	tx := &transaction.Transaction{}

	if err := transaction.DecodeRLP(bytes.NewBuffer(req.Tx), tx); err != nil {
		resp.Error = err.Error()
		return nil
	}
	if err := api.service.AcceptTransaction(tx); err != nil {
		resp.Error = err.Error()
	}
	return nil
}

// AddCheckpoint accepts UID with transaction number for current checkpoint.
func (api *SmartPlasma) AddCheckpoint(req *AddCheckpointReq,
	resp *AddCheckpointResp) error {
	if err := api.service.AcceptUIDState(req.UID, req.Nonce); err != nil {
		resp.Error = err.Error()
	}
	return nil
}
