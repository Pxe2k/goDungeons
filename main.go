package main

import (
	g "github.com/Pxe2k/goDungeons/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(g.ScreenSize())
	ebiten.SetWindowTitle("Hasan Game")
	ebiten.RunGame(g.NewGame())
}
