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
	"log"
	"os"
)

func main() {
	dc := gg.NewContext(4350, 870)

	// 画底图
	dc.SetColor(color.RGBA{0, 166, 94, 255})
	dc.DrawRectangle(0, 0, 4350, 870)
	dc.Fill()

	// 画边框
	dc.SetColor(color.RGBA{0, 0, 0, 255})
	dc.SetLineWidth(10)
	dc.DrawRoundedRectangle(20, 20, 4310, 830, 10)
	dc.Stroke()

	// 画矩形内部渐变图
	grad := gg.NewLinearGradient(2155, 850, 2155, 0)
	grad.AddColorStop(0.1, color.RGBA{43, 181, 121, 255})
	grad.AddColorStop(0.5, color.RGBA{149, 218, 188, 255})
	grad.AddColorStop(1, color.RGBA{233, 247, 241, 255})
	dc.SetFillStyle(grad)
	dc.MoveTo(26, 26)
	dc.LineTo(4324, 26)
	dc.LineTo(4324, 844)
	dc.LineTo(26, 844)
	dc.ClosePath()
	dc.Fill()

	// 添加字体(工号牌)
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./tiff-demo/demo10(后车牌)/FZDHTJW.TTF", 350); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("泉工", 100, 422, 0, 0.5)
	dc.DrawStringAnchored("·", 700, 422, 0, 0.5)

	if err := dc.LoadFontFace("./tiff-demo/demo09(前车牌)/车牌字体DIN1451.ttf", 380); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("C02D001", 950, 422, 0, 0.5)

	// 添加车辆信息二维码
	im, err := gg.LoadImage("./tiff-demo/demo10(后车牌)/img.png")
	if err != nil {
		log.Fatal(err)
	}
	dc.DrawImage(im, 2350, 272)

	// 添加字体(车牌号)
	if err := dc.LoadFontFace("./tiff-demo/demo10(后车牌)/FZDHTJW.TTF", 270); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("闽C", 2800, 422, 0, 0.5)
	dc.DrawStringAnchored("·", 3200, 422, 0, 0.5)
	if err := dc.LoadFontFace("./tiff-demo/demo09(前车牌)/车牌字体DIN1451.ttf", 300); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("53602", 3400, 422, 0, 0.5)

	img := dc.Image()
	f, _ := os.OpenFile("./tiff-demo/demo10(后车牌)/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f, img, nil)
}
