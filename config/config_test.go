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
	c := New(tcfg)
	//~ t.Logf("%#v", c)
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "get raw")
	IsFalse(t, c.HasOption("testing", "t0"), "has option")

	tcfg["default"]["testing"] = "t0"
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "overwrite map src")

	c.sect["default"].opt["testing"] = "t1"
	IsEqual(t, tcfg["default"]["testing"], "t0", "tcfg override")
	IsEqual(t, c.GetRaw("default", "testing"), "t1", "overwrite c.sect")

	c2 := c
	IsEqual(t, c2.GetRaw("default", "testing"), "t1", "c2 get raw")

	c2.sect["default"].opt["testing"] = "t2"
	IsEqual(t, tcfg["default"]["testing"], "t0", "overwrite c2.sect")
	IsEqual(t, c.GetRaw("default", "testing"), "t2", "c get raw")
	IsEqual(t, c2.GetRaw("default", "testing"), "t2", "c2 get raw")

	IsEqual(t, checkArg(c, "t3"), "t2", "c arg")
	IsEqual(t, checkArg(c2, "t4"), "t2", "c2 arg")

	IsTrue(t, c.HasSection("default"), "default section")
	IsFalse(t, c.HasSection("testing"), "testing section")

	IsTrue(t, c.HasOption("default", "testing"), "default option")
	IsFalse(t, c.HasOption("default", "noopt"), "no option")

	IsEqual(t, len(c.Sections()), 0, "sections list")
	IsEqual(t, len(c.Options("default")), 1, "options list")
}

func TestConfigSet(t *testing.T) {
	c := New(tcfg)
	c.Set("set", "t0", "v0")
	IsEqual(t, c.GetRaw("set", "t0"), "v0", "set get")
	IsEqual(t, tcfg["set"]["t0"], "", "set tcfg")
}

func TestConfigUpdate(t *testing.T) {
	c := New(tcfg)
	c.Update("update", "t0", "v0")
	IsEqual(t, c.GetRaw("update", "t0"), "v0", "update get")
	IsEqual(t, tcfg["update"]["t0"], "", "update tcfg")
	c.Update("update", "t0", "v1")
	IsEqual(t, c.GetRaw("update", "t0"), "v1", "update get v1")
	IsEqual(t, tcfg["update"]["t0"], "", "update tcfg v1")
}

func TestConfigGet(t *testing.T) {
	c := New(tcfg)
	c.Update("get", "t0", "v0")
	IsEqual(t, c.Get("get", "t0"), "v0", "get")
}

func TestConfigEvalGet(t *testing.T) {
	c := New(tcfg)
	c.Update("x", "y", "true")
	IsTrue(t, c.GetBool("x", "y"), "get true bool")

	c.Update("x", "y", "false")
	IsFalse(t, c.GetBool("x", "y"), "get false bool")

	c.Update("x", "y", "1.2")
	IsEqual(t, c.GetFloat32("x", "y"), float32(1.2), "get float32")

	c.Update("x", "y", "1.2")
	IsEqual(t, c.GetFloat64("x", "y"), float64(1.2), "get float64")

	c.Update("x", "y", "-9")
	IsEqual(t, c.GetInt("x", "y"), -9, "get int")

	c.Update("x", "y", "-9")
	IsEqual(t, c.GetInt8("x", "y"), int8(-9), "get int8")

	c.Update("x", "y", "-9")
	IsEqual(t, c.GetInt16("x", "y"), int16(-9), "get int16")

	c.Update("x", "y", "-9")
	IsEqual(t, c.GetInt32("x", "y"), int32(-9), "get int32")

	c.Update("x", "y", "-9")
	IsEqual(t, c.GetInt64("x", "y"), int64(-9), "get int64")

	c.Update("x", "y", "9")
	IsEqual(t, c.GetUint("x", "y"), uint(9), "get uint")

	c.Update("x", "y", "9")
	IsEqual(t, c.GetUint8("x", "y"), uint8(9), "get uint8")

	c.Update("x", "y", "9")
	IsEqual(t, c.GetUint16("x", "y"), uint16(9), "get uint16")

	c.Update("x", "y", "9")
	IsEqual(t, c.GetUint32("x", "y"), uint32(9), "get uint32")

	c.Update("x", "y", "9")
	IsEqual(t, c.GetUint64("x", "y"), uint64(9), "get uint64")
}

func TestConfigGetDefault(t *testing.T) {
	c := New(tcfg)
	c.Update("default", "a", "defa")
	c.Update("x", "b", "xb")
	IsEqual(t, c.Get("x", "b"), "xb", "get x")
	IsEqual(t, c.Get("x", "a"), "defa", "get default")
	IsFalse(t, c.HasOption("x", "a"), "has default option")
	IsTrue(t, c.HasOption("x", "b"), "xb option")
}
