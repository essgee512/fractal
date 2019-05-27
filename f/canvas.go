package f

import (
	"image"
)


type CnvParams struct {
	Size    int
	Center  Point
	Scale   float64
}

type Canvas struct {
	Size    int
	Width   int
	Height  int
	Center  Point
	Scale   float64
	Img     *image.NRGBA
	Pixels  []Pixel
}

func NewCanvas(p CnvParams) Canvas {
	w, h := 3*p.Size, 2*p.Size
	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	return Canvas{
		Size: p.Size,
		Width: w,
		Height: h,
		Center: p.Center,
		Scale: p.Scale,
		Img: img,
		Pixels: make([]Pixel, w*h),
	}
}

func (cnv Canvas) Prepare() Canvas {
	w, h := cnv.Width, cnv.Height
	fw, fh := float64(w), float64(h)
	s := float64(cnv.Scale)
	cx := cnv.Center.X
	cy := cnv.Center.Y

	i := 0
	for k := 0; k < w; k++ {
		for l := 0; l < h; l++ {
			fk := float64(k)
			fl := float64(l)
			x := s * ((3.0/fw)*fk + cx - 1.5) + cx
			y := s * ((2.0/fh)*fl + cy - 1.0) + cy

			cnv.Pixels[i] = Pixel{k, l, x, y}

			i++
		}
	}

	return cnv
}

type Pixel struct {
	K int
	L int
	X float64
	Y float64
}

type Point struct {
	X float64
	Y float64
}
