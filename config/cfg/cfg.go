// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// Package cfg implements a global config manager.
package cfg

import (
	"io"

	"github.com/jrmsdev/gojc/config"
)

var cfg *config.Config

func init() {
	cfg = config.New(nil)
}

// Map creates a new config with src as it initial content.
// Panics if there's any error, like empty section names and such.
func Map(src config.Cfg) {
	cfg.Map(src)
}

// Read acts like ReadFile but using a filename instead.
func Read(filename string) error {
	return cfg.Read(filename)
}

// ReadFile parses file's content. Returns an error if any.
func ReadFile(file io.ReadSeeker) error {
	return cfg.ReadFile(file)
}

// HasSection checks if section name exists.
func HasSection(name string) bool {
	return cfg.HasSection(name)
}

// Section returns a pointer to the named section. Panics if section not found.
func Section(name string) *config.Section {
	return cfg.Section(name)
}

// HasOption checks if option exists in named section.
func HasOption(section, option string) bool {
	return cfg.HasOption(section, option)
}

// Set sets option's value in section. If section does not exists it is created.
// Panics if option already exists.
func Set(section, option, value string) {
	cfg.Set(section, option, value)
}

// Update updates option's value in section. If section does not exists it is
// created. It is ok if the option already exists.
func Update(section, option, value string) {
	cfg.Update(section, option, value)
}

// GetRaw returns the raw string value of section's option.
// Panics if section or option are not found.
func GetRaw(section, option string) string {
	return cfg.GetRaw(section, option)
}

// Get returns expanded value of section's option.
// Panics if section or option are not found or if there's any Eval error.
func Get(section, option string) string {
	return cfg.Get(section, option)
}

// GetBool returns bool value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetBool(section, option string) bool {
	return cfg.GetBool(section, option)
}

// GetFloat32 returns float32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetFloat32(section, option string) float32 {
	return cfg.GetFloat32(section, option)
}

// GetFloat64 returns float64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetFloat64(section, option string) float64 {
	return cfg.GetFloat64(section, option)
}

// GetInt returns int value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetInt(section, option string) int {
	return cfg.GetInt(section, option)
}

// GetInt8 returns int8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetInt8(section, option string) int8 {
	return cfg.GetInt8(section, option)
}

// GetInt16 returns int16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetInt16(section, option string) int16 {
	return cfg.GetInt16(section, option)
}

// GetInt32 returns int32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetInt32(section, option string) int32 {
	return cfg.GetInt32(section, option)
}

// GetInt64 returns int64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetInt64(section, option string) int64 {
	return cfg.GetInt64(section, option)
}

// GetUint returns uint value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetUint(section, option string) uint {
	return cfg.GetUint(section, option)
}

// GetUint8 returns uint8 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetUint8(section, option string) uint8 {
	return cfg.GetUint8(section, option)
}

// GetUint16 returns uint16 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetUint16(section, option string) uint16 {
	return cfg.GetUint16(section, option)
}

// GetUint32 returns uint32 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetUint32(section, option string) uint32 {
	return cfg.GetUint32(section, option)
}

// GetUint64 returns uint64 value of section's option.
// Panics if section or option are not found or if there's any parsing error.
func GetUint64(section, option string) uint64 {
	return cfg.GetUint64(section, option)
}
