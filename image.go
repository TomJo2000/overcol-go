package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Export_Cube(cube [][][]OkLAB) []image.Image {
	res := len(cube)

	images := make([]image.Image, res)

	for num, slice := range cube {
		img := image.NewRGBA(image.Rect(0, 0, res, res))
		for row := 0; row < res; row++ {
			for col := 0; col < res; col++ {
				R := uint8(slice[row][col].L * 256)
				G := uint8(slice[row][col].A * 256)
				B := uint8(slice[row][col].B * 256)
				A := uint8(255)

				color := color.RGBA{R, G, B, A}
				img.SetRGBA(col, row, color)
			}
		}
		images[num] = img
	}
	return images
}

func SaveImg(img image.Image, name string) error {
	var err error

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, img)

	return err
}
