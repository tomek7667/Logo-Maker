package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

func createBaseLogoImage(W int, H int, borderRadius float64, lineWidth float64) *gg.Context {
	// The BaseLogoImage is a 1024x1024 image with a transparent background,
	// black font for the abbreviation, and a rounder border with a 10px width and colored
	// with the borderColor variable. It emits a glow for a 3D effect.
	dc := gg.NewContext(W, H)
	dc.SetRGBA(0, 0, 0, 1)
	dc.LoadFontFace("./product-sans.ttf", float64(H)/1.4)
	middleW := float64(W / 2)
	middleH := float64(H / 2)
	dc.DrawStringAnchored(abbreviation, middleW, middleH, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// Border
	dc.SetColor(borderColor)
	dc.SetLineWidth(lineWidth)
	dc.DrawRoundedRectangle(lineWidth/2, lineWidth/2, float64(W)-(lineWidth), float64(H)-(lineWidth), borderRadius)
	dc.Stroke()

	// Background color
	dc.SetColor(color.White)
	dc.DrawRoundedRectangle(lineWidth-15, lineWidth-15, float64(W)-(lineWidth*2)+30, float64(H)-(lineWidth*2)+30, borderRadius)
	dc.Fill()

	// TODO: Add Glow effect

	// Color gradient
	g := gg.NewLinearGradient(0, 0, float64(W), float64(H))
	g.AddColorStop(0, fontColorGradient1)
	g.AddColorStop(1, fontColorGradient2)
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.Fill()

	return dc
}

func drawLogoAsSplash(W int, H int) *gg.Context {
	splashDc := gg.NewContext(W, H)
	splashDc.SetColor(color.Transparent)
	splashDc.DrawRectangle(0, 0, float64(W), float64(H))
	splashDc.Fill()

	// scale the logo much down
	println("creating logo")
	logo := createBaseLogoImage(int(W/3), int(W/3), 32, 60)
	println("creating logo done")

	x := float64(W/2) - float64(logo.Width()/2)
	y := float64(H/2) - float64(logo.Height()/2)
	splashDc.DrawImage(logo.Image(), int(x), int(y))
	return splashDc
}
