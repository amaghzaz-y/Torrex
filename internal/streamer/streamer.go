package streamer

import (
	_ "embed"
	"net/http"
)

type Streamer struct {
	streams []*HlsStream
}

func NewStreamer() *Streamer {
	return &Streamer{}
}

func (s *Streamer) AddStream(stream *HlsStream) {
	s.streams = append(s.streams, stream)
}

func (s *Streamer) Test() http.HandlerFunc {
	mpeg := NewMpegStream("Looney/Looney.mkv", "9000")
	hls := NewHlsStream("Lonney", "9000")
	go mpeg.Stream()
	go hls.Stream()
	return hls.Handler()
}
