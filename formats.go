package main

import (
	"image"
	"os"
	"strings"

	ico "github.com/Kodeworks/golang-image-ico"
	"github.com/jackmordaunt/icns"
)

func pngToIcns(pngPath string) string {
	icnsPath := strings.Replace(pngPath, ".png", ".icns", 1)
	pngf, err := os.Open(pngPath)
	if err != nil {
		panic(err)
	}
	defer pngf.Close()
	srcImg, _, err := image.Decode(pngf)
	if err != nil {
		panic(err)
	}
	dest, err := os.Create(icnsPath)
	if err != nil {
		panic(err)
	}
	defer dest.Close()
	if err := icns.Encode(dest, srcImg); err != nil {
		panic(err)
	}
	return icnsPath
}

func pngToIco(pngPath string) string {
	icoPath := strings.Replace(pngPath, ".png", ".ico", 1)
	pngf, err := os.Open(pngPath)
	if err != nil {
		panic(err)
	}
	defer pngf.Close()
	srcImg, _, err := image.Decode(pngf)
	if err != nil {
		panic(err)
	}
	dest, err := os.Create(icoPath)
	if err != nil {
		panic(err)
	}
	defer dest.Close()
	if err := ico.Encode(dest, srcImg); err != nil {
		panic(err)
	}
	return icoPath
}
