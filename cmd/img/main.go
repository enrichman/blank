package main

import (
	"image"
	"log"
	"os"

	"image/jpeg"
	_ "image/png"
)

func main() {
	imageFile, err := os.Open("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	img, format, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("format:", format)

	imgJpeg, err := os.Create("gopher.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgJpeg.Close()

	err = jpeg.Encode(imgJpeg, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
