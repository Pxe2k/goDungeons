package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	x, y float32 = screenWidth / 2, screenHeight / 2
)

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		x -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		x += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		y += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x80, 0x80, 0xff})
	vector.DrawFilledRect(screen, x, y, 80, 60, color.Black, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Ebiten Game")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
