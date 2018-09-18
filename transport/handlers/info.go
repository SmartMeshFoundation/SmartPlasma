package handlers

// DepositCount returns a deposit counter.
func (api *SmartPlasma) DepositCount(
	req *DepositCountReq, resp *DepositCountResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	count, err := api.service.DepositCount(ctx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Count = count
	return nil
}

// ChallengePeriod returns a period for challenging in seconds.
func (api *SmartPlasma) ChallengePeriod(
	req *ChallengePeriodReq, resp *ChallengePeriodResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	secs, err := api.service.ChallengePeriod(ctx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.ChallengePeriod = secs
	return nil
}

// Operator returns a Plasma Cash operator address.
func (api *SmartPlasma) Operator(req *OperatorReq, resp *OperatorResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	operator, err := api.service.Operator(ctx)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Operator = operator
	return nil
}

// ChildChain returns a block hash by a block number.
func (api *SmartPlasma) ChildChain(
	req *ChildChainReq, resp *ChildChainResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	hash, err := api.service.ChildChain(ctx, req.BlockNumber)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.BlockHash = hash
	return nil
}

// Wallet returns a deposit amount.
func (api *SmartPlasma) Wallet(
	req *WalletReq, resp *WalletResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	amount, err := api.service.Wallet(ctx, req.UID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Amount = amount
	return nil
}

// Exits returns a incomplete exit by UID.
func (api *SmartPlasma) Exits(req *ExitsReq, resp *ExitsResp) error {
	ctx, cancel := api.newContext()
	defer cancel()

	result, err := api.service.Exits(ctx, req.UID)
	if err != nil {
		resp.Error = err.Error()
	}
	resp.State = result.State
	resp.ExitTime = result.ExitTime
	resp.ExitTxBlkNum = result.ExitTxBlkNum
	resp.ExitTx = result.ExitTx
	resp.TxBeforeExitTxBlkNum = result.TxBeforeExitTxBlkNum
	resp.TxBeforeExitTx = result.TxBeforeExitTx
	return nil
}
