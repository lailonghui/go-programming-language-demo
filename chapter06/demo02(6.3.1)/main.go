/*
@Time : 2020/11/30 17:25
@Author : lai
@Description :
@File : main
*/
package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	//red := color.RGBA{255, 0, 0, 255}
	//blue := color.RGBA{0, 0, 255, 255}
	//var p = ColoredPoint{Point{1, 1}, red}
	//var q = ColoredPoint{Point{5, 4}, blue}
	//fmt.Println(p.Distance(q.Point)) // "5"
	//p.ScaleBy(2)
	//q.ScaleBy(2)
	//fmt.Println(p.Distance(q.Point)) // "10"
}
