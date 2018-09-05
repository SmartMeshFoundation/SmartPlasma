package service

import (
	"context"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"

	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

// DepositCount returns a deposit counter.
func (s *Service) DepositCount(ctx context.Context) (*big.Int, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.DepositCount()
}

// ChallengePeriod returns a period for challenging in seconds.
func (s *Service) ChallengePeriod(ctx context.Context) (*big.Int, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.ChallengePeriod()
}

// Operator returns a Plasma Cash operator address.
func (s *Service) Operator(ctx context.Context) (common.Address, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.Operator()
}

// ChildChain returns a block hash by a block number.
func (s *Service) ChildChain(
	ctx context.Context, key *big.Int) (common.Hash, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.ChildChain(key)
}

// Exits returns a incomplete exit by UID.
func (s *Service) Exits(ctx context.Context, key *big.Int) (struct {
	State                *big.Int
	ExitTime             *big.Int
	ExitTxBlkNum         *big.Int
	ExitTx               []byte
	TxBeforeExitTxBlkNum *big.Int
	TxBeforeExitTx       []byte
}, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.Exits(key)
}

// Wallet returns a deposit amount.
func (s *Service) Wallet(ctx context.Context, uid *big.Int) (*big.Int, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.Wallet(common.BigToHash(uid))
}

// ChallengeExists if this is true,
// that a exit is blocked by a transaction of challenge.
func (s *Service) ChallengeExists(
	ctx context.Context, uid *big.Int, challengeTx []byte) (bool, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.ChallengeExists(uid, challengeTx)
}

// CheckpointIsChallenge if this is true,
// that a checkpoint is blocked by a transaction of challenge.
func (s *Service) CheckpointIsChallenge(ctx context.Context,
	uid *big.Int, checkpoint common.Hash, challengeTx []byte) (bool, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.CheckpointIsChallenge(uid, checkpoint, challengeTx)
}

// ChallengesLength returns number of disputes on withdrawal of uid.
func (s *Service) ChallengesLength(
	ctx context.Context, uid *big.Int) (*big.Int, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.ChallengesLength(uid)
}

// CheckpointChallengesLength returns number of disputes
// for checkpoint by a uid.
func (s *Service) CheckpointChallengesLength(ctx context.Context,
	uid *big.Int, checkpoint common.Hash) (*big.Int, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.CheckpointChallengesLength(uid, checkpoint)
}

// GetChallenge returns exit challenge transaction by uid and index.
func (s *Service) GetChallenge(
	ctx context.Context, uid *big.Int, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.GetChallenge(uid, index)
}

// GetCheckpointChallenge Returns checkpoint challenge transaction
// by checkpoint merkle root, uid and index.
func (s *Service) GetCheckpointChallenge(ctx context.Context,
	uid *big.Int, checkpoint common.Hash, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx
	return session.GetCheckpointChallenge(uid, checkpoint, index)
}
