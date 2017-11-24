package main

import (
	"github.com/wwgberlin/godagio/color"
	"strings"
)

type (
	palette []color.Color
)

func (p palette) String() string {
	arr := make([]string, len(p))
	for i, c := range p {
		arr[i] = c.Hex()
	}
	return strings.Join(arr, ",")
}