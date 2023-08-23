package server

import (
	"log"
	"net"
	"time"

	_ "embed"

	"github.com/bluenviron/gohlslib"
	"github.com/bluenviron/gohlslib/pkg/codecs"
	"github.com/bluenviron/mediacommon/pkg/formats/mpegts"
)

func (s *Server) stream() {
	// create the HLS muxer
	mux := &gohlslib.Muxer{
		VideoTrack: &gohlslib.Track{
			Codec: &codecs.H264{},
		},
		AudioTrack: &gohlslib.Track{
			Codec: &codecs.Opus{
				ChannelCount: 2,
			},
		},
	}
	err := mux.Start()
	if err != nil {
		panic(err)
	}

	s.router.HandleFunc("/*", mux.Handle)

	// create a socket to receive MPEG-TS packets
	pc, err := net.ListenPacket("udp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	// create a MPEG-TS reader
	r, err := mpegts.NewReader(mpegts.NewBufferedReader(newPacketConnReader(pc)))
	if err != nil {
		log.Fatalln("error reading mpeg-ts", err)
	}

	var timeDec *mpegts.TimeDecoder

	VideoFound, AudioFound := false, false
	for _, track := range r.Tracks() {
		if _, ok := track.Codec.(*mpegts.CodecH264); ok {
			r.OnDataH26x(track, func(rawPTS int64, _ int64, au [][]byte) error {
				if timeDec == nil {
					timeDec = mpegts.NewTimeDecoder(rawPTS)
				}
				pts := timeDec.Decode(rawPTS)
				mux.WriteH26x(time.Now(), pts, au)
				return nil
			})
			VideoFound = true
		}
		if _, ok := track.Codec.(*mpegts.CodecOpus); ok {
			r.OnDataOpus(track, func(rawPTS int64, aus [][]byte) error {
				if timeDec == nil {
					timeDec = mpegts.NewTimeDecoder(rawPTS)
				}
				pts := timeDec.Decode(rawPTS)
				mux.WriteOpus(time.Now(), pts, aus)
				return nil
			})
			AudioFound = true
		}
		if !VideoFound || !AudioFound {
			log.Println("H264 OR OPUS NOT FOUND")
		} else {
			break
		}
	}
	log.Println("stream started")
	for {
		defer func() {
			if err := recover(); err != nil {
				log.Println("error: stream halted due to", err)
				return
			}
		}()
		err := r.Read()
		if err != nil {
			break
		}
	}
	log.Println("stream finished")
}
