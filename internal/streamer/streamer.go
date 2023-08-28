package streamer

import (
	_ "embed"
	"net/http"
)

type Streamer struct{}

func NewStreamer() *Streamer {
	return &Streamer{}
}

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
