package torrent

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/anacrolix/torrent"
)

type Client struct {
	*torrent.Client
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
	t, err := c.AddMagnet(magnet)
	if err != nil {
		log.Println("error: cannot add magnet to torrent client for", title)
	}
	torr := &Torrent{
		client:   c.Client,
		title:    title,
		filepath: "",
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
func (c *Client) Close() {
	c.Client.Close()
}
func (t *Torrent) UdpPort() string {
	return t.port
}

func (t *Torrent) FilePath() string {
	var target *torrent.File
	var maxSize int64
	if len(t.filepath) != 0 {
		return t.filepath
	}
	for _, file := range t.torrent.Files() {
		if maxSize < file.Length() {
			maxSize = file.Length()
			target = file
		}
	}
	t.filepath = target.Path()
	return target.Path()
}
