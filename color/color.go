package color

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

type (
	Color interface {
		Hex() string
		Hsl() (h, s, l float64)
		RGB255() (h, s, l uint8)
		Rotate(angle float64) Color
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

func (c color) Rotate(angle float64) Color {
	h, s, l := c.Hsl()
	h = math.Mod(h+angle, 360)
	return Hsl(h, s, l)
}
