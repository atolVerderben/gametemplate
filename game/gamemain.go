package gamename

import (
	"os"

	"github.com/atolVerderben/tentsuyu"
)

//GameMain represents the main GameState of our game
type GameMain struct {
	gameStateMsg  tentsuyu.GameStateMsg
	mainCamera    *tentsuyu.Camera
	width, height float64
}

//NewGameMain returns our main gamestate
func NewGameMain(game *tentsuyu.Game) *GameMain {

	g := &GameMain{}

	return g
}

//Update the gamestate
func (g *GameMain) Update(game *tentsuyu.Game) error {
	if game.Input.Button("Escape").JustPressed() {
		os.Exit(0)
	}
	return nil
}

//Draw the gamestate scene
func (g *GameMain) Draw(game *tentsuyu.Game) error {

	return nil
}

//Msg returns the gamestatemsg and achieves the GameState interface
func (g GameMain) Msg() tentsuyu.GameStateMsg {
	return g.gameStateMsg
}

//SetMsg sets the gamestatemsg value
func (g *GameMain) SetMsg(gs tentsuyu.GameStateMsg) {
	g.gameStateMsg = gs
}
