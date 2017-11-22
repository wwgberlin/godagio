package color

import "github.com/lucasb-eyer/go-colorful"

// todo: could not find the implantation of what we need in the core lib is this okay?
// using this library to wrap colorful lib
// if needed, can add RGB255() to interface...

type (
	Color interface {
		Hex() string
		Hsv() (h, s, v float64)
	}
)

func Hex(h string) (Color, error) {
	return colorful.Hex(h)
}
