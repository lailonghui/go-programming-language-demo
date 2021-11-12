/*
@Time : 2021/11/11 19:54
@Author : Administrator
@Description :
@File : main
@Software: GoLand
*/
package main

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/tiff"
	"math"
	"os"
)


type Point struct {
	X, Y float64
}

func Polygon(n int, x, y, r float64) []Point {
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
	}
	return result
}

func main() {
	n := 5
	points := Polygon(n, 512, 512, 400)
	dc := gg.NewContext(1024, 1024)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()
	//dc.SavePNG("./tiff-demo/demo02/out.png")
	//dc.SavePNG("./tiff-demo/demo02/out.tiff")
	img := dc.Image()

	f, _ := os.OpenFile("./tiff-demo/demo02/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f,img,nil)
}

