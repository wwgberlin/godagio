package schemers

import (
	"github.com/wwgberlin/godagio/color"
)

type (
	monochromaticSchemer struct {
	}
)

// todo implemnt this
func (s *monochromaticSchemer) GetPalette(base color.Color, num int) Palette {
	return palette{base}
}
