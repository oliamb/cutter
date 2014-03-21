package cutter

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	_ "testing"
)

func ExampleCrop() {
	f, err := os.Open("fixtures/gopher.jpg")
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}

	cImg, err := Crop(img, Config{
		Height:  500,                 // height in pixel or Y ratio(see Ratio Option below)
		Width:   500,                 // width in pixel or X ratio
		Mode:    TopLeft,             // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{10, 10}, // Position of the top left point
		Options: 0,                   // Accepted Option: Ratio
	})

	if err != nil {
		log.Fatal("Cannot crop image:", err)
	}
	fmt.Println("cImg dimension:", cImg.Bounds())
	// Output: cImg dimension: (10,10)-(510,510)
}
