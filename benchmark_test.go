package cutter

import (
	"image"
	"testing"
)

/*
Initial time on my MacBook Air: 23 ns/op
By inverting x and y in copy loop: 0.09 ns/op
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
	r, _ := Crop(img, c)
	r.Bounds()
}
