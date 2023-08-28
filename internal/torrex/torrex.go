package torrex

import (
	"time"

	model "github.com/amaghzaz-y/torrex/internal/models"
	"github.com/amaghzaz-y/torrex/internal/store"
	"github.com/amaghzaz-y/torrex/internal/streamer"
	"github.com/amaghzaz-y/torrex/internal/torrent"
	"github.com/labstack/echo/v4"
)

type Torrex struct {
	Streamer *streamer.Streamer
	Store    *store.Store
	Torrent  *torrent.Client
}

func New() *Torrex {
	streamer := streamer.NewStreamer()
	store := store.New("torrex.data")
	torrent := torrent.DefaultClient()
	return &Torrex{
		streamer,
		store,
		torrent,
	}
}
func (t *Torrex) Close() {
	t.Store.Close()
	t.Torrent.Close()
}
func (t *Torrex) NewPipelineHandler(room *model.Room) echo.HandlerFunc {
	torr := t.Torrent.NewTorrent(room.Movie.Title, room.Magnet)
	go torr.Download()
	for !torr.Ready() {
		time.Sleep(5 * time.Second)
	}
	handler := t.Streamer.Stream("Asteroid city 2023", torr.FilePath(), torr.UdpPort())
	return echo.WrapHandler(handler)
}
