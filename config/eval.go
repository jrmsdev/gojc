// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"strings"
)

// Eval parses expr string and returns its expanded value.
// Panics if there's any parsing error.
func (c *Config) Eval(expr string) string {
	expr = strings.Replace(expr, "\r", "", -1)
	expr = strings.Replace(expr, "\n", "", -1)
	expr = strings.TrimLeft(expr, " \t")
	return expr
}
