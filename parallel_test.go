package imageutil

import (
	"context"
	"image"
	"testing"
)

func TestParallel1D(t *testing.T) {
	r := image.Rect(100, 100, 200, 200)
	Parallel1D(context.Background(), r, func(ctx context.Context, sub image.Rectangle) {
		if !sub.In(r) {
			t.Fatalf("%s is not in %s", sub, r)
		}
	})
}

func TestParallel2D(t *testing.T) {
	r := image.Rect(100, 100, 200, 200)
	Parallel2D(context.Background(), r, func(ctx context.Context, sub image.Rectangle) {
		if !sub.In(r) {
			t.Fatalf("%s is not in %s", sub, r)
		}
	})
}

func TestParallel2DEmpty(t *testing.T) {
	called := false
	Parallel2D(context.Background(), image.ZR, func(ctx context.Context, sub image.Rectangle) {
		called = true
	})
	if called {
		t.Fatal("function called")
	}
}

func TestParallel2DCtxDone(t *testing.T) {
	r := image.Rect(100, 100, 200, 200)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	called := false
	Parallel2D(ctx, r, func(ctx context.Context, sub image.Rectangle) {
		called = true
	})
	if called {
		t.Fatal("function called")
	}
}
