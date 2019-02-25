package persist

import (
	"errors"
	"unsafe"

	"github.com/stevexnicholls/next/models"
)

// Common errors.
var (
	ErrNotFound         = errors.New("not found")
	ErrGone             = errors.New("gone")
	ErrReadOnly         = errors.New("read-only mode")
	ErrSnapshotReleased = errors.New("snapshot released")
	ErrIterReleased     = errors.New("iterator released")
	ErrClosed           = errors.New("closed")
	ErrVersionMismatch  = errors.New("version mismatch")
)

// UnsafeStringToBytes converts strings to []byte without memcopy
func UnsafeStringToBytes(s string) []byte {
	/* #nosec */
	return *(*[]byte)(unsafe.Pointer(&s))
}

// UnsafeBytesToString converts []byte to string without a memcopy
func UnsafeBytesToString(b []byte) string {
	/* #nosec */
	return *(*string)(unsafe.Pointer(&b))
}

// KeyValue represents an entry with key name
// type KeyValue struct {
// 	Key   string
// 	Value Value
// 	_     struct{}
// }

// Store for values by key
type Store interface {
	Update(*models.KeyValue) error
	Get(string) (*models.KeyValue, error)
	View() ([]*models.KeyValue, error)
	Backup() ([]byte, error)
	// FindByPrefix(string) ([]KeyValue, error)
	// Delete(string) error
	Close() error
}
