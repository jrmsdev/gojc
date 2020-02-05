// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"io"
	//~ "os"
	"strconv"

	"github.com/jrmsdev/gojc/errors"
)

var ErrSection = errors.New("config section not found: %s")
var ErrEmptySection = errors.New("config section empty name")
var ErrOptionParse = errors.New("config section '%s' option '%s': %s parse error")

type Cfg map[string]Option

type Config struct {
	sect Cfg
}

func New(defaults Cfg) *Config {
	c := &Config{make(Cfg)}
	if defaults != nil {
		c.Map(defaults)
	}
	return c
}

// Read parses filename content and returns a new config instance. Or an error,
// if any.
func (c *Config) Read(filename string) (*Config, error) {
	//~ fh, err := os.Open(filename)
	//~ if err != nil {
		//~ return nil, err
	//~ }
	//~ defer fh.Close()
	//~ return ReadFile(fh)
	return nil, nil
}

// ReadFile acts like Read but using a file pointer instead of a filename.
func (c *Config) ReadFile(file io.Reader) (*Config, error) {
	//~ c, err := parseFile(file)
	//~ if err != nil {
		//~ return nil, err
	//~ }
	//~ return &Config{c}, nil
	return nil, nil
}

// Map creates a new config with src as it initial content.
// Panics if there's any error, like empty section names and such.
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
	return &Section{name, s}
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
		c.sect[section] = make(Option)
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
		c.sect[section] = make(Option)
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
	return c.Eval(section, c.Section(section).GetRaw(option))
}

// GetBool returns bool value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetBool(section, option string) bool {
	val, err := strconv.ParseBool(c.Get(section, option))
	if err != nil {
		panic(ErrOptionParse.Format(section, option, "bool"))
	}
	return val
}

// GetFloat returns float64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetFloat(section, option string) float64 {
	val, err := strconv.ParseFloat(c.Get(section, option), 64)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, "float"))
	}
	return val
}

// GetInt returns int value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt(section, option string) int {
	val, err := strconv.ParseInt(c.Get(section, option), 0, 0)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, "int"))
	}
	return int(val)
}

// GetInt8 returns int8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt8(section, option string) int8 {
	val, err := strconv.ParseInt(c.Get(section, option), 0, 8)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return int8(val)
}

// GetInt16 returns int16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt16(section, option string) int16 {
	val, err := strconv.ParseInt(c.Get(section, option), 0, 16)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return int16(val)
}

// GetInt32 returns int32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt32(section, option string) int32 {
	val, err := strconv.ParseInt(c.Get(section, option), 0, 32)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return int32(val)
}

// GetInt64 returns int64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetInt64(section, option string) int64 {
	val, err := strconv.ParseInt(c.Get(section, option), 0, 64)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return int64(val)
}

// GetUint returns uint value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint(section, option string) uint {
	val, err := strconv.ParseUint(c.Get(section, option), 0, 0)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, "uint"))
	}
	return uint(val)
}

// GetUint8 returns uint8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint8(section, option string) uint8 {
	val, err := strconv.ParseUint(c.Get(section, option), 0, 8)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return uint8(val)
}

// GetUint16 returns uint16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint16(section, option string) uint16 {
	val, err := strconv.ParseUint(c.Get(section, option), 0, 16)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return uint16(val)
}

// GetUint32 returns uint32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint32(section, option string) uint32 {
	val, err := strconv.ParseUint(c.Get(section, option), 0, 32)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return uint32(val)
}

// GetUint64 returns uint64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (c *Config) GetUint64(section, option string) uint64 {
	val, err := strconv.ParseUint(c.Get(section, option), 0, 64)
	if err != nil {
		panic(ErrOptionParse.Format(section, option, err))
	}
	return uint64(val)
}
