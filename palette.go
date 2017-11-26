package main

import (
	"fmt"
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
	str := "<body bgcolor='gray' text='white'><table align='center' style='border-spacing:0' width='50%' height='100%'><tr>"
	for _, c := range p {
		str += fmt.Sprintf("<td bgcolor='%s'></td>", c.Hex())
	}
	str += "</tr><tr>"
	for _, c := range p {
		str += fmt.Sprintf("<td align='center' valign='top' height='20px'>%s</td>", c.Hex())
	}
	str += "</tr></table></body>"

	return str
}
