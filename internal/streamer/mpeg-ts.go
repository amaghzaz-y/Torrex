package streamer

import (
	"fmt"
	"log"
	"os/exec"
)

type MpegStream struct {
	uri  string
	port string
	cmd  *exec.Cmd
}

func newMpegStream(uri string, port string) *MpegStream {

	return &MpegStream{
		uri,
		port,
		nil,
	}
}

func (m *MpegStream) stream() {
	if m.uri == "" || m.port == "" {
		log.Println("cannot start mpeg stream : URI or PORT is undefined")
		return
	}
	// pkg_size=1316 is important for mpeg reader
	udp := fmt.Sprintf("udp://127.0.0.1:%s?pkt_size=1316", m.port)
	cmd := exec.Command("ffmpeg",
		"-re", "-i", m.uri, // input file
		"-c:v", "copy", //video config
		"-c:a", "libopus", "-b:a", "92k", "-ac", "2", // audio config
		"-preset", "veryfast", //preset
		"-tune", "film",
		"-maxrate", "4000k", // max upload rate
		"-bufsize", "9200k",
		"-f", "mpegts", udp, // output pipeline
	)
	log.Println("starting mpeg streaming", m.uri, "on port", m.port)
	m.cmd = cmd
	err := cmd.Run()
	if err != nil {
		if err.Error() != "signal: killed" {
			log.Println("error: mpeg streaming stopped")
		}
	}
	log.Println("finished mpeg streaming", m.uri, "on port", m.port)
}

func (m *MpegStream) close() {
	m.cmd.Process.Kill()
}
