/*
@Time : 2021/11/11 19:54
@Author : Administrator
@Description :
@File : main
@Software: GoLand
*/
package main

import (
	"golang.org/x/image/tiff"
	"image"
	"image/color"
	"math/cmplx"
	"math/rand"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2907, 713
	)
	//f, _ := os.OpenFile("d.png", os.O_CREATE|os.O_RDWR, 0666)
	f, _ := os.OpenFile("./tiff-demo/demo02/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	//png.Encode(f, img) // NOTE: ignoring errors
	tiff.Encode(f,img,nil)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
			//return color.RGBA{uint8(rand.Intn(255)),uint8(rand.Intn(255)),uint8(rand.Intn(255)),255}
		}
	}
	return color.RGBA{uint8(rand.Intn(255)), 0, 0, 255}
}
