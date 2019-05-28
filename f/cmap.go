package f

type Cmap func(n int) (r, g, b, a uint8)

func Bool(n int) (r, g, b, a uint8) {
	if n < N {
		r, g, b = 255, 255, 255
	}
	return r, g ,b, 255
}
