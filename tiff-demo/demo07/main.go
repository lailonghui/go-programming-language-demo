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
	"image/color"
	"os"
)

func main() {
	dc := gg.NewContext(2907, 713)

	grad := gg.NewLinearGradient(1453, 713, 1453, 0)

	//grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	//grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	//grad.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})

	grad.AddColorStop(0, color.RGBA{43, 181, 121, 255})
	grad.AddColorStop(1, color.RGBA{149, 218, 188, 255})
	grad.AddColorStop(5, color.RGBA{233, 247, 241, 255})

	//RGB(43,181,121)
	//RGB(149,218,188)
	//RGB(233,247,241)

	//dc.SetColor(color.White)
	//dc.DrawRectangle(20, 20, 400-20, 300)
	//dc.Stroke()
	//
	//dc.SetStrokeStyle(grad)
	//dc.SetLineWidth(4)
	//dc.MoveTo(10, 10)
	//dc.LineTo(410, 10)
	//dc.LineTo(410, 100)
	//dc.LineTo(10, 100)
	//dc.ClosePath()
	//dc.Stroke()

	dc.SetFillStyle(grad)
	//dc.MoveTo(0, 0)
	//dc.LineTo(0, 713)
	//dc.LineTo(2907, 713)
	//dc.LineTo(2907, 0)

	dc.MoveTo(0, 0)
	dc.LineTo(2907, 0)
	dc.LineTo(2907, 713)
	dc.LineTo(0, 713)
	//dc.LineTo(0, 713)

	//dc.LineTo(0, 713)
	dc.ClosePath()
	dc.Fill()
	img := dc.Image()
	f, _ := os.OpenFile("./tiff-demo/demo07/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f, img, nil)
}
