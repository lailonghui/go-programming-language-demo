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
	"os"
)



func main() {
	//n := 5
	//points := Polygon(n, 512, 512, 400)
	//dc := gg.NewContext(1024, 1024)
	//dc.SetHexColor("fff")
	//dc.Clear()
	//for i := 0; i < n+1; i++ {
	//	index := (i * 2) % n
	//	p := points[index]
	//	dc.LineTo(p.X, p.Y)
	//}
	//dc.SetRGBA(0, 0.5, 0, 1)
	//dc.SetFillRule(gg.FillRuleEvenOdd)
	//dc.FillPreserve()
	//dc.SetRGBA(0, 1, 0, 0.5)
	//dc.SetLineWidth(16)
	//dc.Stroke()
	//dc.SavePNG("./tiff-demo/demo02/out.png")
	//dc.SavePNG("./tiff-demo/demo02/out.tiff")

	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	img := dc.Image()

	f, _ := os.OpenFile("./tiff-demo/demo03/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f,img,nil)
}

