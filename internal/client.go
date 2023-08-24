package torrex

import (
	"log"

	"github.com/anacrolix/torrent"
)

type Torrex struct {
	client   *torrent.Client
	torrents []*torrent.Torrent
}

func NewClient() *Torrex {
	client, err := torrent.NewClient(nil)
	if err != nil {
		log.Fatalln(err)
	}
	return &Torrex{
		client,
		nil,
	}
}
