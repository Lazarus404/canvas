package main

import "github.com/vuleetu/canvas"
import "os"
import "io"
import "log"
//import "time"

func main() {
	cv := canvas.New()
	defer cv.Destroy()

    file, err := os.Open("input/1.jpg")
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
        cv1 :=cv.Clone()
		// Photo auto orientation based on EXIF tags.
		cv1.AutoOrientate()
        cv1.Thumbnail(180, 180)
        cv1.SetImageFormat("PNG")
        bin3 := cv1.GetImageBlob()
        tf3, err := os.Create("output/100.png")
        defer tf3.Close()
        if err != nil {
            log.Println("create file failed")
            return
        }
        log.Println("size is", len(bin3))
        tf3.Write(bin3)
        // Saving the thumbnail to disk.
        //cv1.Write("output/example-thumbnail_2012.png")
        log.Println("The format is", cv1.Format())

        log.Println("Width is", cv1.Width(), "height is", cv1.Height())
	} else {
        log.Println("Open failed")
    }
}
