package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x, y, width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(img.x, img.y, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x * y)

	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{0, 0, 300, 300}
	pic.ShowImage(m)
}
