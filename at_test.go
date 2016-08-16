package imageutil

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestNewAtFunc(t *testing.T) {
	bd := image.Rect(0, 0, 3, 3)
	for _, newImageFunc := range []func(image.Rectangle) image.Image{
		func(r image.Rectangle) image.Image {
			return image.NewRGBA(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewRGBA64(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNRGBA(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNRGBA64(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewAlpha(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewAlpha16(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewGray(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewGray16(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewCMYK(r)
		},
		func(r image.Rectangle) image.Image {
			return image.NewPaletted(r, testPalette)
		},
		func(r image.Rectangle) image.Image {
			return image.NewYCbCr(r, image.YCbCrSubsampleRatio444)
		},
		func(r image.Rectangle) image.Image {
			return image.NewYCbCr(r, image.YCbCrSubsampleRatio422)
		},
		func(r image.Rectangle) image.Image {
			return image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
		},
		func(r image.Rectangle) image.Image {
			return image.NewYCbCr(r, image.YCbCrSubsampleRatio440)
		},
		func(r image.Rectangle) image.Image {
			return image.NewYCbCr(r, image.YCbCrSubsampleRatio411)
		},
		func(r image.Rectangle) image.Image {
			return image.NewYCbCr(r, image.YCbCrSubsampleRatio410)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio444)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio422)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio420)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio440)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio411)
		},
		func(r image.Rectangle) image.Image {
			return image.NewNYCbCrA(r, image.YCbCrSubsampleRatio410)
		},
		func(r image.Rectangle) image.Image {
			return image.NewUniform(color.RGBA{})
		},
		func(r image.Rectangle) image.Image {
			return &testImageDefault{image.NewRGBA(r)}
		},
	} {
		p := newImageFunc(bd)
		t.Run(fmt.Sprintf("%T", p), func(t *testing.T) {
			set := newSimpleSetFunc(p)
			at := NewAtFunc(p)
			for _, c := range testColors {
				for y := bd.Min.Y; y < bd.Max.Y; y++ {
					for x := bd.Min.X; x < bd.Max.X; x++ {
						set(x, y, c)
						r1, g1, b1, a1 := at(x, y)
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

func newSimpleSetFunc(p image.Image) func(x, y int, c color.Color) {
	switch p := p.(type) {
	case draw.Image:
		return p.Set
	case *image.YCbCr:
		return func(x, y int, c color.Color) {
			c1 := color.YCbCrModel.Convert(c).(color.YCbCr)
			yi := p.YOffset(x, y)
			ci := p.COffset(x, y)
			p.Y[yi] = c1.Y
			p.Cb[ci] = c1.Cb
			p.Cr[ci] = c1.Cr
		}
	case *image.NYCbCrA:
		return func(x, y int, c color.Color) {
			c1 := color.NYCbCrAModel.Convert(c).(color.NYCbCrA)
			yi := p.YOffset(x, y)
			ci := p.COffset(x, y)
			ai := p.AOffset(x, y)
			p.Y[yi] = c1.Y
			p.Cb[ci] = c1.Cb
			p.Cr[ci] = c1.Cr
			p.A[ai] = c1.A
		}
	case *image.Uniform:
		return func(x, y int, c color.Color) {
			p.C = c
		}
	default:
		return nil
	}
}
