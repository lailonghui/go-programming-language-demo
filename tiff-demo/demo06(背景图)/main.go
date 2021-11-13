package main

import (
	"log"

	"github.com/fogleman/gg"
)

func main() {
	im, err := gg.LoadImage("./tiff-demo/demo04(字体)/baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(1024, 1024)
	dc.DrawRoundedRectangle(0, 0, 1024, 1024, 64)
	dc.Clip()
	dc.DrawImage(im, 200, 0)
	dc.SavePNG("./tiff-demo/demo06(背景图)/out.png")
}
