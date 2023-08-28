package store

import (
	"encoding/json"
	"errors"

	model "github.com/amaghzaz-y/torrex/internal/models"
)

func (s *Store) MovieUpsert(movie *model.Movie) error {
	txn, err := s.kv.Begin(true)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	json, err := json.Marshal(movie)
	if err != nil {
		return err
	}
	err = txn.Bucket([]byte("movies")).Put([]byte(movie.Url), json)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func (s *Store) GetMovieByUrl(url string) (*model.Movie, error) {
	txn, err := s.kv.Begin(false)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	var movie model.Movie
	blob := txn.Bucket([]byte("movies")).Get([]byte(url))
	if blob == nil {
		return nil, errors.New("movie not found")
	}
	err = json.Unmarshal(blob, &movie)
	return &movie, err
}
