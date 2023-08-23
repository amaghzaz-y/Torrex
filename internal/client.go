package torrex

import (
	"log"

	"github.com/anacrolix/torrent"
)

type Torrex struct {
	client   *torrent.Client
	torrents []*torrent.Torrent
	files    []*torrent.File
}

func NewClient() *Torrex {
	client, err := torrent.NewClient(nil)
	if err != nil {
		log.Fatalln(err)
	}
	return &Torrex{
		client,
		nil,
		nil,
	}
}
