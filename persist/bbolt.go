package persist

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"strconv"

	bolt "github.com/etcd-io/bbolt"
	"github.com/stevexnicholls/next/models"
)

// NewBoltStore creates a new store backed by bbolt db
func NewBoltStore(p string, b string) (Store, error) {

	db, err := bolt.Open(p, 0666, nil)
	if err != nil {
		return nil, err
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(b))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return &boltStore{
		DB:     db,
		Bucket: []byte(b),
	}, nil
}

type boltStore struct {
	DB     *bolt.DB
	Bucket []byte
}

// Update updates a value in the persistent store
func (g *boltStore) Update(kv *models.KeyValue) error {
	err := g.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(g.Bucket)
		v := []byte(strconv.FormatInt(kv.Value, 10))
		err := b.Put(UnsafeStringToBytes(kv.Key), v)
		return err
	})
	if err != nil {

	}

	return err
}

// Gets returns a value from the persistent store
func (g *boltStore) Get(key string) (*models.KeyValue, error) {

	// Start a writable transaction.
	tx, err := g.DB.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	b := tx.Bucket(g.Bucket)
	v := b.Get(UnsafeStringToBytes(key))
	// if key is nil then put new key with value 1
	// else increment value
	if v == nil {
		err = b.Put(UnsafeStringToBytes(key), []byte("1"))
		if err != nil {
			log.Fatal(err)
		}
		v = []byte("1")
	} else {
		x, _ := strconv.Atoi(string(v))
		x++
		v = []byte(strconv.Itoa(x))
		err = b.Put(UnsafeStringToBytes(key), v)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Commit the transaction and check for error
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	i, err := strconv.Atoi(string(v))
	if err != nil {
		log.Fatalln(err)
	}

	return &models.KeyValue{Key: key, Value: int64(i)}, err
}

// Gets returns a value from the persistent store
func (g *boltStore) View() ([]*models.KeyValue, error) {

	var kv []*models.KeyValue
	err := g.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(g.Bucket)

		b.ForEach(func(k, v []byte) error {
			x, _ := strconv.Atoi(string(v))
			kv = append(kv, &models.KeyValue{Key: string(k), Value: int64(x)})
			return nil
		})
		return nil
	})
	return kv, err
}

// Backup returns a dump of the persistent store
func (g *boltStore) Backup() ([]byte, error) {

	var v []byte

	err := g.DB.View(func(tx *bolt.Tx) error {
		var b bytes.Buffer
		_, err := tx.WriteTo(io.Writer(&b))
		v = b.Bytes()
		return err
	})
	if err != nil {
		return nil, err
	}

	return v, err
}

// Close closes the store
func (g *boltStore) Close() error {
	return g.DB.Close()
}

// Itob returns an 8-byte big endian representation of v.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
