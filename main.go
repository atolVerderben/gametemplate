package main

import (
	"github.com/atolVerderben/gametemplate/game"
	"github.com/hajimehoshi/ebiten"
)

//Size of the game
const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

func main() {
	game, err := gamename.NewGame(ScreenWidth, ScreenHeight)
	if err != nil {
		panic(err)
	}
	if err := ebiten.Run(game.Loop, ScreenWidth, ScreenHeight, 1, "Game Template"); err != nil {
		panic(err)
	}
}
