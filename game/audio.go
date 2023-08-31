package game

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"os"
)

func PlayMusic() (*audio.Player, error) {
	audioContext := audio.NewContext(44100)

	f, err := os.Open("resource/audio/macan-asphalt-8.wav")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := wav.DecodeWithSampleRate(44100, f)
	if err != nil {
		return nil, err
	}

	audioPlayer, err := audioContext.NewPlayer(d)
	if err != nil {
		return nil, err
	}

	return audioPlayer, nil
}
