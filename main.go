package main

import (
	"fmt"

	"github.com/wwgberlin/godagio/color"
)

func main() {
	//this is example usage of the palette type
	//run go build . && ./godagio > index.html && open index.html
	c1, _ := color.Hex("#CC44CC")
	c2, _ := color.Hex("#44AAAA")
	p := palette([]color.Color{c1, c2})
	fmt.Println(p.Html())
}
