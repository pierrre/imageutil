package imageutil

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func BenchmarkNewSetFunc(b *testing.B) {
	for _, tc := range []struct {
		name     string
		newImage func(r image.Rectangle) draw.Image
		color    color.Color
	}{
		{
			name: "RGBA",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewRGBA(r)
			},
		},
		{
			name: "RGBA64",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewRGBA64(r)
			},
		},
		{
			name: "NRGBAOpaque",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewNRGBA(r)
			},
			color: color.NRGBA{0xff, 0xff, 0xff, 0xff},
		},
		{
			name: "NRGBATransparent",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewNRGBA(r)
			},
			color: color.NRGBA{0xff, 0xff, 0xff, 0x00},
		},
		{
			name: "NRGBATranslucent",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewNRGBA(r)
			},
			color: color.NRGBA{0xff, 0xff, 0xff, 0x80},
		},

		{
			name: "NRGBA64Opaque",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewNRGBA64(r)
			},
			color: color.NRGBA64{0xffff, 0xffff, 0xffff, 0xffff},
		},
		{
			name: "NRGBA64Transparent",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewNRGBA64(r)
			},
			color: color.NRGBA64{0xffff, 0xffff, 0xffff, 0x0000},
		},
		{
			name: "NRGBA64Translucent",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewNRGBA64(r)
			},
			color: color.NRGBA64{0xffff, 0xffff, 0xffff, 0x8000},
		},
		{
			name: "Alpha",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewAlpha(r)
			},
		},
		{
			name: "Alpha16",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewAlpha16(r)
			},
		},
		{
			name: "Gray",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewGray(r)
			},
		},
		{
			name: "Gray16",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewGray16(r)
			},
		},
		{
			name: "Paletted",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewPaletted(r, testPalette)
			},
		},
		{
			name: "CMYK",
			newImage: func(r image.Rectangle) draw.Image {
				return image.NewCMYK(r)
			},
		},
		{
			name: "Default",
			newImage: func(r image.Rectangle) draw.Image {
				return &testImageDefault{image.NewRGBA(image.Rect(0, 0, 1, 1))}
			},
		},
	} {
		b.Run(tc.name, func(b *testing.B) {
			p := tc.newImage(image.Rect(0, 0, 1, 1))
			set := NewSetFunc(p)
			var rr, gg, bb, aa uint32 = 0xffff, 0xffff, 0xffff, 0xffff
			if tc.color != nil {
				rr, bb, gg, aa = tc.color.RGBA()
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set(0, 0, rr, gg, bb, aa)
			}
		})
	}
}
