package gamename

import (
	"image/color"
	"log"

	"github.com/atolVerderben/tentsuyu"
	"github.com/hajimehoshi/ebiten"
)

func loadImages() *tentsuyu.ImageManager {
	imageManager := tentsuyu.NewImageManager()

	//Simple white 1x1 pixel image for manipulation
	whiteImg, err := ebiten.NewImage(1, 1, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	whiteImg.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	imageManager.AddImage("pixel", whiteImg)

	return imageManager
}
