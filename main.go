package main

import (
	"github.com/disintegration/imaging"
	"image/color"
	"log"
)

func main() {
	src, err := imaging.Open("input.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y

	newDim := 0
	if height > width {
		newDim = int(float64(height) * 1.1)
	} else {
		newDim = int(float64(width) * 1.1)
	}
	white := imaging.New(newDim, newDim, color.White)
	dst := imaging.PasteCenter(white, src)
	err = imaging.Save(dst, "output.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
