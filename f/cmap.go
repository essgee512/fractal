package f

func init() {
	registerCoolBlue()
}

type Cmap func(n int, f Fractal) (r, g, b, a uint8)

func Bool(n int, f Fractal) (r, g, b, a uint8) {
	if n < f.N {
		r, g, b = 255, 255, 255
	}
	return r, g, b, 255
}

func CoolBlue(n int, f Fractal)  (r, g, b, a uint8) {
	dc := len(cb)/f.N
	i := dc * (n-1)

	return cb[i].r, cb[i].g, cb[i].b, 255
}

type rgb struct {
	r, g, b uint8
}

var cb []rgb
func registerCoolBlue() {
	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cb = append(cb, rgb{255-k, 255-k, 255-k})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cb = append(cb, rgb{0, 255-k, 255})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cb = append(cb, rgb{255-k, 255-k, 255-k})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cb = append(cb, rgb{0, 255-k, 255})
	}
}
