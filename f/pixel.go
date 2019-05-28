package f

type pixels []pixel

type pixel struct {
	k, l int
	x, y float64
}

func initPixels(f Fractal) pixels {
	s := float64(f.Scale)
	cx, cy := f.Center.X, f.Center.Y
	w, h := 3*f.Size, 2*f.Size
	fw, fh := float64(w), float64(h)

	pixels := make([]pixel, w*h)

	i := 0
	for k := 0; k < w; k++ {
		for l := 0; l < h; l++ {
			fk, fl := float64(k), float64(l)
			x := s * ((3.0/fw)*fk + cx - 1.5) + cx
			y := s * ((2.0/fh)*fl + cy - 1.0) + cy

			pixels[i] = pixel{k, l, x, y}

			i++
		}
	}

	return pixels
}
