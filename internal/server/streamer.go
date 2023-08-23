package server

import (
	"log"
	"net"
	"net/http"
	"time"

	_ "embed"

	"github.com/bluenviron/gohlslib"
	"github.com/bluenviron/gohlslib/pkg/codecs"
	"github.com/bluenviron/mediacommon/pkg/formats/mpegts"
)

//go:embed index.html
var index []byte

func handleIndex(wrapped http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(index))
			return
		}
		wrapped(w, r)
	}
}

func Stream() {
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

	// create an HTTP server and link it to the HLS muxer
	s := &http.Server{
		Addr:    ":8080",
		Handler: handleIndex(mux.Handle),
	}
	log.Println("HTTP server created on :8080")
	go s.ListenAndServe()

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
			// setup a callback that is called once a H264 access unit is received
			r.OnDataH26x(track, func(rawPTS int64, _ int64, au [][]byte) error {
				// decode the time
				if timeDec == nil {
					timeDec = mpegts.NewTimeDecoder(rawPTS)
				}
				pts := timeDec.Decode(rawPTS)
				// pass the access unit to the HLS muxer
				err := mux.WriteH26x(time.Now(), pts, au)
				if err != nil {
					log.Panic(err)
				}
				return nil
			})
			log.Println("H264 FOUND")

			VideoFound = true
		}
		if _, ok := track.Codec.(*mpegts.CodecOpus); ok {
			r.OnDataOpus(track, func(rawPTS int64, aus [][]byte) error {
				if timeDec == nil {
					timeDec = mpegts.NewTimeDecoder(rawPTS)
				}
				pts := timeDec.Decode(rawPTS)
				err := mux.WriteOpus(time.Now(), pts, aus)
				if err != nil {
					log.Panic(err)
				}
				return nil
			})
			AudioFound = true
			log.Println("OPUS FOUND")
		}
		if !VideoFound || !AudioFound {
			log.Println("H264 OR OPUS NOT FOUND")
		} else {
			break
		}
	}

	// read from the MPEG-TS stream
	log.Println("Stream started")
	for {
		err := r.Read()
		if err != nil {
			log.Println("stream disconnected, waiting...")
			time.Sleep(5 * time.Second)
		}
	}
}
