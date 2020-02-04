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
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "get raw")

	tcfg["default"]["testing"] = "t0"
	IsEqual(t, c.GetRaw("default", "testing"), "t0", "get raw override")

	c.sect["default"]["testing"] = "t1"
	IsEqual(t, tcfg["default"]["testing"], "t1", "tcfg override")
	IsEqual(t, c.GetRaw("default", "testing"), "t1", "t1 get raw")

	c2 := c
	IsEqual(t, c2.GetRaw("default", "testing"), "t1", "c2 get raw")

	c2.sect["default"]["testing"] = "t2"
	IsEqual(t, tcfg["default"]["testing"], "t2", "tcfg")
	IsEqual(t, c.GetRaw("default", "testing"), "t2", "c get raw")
	IsEqual(t, c2.GetRaw("default", "testing"), "t2", "c2 get raw")

	IsEqual(t, checkArg(c, "t3"), "t3", "c arg")
	IsEqual(t, checkArg(c2, "t4"), "t4", "c2 arg")

	IsTrue(t, c.HasSection("default"), "default section")
	IsFalse(t, c.HasSection("testing"), "testing section")

	IsTrue(t, c.HasOption("default", "testing"), "default option")
	IsFalse(t, c.HasOption("default", "noopt"), "no option")
}
