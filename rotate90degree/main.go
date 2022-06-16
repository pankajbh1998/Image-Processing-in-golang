package main

import (
	"fmt"
	"image_processing/common"
	"image/png"
	"log"
	"os"
)

func main(){
	fmt.Println("Started the process Image Rotation")
	file,err := os.Open("img.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Started loading the Image")
	// Decode the file into image format
	img,err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	// get the pixels from image
	matrix := common.GetPixels(img)

	fmt.Println("Image loading successful")
	fmt.Println("Started Image Rotation")
	// perform operation to rotate the matrix
	matrix = matrixRotation90(matrix)

	fmt.Println("Image is Rotated Successfully")
	// create a file to save our rotated image
	op, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer op.Close()

	fmt.Println("Saving the new image.", "you can find it in output.png")
	// load/encode the image into the file
	err = png.Encode(op, common.GetImage(matrix))
	if err != nil {
		log.Fatal(err)
	}

}

func matrixRotation90 (input [][]common.Pixel) [][]common.Pixel {
	n,m := common.GetSize(input)
	ret := common.MakeArray(m,n)
	for i:=0;i<n;i++ {
		for j:=0;j<m;j++ {
			ret[j][n-1-i] = input[i][j]
		}
	}
	return ret
}