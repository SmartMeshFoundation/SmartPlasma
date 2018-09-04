package service

import (
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
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
func (s *Service) Close() error {
	err := s.blockBase.Close()
	if err != nil {
		return err
	}
	return s.chptBase.Close()

}
