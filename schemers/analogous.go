package schemers

import (
	"github.com/wwgberlin/godagio/color"
)

func analogousSchemer(bases []color.Color, size int, r randomizer) ([]color.Color, error) {
	if size == len(bases) {
		return bases, nil
	}

	deg := float64(360) / float64(size)

	base := bases[0]
	h, s, l := base.Hsl()
	p := make([]color.Color, size)

	for i := 0; i < size; i++ {
		d := h + deg*float64(size-i)
		if d > 360 {
			d = - (360 - d)
		}
		p[i] = color.Hsl(d, s, l)
	}
	return p, nil
}
