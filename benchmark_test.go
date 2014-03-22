package cutter

import (
	"image"
	"testing"
)

/*
BenchmarkCrop is used to track the Crop with sharing memory.

Result on my laptop: 2000000		948 ns/op
*/
func BenchmarkCrop(b *testing.B) {
	img := getImage()

	c := Config{
		Width:  1000,
		Height: 1000,
		Mode:   TopLeft,
		Anchor: image.Point{100, 100},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Crop(img, c)
	}
}

/*
BenchmarkCropCopy is used to track the Crop with copy performance.

Below are the actual result on my laptop given each
optimization suggested by Nigel Tao: https://groups.google.com/forum/#!topic/golang-nuts/qxSpOOp1QOk

1. initial time on my Laptop: 														10	 210332414 ns/op
2. after inverting x and y in copy loop:                  10	 195377177 ns/op
3. after removing useless call to ColorModel().Convert(): 10	 193589075 ns/op
4. after replacing the two 'pixel' loops by a call to draw.Draw
   to obtains the cropped image:													20	  84960510 ns/op
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
	for i := 0; i < b.N; i++ {
		Crop(img, c)
	}
}
