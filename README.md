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

## Documentation

You can read ``gosexy/canvas`` documentation from a terminal

    $ go doc github.com/gosexy/canvas

Or you can [browse it](http://go.pkgdoc.org/github.com/gosexy/canvas) online.
