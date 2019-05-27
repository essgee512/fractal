package f

import (
	"image/color"
)

type Fractal struct {
	Canvas Canvas
	Fmap   Fmap
	G      float64
	N      int
	Cmap   Cmap
}

func NewFractal(cfg Fractal) Fractal {
	return Fractal{
		Canvas: cfg.Canvas.Prepare(),
		Fmap:   cfg.Fmap,
		G:      cfg.G,
		N:      cfg.N,
		Cmap:   cfg.Cmap,
	}
}

func (f Fractal) Render() {
	var numCPU = 4
	// var numCPU = runtime.NumCPU()
	c := make(chan int, numCPU)  // Buffering optional but sensible.
	for i := 0; i < numCPU; i++ {
		offset := i*len(f.Canvas.Pixels)/numCPU
		n := (i+1)*len(f.Canvas.Pixels)/numCPU
		go f.RenderGroup(offset, n, c)
	}

	// Drain the channel.
	for i := 0; i < numCPU; i++ {
		<-c    // wait for one task to complete
	}
	// All done.
}

func (f Fractal) RenderGroup(i, n int, c chan int) {
	for ; i < n; i++ {
		// f.Pixel[i].Color(f)

		k := f.Canvas.Pixels[i].K
		l := f.Canvas.Pixels[i].L
		x := f.Canvas.Pixels[i].X
		y := f.Canvas.Pixels[i].Y

		// color
		z := complex(x, y)
		r, g, b, a := f.Cmap(z, f.Fmap, f.G, f.N)

		// paint
		f.Canvas.Img.Set(k, l, color.NRGBA{r, g, b, a})
	}

	c <- 1    // signal that this piece is done
}
