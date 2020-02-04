// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/gojc/errors"
)

var ErrOption = errors.New("section option not found")

type Option map[string]string

type Section struct {
	name string
	opt Option
}

// HasOption returns true if the named option exists in this Sectionion.
func (s *Section) HasOption(name string) bool {
	return false
}

// GetRaw returns the raw string value of option from this Sectionion.
// Panics if option does not exists.
func (s *Section) GetRaw(option string) string {
	val, found := s.opt[option]
	if !found {
		panic(ErrOption.Format("%s: %s", s.name, option))
	}
	return val
}
