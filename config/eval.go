// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"regexp"
	"strings"
)

var reOption = regexp.MustCompile(`\${([0-9A-Za-z._-]+)}`)
var reSectionOption = regexp.MustCompile(`\${([0-9A-Za-z._-]+):([0-9A-Za-z._-]+)}`)

// Eval parses expr string and returns its expanded value.
// Panics if there's any parsing error.
// Expression is evaluated in the context of the named section.
func (c *Config) Eval(section, expr string) string {
	dst := reOption.ReplaceAllStringFunc(expr, c.evalOption(section))
	dst = reSectionOption.ReplaceAllStringFunc(dst, c.evalSectionOption(section))
	return dst
}

func (c *Config) evalOption(section string) func(string) string {
	return func(option string) string {
		return c.Get(section, option[2:len(option)-1])
	}
}

func (c *Config) evalSectionOption(section string) func(string) string {
	return func(s string) string {
		opt := ""
		sect := section
		s = s[2:len(s)-1]
		idx := strings.Index(s, ":")
		if idx > 0 {
			sect = s[0:idx]
			opt = s[idx+1:]
		}
		return c.Get(sect, opt)
	}
}
