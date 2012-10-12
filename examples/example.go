package main

import "github.com/vuleetu/canvas"
import "os"
import "io"
import "log"
//import "time"

func main() {
	cv := canvas.New()
	defer cv.Destroy()

    file, err := os.Open("input/example.jpg")
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
        cv1 := cv.Clone()
        cv1.Thumbnail(20, 20)
        cv1.SetImageFormat("PNG")
        bin3 := cv1.GetImageBlob()
        tf3, err := os.Create("output/example_read_from_mem_30_30.png")
        defer tf3.Close()
        if err != nil {
            log.Println("create file failed")
            return
        }
        log.Println("size is", len(bin3))
        tf3.Write(bin3)
        // Saving the thumbnail to disk.
        //cv.Write("output/example-thumbnail_2012.png")
        log.Println("The format is", cv1.Format())

        log.Println("Width is", cv.Width(), "height is", cv.Height())
        cv.Thumbnail(1000, 2000)
        cv.SetImageFormat("PNG")
        bin2 := cv.GetImageBlob()
        tf2, err := os.Create("output/example_read_from_mem_20_20.png")
        defer tf2.Close()
        if err != nil {
            log.Println("create file failed")
            return
        }
        log.Println("size is", len(bin2))
        tf2.Write(bin2)
        // Saving the thumbnail to disk.
        //cv.Write("output/example-thumbnail_2012.png")
        log.Println("The format is", cv.Format())

	} else {
        log.Println("Open failed")
    }
}
