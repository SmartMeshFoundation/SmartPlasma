package memory

import "sync"

// DB object for in memory storage.
// It is not used in production.
type DB struct {
	mtx    sync.Mutex
	last   uint64
	blocks map[uint64][]byte
}

// NewDB creates new database.
func NewDB() *DB {
	return &DB{
		mtx:    sync.Mutex{},
		last:   0,
		blocks: make(map[uint64][]byte),
	}
}

// Set adds value to new block.
func (d *DB) Set(val []byte) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	d.blocks[d.last+1] = val
	d.last++
	return nil
}

// Get gets value by key. Block number 0 is always empty.
func (d *DB) Get(key uint64) ([]byte, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	return d.blocks[key], nil
}

// Current gets current block number.
func (d *DB) Current() (uint64, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	return d.last, nil
}
