package main

import (
	"flag"
	"github.com/disintegration/imaging"
	"image/color"
	"log"
)

func main() {
	// Parsing command line args
	var inputPath = ""
	flag.StringVar(&inputPath, "i", "input.png", "Path to processing image")
	var percent = 0
	flag.IntVar(&percent, "p", 10, "Spaces length in percent from biggest dimension")
	var outputPath = ""
	flag.StringVar(&outputPath, "o", "output.png", "Path to save result")
	flag.Parse()

	if percent < 1 || percent > 100 {
		log.Fatalf("incorrect percent")
	}

	// Opening src file
	src, err := imaging.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Calculating new size
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y
	multiplier := 1 + float64(percent)/100
	newDim := 0
	if height > width {
		newDim = int(float64(height) * multiplier)
	} else {
		newDim = int(float64(width) * multiplier)
	}

	// Combining with white spaces
	white := imaging.New(newDim, newDim, color.White)
	dst := imaging.PasteCenter(white, src)

	// Saving result
	err = imaging.Save(dst, outputPath)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
