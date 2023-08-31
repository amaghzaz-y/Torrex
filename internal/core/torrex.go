package torrex

import (
	"errors"
	"time"

	"github.com/amaghzaz-y/torrex/internal/chat"
	model "github.com/amaghzaz-y/torrex/internal/models"
	"github.com/amaghzaz-y/torrex/internal/store"
	"github.com/amaghzaz-y/torrex/internal/streamer"
	"github.com/amaghzaz-y/torrex/internal/torrent"
	"github.com/labstack/echo/v4"
)

type Torrex struct {
	Store       *store.Store
	streamer    *streamer.Streamer
	torrent     *torrent.Client
	Chat        *chat.Chat
	activeRooms map[string]*model.Room
}

func New() *Torrex {
	streamer := streamer.NewStreamer()
	store := store.New("torrex.data")
	torrent := torrent.DefaultClient()
	chat := chat.New()
	return &Torrex{
		store,
		streamer,
		torrent,
		chat,
		make(map[string]*model.Room, 12),
	}
}
func (t *Torrex) Close() {
	t.Store.Close()
	t.torrent.Close()
}

func (t *Torrex) NewPipelineHandler(room *model.Room) (echo.HandlerFunc, error) {
	if _, ok := t.activeRooms[room.Id]; ok {
		return nil, errors.New("room already exists")
	}
	if len(t.activeRooms) >= 12 {
		return nil, errors.New("unsufficient resources to handle more streams")
	}
	torr := t.torrent.NewTorrent(room.Movie.Title, room.Magnet)
	go torr.Download()
	for !torr.Ready() {
		time.Sleep(1 * time.Second)
	}
	room.Path = torr.FilePath()
	handler, flag := t.streamer.NewRoomStream(room)
	t.activeRooms[room.Id] = room
	go func() {
		<-flag
		delete(t.activeRooms, room.Id)
	}()
	return echo.WrapHandler(handler), nil
}

func (t *Torrex) StopStream(roomId string) {
	t.streamer.StopStream(roomId)
}

func (t *Torrex) Rooms() []*model.Room {
	var rooms []*model.Room
	for k, v := range t.activeRooms {
		rooms = append(rooms, &model.Room{
			Id:    k,
			Movie: v.Movie,
		})
	}
	return rooms
}
