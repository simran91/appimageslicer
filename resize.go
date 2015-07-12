package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	//
	// Define the factors we want for the resizings
	//
	factors := [...]string{"0.5", "0.25", "1"}
	
	//
	// clean auto-dest directory
	//
	err := os.RemoveAll("auto-dest")
	errorCheck(err)

	//
	// foreach factor
	//
	for _, factor := range factors {
		//
		// convert the factor string to a float64
		//
 		factorFloat, err := strconv.ParseFloat(factor, 64)
 		errorCheck(err)

		//
		// Create the factor destination directory
		//
		fmt.Println("Factoring for size", factor)
		destPath := fmt.Sprintf("auto-dest/%sx", factor)
		err = os.MkdirAll(destPath, 0777)
		errorCheck(err)		

		//
		// foreach file in the original directory, resize it!
		//
		files, err := ioutil.ReadDir("orig")
		errorCheck(err)

		for _, file := range files {
			origFilename := file.Name() // eg. surfboard.png
			origFilenameWithDir := fmt.Sprintf("orig/%s", origFilename) // eg. "orig/surfboard.png"
			destFilenameWitDir := fmt.Sprintf("%s/%s", destPath, origFilename)

			//
			// resize the image
			//
			resizeImage(origFilenameWithDir, destFilenameWitDir, factorFloat)
		}

	}


	// resizeImage("orig/monk.png", "result.png", float64(0.2))

}

/*
	resizeImage: resize the image and save it
*/
func resizeImage(inFilename string, outFilename string, factor float64) {
	//
	// print a message about what we are converting
	//
	fmt.Println("\t\t", inFilename, "=>", outFilename, "using factor", factor)

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
