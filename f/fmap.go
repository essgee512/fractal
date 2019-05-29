package f

type Fmap func(z complex128, f Fractal) int

func Mandelbrot(G float64, N int) Fmap {
	return func(c complex128, f Fractal) int {
		z := complex(0, 0)
		i := 0

		for real(z)*real(z)+imag(z)*imag(z) <= G*G {
			z = z*z + c
			if i == N {
				break
			}
			i++
		}

		return i
	}
}

func Julia(Cj complex128, G float64, N int) Fmap {
	return func(z complex128, f Fractal) int {
		i := 0
		for real(z)*real(z)+imag(z)*imag(z) <= G*G {
			z = z*z + Cj
			if i == N {
				break
			}
			i++
		}

		return i
	}
}
