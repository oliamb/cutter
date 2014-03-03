/*
crop contains the Crop function, used to retrieve a cropped version of the input image.
*/
package cutter

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
)

// CropParam struct is used to defined
// the way the crop should be realized.
//
// By default, it crop from the image center.
type Cutter struct {
	Width, Height, Left, Top int
}

// Retrieve an image representation that is a cropped view from the original image
func (c Cutter) Crop(img image.Image) (image.Image, error) {
	cr := image.Rect(c.Left, c.Top, c.Left+c.Width, c.Top+c.Height)
	cr = img.Bounds().Intersect(cr)
	result := image.NewRGBA(cr)
	for x, dx := cr.Min.X, cr.Max.X; x < dx; x += 1 {
		for y, dy := cr.Min.Y, cr.Max.Y; y < dy; y += 1 {
			result.Set(x, y, result.ColorModel().Convert(img.At(x, y)))
		}
	}
	return result, nil
}

func (c *Cutter) CropCenter(img image.Image) (image.Image, error) {
	c.Left = img.Bounds().Min.X + (img.Bounds().Dx()-c.Width)/2
	c.Top = img.Bounds().Min.Y + (img.Bounds().Dy()-c.Height)/2
	return c.Crop(img)
}
