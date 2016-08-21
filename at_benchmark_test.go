package imageutil

import (
	"image"
	"image/color"
	"testing"
)

func BenchmarkNewAtFunc(b *testing.B) {
	for _, tc := range []struct {
		name     string
		newImage func(r image.Rectangle) image.Image
	}{
		{
			"RGBA",
			func(r image.Rectangle) image.Image {
				return image.NewRGBA(r)
			},
		},
		{
			"RGBA64",
			func(r image.Rectangle) image.Image {
				return image.NewRGBA64(r)
			},
		},
		{
			"NRGBAOpaque",
			func(r image.Rectangle) image.Image {
				p := image.NewNRGBA(r)
				p.SetNRGBA(r.Min.X, r.Min.Y, color.NRGBA{0xff, 0xff, 0xff, 0xff})
				return p
			},
		},
		{
			"NRGBATransparent",
			func(r image.Rectangle) image.Image {
				p := image.NewNRGBA(r)
				p.SetNRGBA(r.Min.X, r.Min.Y, color.NRGBA{0xff, 0xff, 0xff, 0x00})
				return p
			},
		},
		{
			"NRGBATranslucent",
			func(r image.Rectangle) image.Image {
				p := image.NewNRGBA(r)
				p.SetNRGBA(r.Min.X, r.Min.Y, color.NRGBA{0xff, 0xff, 0xff, 0x80})
				return p
			},
		},
		{
			"NRGBA64Opaque",
			func(r image.Rectangle) image.Image {
				p := image.NewNRGBA64(r)
				p.SetNRGBA64(r.Min.X, r.Min.Y, color.NRGBA64{0xffff, 0xffff, 0xffff, 0xffff})
				return p
			},
		},
		{
			"NRGBA64Transparent",
			func(r image.Rectangle) image.Image {
				p := image.NewNRGBA64(r)
				p.SetNRGBA64(r.Min.X, r.Min.Y, color.NRGBA64{0xffff, 0xffff, 0xffff, 0x0000})
				return p
			},
		},
		{
			"NRGBA64Translucent",
			func(r image.Rectangle) image.Image {
				p := image.NewNRGBA64(r)
				p.SetNRGBA64(r.Min.X, r.Min.Y, color.NRGBA64{0xffff, 0xffff, 0xffff, 0x8000})
				return p
			},
		},
		{
			"Alpha",
			func(r image.Rectangle) image.Image {
				return image.NewAlpha(r)
			},
		},
		{
			"Alpha16",
			func(r image.Rectangle) image.Image {
				return image.NewAlpha16(r)
			},
		},
		{
			"Gray",
			func(r image.Rectangle) image.Image {
				return image.NewGray(r)
			},
		},
		{
			"Gray16",
			func(r image.Rectangle) image.Image {
				return image.NewGray16(r)
			},
		},
		{
			"Paletted",
			func(r image.Rectangle) image.Image {
				return image.NewPaletted(r, testPalette)
			},
		},
		{
			"Uniform",
			func(r image.Rectangle) image.Image {
				return image.NewUniform(color.RGBA{})
			},
		},
		{
			"YCbCr444",
			func(r image.Rectangle) image.Image {
				return image.NewYCbCr(r, image.YCbCrSubsampleRatio444)
			},
		},
		{
			"YCbCr422",
			func(r image.Rectangle) image.Image {
				return image.NewYCbCr(r, image.YCbCrSubsampleRatio422)
			},
		},
		{
			"YCbCr420",
			func(r image.Rectangle) image.Image {
				return image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
			},
		},
		{
			"YCbCr440",
			func(r image.Rectangle) image.Image {
				return image.NewYCbCr(r, image.YCbCrSubsampleRatio440)
			},
		},
		{
			"YCbCr411",
			func(r image.Rectangle) image.Image {
				return image.NewYCbCr(r, image.YCbCrSubsampleRatio411)
			},
		},
		{
			"YCbCr410",
			func(r image.Rectangle) image.Image {
				return image.NewYCbCr(r, image.YCbCrSubsampleRatio410)
			},
		},
		{
			"NYCbCrA444",
			func(r image.Rectangle) image.Image {
				return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio444)
			},
		},
		{
			"NYCbCrA422",
			func(r image.Rectangle) image.Image {
				return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio422)
			},
		},
		{
			"NYCbCrA420",
			func(r image.Rectangle) image.Image {
				return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio420)
			},
		},
		{
			"NYCbCrA440",
			func(r image.Rectangle) image.Image {
				return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio440)
			},
		},
		{
			"NYCbCrA411",
			func(r image.Rectangle) image.Image {
				return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio411)
			},
		},
		{
			"NYCbCrA410",
			func(r image.Rectangle) image.Image {
				return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio410)
			},
		},
		{
			"NYCbCrAOpaque",
			func(r image.Rectangle) image.Image {
				p := image.NewNYCbCrA(r, image.YCbCrSubsampleRatio444)
				set := newSimpleSetFunc(p)
				set(r.Min.X, r.Min.Y, color.NRGBA64{0xffff, 0xffff, 0xffff, 0xffff})
				return p
			},
		},
		{
			"NYCbCrATransparent",
			func(r image.Rectangle) image.Image {
				p := image.NewNYCbCrA(r, image.YCbCrSubsampleRatio444)
				set := newSimpleSetFunc(p)
				set(r.Min.X, r.Min.Y, color.NRGBA64{0xffff, 0xffff, 0xffff, 0x0000})
				return p
			},
		},
		{
			"NYCbCrATranslucent",
			func(r image.Rectangle) image.Image {
				p := image.NewNYCbCrA(r, image.YCbCrSubsampleRatio444)
				set := newSimpleSetFunc(p)
				set(r.Min.X, r.Min.Y, color.NRGBA64{0xffff, 0xffff, 0xffff, 0x8000})
				return p
			},
		},
		{
			"CMYK",
			func(r image.Rectangle) image.Image {
				return image.NewCMYK(r)
			},
		},
		{
			"Default",
			func(r image.Rectangle) image.Image {
				return &testImageDefault{image.NewRGBA(r)}
			},
		},
	} {
		b.Run(tc.name, func(b *testing.B) {
			p := tc.newImage(image.Rect(0, 0, 1, 1))
			at := NewAtFunc(p)
			b.ResetTimer()
			var resR, resG, resB, resA uint32
			for i := 0; i < b.N; i++ {
				resR, resG, resB, resA = at(0, 0)
			}
			benchResR, benchResG, benchResB, benchResA = resR, resG, resB, resA
		})
	}
}
