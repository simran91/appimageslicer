package main

import (
        "image"
        "image/color"
        "image/png"
        "math"
        "os"
    )

    func main() {
        filename := "orig/monk.png"
        infile, err := os.Open(filename)
        if err != nil {
            // replace this with real error handling
            panic(err.Error())
        }
        defer infile.Close()

        // Decode will figure out what type of image is in the file on its own.
        // We just have to be sure all the image packages we want are imported.
        src, _, err := image.Decode(infile)
        if err != nil {
            // replace this with real error handling
            panic(err.Error())
        }

        // Create a new grayscale image
        bounds := src.Bounds()
        w, h := bounds.Max.X, bounds.Max.Y
        gray := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
        for x := 0; x < w; x++ {
            for y := 0; y < h; y++ {
                oldColor := src.At(x, y)
                r, g, b, _ := oldColor.RGBA()
                avg := 0.2125*float64(r) + 0.7154*float64(g) + 0.0721*float64(b)
                grayColor := color.Gray{uint8(math.Ceil(avg))}
                gray.Set(x, y, grayColor)
            }
        }

        // Encode the grayscale image to the output file
        outfilename := "result.png"
        outfile, err := os.Create(outfilename)
        if err != nil {
            // replace this with real error handling
            panic(err.Error())
        }
        defer outfile.Close()
        png.Encode(outfile, gray)
    }
