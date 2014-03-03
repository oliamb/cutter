package cutter

import (
	"fmt"
	"image"
	"os"
	_ "testing"
)

func ExampleCutter_Crop() {
	img := getGopherImage()

	c := Cutter{512, 512, 0, 0}
	r, err := c.Crop(img)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Bounds())
	// Output: (0,0)-(512,512)
}

func ExampleCutter_CenteredCrop() {
	img := getGopherImage()

	c := Cutter{
		Width:  512,
		Height: 512,
	}
	r, err := c.CropCenter(img)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Bounds())
	// Output: (544,462)-(1056,974)
}

func getGopherImage() image.Image {
	fi, err := os.Open("fixtures/gopher.jpg")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	img, _, err := image.Decode(fi)
	if err != nil {
		panic(err)
	}
	return img
}
