package database

// Database is interface for storage.
type Database interface {
	Set(val []byte) error
	Get(key uint64) ([]byte, error)
	Current() (uint64, error)
}
