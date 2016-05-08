package imageutil

import (
	"image"
	"image/color"
	"math/rand"
)

type testImageDefault struct {
	*image.RGBA
}

var testColors []color.Color

func init() {
	vals := []uint8{0x00, 0x40, 0x80, 0xc0, 0xff}
	for _, r := range vals {
		for _, g := range vals {
			for _, b := range vals {
				for _, a := range vals {
					testColors = append(testColors, color.NRGBA{r, g, b, a})
				}
			}
		}
	}
	for i := 0; i < 100; i++ {
		testColors = append(testColors, testRandomColor())
	}
}

func testRandomColor() color.Color {
	return color.NRGBA{
		R: uint8(rand.Intn(1 << 8)),
		G: uint8(rand.Intn(1 << 8)),
		B: uint8(rand.Intn(1 << 8)),
		A: uint8(rand.Intn(1 << 8)),
	}
}

var testPalette = color.Palette{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 255, 255, 255},
}
