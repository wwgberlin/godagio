package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/wwgberlin/godagio/color"
	"github.com/wwgberlin/godagio/schemers"
)

var (
	fbase   string
	fscheme string
	foutput string
	fsize   int
)

func r(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func init() {
	flag.IntVar(&fsize, "size", 5, "Number of colors in palette")
	flag.StringVar(&fscheme, "scheme", "random", "Color scheme type [random|monochromatic|analogous]")
	flag.StringVar(&foutput, "output", "text", "Output type [html|text]")

	// todo: base color is mandatory? - randomize one
	// how many base colors do we want to allow?
	flag.StringVar(&fbase, "base", "random", "Comma separated base color values. \n\tExample [#FFDD00,#FFCC00]\n\tIf not given or random specified will randomize a base color")
}

func main() {
	flag.Parse()

	if fbase == "random" {
		fbase = fmt.Sprintf("#%02X%02X%02X", r(255), r(255), r(255))
	}

	if foutput != "text" && fbase != "html" {
		flag.Usage()
		return
	}

	if fsize < 0 {
		flag.Usage()
		return
	}

	baseColor, err := color.Hex(fbase) // todo: add proper validation

	if err != nil {
		flag.Usage()
		return
	}

	if schemer, err := schemers.New(r).Get(fscheme); err != nil {
		flag.Usage()
		return
	} else {
		colors := schemer(baseColor, fsize, r)

		p := palette(colors)
		if foutput == "html" {
			fmt.Println(p.Html())
		} else if foutput == "text" {
			fmt.Println(p)
		}
	}

}
