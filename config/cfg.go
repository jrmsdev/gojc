// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/gojc/errors"
)

var ErrSection = errors.New("section not found")

type Cfg map[string]*Section

// Section returns a pointer to the named section. Panics if section not found.
func (c Cfg) Section(name string) *Section {
	s, found := c[name]
	if !found {
		panic(ErrSection.Format("%s", name))
	}
	return s
	//~ return nil
}

// HasSection returns true if section name exists.
func (c *Cfg) HasSection(name string) bool {
	return false
}

// GetRaw returns the raw string value of section's option.
// Panics if section or option are not found or if any parsing error.
func (c *Cfg) GetRaw(section, option string) string {
	return c.Section(section).GetRaw(option)
}
