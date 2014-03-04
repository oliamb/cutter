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
	maxBounds := c.maxBounds(img.Bounds())
	size := c.computeSize(maxBounds, image.Point{c.Width, c.Height})
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

func (c Cutter) maxBounds(bounds image.Rectangle) image.Rectangle {
	if c.Mode == Centered {
		w := min(c.Anchor.X-bounds.Min.X, bounds.Max.X-c.Anchor.X)
		h := min(c.Anchor.Y-bounds.Min.Y, bounds.Max.Y-c.Anchor.Y)
		return image.Rect(c.Anchor.X-w, c.Anchor.Y-h, c.Anchor.X+w, c.Anchor.Y+h)
	} else {
		return image.Rect(c.Anchor.X, c.Anchor.Y, bounds.Max.X, bounds.Max.Y)
	}
}

// computeSize retrieve the effective size of the cropped image.
// It is defined by Height, Width, and Ratio option.
func (c Cutter) computeSize(bounds image.Rectangle, ratio image.Point) image.Point {
	if c.Options&Ratio == Ratio {
		// Ratio option is on, so we take the biggest size available that fit the given ratio.
		if float64(ratio.X)/float64(bounds.Dx()) > float64(ratio.Y)/float64(bounds.Dy()) {
			return image.Point{bounds.Dx(), (bounds.Dx() / ratio.X) * ratio.Y}
		} else {
			return image.Point{(bounds.Dy() / ratio.Y) * ratio.X, bounds.Dy()}
		}
	} else {
		return image.Point{ratio.X, ratio.Y}
	}
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
