// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"

	. "github.com/jrmsdev/gojc/testing/check"
)

var jsonCfg = []byte(`{
	"default": {
		"testing": "ok"
	}
}`)

func TestConfigMapJSON(t *testing.T) {
	c := New(nil)
	c.MapJSON(jsonCfg)
	IsTrue(t, c.HasSection("default"), "default section")
	IsFalse(t, c.HasSection("invalid"), "invalid section")
	IsTrue(t, c.HasOption("default", "testing"), "default option")
	IsFalse(t, c.HasOption("default", "invalid"), "invalid option")
	IsEqual(t, c.Get("default", "testing"), "ok", "json map")
}

func TestConfigReadJSON(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/config.json")
	Fatal(t, IsNil(t, err, "json read error"))
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "json get raw")
	IsEqual(t, c.Get("testing", "opt"), "ok", "json get")
}

func TestConfigRead(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/config.ini")
	Fatal(t, IsNil(t, err, "read error"))
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "get raw")
	IsEqual(t, c.GetRaw("testing", "opt0"), "val0", "get raw opt0")
	IsEqual(t, c.GetRaw("testing", "opt1"), "${default:testing}", "get raw opt1")
	IsEqual(t, c.Get("testing", "opt1"), "ok", "get opt1")
	IsEqual(t, len(c.Sections()), 1, "sections list")
	IsEqual(t, len(c.Options("default")), 1, "default options list")
	IsEqual(t, len(c.Options("testing")), 2, "testing options list")
}

func TestConfigEvalRead(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/eval.ini")
	Fatal(t, IsNil(t, err, "read error"))
	IsTrue(t, c.GetBool("testing", "boolTrue"), "get bool true")
	IsFalse(t, c.GetBool("testing", "boolFalse"), "get bool false")
	IsEqual(t, c.GetFloat32("testing", "float32"), float32(9.9), "get float32")
	IsEqual(t, c.GetFloat64("testing", "float64"), float64(9.9), "get float64")
	IsEqual(t, c.GetInt("testing", "int"), -9, "get int")
	IsEqual(t, c.GetInt8("testing", "int8"), int8(-9), "get int8")
	IsEqual(t, c.GetInt16("testing", "int16"), int16(-9), "get int16")
	IsEqual(t, c.GetInt32("testing", "int32"), int32(-9), "get int32")
	IsEqual(t, c.GetInt64("testing", "int64"), int64(-9), "get int64")
}

func TestConfigReadEmptyFile(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/empty.ini")
	Fatal(t, IsNil(t, err, "read error"))
	IsEqual(t, len(c.Sections()), 0, "sections list")
}

func TestConfigSections(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/sections.ini")
	Fatal(t, IsNil(t, err, "read error"))
	IsEqual(t, len(c.Sections()), 2, "sections list")
	IsTrue(t, c.HasSection("s0"), "section s0")
	IsFalse(t, c.HasOption("s0", "opt"), "s0 option")
	IsTrue(t, c.HasSection("s1"), "section s1")
	IsTrue(t, c.HasOption("s1", "opt"), "s1 option")
}

func TestConfigSectionsError(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/sections-error.ini")
	Fatal(t, NotNil(t, err, "read error"))
	IsEqual(t, err.Error(), "config parse, line 3: '[s1] [s2]'", "error message")
}
