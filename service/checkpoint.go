package service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
)

func (s *Service) AddCheckpoint(uid, number *big.Int) error {
	return s.currentChpt.AddCheckpoint(uid, number)
}

func (s *Service) CreateCheckpointProof(uid *big.Int,
	chptHash common.Hash) ([]byte, error) {
	raw, err := s.RawCheckpointBlock(chptHash)
	if err != nil {
		return nil, err
	}

	blk := checkpoints.NewBlock()
	err = blk.Unmarshal(raw)
	if err != nil {
		return nil, err
	}

	return blk.CreateProof(uid), err
}

func (s *Service) InitCurrentCheckpoint() {
	s.currentChpt = checkpoints.NewBlock()
}

func (s *Service) BuildCheckpoint() (common.Hash, error) {
	return s.currentChpt.Build()
}

func (s *Service) RawCheckpointBlock(hash common.Hash) ([]byte, error) {
	return s.chptBase.Get(hash.Bytes())
}

func (s *Service) SaveCheckpointBlock(blk checkpoints.CheckpointBlock) error {
	raw, err := blk.Marshal()
	if err != nil {
		return err
	}

	return s.chptBase.Set(blk.Hash().Bytes(), raw)
}

func (s *Service) SendCurrentChptBlockHash(ctx context.Context) error {
	tx, err := s.session.NewCheckpoint(s.currentChpt.Hash()) // TODO: check hash
	if err != nil {
		return err
	}
	return s.mineTx(tx, ctx)
}
