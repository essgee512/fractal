package f

import (
	"image"
	"image/color"
	"runtime"
)

const (
	G = float64(2)
	N = 100
)

type Fractal struct {
	Size    int
	Center  Point
	Scale   float64

	Fmap    Fmap
	Cmap    Cmap

	w       int
	h       int
	Img     *image.NRGBA
	pixels  []pixel
}

func NewFractal(cfg Fractal) Fractal {
	w, h := 3*cfg.Size, 2*cfg.Size
	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	return Fractal{
		Size:     cfg.Size,
		Center:   cfg.Center,
		Scale:    cfg.Scale,
		Fmap:     cfg.Fmap,
		Cmap:     cfg.Cmap,

		w:        w,
		h:        h,
		Img:      img,
		pixels:   initPixels(cfg),
	}
}

func (f Fractal) Render() {
	f.RenderParallel()
}

func (f Fractal) RenderSerial() {
	for _, p := range f.pixels {
		f.render(p)
	}
}

func (f Fractal) RenderParallel() {
	nCPU := runtime.NumCPU()
	c := make(chan int, nCPU)

	for i := 0; i < nCPU; i++ {
		off := i*len(f.pixels) / nCPU
		len := (i + 1)*len(f.pixels) / nCPU

		go f.renderSlice(off, len, c)
	}

	for i := 0; i < nCPU; i++ {
		<-c
	}
}

func (f Fractal) renderSlice(i, len int, c chan int) {
	for ; i < len; i++ {
		f.render(f.pixels[i])
	}

	c <- 1
}

func (f Fractal) render(p pixel) {
	r, g, b, a := f.Cmap(f.Fmap(complex(p.x, p.y)))
	f.Img.SetNRGBA(p.k, p.l, color.NRGBA{r, g, b, a})
}

type Point struct {
	X float64
	Y float64
}
