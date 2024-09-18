package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func createBaseLogoImage(W int, H int) *gg.Context {
	borderRadius := float64(W) / 15.0
	lineWidth := float64(W) / 6.5
	// The BaseLogoImage is a 1024x1024 image with a transparent background,
	// black font for the abbreviation, and a rounder border with a 10px width and colored
	// with the borderColor variable. It emits a glow for a 3D effect.
	dc := gg.NewContext(W, H)
	dc.SetRGBA(0, 0, 0, 1)
	var fontSize float64
	if len(abbreviation) == 1 {
		fontSize = pixelToPoints(float64(H) * 0.90)
	} else if len(abbreviation) == 2 {
		fontSize = pixelToPoints(float64(H) * 0.55)
	} else {
		fontSize = pixelToPoints(float64(H) * 0.40)
	}

	err := dc.LoadFontFace("./product-sans.ttf", fontSize)
	if err != nil {
		panic(fmt.Errorf("failed to load/download font 'product-sans.ttf'. Make sure you have internet connection. %v", err))
	}

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
	innerPaddingX := float64(W) * (20.0 / 1024.0)
	innerPaddingY := float64(H) * (20.0 / 1024.0)
	dc.SetColor(color.White)
	dc.DrawRoundedRectangle(
		lineWidth-innerPaddingX,
		lineWidth-innerPaddingY,
		float64(W)-(lineWidth*2)+innerPaddingX*2,
		float64(H)-(lineWidth*2)+innerPaddingY*2,
		borderRadius,
	)
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
	logo := createBaseLogoImage(int(W/3), int(W/3))
	x := float64(W/2) - float64(logo.Width()/2)
	y := float64(H/2) - float64(logo.Height()/2)
	splashDc.DrawImage(logo.Image(), int(x), int(y))
	return splashDc
}
