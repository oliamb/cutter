/*
cutter contains the Crop function, used to retrieve a cropped version of the input image.
*/
package cutter

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

// AnchorMode is an enumeration of the position an anchor can represent.
type AnchorMode int

const (
	// TopLeft defines the Anchor Point as the top left of the cropped picture.
	TopLeft AnchorMode = iota
	// Centered defines the Anchor Point as the center of the cropped picture.
	Centered = iota
)

// List of Option
type Option int

const (
	// Use Width and Height as a ratio and keep as most of the image as possible
	Ratio = 1 << iota
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
	fmt.Println("maxBounds", maxBounds)
	size := c.computeSize(maxBounds, image.Point{c.Width, c.Height})
	fmt.Println("size", size)
	cr := c.computedCropArea(img.Bounds(), size)
	fmt.Println("cr", cr)
	cr = img.Bounds().Intersect(cr)
	fmt.Println("cr2", cr)
	result := image.NewRGBA(cr)
	for x, dx := cr.Min.X, cr.Max.X; x < dx; x += 1 {
		for y, dy := cr.Min.Y, cr.Max.Y; y < dy; y += 1 {
			result.Set(x, y, result.ColorModel().Convert(img.At(x, y)))
		}
	}
	return result, nil
}

func (c Cutter) maxBounds(bounds image.Rectangle) (r image.Rectangle) {
	if c.Mode == Centered {
		fmt.Println(c.Anchor.X-bounds.Min.X, bounds.Max.X-c.Anchor.X)
		anchor := c.centeredMin(bounds)
		w := min(anchor.X-bounds.Min.X, bounds.Max.X-anchor.X)
		h := min(anchor.Y-bounds.Min.Y, bounds.Max.Y-anchor.Y)
		r = image.Rect(anchor.X-w, anchor.Y-h, anchor.X+w, anchor.Y+h)
	} else {
		r = image.Rect(c.Anchor.X, c.Anchor.Y, bounds.Max.X, bounds.Max.Y)
	}
	return
}

// computeSize retrieve the effective size of the cropped image.
// It is defined by Height, Width, and Ratio option.
func (c Cutter) computeSize(bounds image.Rectangle, ratio image.Point) (p image.Point) {
	if c.Options&Ratio == Ratio {
		// Ratio option is on, so we take the biggest size available that fit the given ratio.
		if float64(ratio.X)/float64(bounds.Dx()) > float64(ratio.Y)/float64(bounds.Dy()) {
			p = image.Point{bounds.Dx(), (bounds.Dx() / ratio.X) * ratio.Y}
		} else {
			p = image.Point{(bounds.Dy() / ratio.Y) * ratio.X, bounds.Dy()}
		}
	} else {
		p = image.Point{ratio.X, ratio.Y}
	}
	return
}

// computedCropArea retrieve the theorical crop area.
// It is defined by Height, Width, Mode and
func (c Cutter) computedCropArea(bounds image.Rectangle, size image.Point) (r image.Rectangle) {
	min := bounds.Min
	switch c.Mode {
	case Centered:
		rMin := c.centeredMin(bounds)
		r = image.Rect(rMin.X-size.X/2, rMin.Y-size.Y/2, rMin.X+size.X/2, rMin.Y+size.Y/2)
	default: // TopLeft
		rMin := image.Point{min.X + c.Anchor.X, min.Y + c.Anchor.Y}
		r = image.Rect(rMin.X, rMin.Y, rMin.X+size.X, rMin.Y+size.Y)
	}
	return
}

func (c *Cutter) centeredMin(bounds image.Rectangle) (rMin image.Point) {
	min := bounds.Min
	if c.Anchor.X == 0 && c.Anchor.Y == 0 {
		rMin = image.Point{
			X: min.X + bounds.Dx()/2,
			Y: min.Y + bounds.Dy()/2,
		}
	} else {
		rMin = image.Point{
			X: min.X + c.Anchor.X,
			Y: min.Y + c.Anchor.Y,
		}
	}
	return
}

func min(a, b int) (r int) {
	if a < b {
		r = a
	} else {
		r = b
	}
	return
}
