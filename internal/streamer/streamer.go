package streamer

import (
	_ "embed"
	"net/http"
)

type Streamer struct{}

func NewStreamer() *Streamer {
	return &Streamer{}
}

func (s *Streamer) Stream(title string, path string, port string) http.HandlerFunc {
	mpeg := newMpegStream(path, port)
	hls := newHlsStream(title, port)
	go mpeg.stream()
	go hls.stream()
	return hls.handler()
}
