package main

import (
	"fmt"

	"image/png"
	"log"
	"os"

	"github.com/essgee512/fractal/f"

	"github.com/davecgh/go-spew/spew"
)

func unused() {
	fmt.Println()
	spew.Dump()
	scs.Dump()
}

var (
	scs spew.ConfigState = spew.ConfigState{
		Indent:                  "  ",
		DisableMethods:          true,
		DisablePointerAddresses: true,
	}
)

func main() {
	var fractal f.Fractal
	var fp *os.File
	var err error

	N := 100

	for ps := 3; ps <= 3; ps++ {

		fractal = f.NewFractal(f.Fractal{
			Size:    400,
			Center:  f.Point{0.0, 0.0},
			Scale:   1.0,
			PixSize: ps,

			// Fmap:    f.Mandelbrot(10.0, N),
			Fmap:    f.Julia(complex(-0.8, 0.156), 2.0, N),
			Cmap:    f.CoolBlue(N),

			// Fmap:    f.Mandelbrot(f.CoolBlue, G, N),
			// Fmap:    f.Julia(f.CoolBlue, Cj, G, N),

			// Cmap:    f.CoolBlue(f.Mandelbrot, G, N),
			// Cmap:    f.CoolBlue(f.Julia, Cj, G, N),
		})

		fractal.Render()

		fp, err = os.Create(fmt.Sprintf("output/psize-%d.png", ps))
		if err != nil {
			log.Fatal(err)
		}

		if err = png.Encode(fp, fractal.Img); err != nil {
			fp.Close()
			log.Fatal(err)
		}

		if err = fp.Close(); err != nil {
			log.Fatal(err)
		}

	}

}
