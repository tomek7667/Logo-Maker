package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"strings"
	"sync"
)

var borderColor color.Color
var fontColorGradient1 color.Color
var fontColorGradient2 color.Color
var filename string
var appName string
var abbreviation string
var outputDirPath string

type Resolution struct {
	Width  int
	Height int
	Name   string
}

var folders = [][]string{
	{"iPhone", "iPad", "Universal"},
	{"Android"},
}

var resolutions = []Resolution{
	// Universal
	{1024, 1024, "Universal App Store"},
	// iPhone
	{180, 180, "iPhone App Icon (60pt @3x)"},
	{120, 120, "iPhone App Icon (60pt @2x)"},
	{87, 87, "iPhone App Icon (29pt @3x)"},
	{58, 58, "iPhone App Icon (29pt @2x)"},
	{120, 120, "iPhone App Icon (40pt @3x)"},
	{80, 80, "iPhone App Icon (40pt @2x)"},
	{60, 60, "iPhone App Icon (20pt @3x)"},
	{40, 40, "iPhone App Icon (20pt @2x)"},
	// iPad
	{167, 167, "iPad App Icon (83.5pt @2x)"},
	{152, 152, "iPad App Icon (76pt @2x)"},
	{76, 76, "iPad App Icon (76pt @1x)"},
	{40, 40, "iPad App Icon (20pt @2x)"},
	{60, 60, "iPad App Icon (20pt @3x)"},
	// Android
	{512, 512, "Android App Icon (xxhdpi)"},
	{144, 144, "Android App Icon (hdpi)"},
	{96, 96, "Android App Icon (mdpi)"},
	{72, 72, "Android App Icon (ldpi)"},
	{48, 48, "Android App Icon (drawable-mdpi)"},
}

var fontPath = "./product-sans.ttf"

func main() {
	defaultColor := fmt.Sprintf("#%06x", rand.Intn(0xFFFFFF+1))
	name := flag.String("name", "", "The name (required)")
	path := flag.String("path", ".", "The path to the output directory")
	color := flag.String("color", defaultColor, "The color to use")
	isDebug := flag.Bool("debug", false, "Enable debug mode")
	fontColor1 := flag.String("fontColor1", "#000000", "The first color for the gradient")
	fontColor2 := flag.String("fontColor2", "#000000", "The first color for the gradient")
	include := flag.String("include", "Android,iPhone", "The comma-separated list of resolutions to include")
	help := flag.Bool("help", false, "Show help message")

	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	if *name == "" {
		fmt.Println("The 'name' argument is required.")
		flag.Usage()
		return
	}
	if !isValidColor(*color) {
		fmt.Println("The 'color' argument is not a valid color. It should be in the format #RRGGBB.")
		flag.Usage()
		return
	}
	if !isValidColor(*fontColor1) {
		fmt.Println("The 'fontColor1' argument is not a valid color. It should be in the format #RRGGBB.")
		flag.Usage()
		return
	}
	if !isValidColor(*fontColor2) {
		fmt.Println("The 'fontColor2' argument is not a valid color. It should be in the format #RRGGBB.")
		flag.Usage()
		return
	}

	abbreviation = getAbbreviation(*name)
	borderColor = hexToColor(*color)
	filename = makeBaseFilename(*name)
	appName = strings.Trim(*name, " ")
	outputDirPath = sanitizeOutputDirPath(*path)
	fontColorGradient1 = hexToColor(*fontColor1)
	fontColorGradient2 = hexToColor(*fontColor2)
	if *isDebug {
		fmt.Println("Abbreviation:", abbreviation)
		fmt.Println("Border Color:", borderColor)
		fmt.Println("Filename:", filename)
		fmt.Println("App Name:", appName)
		fmt.Println("Output Directory Path:", outputDirPath)
		fmt.Println("Font Color Gradient 1:", fontColorGradient1)
		fmt.Println("Font Color Gradient 2:", fontColorGradient2)
		fmt.Println("--------------------------------")
	}

	ensureDirExists(outputDirPath)
	ensureFontExists(fontPath)
	baseImagePath := outputDirPath + "/" + filename + ".png"
	logo := createBaseLogoImage(1024, 1024, 50, 128)
	logo.SavePNG(baseImagePath)
	fmt.Println("Success:", *color)
	fmt.Println("Base Image:", baseImagePath)
	fmt.Println("--------------------------------")
	for _, folder := range folders {
		if !strings.Contains(*include, folder[0]) {
			continue
		}
		baseDir := outputDirPath + "/" + folder[0]
		ensureDirExists(baseDir)
		wg := sync.WaitGroup{}
		for _, res := range resolutions {
			wg.Add(1)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered from panic:", r)
					}
					wg.Done()
				}()
				if !strings.Contains(res.Name, folder[0]) {
					return
				}
				imagePath := baseDir + "/" + filename + "__" + makeBaseFilename(res.Name) + ".png"
				logo := createBaseLogoImage(res.Width, res.Height, float64(res.Width)/8, float64(res.Width)/20.48)
				logo.SavePNG(imagePath)
				fmt.Println(imagePath)
				if strings.Contains(folder[0], "iPhone") {
					icnsPath := pngToIcns(imagePath)
					fmt.Println(icnsPath)
				}
				if !strings.Contains(folder[0], "iPhone") {
					icoPath := pngToIco(imagePath)
					fmt.Println(icoPath)
				}
			}()
		}
		wg.Wait()
	}
	splash := drawLogoAsSplash(1179, 2556)
	splashImagePath := outputDirPath + "/" + filename + "__splash.png"
	splash.SavePNG(splashImagePath)
	removeFont(fontPath)
}
