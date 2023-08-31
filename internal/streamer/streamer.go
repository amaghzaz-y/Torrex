package streamer

import (
	"net/http"

	model "github.com/amaghzaz-y/torrex/internal/models"
)

type Streamer struct {
	streams map[string]*Stream
}

func NewStreamer() *Streamer {
	return &Streamer{
		streams: make(map[string]*Stream),
	}
}

// starts room streaming and returns an http handler + stream status as chan
func (s *Streamer) NewRoomStream(room *model.Room) (http.HandlerFunc, chan bool) {
	strm := NewStream(room)
	s.streams[strm.id] = strm
	return strm.Start()
}

func (s *Streamer) StopStream(roomId string) {
	if strm, exists := s.streams[roomId]; exists {
		strm.Stop()
	}
}

// ## DEPRICATED
//
// ## Use NewRoomStream instead
func (s *Streamer) Stream(title string, path string, port string) (http.HandlerFunc, chan bool) {
	mpeg := newMpegStream(path, port)
	hls := newHlsStream(title, port)
	flag := make(chan bool)
	go mpeg.stream()
	go func() {
		hls.stream()
		flag <- true
	}()
	return hls.handler(), flag
}
