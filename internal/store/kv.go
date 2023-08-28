package store

import (
	"log"

	"go.etcd.io/bbolt"
)

type Store struct {
	*bbolt.DB
}

func New(path ...string) *Store {
	if path == nil {
		path[0] = "torrex.data"
	}
	log.Println(path[0])
	db, err := bbolt.Open(path[0], 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	txn, err := db.Begin(true)
	if err != nil {
		log.Fatalln(err)
	}
	defer txn.Rollback()
	_, err = txn.CreateBucketIfNotExists([]byte("movies"))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = txn.CreateBucketIfNotExists([]byte("rooms"))
	if err != nil {
		log.Fatalln(err)
	}
	if txn.Commit() != nil {
		log.Fatalln(err)
	}
	return &Store{db}
}
