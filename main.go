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
	fnum    int
)

func r(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}


func init() {
	flag.IntVar(&fnum, "num", 5, "number of colors in palette")
	flag.StringVar(&fscheme, "scheme", "default", "scheme type")

	// base color is mandatory? - randomize one
	// how many base colors do we want to allow?
	flag.StringVar(&fbase, "base", fmt.Sprintf("#%02X%02X%02X", r(255), r(255), r(255)), "number of colors in palette")
}


func main() {
	flag.Parse()

	//todo: validate num > 0

	baseColor, err := color.Hex(fbase) // todo: add proper validation

	if err != nil {
		log.Fatal(err)
	}

	if fscheme == "default" {
		log.Fatal("must specify scheme") //todo add proper validator
	}

	if schemer, err := schemers.New().Get(fscheme); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(schemer.GetPalette(baseColor, fnum))
	}
}
