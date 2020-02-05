// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"github.com/jrmsdev/gojc/errors"
)

var (
	ErrSection = errors.New("config section not found: %s")
	ErrEmptySection = errors.New("config section empty name")
)

type Cfg map[string]Option

type Config struct {
	sect map[string]*Section
}

// New creates a new config object. If defaults is not nil, it's Map'ed as its
// initial content.
func New(defaults Cfg) *Config {
	c := &Config{make(map[string]*Section)}
	if defaults != nil {
		c.Map(defaults)
	}
	return c
}

// Map updates config data from src Cfg instance.
// Panics if there's any error, like empty section or option names.
func (c *Config) Map(src Cfg) {
	for s, l := range src {
		for o, v := range l {
			c.Update(s, o, v)
		}
	}
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
	return s
}

// HasOption checks if option exists in named section.
func (c *Config) HasOption(section, option string) bool {
	if !c.HasSection(section) {
		return false
	}
	return c.Section(section).HasOption(option)
}

// Set sets option's value in section. If section does not exists it is created.
// Panics if option already exists.
func (c *Config) Set(section, option, value string) {
	if section == "" {
		panic(ErrEmptySection)
	}
	if !c.HasSection(section) {
		c.sect[section] = newSection(c, section)
	}
	c.Section(section).Set(option, value)
}

// Update updates option's value in section. If section does not exists it is
// created. It is ok if the option already exists.
func (c *Config) Update(section, option, value string) {
	if section == "" {
		panic(ErrEmptySection)
	}
	if !c.HasSection(section) {
		c.sect[section] = newSection(c, section)
	}
	c.Section(section).Update(option, value)
}

// GetRaw returns the raw string value of section's option.
// Panics if section or option are not found.
func (c *Config) GetRaw(section, option string) string {
	return c.Section(section).GetRaw(option)
}

// Get returns expanded value of section's option.
// Panics if section or option are not found or if there's any Eval error.
func (c *Config) Get(section, option string) string {
	return c.Section(section).Get(option)
}

// GetBool returns bool value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetBool(section, option string) bool {
	return c.Section(section).GetBool(option)
}

// GetFloat32 returns float32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetFloat32(section, option string) float32 {
	return c.Section(section).GetFloat32(option)
}

// GetFloat64 returns float64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetFloat64(section, option string) float64 {
	return c.Section(section).GetFloat64(option)
}

// GetInt returns int value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt(section, option string) int {
	return c.Section(section).GetInt(option)
}

// GetInt8 returns int8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt8(section, option string) int8 {
	return c.Section(section).GetInt8(option)
}

// GetInt16 returns int16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt16(section, option string) int16 {
	return c.Section(section).GetInt16(option)
}

// GetInt32 returns int32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt32(section, option string) int32 {
	return c.Section(section).GetInt32(option)
}

// GetInt64 returns int64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt64(section, option string) int64 {
	return c.Section(section).GetInt64(option)
}

// GetUint returns uint value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint(section, option string) uint {
	return c.Section(section).GetUint(option)
}

// GetUint8 returns uint8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint8(section, option string) uint8 {
	return c.Section(section).GetUint8(option)
}

// GetUint16 returns uint16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint16(section, option string) uint16 {
	return c.Section(section).GetUint16(option)
}

// GetUint32 returns uint32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint32(section, option string) uint32 {
	return c.Section(section).GetUint32(option)
}

// GetUint64 returns uint64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint64(section, option string) uint64 {
	return c.Section(section).GetUint64(option)
}
