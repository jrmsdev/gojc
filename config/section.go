// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strconv"

	"github.com/jrmsdev/gojc/errors"
)

var (
	ErrOption      = errors.New("config section '%s' option '%s': not found")
	ErrOptionSet   = errors.New("config section '%s' option '%s': already set")
	ErrEmptyOption = errors.New("config section '%s': empty option name")
	ErrOptionParse = errors.New("config section '%s' option '%s': %s")
)

type Option map[string]string

type Section struct {
	cfg  *Config
	name string
	opt  Option
}

func newSection(cfg *Config, name string) *Section {
	return &Section{cfg, name, make(Option)}
}

// HasOption checks if the named option exists in this section.
func (s *Section) HasOption(name string) bool {
	_, found := s.opt[name]
	return found
}

// Options returns a list of all option names in this section.
func (s *Section) Options() []string {
	l := make([]string, 0)
	for n := range s.opt {
		l = append(l, n)
	}
	return l
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
		if s.name != "default" && s.cfg.HasOption("default", option) {
			return s.cfg.GetRaw("default", option)
		} else {
			panic(ErrOption.Format(s.name, option))
		}
	}
	return val
}

// Get returns expanded value of section's option.
// Panics if section or option are not found or if there's any Eval error.
func (s *Section) Get(option string) string {
	return s.cfg.Eval(s.name, s.GetRaw(option))
}

// GetBool returns bool value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetBool(option string) bool {
	val, err := strconv.ParseBool(s.cfg.Get(s.name, option))
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return val
}

// GetFloat32 returns float32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetFloat32(option string) float32 {
	val, err := strconv.ParseFloat(s.cfg.Get(s.name, option), 32)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return float32(val)
}

// GetFloat64 returns float64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetFloat64(option string) float64 {
	val, err := strconv.ParseFloat(s.cfg.Get(s.name, option), 64)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return float64(val)
}

// GetInt returns int value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetInt(option string) int {
	val, err := strconv.ParseInt(s.cfg.Get(s.name, option), 0, 0)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return int(val)
}

// GetInt8 returns int8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetInt8(option string) int8 {
	val, err := strconv.ParseInt(s.cfg.Get(s.name, option), 0, 8)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return int8(val)
}

// GetInt16 returns int16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetInt16(option string) int16 {
	val, err := strconv.ParseInt(s.cfg.Get(s.name, option), 0, 16)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return int16(val)
}

// GetInt32 returns int32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetInt32(option string) int32 {
	val, err := strconv.ParseInt(s.cfg.Get(s.name, option), 0, 32)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return int32(val)
}

// GetInt64 returns int64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetInt64(option string) int64 {
	val, err := strconv.ParseInt(s.cfg.Get(s.name, option), 0, 64)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return int64(val)
}

// GetUint returns uint value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetUint(option string) uint {
	val, err := strconv.ParseUint(s.cfg.Get(s.name, option), 0, 0)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return uint(val)
}

// GetUint8 returns uint8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetUint8(option string) uint8 {
	val, err := strconv.ParseUint(s.cfg.Get(s.name, option), 0, 8)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return uint8(val)
}

// GetUint16 returns uint16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetUint16(option string) uint16 {
	val, err := strconv.ParseUint(s.cfg.Get(s.name, option), 0, 16)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return uint16(val)
}

// GetUint32 returns uint32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetUint32(option string) uint32 {
	val, err := strconv.ParseUint(s.cfg.Get(s.name, option), 0, 32)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return uint32(val)
}

// GetUint64 returns uint64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func (s *Section) GetUint64(option string) uint64 {
	val, err := strconv.ParseUint(s.cfg.Get(s.name, option), 0, 64)
	if err != nil {
		panic(ErrOptionParse.Format(s.name, option, err))
	}
	return uint64(val)
}
