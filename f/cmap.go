package f

type Cmap func(n int, f Fractal) (r, g, b, a uint8)

type rgba struct {
	r, g, b, a uint8
}

func CoolBlue(N int) Cmap {
	cMap := []rgba{}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgba{255-k, 255-k, 255-k, k})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgba{0, 255-k, 255, 255-k})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgba{255-k, 255-k, 255-k, 255})
	}

	for j := 0; j <= 255; j++ {
		k := uint8(j)
		cMap = append(cMap, rgba{0, 255-k, 255, 255})
	}

	return func(n int, f Fractal)  (r, g, b, a uint8) {
		dc := len(cMap)/N
		i := dc * (n-1)

		return cMap[i].r, cMap[i].g, cMap[i].b, cMap[i].a
	}
}

func Bool(N int) Cmap {
	return func(n int, f Fractal)  (r, g, b, a uint8) {
		if n < N {
			r, g, b, a = 255, 255, 255, 255
		}
		return r, g, b, a
	}
}
