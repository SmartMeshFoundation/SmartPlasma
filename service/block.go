package service

import (
	"bytes"
	"context"
	"math/big"
	"strconv"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/merkle"
)

// AcceptTransaction adds a transaction to current transactions block.
func (s *Service) AcceptTransaction(tx *transaction.Transaction) error {
	return s.currentBlock.AddTx(tx)
}

// CreateProof creates merkle proof for particular uid.
// Argument `block` is block number.
func (s *Service) CreateProof(uid *big.Int, block uint64) ([]byte, error) {
	raw, err := s.RawBlockFromDB(block)
	if err != nil {
		return nil, err
	}

	blk := transactions.NewBlock()
	err = buildBlockFromBytes(blk, raw)
	if err != nil {
		return nil, err
	}

	return blk.CreateProof(uid), err
}

// VerifyTxProof returns true if uid was spent in this block.
func (s *Service) VerifyTxProof(uid *big.Int, hash common.Hash,
	block uint64, proof []byte) (bool, error) {
	root, err := s.session.ChildChain(new(big.Int).SetUint64(block))
	if err != nil {
		return false, err
	}

	return merkle.CheckMembership(uid, hash, root, proof), err
}

// InitBlock initializes a new block.
func (s *Service) InitBlock() {
	s.currentBlock = transactions.NewBlock()
}

// BuildBlock build current Plasma block.
func (s *Service) BuildBlock() (common.Hash, error) {
	if s.strongMode { // TODO: not used
		s.ValidateBlock(context.Background()) // TODO: change to context with timeout
	}
	return s.currentBlock.Build()
}

// RawBlockFromDB returns raw Plasma block from database.
func (s *Service) RawBlockFromDB(number uint64) ([]byte, error) {
	return s.blockBase.Get(strconv.AppendUint(nil, number, 10))
}

// SaveBlockToDB saves Plasma Block to database.
func (s *Service) SaveBlockToDB(number uint64,
	blk transactions.TxBlock) error {
	raw, err := blk.Marshal()
	if err != nil {
		return err
	}
	return s.blockBase.Set(strconv.AppendUint(nil, number, 10), raw)
}

// SendBlockHash sends a Plasma block hash to the blockchain.
func (s *Service) SendBlockHash(
	ctx context.Context, hash common.Hash) (*types.Transaction, error) {
	session := rootchain.CopySession(s.session)
	session.TransactOpts.Context = ctx

	return session.NewBlock(hash)
}

// LastBlockNumber gets last block number from blockchain.
func (s *Service) LastBlockNumber(ctx context.Context) (*big.Int, error) {
	session := rootchain.CopySession(s.session)
	session.CallOpts.Context = ctx
	return session.BlockNumber()
}

// CurrentBlock returns current Plasma block.
func (s *Service) CurrentBlock() transactions.TxBlock {
	return s.currentBlock
}

// ValidateBlock returns checked transactions.
func (s *Service) ValidateBlock(ctx context.Context) error {
	ch := s.currentBlock.Transactions(ctx)

	txs := make(map[string]*transaction.Transaction)

	lastBlock, err := s.session.BlockNumber()
	if err != nil {
		return err
	}

	for tx := range ch {
		storedAmount, err := s.session.Wallet(common.BigToHash(tx.UID()))
		if err != nil {
			return err
		}

		// if storedAmount = 0 then the deposit does not exist
		if storedAmount.Int64() == 0 {
			continue
		}

		if lastBlock.Uint64() == 0 {
			sender, err := transaction.Sender(tx)
			if err != nil {
				continue
			}

			if tx.NewOwner().String() != sender.String() {
				continue
			}

			txs[tx.UID().String()] = tx
			continue
		}

		startBlock, err := s.session.Wallet2(tx.UID())
		if err != nil {
			return err
		}

		sender, err := transaction.Sender(tx)
		if err != nil {
			continue
		}

		// if first transaction
		if tx.Nonce().Int64() == 0 {
			if tx.NewOwner().String() != sender.String() {
				continue
			}

			if startBlock.Uint64() > lastBlock.Uint64() {
				continue
			}

			type testTx struct {
				found bool
				err   error
			}

			testTxChan := make(chan *testTx)

			start := int(startBlock.Int64())
			finish := int(lastBlock.Int64())

			go func() {
				for i := start; i <= finish; i++ {
					rawBlock, err := s.RawBlockFromDB(uint64(i))
					if err != nil {
						testTxChan <- &testTx{err: err}
						return
					}

					block := transactions.NewBlock()
					err = block.Unmarshal(rawBlock)
					if err != nil {
						testTxChan <- &testTx{err: err}
						return
					}

					_, err = block.Build()
					if err != nil {
						testTxChan <- &testTx{err: err}
						return
					}

					proof := block.CreateProof(tx.UID())

					found, err := s.VerifyTxProof(
						tx.UID(), tx.Hash(), uint64(i), proof)
					if err != nil {
						testTxChan <- &testTx{err: err}
						return
					}

					if found {
						testTxChan <- &testTx{found: true}
						return
					}
				}
				testTxChan <- &testTx{}
			}()

			select {
			case result := <-testTxChan:
				if result.err != nil {
					return err
				}
				if !result.found {
					txs[tx.UID().String()] = tx
				}
			case <-ctx.Done():
				close(testTxChan)
				return errors.New("timeout")
			}
		} else if tx.Nonce().Int64() > 0 {
			if tx.NewOwner().String() == sender.String() {
				continue
			}

			rawBlock, err := s.RawBlockFromDB(tx.PrevBlock().Uint64())
			if err != nil {
				return err
			}

			prevBlock := transactions.NewBlock()
			err = prevBlock.Unmarshal(rawBlock)
			if err != nil {
				return err
			}

			numTx := prevBlock.NumberOfTX()
			if numTx == 0 {
				continue
			}

			_, err = prevBlock.Build()
			if err != nil {
				return err
			}

			proof := prevBlock.CreateProof(tx.UID())
			prevTx, err := prevBlock.GetTx(tx.UID())
			if err != nil {
				continue
			}

			if tx.Nonce().Uint64() != prevTx.Nonce().Uint64()+1 {
				continue
			}

			if !bytes.Equal(tx.Amount().Bytes(), prevTx.Amount().Bytes()) {
				continue
			}

			found, err := s.VerifyTxProof(
				tx.UID(), prevTx.Hash(), tx.PrevBlock().Uint64(), proof)
			if err != nil {
				return err
			}

			if !found {
				continue
			}
			txs[tx.UID().String()] = tx
		}
	}

	// TODO: revert

	s.currentBlock = transactions.NewBlock()

	for _, v := range txs {
		err = s.currentBlock.AddTx(v)
		if err != nil {
			return err
		}
	}

	return nil
}
