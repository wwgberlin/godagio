package color

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

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
	h = math.Mod(h+180, 360)
	return Hsl(h, s, l)
}
