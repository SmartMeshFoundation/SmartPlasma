package memory

import (
	"math/big"
	"sync"
)

// DB object for in memory storage.
// It is not used in production.
type DB struct {
	mtx    sync.Mutex
	last   uint64
	blocks map[string][]byte
}

// NewDB creates new database.
func NewDB() *DB {
	return &DB{
		mtx:    sync.Mutex{},
		last:   0,
		blocks: make(map[string][]byte),
	}
}

// Close erases data from database.
func (d *DB) Close() error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	d.blocks = make(map[string][]byte)
	return nil
}

func (d *DB) Set(key, val []byte) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	d.blocks[string(key)] = val
	d.last++
	return nil
}

// Get gets value by key. Block number 0 is always empty.
func (d *DB) Get(key []byte) ([]byte, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	return d.blocks[new(big.Int).SetBytes(key).String()], nil
}
