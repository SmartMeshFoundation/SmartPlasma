package service

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database"
)

// Service implements PlasmaCash methods.
type Service struct {
	currentBlock             transactions.TxBlock
	currentChpt              checkpoints.CheckpointBlock
	blockBase                database.Database
	chptBase                 database.Database
	session                  *rootchain.RootChainSession
	backend                  backend.Backend
	rootChainContractWrapper *build.Contract
	mediatorContractWrapper  *build.Contract
}

// NewService creates new PlasmaCash service.
func NewService(session *rootchain.RootChainSession, backend backend.Backend,
	blockBase, chptBase database.Database,
	rootChainContractWrapper *build.Contract,
	mediatorContractWrapper *build.Contract) *Service {

	return &Service{
		currentChpt:              checkpoints.NewBlock(),
		currentBlock:             transactions.NewTxBlock(),
		blockBase:                blockBase,
		chptBase:                 chptBase,
		session:                  session,
		backend:                  backend,
		rootChainContractWrapper: rootChainContractWrapper,
		mediatorContractWrapper:  mediatorContractWrapper,
	}
}

// Close stops service.
func (s *Service) Close() {
	s.blockBase.Close()
	s.chptBase.Close()
}

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
