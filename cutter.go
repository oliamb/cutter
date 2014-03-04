/*
cutter contains the Crop function, used to retrieve a cropped version of the input image.
*/
package cutter

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
)

// An enumeration of the position an anchor can represent.
type AnchorMode int

const (
	TopLeft  AnchorMode = iota
	Centered            = iota
)

type Option int

const (
	Ratio = 1 << iota // Use width and height as a ratio and keep as most of the image as possible
)

// CropParam struct is used to defined
// the way the crop should be realized.
type Cutter struct {
	Width, Height int
	Anchor        image.Point // The Anchor Point in the source image
	Mode          AnchorMode  // Which point in the resulting image the Anchor Point is referring to
	Options       Option
}

// Retrieve an image representation that is a cropped view from the original image
func (c Cutter) Crop(img image.Image) (image.Image, error) {
	var size image.Point
	if c.Options&Ratio == Ratio {
		size = image.Point{img.Bounds().Dx(), (img.Bounds().Dx() / c.Width) * c.Height}
	} else {
		size = image.Point{c.Width, c.Height}
	}
	cr := c.computedCropArea(img, size)
	cr = img.Bounds().Intersect(cr)
	result := image.NewRGBA(cr)
	for x, dx := cr.Min.X, cr.Max.X; x < dx; x += 1 {
		for y, dy := cr.Min.Y, cr.Max.Y; y < dy; y += 1 {
			result.Set(x, y, result.ColorModel().Convert(img.At(x, y)))
		}
	}
	return result, nil
}

// computedCropArea retrieve the theorical crop area.
// It is defined by Height, Width, Mode and
func (c Cutter) computedCropArea(img image.Image, size image.Point) image.Rectangle {
	min := img.Bounds().Min
	switch c.Mode {
	case Centered:
		var rMin image.Point
		if c.Anchor.X == 0 && c.Anchor.Y == 0 {
			rMin = image.Point{
				X: min.X + img.Bounds().Dx()/2,
				Y: min.Y + img.Bounds().Dy()/2,
			}
		} else {
			rMin = image.Point{
				X: min.X + c.Anchor.X,
				Y: min.Y + c.Anchor.Y,
			}
		}
		return image.Rect(rMin.X-size.X/2, rMin.Y-size.Y/2, rMin.X+size.X/2, rMin.Y+size.Y/2)
	default: // TopLeft
		rMin := image.Point{min.X + c.Anchor.X, min.Y + c.Anchor.Y}
		return image.Rect(rMin.X, rMin.Y, rMin.X+size.X, rMin.Y+size.Y)
	}
}
