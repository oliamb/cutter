package cutter

import (
	"image"
	"testing"
)

func BenchmarkCrop(b *testing.B) {
	img := getImage()

	c := Config{
		Width:  1000,
		Height: 1000,
		Mode:   TopLeft,
		Anchor: image.Point{100, 100},
	}
	b.ResetTimer()
	r, _ := Crop(img, c)
	r.Bounds()
}

/*
BenchmarkCrop is used to track the Crop performance.

Below are the actual result on my laptop given each
optimization suggested by Nigel Tao: https://groups.google.com/forum/#!topic/golang-nuts/qxSpOOp1QOk

1. initial time on my Laptop: 23 ns/op
2. after inverting x and y in copy loop: 0.09 ns/op
3. after removing useless call to ColorModel().Convert(): 0.08 ns/op
4. after replacing the two 'pixel' loops by a call to draw.Draw
   to obtains the cropped image: 0.04 ns/op
*/
func BenchmarkCropCopy(b *testing.B) {
	img := getImage()

	c := Config{
		Width:   1000,
		Height:  1000,
		Mode:    TopLeft,
		Anchor:  image.Point{100, 100},
		Options: Copy,
	}
	b.ResetTimer()
	r, _ := Crop(img, c)
	r.Bounds()
}
