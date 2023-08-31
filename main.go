package main

import (
	"fmt"
	g "github.com/Pxe2k/goDungeons/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"os"
)

func main() {
	ebiten.SetWindowSize(g.ScreenSize())
	ebiten.SetWindowTitle("Hasan Game")

	//TODO move this code to audio.go file
	audioContext := audio.NewContext(44100)

	f, err := os.Open("resource/audio/macan-asphalt-8.wav")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	d, err := wav.DecodeWithSampleRate(44100, f)
	if err != nil {
		fmt.Println(err)
	}

	audioPlayer, err := audioContext.NewPlayer(d)
	if err != nil {
		fmt.Println(err)
	}
	audioPlayer.Play()

	ebiten.RunGame(g.NewGame())
}
