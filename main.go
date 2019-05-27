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
	cnv := f.NewCanvas(f.CnvParams{
		Size: 300,
		Center: f.Point{0.0, 0.0},
		Scale: 1.0,
	})

	m := f.NewFractal(f.Fractal{
		Canvas: cnv,
		Fmap: f.Mandelbrot,
		G: 2.0,
		N: 1000,
		Cmap: f.ColorTrue,
	})

	m.Render()

	// j := f.NewFractal(f.Fractal{
	// 	Canvas: cnv,
	// 	Fmap: f.Julia(−0.8 + 0.156i), // Cj
	// 	G: 2.0,
	// 	N: 1000,
	// 	Cmap: f.ColorTrue,
	// })

	// j.Render()

	f, err := os.Create("output/image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, m.Canvas.Img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
