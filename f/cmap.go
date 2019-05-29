package f

type CmapGen func(p Palette) Cmap

type Cmap func(n int, f Fractal) (r, g, b, a uint8)

type Palette func() []rgb

type rgb struct {
	r, g, b uint8
}

func NewCmap(p Palette) Cmap {
	c := p()

	return func (n int, f Fractal)  (r, g, b, a uint8) {
		dc := len(c)/f.N
		i := dc * (n-1)

		return c[i].r, c[i].g, c[i].b, 255
	}
}

func CoolBlue() []rgb {
	cMap := []rgb{}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgb{255-k, 255-k, 255-k})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgb{0, 255-k, 255})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgb{255-k, 255-k, 255-k})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgb{0, 255-k, 255})
	}

	return cMap
}

func Bool(n int, f Fractal) (r, g, b, a uint8) {
	if n < f.N {
		r, g, b = 255, 255, 255
	}
	return r, g, b, 255
}
