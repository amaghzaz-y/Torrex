package torrex

import ffmpeg "github.com/u2takey/ffmpeg-go"

func TranscodeVideoToHLS(input string, output string) error {
	return ffmpeg.Input(input).Output(output).OverWriteOutput().Run()
}
