package torrex

import (
	"time"

	model "github.com/amaghzaz-y/torrex/internal/models"
	"github.com/amaghzaz-y/torrex/internal/store"
	"github.com/amaghzaz-y/torrex/internal/streamer"
	"github.com/amaghzaz-y/torrex/internal/torrent"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type Torrex struct {
	Streamer *streamer.Streamer
	Store    *store.Store
	Torrent  *torrent.Client
}

func New() *Torrex {
	streamer := streamer.NewStreamer()
	store := store.New()
	torrent := torrent.DefaultClient()
	return &Torrex{
		streamer,
		store,
		torrent,
	}
}

func (t *Torrex) NewPipelineHandler(room *model.Room) func(*fiber.Ctx) error {
	torr := t.Torrent.NewTorrent(room.Movie.Title, room.Magnet)
	go torr.Download()
	for !torr.Ready() {
		time.Sleep(5 * time.Second)
	}
	handler := t.Streamer.Stream("Asteroid city 2023", torr.FilePath(), torr.UdpPort())
	return adaptor.HTTPHandlerFunc(handler)
}
