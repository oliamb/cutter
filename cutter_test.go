package cutter

import (
	"image"
	"os"
	"testing"
)

func TestCrop(t *testing.T) {
	img := getImage()

	c := Config{
		Width:  512,
		Height: 400,
	}
	r, err := Crop(img, c)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 512 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 400 {
		t.Error("Bad width should be 400 but is", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 0 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func TestCrop_Centered(t *testing.T) {
	img := getImage()

	c := Config{
		Width:  512,
		Height: 400,
		Mode:   Centered,
	}
	r, err := Crop(img, c)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 512 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 400 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 544 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 518 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func TestCrop_Centered_Ratio_WithoutAnchorPosition(t *testing.T) {
	// (0,0)-(64,64) 32 64 (32,32)-(32,32)
	img := image.NewGray(image.Rect(0, 0, 64, 64))
	c := Config{
		Width:   32,
		Height:  64,
		Mode:    Centered,
		Options: Ratio,
	}
	r, err := Crop(img, c)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 32 {
		t.Error("Bad Width", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 64 {
		t.Error("Bad Height", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 16 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func TestCutter_Crop_TooBigArea(t *testing.T) {
	img := getImage()

	c := Config{
		Width:  2000,
		Height: 2000,
		Anchor: image.Point{100, 100},
	}
	r, err := Crop(img, c)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 1500 {
		t.Error("Bad width should be 1500 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 1337 {
		t.Error("Bad width should be 1337 but is", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 100 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 100 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func TestCrop_TooBigAreaFromCenter(t *testing.T) {
	img := getImage()

	c := Config{
		Width:  1000,
		Height: 2000,
		Anchor: image.Point{1200, 100},
		Mode:   Centered,
	}
	r, err := Crop(img, c)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 900 {
		t.Error("Bad width should be 900 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 1100 {
		t.Error("Bad width should be 1100 but is", r.Bounds().Dy(), r.Bounds())
	}
	if r.Bounds().Min.X != 700 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func TestCrop_OptionRatio(t *testing.T) {
	img := getImage()

	c := Config{
		Width:   4,
		Height:  3,
		Anchor:  image.Point{},
		Mode:    TopLeft,
		Options: Ratio,
	}

	r, err := Crop(img, c)
	if err != nil {
		t.Error(err)
	}
	if r.Bounds().Min.X != 0 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
	if r.Bounds().Dx() != 1600 {
		t.Error("Bad Width", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 1200 {
		t.Error("Bad Height", r.Bounds().Dy(), r.Bounds())
	}
}

func TestCutter_Crop_OptionRatio_Inverted(t *testing.T) {
	img := getImage()

	c := Config{
		Width:   3,
		Height:  4,
		Anchor:  image.Point{},
		Mode:    TopLeft,
		Options: Ratio,
	}

	r, err := Crop(img, c)
	if err != nil {
		t.Error(err)
	}
	if r.Bounds().Min.X != 0 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
	if r.Bounds().Dy() != 1437 {
		t.Error("Bad Height", r.Bounds().Dy(), r.Bounds())
	}
	if r.Bounds().Dx() != 1077 {
		t.Error("Bad Width", r.Bounds().Dx())
	}
}

func TestCutter_Crop_OptionRatio_DecentredAnchor_Overflow(t *testing.T) {
	img := getImage()
	c := Config{
		Width:   3,
		Height:  4,
		Anchor:  image.Point{100, 80},
		Mode:    Centered,
		Options: Ratio,
	}

	r, err := Crop(img, c)
	if err != nil {
		t.Error(err)
	}
	if r.Bounds().Min.X != 40 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
	if r.Bounds().Dy() != 160 {
		t.Error("Bad Height", r.Bounds().Dy(), r.Bounds())
	}
	if r.Bounds().Dx() != 120 {
		t.Error("Bad Width", r.Bounds().Dx())
	}
}

func TestCropForceCopy(t *testing.T) {
	img := getImage()

	c := Config{
		Width:   512,
		Height:  400,
		Options: Copy,
	}
	r, err := Crop(img, c)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 512 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 400 {
		t.Error("Bad width should be 400 but is", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 0 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func getImage() image.Image {
	return image.NewGray(image.Rect(0, 0, 1600, 1437))
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
