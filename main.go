package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"strings"
)

var borderColor color.Color
var filename string
var appName string
var abbreviation string
var outputDirPath string

func main() {
	defaultColor := fmt.Sprintf("#%06x", rand.Intn(0xFFFFFF+1))

	name := flag.String("name", "", "The name (required)")
	path := flag.String("path", ".", "The path to the output directory")
	color := flag.String("color", defaultColor, "The color to use")
	isDebug := flag.Bool("debug", false, "Enable debug mode")

	flag.Parse()

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

	abbreviation = getAbbreviation(*name)
	borderColor = hexToColor(*color)
	filename = makeBaseFilename(*name)
	appName = strings.Trim(*name, " ")
	outputDirPath = sanitizeOutputDirPath(*path)
	if *isDebug {
		fmt.Println("Abbreviation:", abbreviation)
		fmt.Println("Border Color:", borderColor)
		fmt.Println("Filename:", filename)
		fmt.Println("App Name:", appName)
		fmt.Println("Output Directory Path:", outputDirPath)
	}

	ensureDirExists(outputDirPath)
	logo := createBaseLogoImage()
	saveImageToFile(logo, filename)
}
