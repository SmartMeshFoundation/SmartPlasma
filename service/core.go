package service

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database"
)

type Service struct {
	currentBlock transactions.TxBlock
	currentChpt  checkpoints.CheckpointBlock
	blockBase    database.Database
	chptBase     database.Database
	session      *rootchain.RootChainSession
	backend      backend.Backend

	mtx         sync.Mutex
	blockNumber uint64
}

func NewService(session *rootchain.RootChainSession, backend backend.Backend,
	blockBase, chptBase database.Database) (*Service, error) {
	blk, err := session.BlockNumber()
	if err != nil {
		return nil, err
	}

	return &Service{
		currentChpt:  checkpoints.NewBlock(),
		currentBlock: transactions.NewTxBlock(),
		blockBase:    blockBase,
		chptBase:     chptBase,
		session:      session,
		backend:      backend,
		mtx:          sync.Mutex{},
		blockNumber:  blk.Uint64(),
	}, nil
}

func (s *Service) mineTx(tx *types.Transaction, ctx context.Context) error {
	tr, err := s.backend.Mine(tx, ctx)
	if err != nil {
		return err
	}

	if tr.Status == types.ReceiptStatusFailed {
		return errors.New("transaction execution failed")
	}
	return nil
}
