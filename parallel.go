package imageutil

import (
	"context"
	"image"
	"runtime"
	"sync"
)

// Parallel1D dispatches tasks concurrently for a Rectangle.
//
// It splits the image horizontally in GOMAXPROCS parts and runs GOMAXPROCS workers.
//
// It should be used if all the pixels of the image have the same process cost.
func Parallel1D(ctx context.Context, r image.Rectangle, f func(context.Context, image.Rectangle)) {
	p := runtime.GOMAXPROCS(0)
	wg := new(sync.WaitGroup)
	for y := 0; y < p; y++ {
		rr := image.Rect(
			r.Min.X,
			r.Min.Y+(r.Dy()*y/p),
			r.Max.X,
			r.Min.Y+(r.Dy()*(y+1)/p),
		)
		if !rr.Empty() {
			wg.Add(1)
			go func(rr image.Rectangle) {
				f(ctx, rr)
				wg.Done()
			}(rr)
		}
	}
	wg.Wait()
}

// Parallel2D dispatches tasks concurrently for a Rectangle.
//
// It splits the image in a GOMAXPROCS x GOMAXPROCS grid
// and runs GOMAXPROCS workers.
//
// It should be used if all the pixels of the image don't have the same process cost.
func Parallel2D(ctx context.Context, r image.Rectangle, f func(context.Context, image.Rectangle)) {
	p := runtime.GOMAXPROCS(0)
	rc := make(chan image.Rectangle)
	wg := new(sync.WaitGroup)
	wg.Add(p)
	for i := 0; i < p; i++ {
		go func() {
			for rr := range rc {
				f(ctx, rr)
			}
			wg.Done()
		}()
	}
	for y := 0; y < p; y++ {
		for x := 0; x < p; x++ {
			rr := image.Rect(
				r.Min.X+(r.Dx()*x/p),
				r.Min.Y+(r.Dy()*y/p),
				r.Min.X+(r.Dx()*(x+1)/p),
				r.Min.Y+(r.Dy()*(y+1)/p),
			)
			if rr.Empty() {
				continue
			}
			select {
			case rc <- rr:
			case <-ctx.Done():
				goto END
			}
		}
	}
END:
	close(rc)
	wg.Wait()
}
