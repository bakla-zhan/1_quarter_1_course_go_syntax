package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var red color.Color = color.RGBA{200, 30, 30, 255}

func main() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}

	for x := 0; x <= 200; x++ {
		y := 20
		rectImg.Set(x, y, red)
	}

	for x := 0; x <= 200; x++ {
		y := 180
		rectImg.Set(x, y, red)
	}

	for y := 0; y <= 200; y++ {
		x := 20
		rectImg.Set(x, y, red)
	}

	for y := 0; y <= 200; y++ {
		x := 180
		rectImg.Set(x, y, red)
	}

	defer file.Close()
	png.Encode(file, rectImg)
}
