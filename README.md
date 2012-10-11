## gosexy/canvas

``gosexy/canvas`` is an image processing library based on ImageMagick's MagickWand, for the Go programming language.

This branch is compatible with ImageMagick 6.6.x, like the one included with Debian Squeeze.

Please check the [master branch](https://github.com/gosexy/canvas/tree/master) for ImageMagick 6.7.x.

## Requeriments

Get the MagickWand development package

    $ sudo aptitude install libmagickwand-dev

## Installation

    $ cd $GOPATH
    $ mkdir -p github.com/gosexy
    $ cd github.com/gosexy
    $ git clone -b squeeze git://github.com/gosexy/canvas
    $ go install github.com/gosexy/canvas

## Usage

    package main

    import "github.com/gosexy/canvas"

    func main() {
      cv := canvas.New()
      defer cv.Destroy()

      // Opening some image from disk.
      opened := cv.Open("examples/input/example.png")

      if opened {

        // Photo auto orientation based on EXIF tags.
        cv.AutoOrientate()

        // Creating a squared thumbnail
        cv.Thumbnail(100, 100)

        // Saving the thumbnail to disk.
        cv.Write("examples/output/example-thumbnail.png")

      }
    }
=======
## Requeriments

### Mac OSX

The ImageMagick's header files are required. If you're using ``brew`` the installation is straightforward.

```sh
$ brew install imagemagick
```

### Debian

Debian has an old version of MagickWand (6.6.x), this binding was built against 6.7.x. Please check out the
[squeeze branch](https://github.com/gosexy/canvas/tree/squeeze) to get a version that works on Debian Squeeze and
probably other debian-based distros. This may not be required for Ubuntu.

### Arch Linux

Arch Linux already has a recent version of MagickWand.

```sh
$ sudo pacman -S extra/imagemagick
```

### Other Operative Systems

Please, follow the [install from source](http://imagemagick.com/script/install-source.php?ImageMagick=9uv1bcgofrv21mhftmlk4v1465) tutorial.

## Installation

After installing ImageMagick's header files, pull gocanvas from github:

```sh
$ go get github.com/gosexy/canvas
```

## Updating

After installing, you can use `go get -u github.com/gosexy/canvas` to keep up to date.

## Usage

```go
package main

import "github.com/gosexy/canvas"

func main() {
  cv := canvas.New()
  defer cv.Destroy()

  // Opening some image from disk.
  opened := cv.Open("examples/input/example.png")

  if opened {

    // Photo auto orientation based on EXIF tags.
    cv.AutoOrientate()

    // Creating a squared thumbnail
    cv.Thumbnail(100, 100)

    // Saving the thumbnail to disk.
    cv.Write("examples/output/example-thumbnail.png")

  }
}
```
    $ go doc github.com/gosexy/canvas

Or you can [browse it](http://go.pkgdoc.org/github.com/gosexy/canvas) online.
