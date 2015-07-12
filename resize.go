package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	resizeImage("orig/monk.png", "result.png", 0.2)

}

/*
	resizeImage: resize the image and save it
*/
func resizeImage(inFilename string, outFilename string, factor float64) {
	//
	// print a message about what we are converting
	//
	fmt.Println("Converting", inFilename, "to", outFilename, "using factor", factor)

	//
	// Open the file and read in the image
	//
	infile, err := os.Open(inFilename)
	errorCheck(err)
	defer infile.Close()

	//
	//  load in the actual image data (decode image) and work out the new dimensions
	//
	srcImage, _, err := image.Decode(infile)
	errorCheck(err)

	srcImageWidth := srcImage.Bounds().Max.X
	newImageWidth := uint(int(factor * float64(srcImageWidth)))

	//
	// Resize the image
	//
	imgResized := resize.Resize(newImageWidth, 0, srcImage, resize.MitchellNetravali)

	//
	// save the new image
	//
	outfile, err := os.Create(outFilename)
	errorCheck(err)
	defer outfile.Close()

	png.Encode(outfile, imgResized)
}

/*
	errorCheck: Helper function to check for errors and log/panic on fail
*/
func errorCheck(e error) {
	if e != nil {
		log.Fatalf("%s", e)
		panic(e)
	}
}
