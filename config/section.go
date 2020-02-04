// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/gojc/errors"
)

var ErrOption = errors.New("config section '%s' option '%s': not found")
var ErrOptionSet = errors.New("config section '%s' option '%s': already set")
var ErrEmptyOption = errors.New("config section '%s': empty option name")

type Option map[string]string

type Section struct {
	name string
	opt  Option
}

// HasOption checks if the named option exists in this section.
func (s *Section) HasOption(name string) bool {
	_, found := s.opt[name]
	return found
}

// Set sets option's value in this section. Panics if option already exists.
func (s *Section) Set(option, value string) {
	if option == "" {
		panic(ErrEmptyOption.Format(s.name))
	}
	_, found := s.opt[option]
	if found {
		panic(ErrOptionSet.Format(s.name, option))
	} else {
		s.opt[option] = value
	}
}

// Update updates option's value in this section.
// If the option already exists, its value will be replaced. It will be created
// otherwise.
// Panics if option name is an empty string.
func (s *Section) Update(option, value string) {
	if option == "" {
		panic(ErrEmptyOption.Format(s.name))
	}
	s.opt[option] = value
}

// GetRaw returns the raw value of option from this section.
// Panics if option is not found.
func (s *Section) GetRaw(option string) string {
	val, found := s.opt[option]
	if !found {
		panic(ErrOption.Format(s.name, option))
	}
	return val
}
