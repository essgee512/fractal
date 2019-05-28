package f

type Fmap func(z complex128) int

func Mandelbrot(c complex128) int {
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

func Julia(z complex128) int {
	Cj := complex(-0.8, 0.156)
	rMax := G*G
	iMax := N

	i := 0
	for real(z)*real(z) + imag(z)*imag(z) <= rMax {
		z = z*z + Cj
		if i == iMax {
			break
		}
		i++
	}

	return i
}
