package main

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/wwgberlin/godagio/color"
)

type (
	palette []color.Color
)

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
