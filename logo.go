package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func createBaseLogoImage() *image.RGBA {
	// The BaseLogoImage is a 1024x1024 image with a transparent background,
	// black font for the abbreviation, and a rounder border with a 10px width and colored
	// with the borderColor variable. It emits a glow for a 3D effect.

	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	// Make entire background transparent
	clearColor := color.RGBA{0, 0, 0, 0}
	draw.Draw(img, img.Bounds(), &image.Uniform{clearColor}, image.Point{}, draw.Src)

	// Fill inner area with white color
	// white := color.RGBA{255, 255, 255, 255}
	borderWidth := 10
	borderRadius := 100 // Set border-radius to 100 for a rounder effect
	bounds := img.Bounds()
	innerRect := image.Rect(borderWidth, borderWidth, bounds.Dx()-borderWidth, bounds.Dy()-borderWidth)
	draw.Draw(img, innerRect, &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Draw rounded rectangle for the border
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Check if pixel is part of the border
			isBorder := (x < innerRect.Min.X || x >= innerRect.Max.X || y < innerRect.Min.Y || y >= innerRect.Max.Y)

			// Determine corner radius and calculate distance for the rounded corners
			isCorner := (x < innerRect.Min.X+borderRadius && y < innerRect.Min.Y+borderRadius) ||
				(x >= innerRect.Max.X-borderRadius && y < innerRect.Min.Y+borderRadius) ||
				(x < innerRect.Min.X+borderRadius && y >= innerRect.Max.Y-borderRadius) ||
				(x >= innerRect.Max.X-borderRadius && y >= innerRect.Max.Y-borderRadius)

			if isCorner {
				// left top:
				ltX := innerRect.Min.X + borderRadius
				ltY := innerRect.Min.Y + borderRadius
				lefttop := math.Sqrt(float64((x-ltX)*(x-ltX) + (y-ltY)*(y-ltY)))
				// right top:
				rtX := innerRect.Max.X - borderRadius
				rtY := innerRect.Min.Y + borderRadius
				righttop := math.Sqrt(float64((x-rtX)*(x-rtX) + (y-rtY)*(y-rtY)))
				// left bottom:
				lbX := innerRect.Min.X + borderRadius
				lbY := innerRect.Max.Y - borderRadius
				leftbottom := math.Sqrt(float64((x-lbX)*(x-lbX) + (y-lbY)*(y-lbY)))
				// right bottom:
				rbX := innerRect.Max.X - borderRadius
				rbY := innerRect.Max.Y - borderRadius
				rightbottom := math.Sqrt(float64((x-rbX)*(x-rbX) + (y-rbY)*(y-rbY)))

				if lefttop <= float64(borderRadius) && lefttop >= float64(borderRadius-borderWidth) {
					img.Set(x, y, borderColor)
				} else if righttop <= float64(borderRadius) && righttop >= float64(borderRadius-borderWidth) {
					img.Set(x, y, borderColor)
				} else if leftbottom <= float64(borderRadius) && leftbottom >= float64(borderRadius-borderWidth) {
					img.Set(x, y, borderColor)
				} else if rightbottom <= float64(borderRadius) && rightbottom >= float64(borderRadius-borderWidth) {
					img.Set(x, y, borderColor)
				} else {
					img.Set(x, y, clearColor)
				}
			} else if isBorder {
				// if it's the top border then y+ border width
				if y < innerRect.Min.Y {
					img.Set(x, y+borderWidth, borderColor)
				} else if y >= innerRect.Max.Y {
					img.Set(x, y-borderWidth, borderColor)
				} else if x < innerRect.Min.X {
					img.Set(x+borderWidth, y, borderColor)
				} else if x >= innerRect.Max.X {
					img.Set(x-borderWidth, y, borderColor)
				}
			}
		}
	}

	// Draw the abbreviation in the center with a large font
	col := color.Black
	//openfont

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  fixed.P(bounds.Dx()/2, bounds.Dy()/2),
	}
	d.DrawString(abbreviation)

	return img
}

func saveImageToFile(img *image.RGBA, filename string) {
	f, err := os.Create(outputDirPath + "/" + filename + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
