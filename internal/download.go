package torrex

import (
	"log"
	"time"
)

func (t *Torrex) RegisterTorrentFile(path string) error {
	torrent, err := t.client.AddTorrentFromFile(path)
	if err != nil {
		log.Println("error: cannot load torrent from filepath : ", path)
		return err
	}
	select {
	case <-torrent.GotInfo():
		t.torrents = append(t.torrents, torrent)
	case <-time.After(10 * time.Second):
		log.Println("timeout: cannot load torrent from filepath : ", path)
	}
	return nil
}

func (t *Torrex) RegisterTorrentMagnet(magnet string) error {
	torrent, err := t.client.AddMagnet(magnet)
	if err != nil {
		log.Println("error: cannot load torrent from magnet : ", magnet)
		return err
	}
	select {
	case <-torrent.GotInfo():
		t.torrents = append(t.torrents, torrent)
	case <-time.After(10 * time.Second):
		log.Println("timeout: cannot load torrent from filepath : ", magnet)
	}
	return nil
}

func (t *Torrex) DownloadFiles() {
	for _, t := range t.torrents {
		t.DownloadAll()
	}
}

func (t *Torrex) IsDownloadComplete() bool {
	state := false
	for _, torr := range t.torrents {
		if torr.Stats().PiecesComplete == torr.NumPieces() {
			state = true
		} else {
			state = false
		}
	}
	return state
}

func (t *Torrex) LogInfo() {
	for _, torr := range t.torrents {
		log.Println(torr.String(), "peers:", torr.Stats().ActivePeers,
			"status:", torr.BytesCompleted(), "/", torr.Length(),
			"pieces:", torr.NumPieces(), "/", torr.Stats().PiecesComplete,
		)
	}
}
