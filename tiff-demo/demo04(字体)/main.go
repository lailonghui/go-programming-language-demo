package main

import (
	"github.com/fogleman/gg"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	//if err := dc.LoadFontFace("./tiff-demo/demo04(字体)/JingDianCuHeiJian-1.ttf", 96); err != nil {
	if err := dc.LoadFontFace("./tiff-demo/demo04(字体)/JingDianZongYiTiJian-1.ttf", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("泉工-C02D001", 512, 512, 0.5, 0.5)
	dc.SavePNG("./tiff-demo/demo04(字体)/out2.png")
}
