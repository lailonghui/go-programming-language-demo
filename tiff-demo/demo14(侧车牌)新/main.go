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

	// 添加字体(企业简称)
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./tiff-demo/demo14(侧车牌)新/FZDHTJW.TTF", 250); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("祥茂", 34, 60, 0, 1)

	// 添加字体(自编号)
	if err := dc.LoadFontFace("./tiff-demo/demo14(侧车牌)新/车牌字体DIN1451.ttf", 250); err != nil {
		panic(err)
	}
	var numbers = []string{"0", "0", "8"}
	thick := 2 // 字体厚度

	for i, s := range numbers {
		for dy := -thick; dy <= thick; dy++ {
			for dx := -thick; dx <= thick; dx++ {
				if dx*dx+dy*dy >= thick*thick {
					// give it rounded corners
					continue
				}
				x := 0 + float64(dx)
				y := 310 + float64(dy)
				dc.DrawStringAnchored(s, x, y, -1.7, 2+1.3*float64(i))
				//dc.DrawStringAnchored(s, 0, 310, -1.7, 2+1.3*float64(i))
			}
		}
	}
	//dc.DrawStringAnchored("0", 0, 310, -1.7, 4.3)
	//dc.DrawStringAnchored("2", 0, 310, -1.7, 5.4)

	// 添加车辆信息二维码
	im, err := gg.LoadImage("./tiff-demo/demo14(侧车牌)新/img.png")
	if err != nil {
		log.Fatal(err)
	}
	dc.DrawImage(im, 159, 1400)

	img := dc.Image()
	f, _ := os.OpenFile("./tiff-demo/demo14(侧车牌)新/out.tiff", os.O_CREATE|os.O_RDWR, 0666)
	tiff.Encode(f, img, nil)

	//uploader := &SimpleFileServerUploader{
	//	Client: http.DefaultClient,
	//	Url:    "http://120.37.177.122:61438/file/upload",
	//	Hooks:  nil,
	//}
	//file, err := os.Open("./tiff-demo/demo14(侧车牌)新/out.tiff")
	//
	//uploadResult, err := uploader.doUpload(context.Background(), file)
	//
	//fmt.Println(uploadResult.Url)
}
