package service

import (
	"context"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
)

func (s *Service) AcceptTransaction(tx *transaction.Transaction) error {
	return s.currentBlock.AddTx(tx)
}

func (s *Service) CreateTxProof(uid *big.Int, block uint64) ([]byte, error) {
	raw, err := s.RawTxBlock(block)
	if err != nil {
		return nil, err
	}

	blk := transactions.NewTxBlock()
	err = blk.Unmarshal(raw)
	if err != nil {
		return nil, err
	}
	_, err = blk.Build()
	if err != nil {
		return nil, err
	}

	return blk.CreateProof(uid), err
}

func (s *Service) InitNewTxBlock() {
	s.currentBlock = transactions.NewTxBlock()
}

func (s *Service) BuildTxBlock() (common.Hash, error) {
	return s.currentBlock.Build()
}

func (s *Service) RawTxBlock(number uint64) ([]byte, error) {
	return s.blockBase.Get(strconv.AppendUint(nil, number, 10))
}

func (s *Service) SaveTxBlock(number uint64, blk transactions.TxBlock) error {
	raw, err := blk.Marshal()
	if err != nil {
		return err
	}
	return s.blockBase.Set(strconv.AppendUint(nil, number, 10), raw)
}

func (s *Service) SaveCurrentTxBlock() error {
	hash := s.currentBlock.Hash()
	if (hash == common.Hash{}) {
		return errors.New(" hash of block is empty")
	}

	return s.SaveTxBlock(s.getBlockNumber(), s.currentBlock)
}

func (s *Service) SendCurrentTxBlock(ctx context.Context) error {
	hash := s.currentBlock.Hash()
	if (hash == common.Hash{}) {
		return errors.New(" hash of block is empty")
	}

	err := s.SendTxBlock(ctx, hash)
	if err != nil {
		return err
	}

	s.setBlockNumber(s.getBlockNumber() + 1)
	return nil
}

func (s *Service) SendTxBlock(ctx context.Context, hash common.Hash) error {
	tx, err := s.session.NewBlock(hash)
	if err != nil {
		return err
	}
	err = s.mineTx(tx, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LastTxBlockNumber() (*big.Int, error) {
	return s.session.BlockNumber()
}

func (s *Service) CurrentTxBlockHash() common.Hash {
	return s.currentBlock.Hash()
}

func (s *Service) SyncCurrentTxBlockNumber() error {
	blk, err := s.session.BlockNumber()
	if err != nil {
		return err
	}

	s.setBlockNumber(blk.Uint64() + 1)
	return nil
}

func (s *Service) getBlockNumber() uint64 {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.blockNumber
}

func (s *Service) setBlockNumber(number uint64) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.blockNumber = number
}
