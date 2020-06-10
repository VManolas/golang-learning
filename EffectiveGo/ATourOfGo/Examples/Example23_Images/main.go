// Example 23a Images
// [Package image](https://golang.org/pkg/image/#Image) defines the Image interface:
// package image

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle //Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.
//     At(x, y int) color.Color
// }
// The color.Color and color.Model types are also interfaces, but we'll ignore that
// by using the predefined implementations color.RGBA and color.RGBAModel.
// These interfaces and types are specified by the [image/color package](https://golang.org/pkg/image/color/)
// package main

// import (
// 	"fmt"
// 	"image"
// )

// func main() {
// 	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
// 	fmt.Println(m.Bounds())
// 	fmt.Println(m.At(0, 0).RGBA())
// }

// Example 23b Exercise Images
// Define your own Image type, implement the necessary methods, and call pic.ShowImage.
// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
// ColorModel should return color.RGBAModel.
// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height  int
	colorB, colorA uint8
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img *Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		//return uint8(x*y)
		//return uint8((x+y) / 2)
		return uint8(x ^ y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}

	// 	return color.RGBA{uint8(x), uint8(y), img.colorB, img.colorA}

	// 	return color.RGBA{uint8(x * y), uint8(x ^ y), uint8((x + y) / 2), 255}

	// 	return color.RGBA{uint8(x), uint8(x), uint8(y), uint8(y)}
}

func main() {
	m := &Image{256, 64, 255, 255}
	pic.ShowImage(m)
}
