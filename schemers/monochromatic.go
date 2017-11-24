package schemers

import (
	"github.com/wwgberlin/godagio/color"
)

func monochromaticSchemer(bases []color.Color, size int, r randomizer) ([]color.Color, error) {
	if size == 1 {
		return bases, nil
	}

	base := bases[0]

	h, _, _ := base.Hsl()
	p := make([]color.Color, size)

	for i := 0; i < size-1; i++ {
		ds := float64(r(100))
		dl := float64(r(100))
		s := float64(ds / float64(100.0))
		l := float64(dl / float64(100.0))
		p[i] = color.Hsl(h, s, l)
	}
	p[size-1] = base
	return p, nil
}
