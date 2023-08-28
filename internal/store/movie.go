package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/adler32"

	model "github.com/amaghzaz-y/torrex/internal/models"
)

// the key is the magnet since it's faster to scrape it
func (s *Store) UpsertMovie(magnet string, movie *model.Movie) error {
	txn, err := s.Begin(true)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	json, err := json.Marshal(movie)
	if err != nil {
		return err
	}
	crc := fmt.Sprint(adler32.Checksum([]byte(magnet)))
	err = txn.Bucket([]byte("movies")).Put([]byte(crc), json)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func (s *Store) GetMovieByUrl(magnet string) (*model.Movie, error) {
	txn, err := s.Begin(false)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	var movie model.Movie
	crc := fmt.Sprint(adler32.Checksum([]byte(magnet)))
	blob := txn.Bucket([]byte("movies")).Get([]byte(crc))
	if blob == nil {
		return nil, errors.New("movie not found")
	}
	err = json.Unmarshal(blob, &movie)
	return &movie, err
}
