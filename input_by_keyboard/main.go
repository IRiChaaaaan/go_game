package main

import (
	"log"

	"game/tiping/system"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(system.ScreenWidth*2, system.ScreenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&system.Game{}); err != nil {
		log.Fatal(err)
	}
}
