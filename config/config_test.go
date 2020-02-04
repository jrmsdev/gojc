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
	IsFalse(t, c.HasOption("testing", "t0"), "has option")

	tcfg["default"]["testing"] = "t0"
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "overwrite map src")

	c.sect["default"]["testing"] = "t1"
	IsEqual(t, tcfg["default"]["testing"], "t0", "tcfg override")
	IsEqual(t, c.GetRaw("default", "testing"), "t1", "overwrite c.sect")

	c2 := c
	IsEqual(t, c2.GetRaw("default", "testing"), "t1", "c2 get raw")

	c2.sect["default"]["testing"] = "t2"
	IsEqual(t, tcfg["default"]["testing"], "t0", "overwrite c2.sect")
	IsEqual(t, c.GetRaw("default", "testing"), "t2", "c get raw")
	IsEqual(t, c2.GetRaw("default", "testing"), "t2", "c2 get raw")

	IsEqual(t, checkArg(c, "t3"), "t2", "c arg")
	IsEqual(t, checkArg(c2, "t4"), "t2", "c2 arg")

	IsTrue(t, c.HasSection("default"), "default section")
	IsFalse(t, c.HasSection("testing"), "testing section")

	IsTrue(t, c.HasOption("default", "testing"), "default option")
	IsFalse(t, c.HasOption("default", "noopt"), "no option")
}

func TestConfigSet(t *testing.T) {
	c := Map(tcfg)
	c.Set("set", "t0", "v0")
	IsEqual(t, c.GetRaw("set", "t0"), "v0", "set get")
	IsEqual(t, tcfg["set"]["t0"], "", "set tcfg")
}

func TestConfigUpdate(t *testing.T) {
	c := Map(tcfg)
	c.Update("update", "t0", "v0")
	IsEqual(t, c.GetRaw("update", "t0"), "v0", "update get")
	IsEqual(t, tcfg["update"]["t0"], "", "update tcfg")
	c.Update("update", "t0", "v1")
	IsEqual(t, c.GetRaw("update", "t0"), "v1", "update get v1")
	IsEqual(t, tcfg["update"]["t0"], "", "update tcfg v1")
}

func TestConfigGet(t *testing.T) {
	c := Map(tcfg)
	c.Update("get", "t0", "v0")
	IsEqual(t, c.Get("get", "t0"), "v0", "get")
}
