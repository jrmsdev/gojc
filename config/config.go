// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"io"
	//~ "os"
)

type Config struct {
	cfg *Cfg
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
func Map(src *Cfg) *Config {
	return &Config{src}
}

// Section returns a pointer to the named section. Panics if section not found.
func (c *Config) Section(name string) *Section {
	return c.cfg.Section(name)
}

// HasSection returns true if section name exists.
func (c *Config) HasSection(name string) bool {
	return c.cfg.HasSection(name)
}

// GetRaw returns the raw string value of section's option.
// Panics if section or option are not found or if any parsing error.
func (c *Config) GetRaw(section, option string) string {
	return c.cfg.GetRaw(section, option)
}