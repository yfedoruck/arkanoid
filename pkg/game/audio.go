package game

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/yfedoruck/arkanoid/pkg/env"
	"github.com/yfedoruck/arkanoid/pkg/fail"
	"os"
	"path/filepath"
	"time"
)

func LoadSound(path string) *beep.Buffer {
	file, err := os.Open(env.BasePath() + filepath.FromSlash("/static/audio/"+path))
	fail.Check(err)

	streamer, format, err := wav.Decode(file)
	fail.Check(err)

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	fail.Check(err)

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	// streamer close file
	err = streamer.Close()
	fail.Check(err)

	return buffer
}

func InitSpeaker(format beep.Format) {
	var err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	fail.Check(err)
}

func PlaySound(buffer *beep.Buffer) {
	shot := buffer.Streamer(0, buffer.Len())
	speaker.Play(shot)
}
