package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
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

type (
	constraints []func() error
)

func (c constraints) validate() (err error) {
	var errstrings []string
	for _, constraint := range c {
		if err := constraint(); err != nil {
			errstrings = append(errstrings, err.Error())
		}
	}
	if len(errstrings) > 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}
	return nil
}

func r(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func init() {
	flag.IntVar(&fsize, "size", 5, "Number of colors in palette")
	flag.StringVar(&fscheme, "scheme", "random",
		"Color scheme type [random|monochromatic|analogous]")
	flag.StringVar(&foutput, "output", "text", "Output type [html|text]")

	// todo: base color is mandatory? - randomize one
	// how many base colors do we want to allow?
	flag.StringVar(&fbase, "base", "random",
		"Comma separated base color values. "+
			"\n\tExample [#FFDD00,#FFCC00]\n\tIf not given or random specified will randomize a base color")
}

func main() {
	flag.Parse()

	if fbase == "random" {
		fbase = fmt.Sprintf("#%02X%02X%02X", r(255), r(255), r(255))
	}

	var baseColor color.Color
	var schemer schemers.Schemer

	c := constraints{
		func() (err error) {
			baseColor, err = color.Hex(fbase)
			return err
		},
		func() error {
			if fsize < 0 {
				return fmt.Errorf("size must be > 0")
			}
			return nil
		},
		func() error {
			if foutput != "text" && foutput != "html" {
				return fmt.Errorf("unrecognized output type")
			}
			return nil
		},
		func() (err error) {
			schemer, err = schemers.New(r).Get(fscheme)
			return err
		},
	}

	if err := c.validate(); err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	colors := schemer(baseColor, fsize, r)

	p := palette(colors)
	if foutput == "html" {
		fmt.Println(p.Html())
	} else if foutput == "text" {
		fmt.Println(p)
	}

}
