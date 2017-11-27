package schemers

import (
	"github.com/wwgberlin/godagio/color"
)

func analogousSchemer(base color.Color, size int, r randomizer) []color.Color {
	deg := float64(360) / float64(size)

	p := make([]color.Color, size)

	for i := 0; i < size; i++ {
		h, _, _ := base.Rotate(deg * float64(size-i)).Hsl()
		dl := float64(r(100))
		l := float64(dl / float64(100.0))
		ds := float64(r(100))
		s := float64(ds / float64(100.0))
		p[i] = color.Hsl(h, s, l)
	}
	return p
}
