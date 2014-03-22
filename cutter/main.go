package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {
	s := len(os.Args)
	if s < 4 {
		usage()
		return
	}
	fmt.Println("Args", os.Args)

	inPath := os.Args[s-2]
	fi, err := os.Open(inPath)
	if err != nil {
		log.Fatal("Cannot open input file '", inPath, "':", err)
	}

	outPath := os.Args[s-1]
	fo, err := os.Create(outPath)
	if err != nil {
		log.Fatal("Cannot create output file '", outPath, "':", err)
	}

	img, _, err := image.Decode(fi)
	if err != nil {
		log.Fatal("Cannot decode image at '", inPath, "':", err)
	}

	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  1000,                  // height in pixel or Y ratio(see Ratio Option below)
		Width:   1000,                  // width in pixel or X ratio
		Mode:    cutter.TopLeft,        // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{100, 100}, // Position of the top left point
		Options: 0,                     // Accepted Option: Ratio
	})
	if err != nil {
		log.Fatal("Cannot crop image:", err)
	}

	switch filepath.Ext(outPath) {
	case ".png":
		err = png.Encode(fo, cImg)
	case ".jpg":
		err = jpeg.Encode(fo, cImg, &jpeg.Options{})
	default:
		err = errors.New("Unsupported format: " + filepath.Ext(outPath))
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Image saved to", outPath)
}

func usage() {
	flag.Usage()
}
