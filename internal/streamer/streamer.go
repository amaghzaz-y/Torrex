package streamer

import (
	_ "embed"
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