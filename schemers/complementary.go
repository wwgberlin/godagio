package schemers

import (
	"github.com/wwgberlin/godagio/color"
)

func complementarySchemer(base color.Color, size int, r randomizer) []color.Color {
	var bm []color.Color

	if size%2 == 1 {
		bm = monochromaticSchemer(base, 1+size/2, r)
	} else {
		bm = monochromaticSchemer(base, size/2, r)
	}

	cm := monochromaticSchemer(base.Complement(), size/2, r)

	return append(bm, cm...)
}
