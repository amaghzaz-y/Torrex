package store

import (
	"encoding/json"
	"errors"

	model "github.com/amaghzaz-y/torrex/internal/models"
)

func (s *Store) UpsertRoom(room *model.Room) error {
	txn, err := s.Begin(true)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	json, err := json.Marshal(room)
	if err != nil {
		return err
	}
	err = txn.Bucket([]byte("rooms")).Put([]byte(room.Id), json)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func (s *Store) GetRoom(id string) (*model.Room, error) {
	txn, err := s.Begin(false)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	var room model.Room

	blob := txn.Bucket([]byte("rooms")).Get([]byte(id))
	if blob == nil {
		return nil, errors.New("room not found")
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(blob, &room)
	return &room, err
}
