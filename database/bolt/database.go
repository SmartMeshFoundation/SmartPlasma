package bolt

import (
	"strconv"
	"sync"

	"github.com/coreos/bbolt"
	"github.com/pkg/errors"
)

const (
	fileMode = 0600
)

var (
	BlocksBucket      = "blocks"
	CheckpointsBucket = "checkpoints"
)

// DB object for storage data to filesystem.
type DB struct {
	bucket []byte

	mtx      sync.Mutex
	database *bolt.DB
}

// NewDB creates new database.
func NewDB(file string, bucket string, options *bolt.Options) (*DB, error) {
	var opt *bolt.Options

	if options == nil {
		opt = bolt.DefaultOptions
	} else {
		opt = options
	}

	dBase, err := bolt.Open(file, fileMode, opt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}

	if err := dBase.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	}); err != nil {
		return nil, err
	}

	return &DB{
		database: dBase,
		bucket:   []byte(bucket),
	}, nil
}

// Close closes database file.
func (d *DB) Close() error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	return d.database.Close()
}

// Set adds value to new block.
func (d *DB) Set(val []byte) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	return d.database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(d.bucket)
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		number, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		return bucket.Put(strconv.AppendUint(nil, number, 10), val)
	})
}

// Get gets value by key.
func (d *DB) Get(key uint64) ([]byte, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	var val []byte

	err := d.database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(d.bucket)
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		val = bucket.Get(strconv.AppendUint(nil, key, 10))
		return nil
	})
	return val, err
}

// Current gets current block number.
func (d *DB) Current() (uint64, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	var sequence uint64

	err := d.database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(d.bucket)
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		sequence = bucket.Sequence()
		return nil
	})
	return sequence, err
}
