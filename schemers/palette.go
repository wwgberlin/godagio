package schemers

import (
	"github.com/wwgberlin/godagio/color"
	"strings"
	"sort"
)

type (
	Palette interface {
		String() string
	}

	palette []color.Color
	by func(c1, c2 *color.Color) bool

	colorSorter struct {
		colors []color.Color
		by      by
	}
)

func (s *colorSorter) Less(i, j int) bool {
	return s.by(&s.colors[i], &s.colors[j])
}

func (s *colorSorter) Len() int {
	return len(s.colors)
}

func (s *colorSorter) Swap(i, j int) {
	s.colors[i], s.colors[j] = s.colors[j], s.colors[i]
}
func (p palette) String() string {
	arr := make([]string, len(p))
	for i, c := range p {
		arr[i] = c.Hex()
	}
	return strings.Join(arr, ",")
}

func saturation(c1, c2 *color.Color) bool {
	_, s1, _ := (*c1).Hsl()
	_, s2, _ := (*c2).Hsl()
	return s1 > s2
}
func lightness(c1, c2 *color.Color) bool {
	_, _, l1 := (*c1).Hsl()
	_, _, l2 := (*c2).Hsl()
	return l1 > l2
}
func (by by) Sort(colors []color.Color) {
	ps := &colorSorter{
		colors: colors,
		by:      by,
	}
	sort.Sort(ps)
}