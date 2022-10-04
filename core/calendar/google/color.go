package google

import (
	"image/color"
)

// Sourced from Google Calendar CSS 2022-08-02
var knownColors = map[string]color.Color{
	"tomato":    color.RGBA{A: 255, R: 213, G: 0, B: 0},     // A deep red
	"flamingo":  color.RGBA{A: 255, R: 230, G: 124, B: 115}, // A pale red
	"tangerine": color.RGBA{A: 255, R: 244, G: 81, B: 30},   // A vibrant orange
	"banana":    color.RGBA{A: 255, R: 246, G: 191, B: 38},  // A simple yellow
	"sage":      color.RGBA{A: 255, R: 51, G: 182, B: 121},  // A pale green
	"basil":     color.RGBA{A: 255, R: 11, G: 128, B: 67},   // A deep green
	"peacock":   color.RGBA{A: 255, R: 3, G: 155, B: 229},   // A light blue
	"blueberry": color.RGBA{A: 255, R: 63, G: 81, B: 181},   // A dark blue
	"lavender":  color.RGBA{A: 255, R: 121, G: 134, B: 203}, // A pale purple
	"grape":     color.RGBA{A: 255, R: 142, G: 36, B: 170},  // A deep purple
	"graphite":  color.RGBA{A: 255, R: 97, G: 97, B: 97},    // A smooth gray
}
