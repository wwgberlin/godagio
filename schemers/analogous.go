package schemers

import (
	"github.com/wwgberlin/godagio/color"
)

func analogousSchemer(base color.Color, size int, r randomizer) []color.Color {
	deg := float64(360) / float64(size)

	h, s, _ := base.Hsl()
	p := make([]color.Color, size)

	for i := 0; i < size; i++ {
		d := h + deg*float64(size-i)

		if d > 360 {
			d = - (360 - d)
		}
		dl := float64(r(100))
		l := float64(dl / float64(100.0))
		p[i] = color.Hsl(d, s, l)
	}
	return p
}
