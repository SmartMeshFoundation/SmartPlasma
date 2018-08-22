package service

import (
	"context"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/merkle"
)

// AcceptTransaction adds a transaction to current transactions block.
func (s *Service) AcceptTransaction(tx *transaction.Transaction) error {
	return s.currentBlock.AddTx(tx)
}

// CreateProof creates merkle proof for particular uid.
// Argument `block` is block number.
func (s *Service) CreateProof(uid *big.Int, block uint64) ([]byte, error) {
	raw, err := s.RawBlockFromDB(block)
	if err != nil {
		return nil, err
	}

	blk := transactions.NewTxBlock()
	err = buildBlockFromBytes(blk, raw)
	if err != nil {
		return nil, err
	}

	return blk.CreateProof(uid), err
}

// VerifyTxProof returns true if uid was spent in this block.
func (s *Service) VerifyTxProof(uid *big.Int, hash common.Hash,
	block uint64, proof []byte) (bool, error) {
	root, err := s.session.ChildChain(new(big.Int).SetUint64(block))
	if err != nil {
		return false, err
	}

	return merkle.CheckMembership(uid, hash, root, proof), err
}

// InitBlock initializes a new block.
func (s *Service) InitBlock() {
	s.currentBlock = transactions.NewTxBlock()
}

// BuildBlock build current Plasma block.
func (s *Service) BuildBlock() (common.Hash, error) {
	return s.currentBlock.Build()
}

// RawBlockFromDB returns raw Plasma block from database.
func (s *Service) RawBlockFromDB(number uint64) ([]byte, error) {
	return s.blockBase.Get(strconv.AppendUint(nil, number, 10))
}

// SaveBlockToDB saves Plasma Block to database.
func (s *Service) SaveBlockToDB(number uint64,
	blk transactions.TxBlock) error {
	raw, err := blk.Marshal()
	if err != nil {
		return err
	}
	return s.blockBase.Set(strconv.AppendUint(nil, number, 10), raw)
}

// SendBlockHash sends a Plasma block hash to the blockchain.
func (s *Service) SendBlockHash(ctx context.Context, hash common.Hash) error {
	tx, err := s.session.NewBlock(hash)
	if err != nil {
		return err
	}
	return s.mineTx(ctx, tx)
}

// LastBlockNumber gets last block number from blockchain.
func (s *Service) LastBlockNumber() (*big.Int, error) {
	return s.session.BlockNumber()
}

// CurrentBlock returns current Plasma block.
func (s *Service) CurrentBlock() transactions.TxBlock {
	return s.currentBlock
}
