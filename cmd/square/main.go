package main

import (
	"image"
	"log"
	"os"

	"image/jpeg"

	_ "github.com/enrichman/blank/pkg/image/square"
)

func main() {
	squareFile, err := os.Open("red.square")
	if err != nil {
		log.Fatal(err)
	}

	img, format, err := image.Decode(squareFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("format:", format)

	jpegFile, err := os.Create("red.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer jpegFile.Close()

	err = jpeg.Encode(jpegFile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
