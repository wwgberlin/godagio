package color

import (
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

// todo: could not find the implantation of what we need in the core lib is this okay?
// using this library to wrap colorful lib
// if needed, can add RGB255() to interface...

type (
	Color interface {
		Hex() string
		Hsl() (h, s, l float64)
		Complement() Color
	}
	color struct {
		colorful.Color
	}
)

func Hex(h string) (Color, error) {
	c, err := colorful.Hex(h)
	return color{c}, err
}

func Hsl(h, s, l float64) Color {
	c := colorful.Hsl(h, s, l)
	return color{c}
}

func (c color) Complement() Color {
	h, s, l := c.Hsl()
	h = math.Mod((h + 180), 360)
	return Hsl(h, s, l)
}
