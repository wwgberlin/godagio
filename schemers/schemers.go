package schemers

import (
	"fmt"
	"github.com/wwgberlin/godagio/color"
)

type (
	randomizer func(int) int
	Schemer func(base []color.Color, size int, r randomizer) ([]color.Color, error)

	Schemers interface {
		Get(string) (Schemer, error)
		Register(id string, schemer Schemer)
	}

	undefinedSchemerError struct {
		schemerID string
	}

	schemers map[string]Schemer
)

func randomSchemer(m schemers, r randomizer) func([]color.Color, int, randomizer) ([]color.Color, error) {
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
		"monochromatic": monochromaticSchemer,
		"analogous":     analogousSchemer,
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
