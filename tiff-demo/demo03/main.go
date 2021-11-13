package main

import "github.com/fogleman/gg"

func main() {
	//const S = 4096 * 2
	//const T = 16 * 2
	//const F = 28
	//dc := gg.NewContext(S, S)
	//dc.SetRGB(1, 1, 1)
	//dc.Clear()
	//dc.SetRGB(0, 0, 0)
	//if err := dc.LoadFontFace("./tiff-demo/demo03/XoloniumRegular.woff.ttf", F); err != nil {
	//	panic(err)
	//}
	//for r := 0; r < 256; r++ {
	//	for c := 0; c < 256; c++ {
	//		i := r*256 + c
	//		x := float64(c*T) + T/2
	//		y := float64(r*T) + T/2
	//		dc.DrawStringAnchored(string(rune(i)), x, y, 0.5, 0.5)
	//	}
	//}
	//const S = 1024
	//dc := gg.NewContext(S, S)
	//dc.SetRGB(1, 1, 1)
	//dc.Clear()
	//dc.SetRGB(0, 0, 0)
	//if err := dc.LoadFontFace("./tiff-demo/demo03/JingDianCuHeiJian-1.ttf", 96); err != nil {
	//	panic(err)
	//}
	//dc.DrawStringAnchored("Hello发多少, world!", S/2, S/2, 0.5, 0.5)
	//dc.SavePNG("./tiff-demo/demo03/out.png")

	//const NX = 4
	//const NY = 3
	//im, err := gg.LoadPNG("./tiff-demo/demo09/gopher.png")
	//if err != nil {
	//	panic(err)
	//}
	//w := im.Bounds().Size().X
	//h := im.Bounds().Size().Y
	//dc := gg.NewContext(w*NX, h*NY)
	//for y := 0; y < NY; y++ {
	//	for x := 0; x < NX; x++ {
	//		dc.DrawImage(im, x*w, y*h)
	//	}
	//}

	const W = 400
	const H = 500
	im, err := gg.LoadPNG("./tiff-demo/demo09/gopher.png")
	if err != nil {
		panic(err)
	}
	iw, ih := im.Bounds().Dx(), im.Bounds().Dy()
	dc := gg.NewContext(W, H)
	// draw outline
	dc.SetHexColor("#ff0000")
	dc.SetLineWidth(1)
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.Stroke()
	// draw full image
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.DrawRectangle(100, 210, float64(iw), float64(ih))
	dc.Stroke()
	dc.DrawImage(im, 100, 210)
	// draw image with current matrix applied
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 0, float64(iw), float64(ih)/2+20.0)
	dc.StrokePreserve()
	dc.Clip()
	dc.DrawImageAnchored(im, 100, 0, 0.2, 0.0)

	dc.SavePNG("./tiff-demo/demo03/out.png")
}
