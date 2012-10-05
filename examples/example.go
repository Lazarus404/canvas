package main

import "github.com/vuleetu/canvas"
import "log"

func main() {
	cv := canvas.New()
	defer cv.Destroy()

	// Opening some image from disk.
	opened := cv.Open("input/example.jpg")

	if opened {

		// Photo auto orientation based on EXIF tags.
		//cv.AutoOrientate()

		// Creating a squared thumbnail
		cv.Thumbnail(100, 100)

		// Saving the thumbnail to disk.
		cv.Write("output/example-thumbnail_2012.png")
        log.Println("The format is", cv.Format())
	}
}
