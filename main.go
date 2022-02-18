package main

import (
	"flag"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"path/filepath"
	"strings"
)

func createOutputPath(inputPath string) string {
	ext := filepath.Ext(inputPath)
	return strings.TrimSuffix(inputPath, ext) + "_processed" + ext
}

func calcNewSize(src image.Image, percent int) int {
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y
	multiplier := 1 + float64(percent)/100
	newSize := 0
	if height > width {
		newSize = int(float64(height) * multiplier)
	} else {
		newSize = int(float64(width) * multiplier)
	}
	return newSize
}

func processFile(inputPath string, outputPath string, percent int) (err error) {
	// Opening src file
	src, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	// Calculating new size
	newSize := calcNewSize(src, percent)

	// Combining with white spaces
	white := imaging.New(newSize, newSize, color.White)
	dst := imaging.PasteCenter(white, src)

	// Saving result
	err = imaging.Save(dst, outputPath)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Parsing command line args
	percent := flag.Int("p", 10, "Spaces size, percentage from biggest dimension")
	outputPath := flag.String("o", "", "Path to save result, allowed only with one input file")
	flag.Parse()
	inputArgs := flag.Args()
	if len(inputArgs) < 1 {
		panic("No input files provided")
	}
	if len(inputArgs) > 1 && *outputPath != "" {
		panic("-o flag is allowed only with one input file")
	}
	if *percent < 1 || *percent > 100 {
		panic("incorrect percent")
	}

	for _, inputPath := range inputArgs {
		currentOutputPath := createOutputPath(inputPath)
		if *outputPath != "" {
			currentOutputPath = *outputPath
		}
		err := processFile(inputPath, currentOutputPath, *percent)
		if err != nil {
			panic(err)
		}
	}
}
