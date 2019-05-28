package f

type Fmap func(z complex128, f Fractal) int

func Mandelbrot(c complex128, f Fractal) int {
	z := complex(0, 0)
	i := 0
	rMax := f.G * f.G
	iMax := f.N

	for real(z)*real(z)+imag(z)*imag(z) <= rMax {
		z = z*z + c
		if i == iMax {
			break
		}
		i++
	}

	return i
}

func Julia(z complex128, f Fractal) int {
	Cj := complex(-0.8, 0.156)
	rMax := f.G * f.G
	iMax := f.N

	i := 0
	for real(z)*real(z)+imag(z)*imag(z) <= rMax {
		z = z*z + Cj
		if i == iMax {
			break
		}
		i++
	}

	return i
}
