package main

import (
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func isValidColor(color string) bool {
	// #RRGGBB
	if len(color) != 7 {
		return false
	}
	if color[0] != '#' {
		return false
	}
	for i := 1; i < 7; i++ {
		if !((color[i] >= '0' && color[i] <= '9') || (color[i] >= 'A' && color[i] <= 'F') || (color[i] >= 'a' && color[i] <= 'f')) {
			return false
		}
	}
	return true
}

func hexToDec(hex string) int {
	dec := 0
	for i := 0; i < len(hex); i++ {
		dec *= 16
		if hex[i] >= '0' && hex[i] <= '9' {
			dec += int(hex[i] - '0')
		} else if hex[i] >= 'A' && hex[i] <= 'F' {
			dec += int(hex[i] - 'A' + 10)
		} else {
			dec += int(hex[i] - 'a' + 10)
		}
	}
	return dec
}

func hexToColor(hex string) color.Color {
	// #RRGGBB
	r := uint8(hexToDec(hex[1:3]))
	g := uint8(hexToDec(hex[3:5]))
	b := uint8(hexToDec(hex[5:7]))
	return color.RGBA{r, g, b, 255}
}

func getAbbreviation(name string) string {
	words := strings.Split(name, " ")
	if len(words) == 0 {
		panic("No words found in name. Cannot create an abbreviation.")
	}
	abbr := ""
	if len(words) == 1 {
		s := string(words[0][0])
		abbr = strings.ToUpper(s)
	} else if len(words) == 3 {
		first := strings.ToUpper(string(words[0][0]))
		second := strings.ToUpper(string(words[1][0]))
		third := strings.ToUpper(string(words[2][0]))
		abbr = first + second + third
	} else {
		// if len(words) == 2
		first := strings.ToUpper(string(words[0][0]))
		second := strings.ToLower(string(words[1][0]))
		abbr = first + second
	}
	return abbr
}

func makeBaseFilename(name string) string {
	// only a-zA-Z0-9 and white space to `_`
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	sanitized := regex.ReplaceAllString(name, "_")
	// while __ in the string, replace with _
	for strings.Contains(sanitized, "__") {
		sanitized = strings.Replace(sanitized, "__", "_", -1)
	}
	// if the string starts with _, remove it
	for strings.HasPrefix(sanitized, "_") {
		sanitized = sanitized[1:]
	}
	if len(sanitized) == 0 {
		panic("The base filename is empty. Cannot proceed. Please provide a valid name.")
	}
	return sanitized
}

func ensureDirExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func sanitizeOutputDirPath(path string) string {
	// remove trailing slashes
	for strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	return path
}

func ensureFontExists(fontPath string) {
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		// Font file doesn't exist, attempt to download it
		resp, err := http.Get("https://raw.githubusercontent.com/khanhas/mnmlUI/master/%40Resources/Fonts/Product%20Sans%20Regular.ttf")
		if err != nil {
			log.Fatalf("Failed to fetch font: %v", err)
		}
		defer resp.Body.Close()

		// Check if the response status is OK
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Failed to download font, status code: %d", resp.StatusCode)
		}

		file, err := os.Create(fontPath)
		if err != nil {
			log.Fatalf("Failed to create font file: %v", err)
		}
		defer file.Close()

		// Copy the response body to the file
		if _, err := io.Copy(file, resp.Body); err != nil {
			log.Fatalf("Failed to write font to file: %v", err)
		}
	}
}

func removeFont(fontPath string) {
	err := os.Remove(fontPath)
	if err != nil {
		panic(err)
	}
}

func pixelToPoints(pixel float64) float64 {
	return pixel * (12.0 / 16.0)
}
