package imageutil

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestNewSetFunc(t *testing.T) {
	bd := image.Rect(0, 0, 3, 3)
	for _, newImageDrawFunc := range []func(image.Rectangle) draw.Image{
		func(r image.Rectangle) draw.Image {
			return image.NewRGBA(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewRGBA64(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewNRGBA(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewNRGBA64(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewAlpha(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewAlpha16(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewGray(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewGray16(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewCMYK(r)
		},
		func(r image.Rectangle) draw.Image {
			return image.NewPaletted(r, testPalette)
		},
		func(r image.Rectangle) draw.Image {
			return &testImageDefault{image.NewRGBA(r)}
		},
	} {
		p := newImageDrawFunc(bd)
		t.Run(fmt.Sprintf("%T", p), func(t *testing.T) {
			set := NewSetFunc(p)
			for _, c := range testColors {
				c := color.RGBA64Model.Convert(c).(color.RGBA64)
				for y := bd.Min.Y; y < bd.Max.Y; y++ {
					for x := bd.Min.X; x < bd.Max.X; x++ {
						set(x, y, uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A))
						r1, g1, b1, a1 := p.At(x, y).RGBA()
						p.Set(x, y, c)
						r2, g2, b2, a2 := p.At(x, y).RGBA()
						if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
							t.Fatalf("different color: pixel %dx%d, color %#v: got {%d %d %d %d}, want {%d %d %d %d}", x, y, c, r1, g1, b1, a1, r2, g2, b2, a2)
						}
					}
				}
			}
		})
	}
}
