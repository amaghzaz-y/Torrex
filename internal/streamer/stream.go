package streamer

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/bluenviron/gohlslib"
	"github.com/bluenviron/gohlslib/pkg/codecs"
	"github.com/bluenviron/mediacommon/pkg/formats/mpegts"
)

type Stream struct {
	name string
	port string
	hls  *gohlslib.Muxer
	mpeg *mpegts.Reader
	pc   net.PacketConn
}

func NewStream(name string, port string) *Stream {
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
	return &Stream{
		name,
		port,
		mux,
		nil,
		nil,
	}
}

func (s *Stream) openHlsMuxer() {
	err := s.hls.Start()
	if err != nil {
		panic(err)
	}
}

func (s *Stream) openMpegReader() {
	uri := fmt.Sprintf("127.0.0.1:%s", s.port)
	pc, err := net.ListenPacket("udp", uri)
	if err != nil {
		log.Fatalln("error listening to socket", err)
	}
	r, err := mpegts.NewReader(mpegts.NewBufferedReader(newPacketConnReader(pc)))
	if err != nil {
		log.Fatalln("error reading mpeg-ts", err)
	}
	s.mpeg = r
	s.pc = pc
}

func (s *Stream) openMpegDecoder() {
	var timeDec *mpegts.TimeDecoder
	VideoFound, AudioFound := false, false
	for _, track := range s.mpeg.Tracks() {
		if _, ok := track.Codec.(*mpegts.CodecH264); ok {
			s.mpeg.OnDataH26x(track, func(rawPTS int64, _ int64, au [][]byte) error {
				if timeDec == nil {
					timeDec = mpegts.NewTimeDecoder(rawPTS)
				}
				pts := timeDec.Decode(rawPTS)
				s.hls.WriteH26x(time.Now(), pts, au)
				return nil
			})
			VideoFound = true
		}
		if _, ok := track.Codec.(*mpegts.CodecOpus); ok {
			s.mpeg.OnDataOpus(track, func(rawPTS int64, aus [][]byte) error {
				if timeDec == nil {
					timeDec = mpegts.NewTimeDecoder(rawPTS)
				}
				pts := timeDec.Decode(rawPTS)
				s.hls.WriteOpus(time.Now(), pts, aus)
				return nil
			})
			AudioFound = true
		}
		if VideoFound && AudioFound {
			break
		}
	}
}

func (s *Stream) readMpegStream() {
	for {
		defer func() {
			if err := recover(); err != nil {
				log.Println("error: stream halted due to", err)
				return
			}
		}()
		err := s.mpeg.Read()
		if err != nil {
			return
		}
	}
}

func (s *Stream) Start() {
	s.openHlsMuxer()
	s.openMpegReader()
	s.openMpegDecoder()
	s.readMpegStream()
}

func (s *Stream) Close() {
	s.hls.Close()
	s.pc.Close()
}
