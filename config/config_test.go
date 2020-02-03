// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"

	. "github.com/jrmsdev/gojc/testing/check"
)

var tcfg = &Cfg{
	"default": &Section{
		"testing": "ok",
	},
}

func TestConfig(t *testing.T) {
	c := Map(tcfg)
	//~ t.Logf("%#v", c.cfg)
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "get")
}
