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
	dc := gg.NewContext(2800, 900)

	// 画底图
	dc.SetColor(color.RGBA{0, 166, 94, 255})
	dc.DrawRectangle(0, 0, 2800, 900)
	dc.Fill()

	// 画边框
	dc.SetColor(color.RGBA{0, 0, 0, 255})
	dc.SetLineWidth(10)
	dc.DrawRoundedRectangle(20, 20, 2760, 860, 10)
	dc.Stroke()

	// 画矩形内部渐变图
	grad := gg.NewLinearGradient(1387, 875, 1387, 0)
	grad.AddColorStop(0.1, color.RGBA{43, 181, 121, 255})
	grad.AddColorStop(0.5, color.RGBA{149, 218, 188, 255})
	grad.AddColorStop(1, color.RGBA{233, 247, 241, 255})
	dc.SetFillStyle(grad)
	dc.MoveTo(25, 25)
	dc.LineTo(2775, 25)
	dc.LineTo(2775, 875)
	dc.LineTo(25, 875)
	dc.ClosePath()
	dc.Fill()

	// 添加字体(车牌号)
	if err := dc.LoadFontFace("./tiff-demo/demo13(后车牌)新/FZDHTJW.TTF", 500); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("闽", 200, 422, 0, 0.5)
	if err := dc.LoadFontFace("./tiff-demo/demo13(后车牌)新/车牌字体DIN1451.ttf", 500); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("C53602", 700, 422, 0, 0.5)

	img := dc.Image()
	f, _ := os.OpenFile("./tiff-demo/demo13(后车牌)新/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f, img, nil)
}
