package schemers

import (
	"fmt"

	"github.com/wwgberlin/godagio/color"
)

type (
	randomizer func(int) int
	schemer    func(base color.Color, size int, r randomizer) []color.Color

	Schemer interface {
		Run(base color.Color, size int, r randomizer) []color.Color
	}

	Schemers interface {
		Get(string) (Schemer, error)
		Register(id string, schemer Schemer)
	}

	undefinedSchemerError struct {
		schemerID string
	}

	schemers map[string]Schemer
)

func (s schemer) Run(base color.Color, size int, r randomizer) []color.Color {
	return s(base, size, r)
}

func randomSchemer(m schemers, r randomizer) Schemer {
	idx := r(len(m))
	i := 0
	for _, v := range m {
		if idx == i {
			return v
		}
		i++
	}
	return nil
}

func New(r randomizer) Schemers {
	m := schemers{
		"monochromatic": schemer(monochromaticSchemer),
		"complementary": schemer(complementarySchemer),
		"analogous":     schemer(analogousSchemer),
	}
	m["random"] = randomSchemer(m, r)
	return &m
}

func (s schemers) Register(id string, schemer Schemer) {
	s[id] = schemer
}

func (s schemers) Get(id string) (Schemer, error) {
	if schemer, ok := s[id]; ok {
		return schemer, nil
	}
	return nil, &undefinedSchemerError{id}
}

func (e undefinedSchemerError) Error() string {
	return fmt.Sprintf("Schemer %s doesn't exist", e.schemerID)
}
