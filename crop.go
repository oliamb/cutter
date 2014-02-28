/*
crop contains the Crop function, used to retrieve a cropped version of the input image.
 */
package crop

import (
	"image"
)

// CropParam struct is used to defined
// the way the crop should be realized.
//
// By default, it crop from the image center.
type CropParam struct {
  Width int
  Height int
}

// Retrieve an image representation that is a cropped view from the original image
func Crop(img image.Image, p CropParam) image.Image, error {
  
}
