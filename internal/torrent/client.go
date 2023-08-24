package torrent

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/anacrolix/torrent"
)

type Client struct {
	torrentClient *torrent.Client
}

func DefaultClient() *Client {
	c, err := torrent.NewClient(nil)
	if err != nil {
		log.Fatalln(err)
	}
	return &Client{c}
}

type Torrent struct {
	client   *torrent.Client
	torrent  *torrent.Torrent
	title    string
	filepath string
	port     string
}

func (c *Client) NewTorrent(title, magnet string) *Torrent {
	t, err := c.torrentClient.AddMagnet(magnet)
	if err != nil {
		log.Println("error: cannot add magnet to torrent client for", title)
	}
	torr := &Torrent{
		client:   c.torrentClient,
		title:    title,
		filepath: title,
		port:     fmt.Sprint(rand.Intn(65536-1024) + 1024),
		torrent:  nil,
	}
	select {
	case <-t.GotInfo():
		torr.torrent = t
		return torr
	case <-time.After(10 * time.Second):
		log.Println("timeout: cannot load torrent from magnet : ", title)
		return nil
	}
}
