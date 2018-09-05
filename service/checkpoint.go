package service

import (
	"context"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/merkle"
)

// AcceptUIDState accept uid with transaction number for current checkpoint.
func (s *Service) AcceptUIDState(uid, number *big.Int) error {
	return s.currentChpt.AddCheckpoint(uid, number)
}

// CreateUIDStateProof creates merkle proof for particular uid.
// Argument `chptHash` is checkpoint hash.
func (s *Service) CreateUIDStateProof(uid *big.Int,
	chptHash common.Hash) ([]byte, error) {
	raw, err := s.RawCheckpointFromDB(chptHash)
	if err != nil {
		return nil, err
	}

	blk := checkpoints.NewBlock()
	err = buildBlockFromBytes(blk, raw)
	if err != nil {
		return nil, err
	}

	return blk.CreateProof(uid), err
}

// CurrentCheckpoint returns current checkpoint.
func (s *Service) CurrentCheckpoint() checkpoints.CheckpointBlock {
	return s.currentChpt
}

// InitCheckpoint initializes a new checkpoint.
func (s *Service) InitCheckpoint() {
	s.currentChpt = checkpoints.NewBlock()
}

// BuildCheckpoint build current Checkpoint block.
func (s *Service) BuildCheckpoint() (common.Hash, error) {
	return s.currentChpt.Build()
}

// RawCheckpointFromDB returns raw Checkpoint block from database.
func (s *Service) RawCheckpointFromDB(hash common.Hash) ([]byte, error) {
	return s.chptBase.Get(hash.Bytes())
}

// SaveCheckpointToDB saves Checkpoint Block to database.
func (s *Service) SaveCheckpointToDB(chpt checkpoints.CheckpointBlock) error {
	raw, err := chpt.Marshal()
	if err != nil {
		return err
	}

	return s.chptBase.Set(chpt.Hash().Bytes(), raw)
}

// SendChptHash sends a Checkpoint block hash to the blockchain.
func (s *Service) SendChptHash(
	ctx context.Context, hash common.Hash) (*types.Transaction, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.NewCheckpoint(hash)
}

// IsValidCheckpoint returns true if the uid is fixed at the checkpoint
// and the number is correct.
func (s *Service) IsValidCheckpoint(ctx context.Context,
	uid *big.Int, number *big.Int, checkpoint common.Hash,
	proof []byte) (bool, error) {
	session := rootchain.CopySession(s.session)
	session.CallOpts.Context = ctx

	challenges, err := session.CheckpointChallengesLength(uid, checkpoint)
	if err != nil {
		return false, err
	}
	if challenges.Int64() > 0 {
		return false, nil
	}

	return merkle.CheckMembership(
		uid, common.BigToHash(number), checkpoint, proof), nil
}
