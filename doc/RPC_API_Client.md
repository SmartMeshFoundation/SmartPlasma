# RPC API Client Reference

***

## Challenge

### ChallengeExit

Invokes `challengeExit` method on RootChain contract from a specific account.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `[]byte` - raw Smart Plasma transaction that disputes an exit.
3. `[]byte` - proof of inclusion of the challenge transaction in a Smart Plasma block.
4. `*big.Int` - the number of the block in which the challenge transaction is included.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### ChallengeCheckpoint

Invokes `challengeCheckpoint` method on RootChain contract from a specific account.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `common.Hash` - hash for checkpoint block.
3. `[]byte` - proof of inclusion of a uid in a checkpoint block.
4. `*big.Int` - invalid transaction nonce, according to caller.
5. `[]byte` - raw last deposit transaction, according to caller.
6. `[]byte` - proof of inclusion of a last transaction (according to caller) in a Smart Plasma block.
7. `*big.Int` - the number of the block in which the last transaction (according to caller) is included.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### RespondChallengeExit

Invokes `respondChallengeExit` method on RootChain contract from a specific account.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `[]byte` - raw Smart Plasma transaction that disputes an exit.
3. `[]byte` - raw transaction that answers to a dispute transaction.
4. `[]byte` - proof of inclusion of the respond transaction in a Smart Plasma block.
5. `*big.Int` - the number of the block in which the respond transaction is included.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### RespondChallengeExit

Invokes `respondCheckpointChallenge` method on RootChain contract from a specific account.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `common.Hash` - hash for checkpoint block.
3. `[]byte` - raw transaction that disputes an exit.
4. `[]byte` - raw transaction that answers to a dispute transaction.
5. `[]byte` - proof of inclusion of the respond transaction in a Smart Plasma block.
6. `*big.Int` - the number of the block in which the respond transaction is included.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### RespondWithHistoricalCheckpoint

Invokes `respondWithHistoricalCheckpoint` method on RootChain contract from a specific account.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `common.Hash` - hash for checkpoint block.
3. `[]byte` -  proof of inclusion of the uid in a checkpoint block.
4. `common.Hash` - hash for historical checkpoint block. (historical checkpoint before challenge checkpoint)
5. `[]byte` - proof of inclusion of the uid in a historical checkpoint block.
6. `[]byte` - raw transaction that disputes a checkpoint.
7. `*big.Int` - transaction nonce which is more than challengeTx nonce. This nonce is present in the historical checkpoint.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### ChallengeExists

Checks if a request to exit this transaction is blocked.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `[]byte` - raw transaction that disputes an exit.

#### Returns

1. `bool` - if this is true, that a exit is blocked by a transaction of challenge.
2. `error` - standard error.

### CheckpointIsChallenge

Checks if a checkpoint is blocked by a transaction of challenge.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `common.Hash` - hash for checkpoint block.
3. `[]byte` - raw transaction that disputes a checkpoint.

#### Returns

1. `bool` - if this is true, that a checkpoint is blocked by a transaction of challenge.
2. `error` - standard error.

### ChallengesLength

Returns number of disputes on withdrawal of uid.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).

#### Returns

1. `*big.Int` - number of disputes on withdrawal of uid.
2. `error` - standard error.

### CheckpointChallengesLength

Returns number of disputes for checkpoint by a uid.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).

#### Returns

1. `*big.Int` - number of disputes for checkpoint by a uid.
2. `error` - standard error.

### GetChallenge

Returns exit challenge transaction by uid and index.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `*big.Int` - unique index of exit challenge transaction.

#### Returns

1. `struct`:
    - `[]byte` - raw transaction that disputes an exit.
    - `*big.Int` - the number of the block in which the challenge transaction is included.
    - `error` - error message.
2. `error` - standard error.

### GetCheckpointChallenge

Returns checkpoint challenge transaction by checkpoint hash, uid and index.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `common.Hash` - a checkpoint hash.
2. `*big.Int` - unique index of exit challenge transaction.

#### Returns

1. `struct`:
    - `[]byte` - raw transaction that disputes an exit.
    - `*big.Int` - the number of the block in which the challenge transaction is included.
    - `error` - error message.
2. `error` - standard error.

## Checkpoint

### BuildCheckpoint

Builds current checkpoint block on the server side.

#### Parameters

null

#### Returns

1. `common.Hash` - checkpoint block hash. 
2. `error` - standard error.

### SendCheckpointHash

Sends new checkpoints block hash to RootChain contract.

#### Parameters

1. `common.Hash` - checkpoint block hash. 

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### CurrentCheckpoint

Returns raw current checkpoints block.

#### Parameters

null

#### Returns

1. `[]byte` - raw checkpoint block.
2. `error` - standard error.

### SaveCheckpointToDB

Saves raw checkpoints block in database on server side.

#### Parameters

1. `[]byte` - raw checkpoint block.

#### Returns

1. `error` - standard error.

### InitCheckpoint

Initializes new current checkpoints block on server side.
 
#### Parameters

null

#### Returns

1. `error` - standard error.

### SaveCurrentCheckpointBlock

Saves current checkpoints block in database on server side.
 
#### Parameters

null

#### Returns

1. `error` - standard error.

### GetCheckpointsBlock

Gets and builds checkpoints block.
 
#### Parameters

1. `common.Hash` - checkpoint block hash. 

#### Returns

1. `checkpoints.CheckpointBlock` - checkpoint block object. 
2. `error` - standard error.

## Client

### NewClient

Creates new Smart Plasma client.
The Client must initialize `RemoteEthereumClient` or `DirectEthereumClient`.
 
#### Parameters

1. `uint64` - timeout for request.
2. `*account.PlasmaTransactOpts` - a collection of authorization data required to create a valid Smart Plasma transaction.

#### Returns

1. `*Client` - a RPC API Client instance.

### RemoteEthereumClient

Initializes work with remote ethereum client. Ethereum transactions are generated locally, signed locally, packaged and sent to the server.
 If this function is not called, then all transactions are sent directly to the Ethereum.
 
#### Parameters

1. `*build.Contract` - wrapper object that reflects a Root chain contract on the Spectrum network.
2. `*build.Contract` - wrapper object that reflects a Mediator contract on the Spectrum network.

#### Returns

null

### DirectEthereumClient

Initializes work with direct Spectrum client.

#### Parameters

1. `bind.TransactOpts` - a collection of authorization data required to create a valid Spectrum transaction.
2. `common.Address` - Mediator contract address.
3. `common.Address` - Root chain contract address.
4. `backend.Backend` - backend interface for interaction with a Spectrum network.

#### Returns

null

### Connect

Tries to connect to a Smart Plasma RPC server.

#### Parameters

1. `string` - a network address or hostname for connection Smart Plasma RPC server.
2. `uint16` - a network port for connection Smart Plasma RPC server.

#### Returns

1. `error` - standard error.

### ConnectString

Tries to connect to a Smart Plasma RPC server.

#### Parameters

1. `string` - a network address or hostname and port for connection Smart Plasma RPC server. Format: `hostname:1234`

#### Returns

1. `error` - standard error.

### Close

Closes connection to Smart Plasma RPC server.

#### Parameters

null

#### Returns

1. `error` - standard error.

### Opts

Returns Smart Plasma transact options.

#### Parameters

null

#### Returns

1. `*account.PlasmaTransactOpts` - a collection of authorization data required to create a valid Smart Plasma transaction.

## Info

### DepositCount

Returns a deposit counter from Root chain contract.

#### Parameters

null

#### Returns

1. `*big.Int` - a deposit counter.
2. `error` - standard error.

## ChallengePeriod

### DepositCount

Returns a period for challenging in seconds.

#### Parameters

null

#### Returns

1. `*big.Int` - a period for challenging in seconds.
2. `error` - standard error.

### Operator

Returns a Smart Plasma operator address.

#### Parameters

null

#### Returns

1. `common.Address` - operator address.
2. `error` - standard error.

### ChildChain

Returns a block hash by a block number.

#### Parameters

1. `*big.Int` - block number.

#### Returns

1. `common.Hash` - block hash.
2. `error` - standard error.

### Exits

Returns a incomplete exit by UID.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).

#### Returns

1. `struct`:
    - `*big.Int` - state of exit. 0 - did not request to exit, 1 - in challenge proceeding, it blocks a exit, 2 - in anticipation of exit, 3 - a exit was made.
    - `*big.Int` - unix timestamp of start exit.
    - `*big.Int` - block number of last transaction.
    - `[]byte` - decoded last Smart Plasma transaction.
    - `*big.Int` - block number of penultimate transaction.
    - `[]byte` - decoded penultimate Smart Plasma transaction.
    - `string` - error message.
2. `error` - standard error.

### Wallet

Returns a deposit amount.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).

#### Returns

1. `*big.Int` - deposit amount.
2. `error` - standard error.

### Wallet2

Returns a deposit Smart Plasma block number.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).

#### Returns

1. `*big.Int` - block number.
2. `error` - standard error.

## Plasma Block

### BuildBlock

Builds current transactions block on the server side.

#### Parameters

null

#### Returns

1. `common.Haash` - block hash.
2. `error` - standard error.

### SendBlockHash

Sends new transactions block hash to RootChain contract.

#### Parameters

1. `common.Haash` - block hash.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### LastBlockNumber

Returns last transactions block number from RootChain contract.

#### Parameters

null

#### Returns

1. `*big.Int` - block number.
2. `error` - standard error.

### CurrentBlock

Returns raw current transactions block.

#### Parameters

null

#### Returns

1. `*big.Int` - raw Smart Plasma block.
2. `error` - standard error.

### SaveBlockToDB

Saves raw transactions block in database on server side.

#### Parameters

1. `uint64` - Smart Plasma block number.
2. `[]byte` - raw Smart Plasma block.

#### Returns

1. `error` - standard error.

### InitBlock

Initializes new current transactions block on server side.

#### Parameters

null

#### Returns

1. `error` - standard error.

### SaveCurrentBlock

Saves current transactions block in database on server side.

#### Parameters

1. `uint64` - Smart Plasma block number.

#### Returns

1. `error` - standard error.

### GetTransactionsBlock

Gets and builds transactions block.

#### Parameters

1. `uint64` - Smart Plasma block number.

#### Returns

1. `transactions.TxBlock` - Smart Plasma transactions block object.
2. `error` - standard error.

### ValidateBlock

Validates current block and remove bad transactions.

#### Parameters

null

#### Returns

1. `error` - standard error.

## Proof

### CreateProof

Sends UID and Block number to Smart Plasma RPC server. Returns merkle Proof for a UID.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `uint64` - Smart Plasma block number.

#### Returns

1. `[]byte` - a proof for UID.
2. `error` - standard error.

### VerifyTxProof

Checks whether the transaction is included in the transactions block.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `common.Hash` - a Smart Plasma transaction hash.
3. `uint64` - Smart Plasma block number.
4. `[]byte` - the proof for a Smart Plasma transaction.

#### Returns

1. `bool` - if it is `true`, then transaction is included in the transactions block.
2. `error` - standard error.

### CreateUIDStateProof

sends UID and checkpoint Hash to Smart Plasma RPC server. Returns merkle Proof for a UID.
#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `common.Hash` - a checkpoint hash.

#### Returns

1. `[]byte` - a proof for UID.
2. `error` - standard error.

### VerifyCheckpointProof

Checks whether the UID is included in the checkpoints block.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `*big.Int` - a nonce.
3. `common.Hash` - a checkpoint transaction hash.
4. `[]byte` - the proof for a checkpoint.

#### Returns

1. `bool` - if it is `true`, then nonce for UID is included in the checkpoint block.
2. `error` - standard error.

## Transactor

### PendingCodeAt

Returns the code of the given Account in the pending state.

#### Parameters

1. `context.Context` - golang context.
2. `common.Address` - a Spectrum address.

#### Returns

1. `[]byte` - the code of the given Account in the pending state.
2. `error` - standard error.

### PendingNonceAt

Retrieves the current pending nonce (not Smart Plasma nonce) associated with an Spectrum Account.
#### Parameters

1. `context.Context` - golang context.
2. `common.Address` - a Spectrum address.

#### Returns

1. `uint64` - nonce for Spectrum Account.
2. `error` - standard error.

### SuggestGasPrice

Retrieves the currently suggested gas price to allow a timely execution of a transaction.

#### Parameters

1. `context.Context` - golang context.

#### Returns

1. `*big.Int` - gas price.
2. `error` - standard error.

### EstimateGas

Tries to estimate the gas needed to execute a specific Spectrum transaction based on the current pending state of the backend blockchain.

#### Parameters

1. `context.Context` - golang context.
2. `ethereum.CallMsg` - message to Spectrum blockchain.

#### Returns

1. `*big.Int` - gas amount.
2. `error` - standard error.

### WaitMined

Waits mining of a block with the Spectrum transaction.

#### Parameters

1. `context.Context` - golang context.
2. `*types.Transaction` - Spectrum transaction.

#### Returns

1. `*types.Receipt` - the result of a transaction.
2. `error` - standard error.

## Transfer

### Deposit

Transacts `deposit` function from Mediator contract.

#### Parameters

1. `common.Address` - currency address.
2. `*big.Int` - amount of deposit.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### Withdraw

Transacts `withdraw` function from Mediator contract.

#### Parameters

1. `[]byte` - penultimate deposit transaction.
2. `[]byte` - proof of inclusion of a penultimate transaction in a Smart Plasma block.
3. `*big.Int` - the number of the block in which the penultimate transaction is included.
4. `[]byte` - last deposit transaction.
5. `[]byte` - proof of inclusion of a last transaction in a Smart Plasma block.
6. `*big.Int` - the number of the block in which the last transaction is included.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### StartExit

Transacts `startExit` function from RootChain contract.

#### Parameters

1. `[]byte` - penultimate deposit transaction.
2. `[]byte` - proof of inclusion of a penultimate transaction in a Smart Plasma block.
3. `*big.Int` - the number of the block in which the penultimate transaction is included.
4. `[]byte` - last deposit transaction.
5. `[]byte` - proof of inclusion of a last transaction in a Smart Plasma block.
6. `*big.Int` - the number of the block in which the last transaction is included.

#### Returns

1. `*types.Transaction` - Spectrum transaction.
2. `error` - standard error.

### AcceptTransaction

Sends raw Smart Plasma transaction to Smart Plasma RPC server.

#### Parameters

1. `[]byte` - raw Smart Plasma transaction.

#### Returns

1. `error` - standard error.

### AddCheckpoint

Sends UID and current transaction nonce for inclusion in a checkpoint.

#### Parameters

1. `*big.Int` - unique identifier of a deposit (uid).
2. `*big.Int` - a nonce.
3. `uint64` - a block number in which this transaction is present.

#### Returns

1. `error` - standard error.