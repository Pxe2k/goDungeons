package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
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
	screen.Fill(color.NRGBA{G: 0x80, B: 0x80, A: 0xff})
	vector.DrawFilledRect(screen, x, y, 80, 60, color.Black, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// TODO Make param for screensize
func ScreenSize() (int, int) {
	return 960, 720
}

func NewGame() *Game {
	return &Game{}
}
