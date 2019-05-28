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
	// fractal := f.NewFractal(f.Fractal{
	// 	Size: 300,
	// 	Center: f.Point{0.0, 0.0},
	// 	Scale: 1.0,

	// 	Fmap: f.Mandelbrot,
	// 	Cmap: f.Bool,
	// })

	fractal := f.NewFractal(f.Fractal{
		Size:   400,
		Center: f.Point{0.0, 0.0},
		Scale:  1.0,
		G:      5.0,
		N:      100,
		Fmap:   f.Julia,
		Cmap:   f.Bool,
	})

	fractal.Render()

	f, err := os.Create("output/image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, fractal.Img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
