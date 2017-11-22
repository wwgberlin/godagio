package schemers

import (
	"github.com/wwgberlin/godagio/color"
	"math/rand"
	"time"
)

func r(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func monochromaticSchemer(base color.Color, size int) Palette {
	if size == 1 {
		return palette{base}
	}

	h, _, v := base.Hsv()
	p := make([]color.Color, size)

	margin := 100/size
	for i := 0; i < size; i++ {
		d := float64(r(margin) + margin*i)
		s := float64(d / float64(100.0))
		p[i] = color.Hsv(h, s, v)

	}
	p[size/2+1] = base
	by(saturation).Sort(p)
	return palette(p)
}
