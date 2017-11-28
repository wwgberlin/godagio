package palette

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/wwgberlin/godagio/color"
)

type (
	Palette interface {
		Html() string
	}

	palette []color.Color
)

func New(colors []color.Color) Palette {
	return palette(colors)
}

func (p palette) String() string {
	arr := make([]string, len(p))
	for i, c := range p {
		arr[i] = c.Hex()
	}
	return strings.Join(arr, ", ")
}

func (p palette) Html() string {
	var tpl bytes.Buffer

	t := template.New("palette.tmpl")
	t.ParseFiles("templates/palette.tmpl")
	t.Execute(&tpl, p)

	return tpl.String()
}
