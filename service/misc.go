package service

import (
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

// DepositCount returns a deposit counter.
func (s *Service) DepositCount() (*big.Int, error) {
	return s.session.DepositCount()
}

// ChallengePeriod returns a period for challenging in seconds.
func (s *Service) ChallengePeriod() (*big.Int, error) {
	return s.session.ChallengePeriod()
}

// Operator returns a Plasma Cash operator address.
func (s *Service) Operator() (common.Address, error) {
	return s.session.Operator()
}

// ChildChain returns a block hash by a block number.
func (s *Service) ChildChain(key *big.Int) (common.Hash, error) {
	return s.session.ChildChain(key)
}

// Exits returns a incomplete exit by UID.
func (s *Service) Exits(key *big.Int) (struct {
	State                *big.Int
	ExitTime             *big.Int
	ExitTxBlkNum         *big.Int
	ExitTx               []byte
	TxBeforeExitTxBlkNum *big.Int
	TxBeforeExitTx       []byte
}, error) {
	return s.session.Exits(key)
}

// Wallet returns a deposit amount.
func (s *Service) Wallet(uid *big.Int) (*big.Int, error) {
	return s.session.Wallet(common.BigToHash(uid))
}

// ChallengeExists if this is true,
// that a exit is blocked by a transaction of challenge.
func (s *Service) ChallengeExists(
	uid *big.Int, challengeTx []byte) (bool, error) {
	return s.session.ChallengeExists(uid, challengeTx)
}

// CheckpointIsChallenge if this is true,
// that a checkpoint is blocked by a transaction of challenge.
func (s *Service) CheckpointIsChallenge(
	uid *big.Int, checkpoint common.Hash, challengeTx []byte) (bool, error) {
	return s.session.CheckpointIsChallenge(uid, checkpoint, challengeTx)
}

// ChallengesLength returns number of disputes on withdrawal of uid.
func (s *Service) ChallengesLength(uid *big.Int) (*big.Int, error) {
	return s.session.ChallengesLength(uid)
}

// CheckpointChallengesLength returns number of disputes
// for checkpoint by a uid.
func (s *Service) CheckpointChallengesLength(uid *big.Int,
	checkpoint common.Hash) (*big.Int, error) {
	return s.session.CheckpointChallengesLength(uid, checkpoint)
}

// GetChallenge returns exit challenge transaction by uid and index.
func (s *Service) GetChallenge(uid *big.Int, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	return s.session.GetChallenge(uid, index)
}

// GetCheckpointChallenge Returns checkpoint challenge transaction
// by checkpoint merkle root, uid and index.
func (s *Service) GetCheckpointChallenge(
	uid *big.Int, checkpoint common.Hash, index *big.Int) (struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
}, error) {
	return s.session.GetCheckpointChallenge(uid, checkpoint, index)
}
