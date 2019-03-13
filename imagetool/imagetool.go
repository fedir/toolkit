package imagetool

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
)

func ResizeImageJpgFileToFile(pathFrom string, pathTo string, width uint, height uint) {

	// Open the original image
	reader, err := os.Open(pathFrom)
	if err != nil {
		fmt.Println("Error image decode")
	}
	defer reader.Close()
	originalImage, _, err := image.Decode(reader)

	// Create the resized image
	newImage := resize.Resize(width, height, originalImage, resize.Lanczos3)
	f, err := os.Create(pathTo)
	if err != nil {
		fmt.Println("Error encoding image")
	}
	defer f.Close()
	err = jpeg.Encode(f, newImage, nil)

}
