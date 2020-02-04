// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"io"
	//~ "os"

	"github.com/jrmsdev/gojc/errors"
)

var ErrSection = errors.New("section not found: %s")
var ErrEmptySection = errors.New("empty section name")

type Cfg map[string]Option

type Config struct {
	sect Cfg
}

// Read parses filename content and returns a new config instance. Or an error,
// if any.
func Read(filename string) (*Config, error) {
	//~ fh, err := os.Open(filename)
	//~ if err != nil {
		//~ return nil, err
	//~ }
	//~ defer fh.Close()
	//~ return ReadFile(fh)
	return nil, nil
}

// ReadFile acts like Read but using a file pointer instead of a filename.
func ReadFile(file io.Reader) (*Config, error) {
	//~ c, err := parseFile(file)
	//~ if err != nil {
		//~ return nil, err
	//~ }
	//~ return &Config{c}, nil
	return nil, nil
}

// Map creates a new config with src as it initial content.
// Panics if there's any error, like empty section names and such.
func Map(src Cfg) *Config {
	dst := make(Cfg)
	for sn, s := range src {
		if sn == "" {
			panic(ErrEmptySection)
		}
		dst[sn] = make(Option)
		for on, o := range s {
			if on == "" {
				panic(ErrEmptyOption.Format(sn))
			}
			dst[sn][on] = o
		}
	}
	return &Config{dst}
}

// HasSection checks if section name exists.
func (c *Config) HasSection(name string) bool {
	_, found := c.sect[name]
	return found
}

// Section returns a pointer to the named section. Panics if section not found.
func (c *Config) Section(name string) *Section {
	s, found := c.sect[name]
	if !found {
		panic(ErrSection.Format(name))
	}
	return &Section{name, s}
}

// HasOption checks if option exists in named section.
func (c *Config) HasOption(section, option string) bool {
	if !c.HasSection(section) {
		return false
	}
	return c.Section(section).HasOption(option)
}

// GetRaw returns the raw string value of section's option.
// Panics if section or option are not found.
func (c *Config) GetRaw(section, option string) string {
	return c.Section(section).GetRaw(option)
}

// Set sets option's value in section. If section does not exists it is created.
// Panics if option already exists.
func (c *Config) Set(section, option, value string) {
	if !c.HasSection(section) {
		c.sect[section] = make(Option)
	}
	c.Section(section).Set(option, value)
}

// Update updates option's value in section. If section does not exists it is
// created. It is ok if the option already exists.
func (c *Config) Update(section, option, value string) {
	if !c.HasSection(section) {
		c.sect[section] = make(Option)
	}
	c.Section(section).Update(option, value)
}

// Get returns expanded value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) Get(section, option string) string {
	return c.Eval(section, c.Section(section).GetRaw(option))
}
