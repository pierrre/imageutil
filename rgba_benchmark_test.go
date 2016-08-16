package imageutil

import "testing"

var benchResR, benchResG, benchResB, benchResA uint32

func BenchmarkRGBAToNRGBA(b *testing.B) {
	for _, tc := range []struct {
		name       string
		r, g, b, a uint32
	}{
		{"Opaque", 0xffff, 0x8000, 0x0000, 0xffff},
		{"Transparent", 0x0000, 0x0000, 0x0000, 0x0000},
		{"Translucent", 0x8000, 0x4000, 0x0000, 0x8000},
	} {
		b.Run(tc.name, func(b *testing.B) {
			var resR, resG, resB, resA uint32
			for i := 0; i < b.N; i++ {
				resR, resG, resB, resA = RGBAToNRGBA(tc.r, tc.g, tc.b, tc.a)
			}
			benchResR, benchResG, benchResB, benchResA = resR, resG, resB, resA
		})
	}
}

func BenchmarkNRGBAToRGBA(b *testing.B) {
	for _, tc := range []struct {
		name       string
		r, g, b, a uint32
	}{
		{"Opaque", 0xffff, 0x8000, 0x0000, 0xffff},
		{"Transparent", 0xffff, 0x8000, 0x0000, 0x0000},
		{"Translucent", 0xffff, 0x8000, 0x0000, 0x8000},
	} {
		b.Run(tc.name, func(b *testing.B) {
			var resR, resG, resB, resA uint32
			for i := 0; i < b.N; i++ {
				resR, resG, resB, resA = NRGBAToRGBA(tc.r, tc.g, tc.b, tc.a)
			}
			benchResR, benchResG, benchResB, benchResA = resR, resG, resB, resA
		})
	}
}
