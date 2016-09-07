package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type BoltDBStore struct {
	db *bolt.DB
}

var bucketID = []byte("storage")

type BoltStoreOptions struct {
	filename string
}

func NewBoltStore(options BoltStoreOptions) (*BoltDBStore, error) {
	db, err := bolt.Open(options.filename, 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucketID)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &BoltDBStore{db: db}, nil
}

func NewBoltStoreFromMap(options map[string]interface{}) (*BoltDBStore, error) {
	opts := BoltStoreOptions{}
	opts.filename = options["filename"].(string)
	return NewBoltStore(opts)
}

func (b *BoltDBStore) Get(key string) (string, error) {
	var val []byte
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketID)
		val = bucket.Get([]byte(key))
		return nil
	})
	if err != nil {
		return "", err
	}
	if val == nil {
		return "", NoSuchKey
	}
	return string(val), nil

}
func (b *BoltDBStore) Set(key string, value string) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketID)
		err := bucket.Put([]byte(key), []byte(value))
		return err
	})
}

func (b *BoltDBStore) Delete(key string) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketID)
		err := bucket.Delete([]byte(key))
		return err
	})
}
