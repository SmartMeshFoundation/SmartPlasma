package service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block"
)

func (s *Service) mineTx(ctx context.Context, tx *types.Transaction) error {
	tr, err := s.backend.Mine(ctx, tx)
	if err != nil {
		return err
	}

	if tr.Status == types.ReceiptStatusFailed {
		return errors.New("transaction execution failed")
	}
	return nil
}

// MediatorTransaction decode and sends Ethereum transaction
// to mediator contract.
// methods: deposit, withdraw.
func (s *Service) MediatorTransaction(rawTx []byte) error {
	tx, err := s.mediatorContractWrapper.UnmarshalTransaction(rawTx)
	if err != nil {
		return err
	}
	return s.transact(tx)
}

// RootChainTransaction decode and sends Ethereum transaction
// to root chain contract.
// methods: startExit, challengeExit, challengeCheckpoint,
// respondChallengeExit, respondCheckpointChallenge,
// respondWithHistoricalCheckpoint.
func (s *Service) RootChainTransaction(rawTx []byte) error {
	tx, err := s.rootChainContractWrapper.UnmarshalTransaction(rawTx)
	if err != nil {
		return err
	}
	return s.transact(tx)
}

func buildBlockFromBytes(blk block.Block, raw []byte) error {
	err := blk.Unmarshal(raw)
	if err != nil {
		return err
	}

	_, err = blk.Build()
	return err
}

func (s *Service) transact(tx *types.Transaction) error {
	return s.backend.Connect().SendTransaction(context.Background(), tx)
}

// PendingCodeAt returns the code of the given account in the pending state.
func (s *Service) PendingCodeAt(ctx context.Context,
	account common.Address) ([]byte, error) {
	return s.backend.Connect().PendingCodeAt(ctx, account)
}

// PendingNonceAt retrieves the current pending nonce associated with an account.
func (s *Service) PendingNonceAt(ctx context.Context,
	account common.Address) (uint64, error) {
	return s.backend.Connect().PendingNonceAt(ctx, account)
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (s *Service) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return s.backend.Connect().SuggestGasPrice(ctx)
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.
func (s *Service) EstimateGas(ctx context.Context,
	call ethereum.CallMsg) (gas uint64, err error) {
	return s.backend.Connect().EstimateGas(ctx, call)
}
