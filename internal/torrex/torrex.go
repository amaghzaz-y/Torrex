package torrex

import (
	"errors"
	"time"

	model "github.com/amaghzaz-y/torrex/internal/models"
	"github.com/amaghzaz-y/torrex/internal/store"
	"github.com/amaghzaz-y/torrex/internal/streamer"
	"github.com/amaghzaz-y/torrex/internal/torrent"
	"github.com/labstack/echo/v4"
)

type Torrex struct {
	Streamer    *streamer.Streamer
	Store       *store.Store
	Torrent     *torrent.Client
	ActiveRooms map[string]*model.Room
}

func New() *Torrex {
	streamer := streamer.NewStreamer()
	store := store.New("torrex.data")
	torrent := torrent.DefaultClient()
	return &Torrex{
		streamer,
		store,
		torrent,
		make(map[string]*model.Room, 10),
	}
}
func (t *Torrex) Close() {
	t.Store.Close()
	t.Torrent.Close()
}

func (t *Torrex) NewPipelineHandler(room *model.Room) (echo.HandlerFunc, error) {
	if _, ok := t.ActiveRooms["foo"]; ok {
		return nil, errors.New("room already exists")
	}
	if len(t.ActiveRooms) >= 10 {
		return nil, errors.New("unsufficient resources to handle more streams")
	}
	torr := t.Torrent.NewTorrent(room.Movie.Title, room.Magnet)
	go torr.Download()
	for !torr.Ready() {
		time.Sleep(1 * time.Second)
	}
	handler, flag := t.Streamer.Stream(room.Movie.Title, torr.FilePath(), torr.UdpPort())
	t.ActiveRooms[room.Id] = room
	go func() {
		<-flag
		delete(t.ActiveRooms, room.Id)
	}()
	return echo.WrapHandler(handler), nil
}
