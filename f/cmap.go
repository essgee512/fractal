package f

type Cmap func(n int, f Fractal) (r, g, b, a uint8)

func Bool(n int, f Fractal) (r, g, b, a uint8) {
	if n < f.N {
		r, g, b = 255, 255, 255
	}
	return r, g, b, 255
}
