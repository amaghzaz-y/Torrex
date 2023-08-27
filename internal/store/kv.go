package kv

import (
	"encoding/json"

	model "github.com/amaghzaz-y/torrex/internal/models"
	"go.etcd.io/bbolt"
)

type Store struct {
	movies  *bbolt.Bucket
	magnets *bbolt.Bucket
}

func New() (*Store, error) {
	db, err := bbolt.Open("torrex.data", 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	txn, err := db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	mb, err := txn.Bucket([]byte("movies")).CreateBucketIfNotExists([]byte("movies"))
	if err != nil {
		return nil, err
	}
	mg, err := txn.Bucket([]byte("magnets")).CreateBucketIfNotExists([]byte("magnets"))
	if err != nil {
		return nil, err
	}
	return &Store{mb, mg}, txn.Commit()
}

func (s *Store) MovieUpsert(movie *model.Movie) error {
	json, err := json.Marshal(movie)
	if err != nil {
		return err
	}
	return s.movies.Put([]byte(movie.Url), json)
}
