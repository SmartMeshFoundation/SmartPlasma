package database

// Database is interface for storage.
type Database interface {
	Set(key, val []byte) error
	Get(key []byte) ([]byte, error)
	Close() error
}
