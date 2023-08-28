package store

import (
	"log"

	"go.etcd.io/bbolt"
)

type Store struct {
	kv *bbolt.DB
}

func New(path ...string) (*Store, error) {
	if path == nil {
		path[0] = "torrex.data"
	}
	log.Println(path[0])
	db, err := bbolt.Open(path[0], 0600, nil)
	if err != nil {
		return nil, err
	}
	txn, err := db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	_, err = txn.CreateBucketIfNotExists([]byte("movies"))
	if err != nil {
		return nil, err
	}
	_, err = txn.CreateBucketIfNotExists([]byte("magnets"))
	if err != nil {
		return nil, err
	}
	return &Store{db}, txn.Commit()
}
