Cutter
======

[![Build Status](https://travis-ci.org/oliamb/cutter.png?branch=master)](https://travis-ci.org/oliamb/cutter)

A Go library to crop images, works great in combination with [github.com/nfnt/resize](https://github.com/nfnt/resize)

Usage
=====
Import package with

    import "github.com/oliamb/cutter"

Package cutter provides a function to crop image.

By default, the original image will be cropped at the
given size from the top left corner.

    croppedImg, err := cutter.Crop(img, cutter.Config{
      Width:  250,
      Height: 500,
    })

It is possible to specify the top left position:

    croppedImg, err := cutter.Crop(img, cutter.Config{
      Width:  250,
      Height: 500,
      Anchor: image.Point{100, 100},
      Mode:   TopLeft, // optional, default value
    })

The Anchor property can represents the center of the cropped image
instead of the top left corner:


    croppedImg, err := cutter.Crop(img, cutter.Config{
      Width: 250,
      Height: 500,
      Mode: Centered,
    })

The default crop use the specified dimension, but it is possible
to use Width and Heigth as a ratio instead. In this case,
the resulting image will be as big as possible to fit the asked ratio
from the anchor position.

    croppedImg, err := cutter.Crop(baseImage, cutter.Config{
      Width: 4,
      Height: 3,
      Mode: Centered,
      Options: Ratio,
    })

Contributing
============

Fork it, code and make a pull request. If you plan to modify the API, let's disscuss it first.

License
=======

Copyright © 2014 Urturn SA, Olivier Amblet, Cutter is released under the MIT license

    The MIT License (MIT)

    Copyright (c) 2014 Olivier Amblet

    Permission is hereby granted, free of charge, to any person obtaining a copy of
    this software and associated documentation files (the "Software"), to deal in
    the Software without restriction, including without limitation the rights to
    use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
    the Software, and to permit persons to whom the Software is furnished to do so,
    subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
    FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
    COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
    IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
    CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

Credits
=======

Test Picture: Gopher picture from Heidi Schuyt, http://www.flickr.com/photos/hschuyt/7674222278/,
© copyright Creative Commons(http://creativecommons.org/licenses/by-nc-sa/2.0/)

Thanks to Urturn for the time allocated to develop the library.
