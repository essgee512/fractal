package f

type Fmap func(z complex128, G float64, N int) int

func Mandelbrot(c complex128, G float64, N int) int {
	z := complex(0, 0)
	i := 0
	rMax := G*G
	iMax := N

	for real(z)*real(z) + imag(z)*imag(z) <= rMax {
		z = z*z + c
		if i == iMax {
			break
		}
		i++
	}

	return i
}

func Julia(z, cj complex128, G float64, N int) int {
	i := 0
	rMax := G*G
	iMax := N

	for real(z)*real(z) + imag(z)*imag(z) <= rMax {
		z = z*z + cj
		if i == iMax {
			break
		}
		i++
	}

	return i
}
