package streamer

import (
	"fmt"
	"math/rand"
	"net/http"

	model "github.com/amaghzaz-y/torrex/internal/models"
)

type Stream struct {
	id   string
	room *model.Room
	path string
	port string
	mpeg *MpegStream
	hls  *HlsStream
}

func NewStream(room *model.Room) *Stream {
	return &Stream{
		id:   room.Id,
		room: room,
		path: room.Path,
		port: fmt.Sprint(rand.Intn(65536-1024) + 1024),
		mpeg: nil,
		hls:  nil,
	}
}

func (s *Stream) Start() (http.HandlerFunc, chan bool) {
	s.mpeg = newMpegStream(s.path, s.port)
	s.hls = newHlsStream(s.room.Movie.Title, s.port)
	flag := make(chan bool)
	go s.mpeg.stream()
	go func() {
		s.hls.stream()
		flag <- true
	}()
	return s.hls.handler(), flag
}

func (s *Stream) Stop() {
	s.hls.close()
	s.mpeg.close()
}
