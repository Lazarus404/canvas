package main

import "github.com/vuleetu/canvas"
import "os"
import "io"
import "log"

func main() {
	cv := canvas.New()
	defer cv.Destroy()

    file, err := os.Open("input/example.png")
    if err != nil {
        log.Println("Open file failed")
        return
    }
    data := make([]byte, 0)
    buf := make([]byte, 1024)

    for {
        n, err := file.Read(buf)
        data = append(data, buf[0:n]...)
        if err == io.EOF {
            //done
            break
        } else if err != nil {
            log.Println("error is", err.Error())
            break;
        }
    }
    log.Println("File size is", len(data))
	opened := cv.OpenBlob(data)
	// Opening some image from disk.
	//opened := cv.Open("input/example.jpg")

	if opened {

		// Photo auto orientation based on EXIF tags.
		//cv.AutoOrientate()

		// Creating a squared thumbnail
		cv.Thumbnail(100, 100)

        bin := cv.GetImageBlob()
        tf, err := os.Create("output/example_read_from_mem.png")
        if err != nil {
            log.Println("create file failed")
            return
        }
        tf.Write(bin)
		// Saving the thumbnail to disk.
		//cv.Write("output/example-thumbnail_2012.png")
        log.Println("The format is", cv.Format())
	} else {
        log.Println("Open failed")
    }
}
