// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"

	. "github.com/jrmsdev/gojc/testing/check"
)

var tcfg = Cfg{
	"default": Option{
		"testing": "ok",
	},
}

func checkArg(c *Config, arg string) string {
	tcfg["default"]["testing"] = arg
	return c.GetRaw("default", "testing")
}

func TestConfig(t *testing.T) {
	c := Map(tcfg)
	//~ t.Logf("%#v", c)
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "get")

	tcfg["default"]["testing"] = "t0"
	IsEqual(t, c.GetRaw("default", "testing"), "t0", "get")

	c.sect["default"]["testing"] = "t1"
	IsEqual(t, tcfg["default"]["testing"], "t1", "get")
	IsEqual(t, c.GetRaw("default", "testing"), "t1", "get")

	c2 := c
	IsEqual(t, c2.GetRaw("default", "testing"), "t1", "get")

	c2.sect["default"]["testing"] = "t2"
	IsEqual(t, tcfg["default"]["testing"], "t2", "get")
	IsEqual(t, c.GetRaw("default", "testing"), "t2", "get")
	IsEqual(t, c2.GetRaw("default", "testing"), "t2", "get")

	IsEqual(t, checkArg(c, "t3"), "t3", "get")
	IsEqual(t, checkArg(c2, "t4"), "t4", "get")
}
