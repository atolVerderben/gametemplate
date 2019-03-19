package gamename

import (
	"github.com/atolVerderben/tentsuyu"
)

func loadAudio() *tentsuyu.AudioPlayer {
	audioPlayer, _ := tentsuyu.NewAudioPlayer()

	return audioPlayer
}
