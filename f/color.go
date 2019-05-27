package f

type Cmap func(z complex128, fmap Fmap, G float64, N int) (r, g, b, a uint8)

func ColorTrue(z complex128, fmap Fmap, G float64, N int) (r, g, b, a uint8) {
	n := fmap(z, G, N)

	if n < N {
		r, g, b = 255, 255, 255
	}
	return r, g ,b, 255
}
