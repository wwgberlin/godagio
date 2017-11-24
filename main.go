package main

import (
	"log"
	"flag"
	"math/rand"
	"time"
	"fmt"

	"github.com/wwgberlin/godagio/schemers"
	"github.com/wwgberlin/godagio/color"
)

var (
	fbase   string
	fscheme string
	fsize   int
)

func r(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func init() {
	flag.IntVar(&fsize, "size", 5, "number of colors in palette")
	flag.StringVar(&fscheme, "scheme", "random", "scheme type - default is random")

	// todo: base color is mandatory? - randomize one
	// how many base colors do we want to allow?
	flag.StringVar(&fbase, "base", fmt.Sprintf("#%02X%02X%02X", r(255), r(255), r(255)), "optional: base color")
}

func main() {
	flag.Parse()

	//todo: validate size > 0
	//todo: validate size > number of base colors

	baseColor, err := color.Hex(fbase) // todo: add proper validation

	if err != nil {
		log.Fatal(err)
	}

	if schemer, err := schemers.New(r).Get(fscheme); err != nil {
		log.Fatal(err)
	} else {
		if colors, err := schemer([]color.Color{baseColor}, fsize, r); err == nil {
			fmt.Println(palette(colors))
		}else{
			log.Fatal(err)
		}
	}
}
