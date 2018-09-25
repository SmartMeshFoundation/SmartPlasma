package handlers

import (
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum"
	"github.com/SmartMeshFoundation/Spectrum/common"
)

// AcceptTransactionReq is request for send Plasma transaction to PRC server.
type AcceptTransactionReq struct {
	Tx []byte
}

// AcceptTransactionResp is response for send Plasma transaction to PRC server.
type AcceptTransactionResp struct {
	Error string
}

// CreateProofReq is request for CreateProof method.
type CreateProofReq struct {
	UID   *big.Int
	Block uint64
}

// CreateProofResp is response for CreateProof method.
type CreateProofResp struct {
	Proof []byte
	Error string
}

// AddCheckpointReq is request for AddCheckpoint method.
type AddCheckpointReq struct {
	UID         *big.Int
	Nonce       *big.Int
	BlockNumber uint64
}

// AddCheckpointResp is response for AddCheckpoint method.
type AddCheckpointResp struct {
	Error string
}

// CreateUIDStateProofReq is request for CreateUIDStateProof method.
type CreateUIDStateProofReq struct {
	UID            *big.Int
	CheckpointHash common.Hash
}

// CreateUIDStateProofResp is response for CreateUIDStateProof method.
type CreateUIDStateProofResp struct {
	Nonce *big.Int
	Proof []byte
	Error string
}

// PendingCodeAtReq is request for PendingCodeAt method.
type PendingCodeAtReq struct {
	Account common.Address
}

// PendingCodeAtResp is response for PendingCodeAt method.
type PendingCodeAtResp struct {
	Code  []byte
	Error string
}

// PendingNonceAtReq is request for PendingNonceAt method.
type PendingNonceAtReq struct {
	Account common.Address
}

// PendingNonceAtResp is response for PendingNonceAt method.
type PendingNonceAtResp struct {
	Nonce uint64
	Error string
}

// SuggestGasPriceReq is request for SuggestGasPrice method.
type SuggestGasPriceReq struct {
}

// SuggestGasPriceResp is response for SuggestGasPrice method.
type SuggestGasPriceResp struct {
	Price *big.Int
	Error string
}

// EstimateGasReq is request for EstimateGas method.
type EstimateGasReq struct {
	Call ethereum.CallMsg
}

// EstimateGasResp is response for EstimateGas method.
type EstimateGasResp struct {
	Gas   *big.Int
	Error string
}

// WaitMinedReq is request for WaitMined method.
type WaitMinedReq struct {
	Tx []byte
}

// WaitMinedResp is response for WaitMined method.
type WaitMinedResp struct {
	Tr    []byte
	Error string
}

// RawReq is request for methods that works raw transactions.
type RawReq struct {
	RawTx []byte
}

// RawResp is response for methods that works raw transactions.
type RawResp struct {
	Error string
}

// BuildBlockReq is request for BuildBlock method.
type BuildBlockReq struct{}

// BuildBlockResp is response for BuildBlock method.
type BuildBlockResp struct {
	Hash  common.Hash
	Error string
}

// BuildCheckpointReq is request for BuildCheckpoint method.
type BuildCheckpointReq struct{}

// BuildCheckpointResp is response for BuildCheckpoint method.
type BuildCheckpointResp struct {
	Hash  common.Hash
	Error string
}

// SendBlockHashReq is request for SendBlockHash method.
type SendBlockHashReq struct {
	Hash common.Hash
}

// SendBlockHashResp is response for SendBlockHash method.
type SendBlockHashResp struct {
	Tx    []byte
	Error string
}

// SendCheckpointHashReq is request for SendCheckpointHash method.
type SendCheckpointHashReq struct {
	Hash common.Hash
}

// SendCheckpointHashResp is response for SendCheckpointHash method.
type SendCheckpointHashResp struct {
	Tx    []byte
	Error string
}

// LastBlockNumberReq is request for LastBlockNumber method.
type LastBlockNumberReq struct{}

// LastBlockNumberResp is response for LastBlockNumber method.
type LastBlockNumberResp struct {
	Number *big.Int
	Error  string
}

// CurrentBlockReq is request for CurrentBlock method.
type CurrentBlockReq struct{}

// CurrentBlockResp is response for CurrentBlock method.
type CurrentBlockResp struct {
	Block []byte
	Error string
}

// CurrentCheckpointReq is request for CurrentCheckpoint method.
type CurrentCheckpointReq struct{}

// CurrentCheckpointResp is response for CurrentCheckpoint method.
type CurrentCheckpointResp struct {
	Checkpoint []byte
	Error      string
}

// SaveBlockToDBReq is request for SaveBlockToDB method.
type SaveBlockToDBReq struct {
	Number uint64
	Block  []byte
}

// SaveBlockToDBResp is response for SaveBlockToDB method.
type SaveBlockToDBResp struct {
	Error string
}

// SaveCheckpointToDBReq is request for SaveCheckpointToDB method.
type SaveCheckpointToDBReq struct {
	Block []byte
}

// SaveCheckpointToDBResp is response for SaveCheckpointToDB method.
type SaveCheckpointToDBResp struct {
	Error string
}

// InitBlockReq is request for InitBlock method.
type InitBlockReq struct {
}

// InitBlockResp is response for InitBlock method.
type InitBlockResp struct {
	Error string
}

// InitCheckpointReq is request for InitCheckpoint method.
type InitCheckpointReq struct {
}

// InitCheckpointResp is response for InitCheckpoint method.
type InitCheckpointResp struct {
	Error string
}

// VerifyTxProofReq is request for VerifyTxProof method.
type VerifyTxProofReq struct {
	UID   *big.Int
	Hash  common.Hash
	Block uint64
	Proof []byte
}

// VerifyTxProofResp is response for VerifyTxProof method.
type VerifyTxProofResp struct {
	Exists bool
	Error  string
}

// VerifyCheckpointProofReq is request for VerifyCheckpointProof method.
type VerifyCheckpointProofReq struct {
	UID        *big.Int
	Number     *big.Int
	Checkpoint common.Hash
	Proof      []byte
}

// VerifyCheckpointProofResp is response for VerifyCheckpointProof method.
type VerifyCheckpointProofResp struct {
	Exists bool
	Error  string
}

// DepositCountReq is request for DepositCount method.
type DepositCountReq struct{}

// DepositCountResp is response for DepositCount method.
type DepositCountResp struct {
	Count *big.Int
	Error string
}

// ChallengePeriodReq is request for ChallengePeriod method.
type ChallengePeriodReq struct{}

// ChallengePeriodResp is response for ChallengePeriod method.
type ChallengePeriodResp struct {
	ChallengePeriod *big.Int
	Error           string
}

// OperatorReq is request for Operator method.
type OperatorReq struct{}

// OperatorResp is response for Operator method.
type OperatorResp struct {
	Operator common.Address
	Error    string
}

// ChildChainReq is request for ChildChain method.
type ChildChainReq struct {
	BlockNumber *big.Int
}

// ChildChainResp is response for ChildChain method.
type ChildChainResp struct {
	BlockHash common.Hash
	Error     string
}

// ExitsReq is request for Exits method.
type ExitsReq struct {
	UID *big.Int
}

// ExitsResp is response for Exits method.
type ExitsResp struct {
	State                *big.Int
	ExitTime             *big.Int
	ExitTxBlkNum         *big.Int
	ExitTx               []byte
	TxBeforeExitTxBlkNum *big.Int
	TxBeforeExitTx       []byte
	Error                string
}

// WalletReq is request for Wallet method.
type WalletReq struct {
	UID *big.Int
}

// WalletResp is response for Wallet method.
type WalletResp struct {
	Amount *big.Int
	Error  string
}

// Wallet2Req is request for Wallet2 method.
type Wallet2Req struct {
	UID *big.Int
}

// Wallet2Resp is response for Wallet2 method.
type Wallet2Resp struct {
	BlockNumber *big.Int
	Error       string
}

// ChallengeExistsReq is request for ChallengeExists method.
type ChallengeExistsReq struct {
	UID         *big.Int
	ChallengeTx []byte
}

// ChallengeExistsResp is response for ChallengeExists method.
type ChallengeExistsResp struct {
	Exists bool
	Error  string
}

// CheckpointIsChallengeReq is request for CheckpointIsChallenge method.
type CheckpointIsChallengeReq struct {
	UID         *big.Int
	Checkpoint  common.Hash
	ChallengeTx []byte
}

// CheckpointIsChallengeResp is response for CheckpointIsChallenge method.
type CheckpointIsChallengeResp struct {
	Exists bool
	Error  string
}

// ChallengesLengthReq is request for ChallengesLength method.
type ChallengesLengthReq struct {
	UID *big.Int
}

// ChallengesLengthResp is response for ChallengesLength method.
type ChallengesLengthResp struct {
	Length *big.Int
	Error  string
}

// CheckpointChallengesLengthReq is request
// for CheckpointChallengesLength method.
type CheckpointChallengesLengthReq struct {
	UID        *big.Int
	Checkpoint common.Hash
}

// CheckpointChallengesLengthResp is response
// for CheckpointChallengesLength method.
type CheckpointChallengesLengthResp struct {
	Length *big.Int
	Error  string
}

// GetChallengeReq is request for GetChallenge method.
type GetChallengeReq struct {
	UID   *big.Int
	Index *big.Int
}

// GetChallengeResp is response for GetChallenge method.
type GetChallengeResp struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
	Error          string
}

// GetCheckpointChallengeReq is request for GetCheckpointChallenge method.
type GetCheckpointChallengeReq struct {
	UID        *big.Int
	Checkpoint common.Hash
	Index      *big.Int
}

// GetCheckpointChallengeResp is response for GetCheckpointChallenge method.
type GetCheckpointChallengeResp struct {
	ChallengeTx    []byte
	ChallengeBlock *big.Int
	Error          string
}

// SaveCurrentBlockReq is request for SaveCurrentBlock method.
type SaveCurrentBlockReq struct {
	Number uint64
}

// SaveCurrentBlockResp is response for SaveCurrentBlock method.
type SaveCurrentBlockResp struct {
	Error string
}

// SaveCurrentCheckpointBlockReq is request for SaveCurrentBlock method.
type SaveCurrentCheckpointBlockReq struct {
}

// SaveCurrentCheckpointBlockResp is response for SaveCurrentBlock method.
type SaveCurrentCheckpointBlockResp struct {
	Error string
}

// GetTransactionsBlockReq is request for GetTransactionsBlock method.
type GetTransactionsBlockReq struct {
	Number uint64
}

// GetTransactionsBlockResp is response for GetTransactionsBlock method.
type GetTransactionsBlockResp struct {
	Block []byte
	Error string
}

// GetCheckpointsBlockReq is request for GetCheckpointsBlock method.
type GetCheckpointsBlockReq struct {
	Hash common.Hash
}

// GetCheckpointsBlockResp is response for GetCheckpointsBlock method.
type GetCheckpointsBlockResp struct {
	Block []byte
	Error string
}

// ValidateBlockReq is request for ValidateBlock method.
type ValidateBlockReq struct {
}

// ValidateBlockResp is response for ValidateBlock method.
type ValidateBlockResp struct {
	Error string
}
