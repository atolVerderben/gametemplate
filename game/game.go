package gamename

import (
	"math/rand"
	"time"

	"github.com/atolVerderben/tentsuyu"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var (
	//ScreenWidth of the game
	ScreenWidth float64
	//ScreenHeight of the game
	ScreenHeight float64
)

//NewGame returns a new Game while setting the width and height of the screen
func NewGame(w, h float64) (game *tentsuyu.Game, err error) {
	ScreenWidth = w
	ScreenHeight = h
	game, err = tentsuyu.NewGame(w, h)

	game.LoadImages(func() *tentsuyu.ImageManager {
		return loadImages()
	})
	game.LoadAudio(func() *tentsuyu.AudioPlayer {
		return loadAudio()
	})

	game.SetGameStateLoop(func() error {
		switch game.GetGameState().Msg() {
		case GameStateMsgMain:
			game.SetGameState(NewGameMain(game))
		case tentsuyu.GameStateMsgNotStarted: //Use this message to load the initial game
			game.SetGameState(NewGameMain(game))
		default:

		}
		return nil
	})

	return
}

//GameState Messages used for this game
var (
	GameStateMsgMain tentsuyu.GameStateMsg = "MainGame"
)
