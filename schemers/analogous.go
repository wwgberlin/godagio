package schemers

import (
	"github.com/wwgberlin/godagio/color"
)

func analogousSchemer(base color.Color, size int, r randomizer) Palette {
	if size == 1 {
		return palette{base}
	}

	deg := float64(r(360)) / float64(size)

	h, s, l := base.Hsl()
	p := make([]color.Color, size)

	for i := 0; i < size; i++ {
		d := h + deg*float64(size-i)
		if d > 360 {
			d = 360 - d
		}
		p[i] = color.Hsl(d, s, l)
	}
	return palette(p)
}
