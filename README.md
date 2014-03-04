Cutter
======

A Go library to crop images, works great in combination with [github.com/nfnt/resize](https://github.com/nfnt/resize)

Usage
=====
Import package with

```go
import "github.com/oliamb/cutter"
```

The cutter package provides a ```Cutter``` which in turn provide a ```Crop``` function that realize the crop operation on the passed image.

Realize a cut from top left corner:

```go
c := Cutter{
  Width: 250,
  Height: 500,
}
img, err := c.Crop(baseImage)
```

Specify the top left position:

```go
c := Cutter{
  Width: 250,
  Height: 500,
  Anchor: image.Point{100, 100},
}
img, err := c.Crop(baseImage)
```

Make a centered crop:
```go
c := Cutter{
  Width: 250,
  Height: 500,
  Mode: Centered,
}
img, err := c.Crop(baseImage)
```

Use ratio instead of specific width and height:
```go
c := Cutter{
  Width: 4,
  Height: 3,
  Mode: Centered,
  Options: Ratio,
}
img, err := c.Crop(baseImage)
```

Contributing
============

Fork it, and make a pull request. If you plan to modify the API, let's disscuss it first.

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
