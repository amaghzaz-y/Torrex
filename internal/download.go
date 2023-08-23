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

// func GetFilesWithSuffix(t *torrent.Torrent, suffixes []string) []*torrent.File {
// 	var files []*torrent.File
// 	for _, file := range t.Files() {
// 		for _, suffix := range suffixes {
// 			if strings.HasSuffix(file.DisplayPath(), suffix) {
// 				files = append(files, file)
// 			}
// 		}
// 	}
// 	return files
// }

// func (t *Torrex) DownloadFilesWithFormat(formats []string) {
// 	for _, torrent := range t.torrents {
// 		files := GetFilesWithSuffix(torrent, formats)
// 		for _, file := range files {
// 			t.files = append(t.files, file)
// 			file.Download()
// 		}
// 	}
// }

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
		}
		state = false
	}
	return state
}
