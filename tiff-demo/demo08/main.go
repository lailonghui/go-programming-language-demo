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
	dc := gg.NewContext(2907, 713)

	// 画底图
	dc.SetColor(color.RGBA{0, 166, 94, 255})
	dc.DrawRectangle(0, 0, 2907, 713)
	dc.Fill()

	// 画边框
	dc.SetColor(color.RGBA{0, 0, 0, 255})
	dc.SetLineWidth(10)
	dc.DrawRoundedRectangle(20, 20, 2866, 673, 10)
	//dc.MoveTo(20, 20)
	//dc.LineTo(2887, 20)
	//dc.LineTo(2887, 693)
	//dc.LineTo(20, 693)
	//dc.ClosePath()
	dc.Stroke()

	// 画矩形内部渐变图
	grad := gg.NewLinearGradient(1433, 673, 1433, 0)
	grad.AddColorStop(0, color.RGBA{43, 181, 121, 255})
	grad.AddColorStop(1, color.RGBA{149, 218, 188, 255})
	grad.AddColorStop(5, color.RGBA{233, 247, 241, 255})
	dc.SetFillStyle(grad)
	dc.MoveTo(26, 26)
	dc.LineTo(2880, 26)
	dc.LineTo(2880, 686)
	dc.LineTo(26, 686)
	dc.ClosePath()
	dc.Fill()

	// 添加字体
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./tiff-demo/demo08/JingDianZongYiTiJian-1.ttf", 300); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("泉工-C02D001", 100, 343, 0, 0.5)

	// 添加车辆信息二维码
	im, err := gg.LoadImage("./tiff-demo/demo08/img.png")
	if err != nil {
		log.Fatal(err)
	}
	//dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	//dc.Clip()
	dc.DrawImage(im, 2300, 0)

	img := dc.Image()
	f, _ := os.OpenFile("./tiff-demo/demo08/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f, img, nil)
}
