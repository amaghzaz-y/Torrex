package torrent

import "github.com/anacrolix/torrent"

func (t *Torrent) Download() {
	t.torrent.DownloadAll()
	var target *torrent.File
	var maxSize int64
	for _, file := range t.torrent.Files() {
		if maxSize < file.Length() {
			maxSize = file.Length()
			target = file
		}
	}
	t.filepath = target.Path()
	//starting index
	sidx := target.Offset() * int64(t.torrent.NumPieces()) / t.torrent.Length()
	//ending index
	eidx := (target.Offset() + target.Length()) * int64(t.torrent.NumPieces()) / t.torrent.Length()
	// Prioritize first 5% of the file.
	for idx := sidx; idx <= eidx*5/100; idx++ {
		t.torrent.Piece(int(idx)).SetPriority(torrent.PiecePriorityNow)
	}
}

func (t *Torrent) Completion() float64 {
	info := t.torrent.Info()
	if info == nil {
		return 0
	}
	return float64(t.torrent.BytesCompleted()) / float64(info.TotalLength()) * 100
}

func (t *Torrent) Ready() bool {
	return t.Completion() > 5
}
