package schemers

import (
	"fmt"
	"github.com/wwgberlin/godagio/color"
)

type (
	Schemer func(base color.Color, size int) Palette

	Schemers interface {
		Get(string) (Schemer, error)
	}

	undefinedSchemerError struct {
		schemerID string
	}

	schemers map[string]Schemer // modify this to schemer builder? is there a difference in construction of schemers?
)

func New() Schemers {
	return &schemers{
		"monochromatic": monochromaticSchemer,
	}
}

func (s *schemers) Get(id string) (Schemer, error) {
	if schemer, ok := (*s)[id]; ok {
		return schemer, nil
	}
	return nil, &undefinedSchemerError{id}
}

func (e *undefinedSchemerError) Error() string {
	return fmt.Sprintf("Schemer %s doesn't exist", e.schemerID)
}
