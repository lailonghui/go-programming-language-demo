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
	dc := gg.NewContext(568, 1700)

	// 画底图
	dc.SetColor(color.RGBA{0, 166, 94, 255})
	dc.DrawRectangle(0, 0, 568, 1700)
	dc.Fill()

	// 画边框
	dc.SetColor(color.RGBA{0, 0, 0, 255})
	dc.SetLineWidth(10)
	dc.DrawRoundedRectangle(20, 20, 528, 1660, 10)
	dc.Stroke()

	// 画矩形内部渐变图
	grad := gg.NewLinearGradient(264, 1680, 264, 0)
	grad.AddColorStop(0.1, color.RGBA{43, 181, 121, 255})
	grad.AddColorStop(0.5, color.RGBA{149, 218, 188, 255})
	grad.AddColorStop(1, color.RGBA{233, 247, 241, 255})
	dc.SetFillStyle(grad)
	dc.MoveTo(26, 26)
	dc.LineTo(542, 26)
	dc.LineTo(542, 1674)
	dc.LineTo(26, 1674)
	dc.ClosePath()
	dc.Fill()

	// 添加字体(泉工)
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./tiff-demo/demo11(侧车牌)/FZDHTJW.TTF", 250); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("泉工", 34, 60, 0, 1)

	// 添加字体(企业简称)
	if err := dc.LoadFontFace("./tiff-demo/demo09(前车牌)/车牌字体DIN1451.ttf", 250); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("C02", 0, 310, -0.2, 1)

	// 添加字体(工号牌)
	//var numbers = []string{"D", "0", "0", "1"}
	dc.DrawStringAnchored("D", 0, 310, -1.2, 2.1)
	dc.DrawStringAnchored("0", 0, 310, -1.7, 3.2)
	dc.DrawStringAnchored("0", 0, 310, -1.7, 4.3)
	dc.DrawStringAnchored("2", 0, 310, -1.7, 5.4)

	// 添加车辆信息二维码
	im, err := gg.LoadImage("./tiff-demo/demo11(侧车牌)/img.png")
	if err != nil {
		log.Fatal(err)
	}
	dc.DrawImage(im, 159, 1400)

	img := dc.Image()
	f, _ := os.OpenFile("./tiff-demo/demo11(侧车牌)/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f, img, nil)

	//uploader := &SimpleFileServerUploader{
	//	Client: http.DefaultClient,
	//	Url:    "http://120.37.177.122:61438/file/upload",
	//	Hooks:  nil,
	//}
	//file, err := os.Open("./tiff-demo/demo11(侧车牌)/out.tiff")
	//
	//uploadResult, err := uploader.doUpload(context.Background(), file)
	//
	//fmt.Println(uploadResult.Url)
}
