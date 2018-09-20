package transport

// Smart Plasma RPC Methods.
const (
	// transfer methods
	DepositMethod           = "SmartPlasma.Deposit"
	WithdrawMethod          = "SmartPlasma.Withdraw"
	StartExitMethod         = "SmartPlasma.StartExit"
	AcceptTransactionMethod = "SmartPlasma.AcceptTransaction"
	AddCheckpointMethod     = "SmartPlasma.AddCheckpoint"

	// proof methods
	CreateProofMethod           = "SmartPlasma.CreateProof"
	CreateUIDStateProofMethod   = "SmartPlasma.CreateUIDStateProof"
	VerifyTxProofMethod         = "SmartPlasma.VerifyTxProof"
	VerifyCheckpointProofMethod = "SmartPlasma.VerifyCheckpointProof"

	// transactor methods
	PendingCodeAtMethod   = "SmartPlasma.PendingCodeAt"
	PendingNonceAtMethod  = "SmartPlasma.PendingNonceAt"
	SuggestGasPriceMethod = "SmartPlasma.SuggestGasPrice"
	EstimateGasMethod     = "SmartPlasma.EstimateGas"
	WaitMinedMethod       = "SmartPlasma.WaitMined"

	// challenge methods
	ChallengeExitMethod                   = "SmartPlasma.ChallengeExit"
	ChallengeCheckpointMethod             = "SmartPlasma.ChallengeCheckpoint"
	RespondChallengeExitMethod            = "SmartPlasma.RespondChallengeExit"
	RespondCheckpointChallengeMethod      = "SmartPlasma.RespondCheckpointChallenge"
	RespondWithHistoricalCheckpointMethod = "SmartPlasma.RespondWithHistoricalCheckpoint"
	ChallengeExistsMethod                 = "SmartPlasma.ChallengeExists"
	CheckpointIsChallengeMethod           = "SmartPlasma.CheckpointIsChallenge"
	ChallengesLengthMethod                = "SmartPlasma.ChallengesLength"
	CheckpointChallengesLengthMethod      = "SmartPlasma.CheckpointChallengesLength"
	GetChallengeMethod                    = "SmartPlasma.GetChallenge"
	GetCheckpointChallengeMethod          = "SmartPlasma.GetCheckpointChallenge"

	// Plasma Cash blocks methods
	BuildBlockMethod           = "SmartPlasma.BuildBlock"
	SendBlockHashMethod        = "SmartPlasma.SendBlockHash"
	LastBlockNumberMethod      = "SmartPlasma.LastBlockNumber"
	CurrentBlockMethod         = "SmartPlasma.CurrentBlock"
	SaveBlockToDBMethod        = "SmartPlasma.SaveBlockToDB"
	InitBlockMethod            = "SmartPlasma.InitBlock"
	SaveCurrentBlockMethod     = "SmartPlasma.SaveCurrentBlock"
	GetTransactionsBlockMethod = "SmartPlasma.GetTransactionsBlock"
	ValidateBlockMethod        = "SmartPlasma.ValidateBlock"

	// checkpoints methods
	BuildCheckpointMethod            = "SmartPlasma.BuildCheckpoint"
	SendCheckpointHashMethod         = "SmartPlasma.SendCheckpointHash"
	CurrentCheckpointMethod          = "SmartPlasma.CurrentCheckpoint"
	SaveCheckpointToDBMethod         = "SmartPlasma.SaveCheckpointToDB"
	InitCheckpointMethod             = "SmartPlasma.InitCheckpoint"
	SaveCurrentCheckpointBlockMethod = "SmartPlasma.SaveCurrentCheckpointBlock"
	GetCheckpointsBlockMethod        = "SmartPlasma.GetCheckpointsBlock"

	// info methods
	DepositCountMethod    = "SmartPlasma.DepositCount"
	ChallengePeriodMethod = "SmartPlasma.ChallengePeriod"
	OperatorMethod        = "SmartPlasma.Operator"
	ChildChainMethod      = "SmartPlasma.ChildChain"
	ExitsMethod           = "SmartPlasma.Exits"
	WalletMethod          = "SmartPlasma.Wallet"
	Wallet2Method         = "SmartPlasma.Wallet2"
)
