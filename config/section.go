// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/gojc/errors"
)

var ErrOption = errors.New("section option not found")

type Section map[string]string

// HasOption returns true if the named option exists in this section.
func (s *Section) HasOption(name string) bool {
	return false
}

// GetRaw returns the raw string value of option from this section.
// Panics if option does not exists.
func (s Section) GetRaw(option string) string {
	val, found := s[option]
	if !found {
		panic(ErrOption.Format("%s", option))
	}
	return val
}
