package imageutil

import (
	"image"
	"testing"
)

func TestParallel1D(t *testing.T) {
	r := image.Rect(100, 100, 200, 200)
	Parallel1D(r, func(sub image.Rectangle) {
		if !sub.In(r) {
			t.Fatalf("%s is not in %s", sub, r)
		}
	})
}

func TestParallel2D(t *testing.T) {
	r := image.Rect(100, 100, 200, 200)
	Parallel2D(r, func(sub image.Rectangle) {
		if !sub.In(r) {
			t.Fatalf("%s is not in %s", sub, r)
		}
	})
}
