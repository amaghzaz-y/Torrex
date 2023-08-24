package streamer

import (
	"fmt"
	"log"
	"os/exec"
)

type MpegStream struct {
	uri  string
	port string
}

func newMpegStream(uri string, port string) *MpegStream {
	return &MpegStream{
		uri,
		port,
	}
}

func (m *MpegStream) stream() {
	// pkg_size=1316 is important for mpeg reader
	udp := fmt.Sprintf("udp://127.0.0.1:%s?pkt_size=1316", m.port)
	cmd := exec.Command("ffmpeg",
		"-re", "-i", m.uri, // input file
		"-c:v", "libx264", "-b:v", "1000k", //video config
		"-video_size", "1280x720", //video output
		"-c:a", "libopus", "-b:a", "72k", "-ac", "2", // audio config
		"-bufsize", "1500k",
		"-preset", "ultrafast", //preset
		"-tune", "zerolatency",
		"-maxrate", "1200k", // max upload rate
		"-f", "mpegts", udp, // output pipeline
	)
	defer cmd.Cancel()
	log.Println("running:", cmd.String())
	log.Println("starting mpeg streaming", m.uri, "on port", m.port)
	err := cmd.Run()
	if err != nil {
		log.Fatalln("error: cannot start mpeg-ts udp streaming :", err)
	}
	log.Println("finished mpeg streaming", m.uri, "on port", m.port)
}
