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
}

func TestConfigEvalRead(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/eval.ini")
	Fatal(t, IsNil(t, err, "read error"))
	IsEqual(t, c.GetInt("testing", "int"), -9, "get int")
}

func TestConfigReadEmptyFile(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/empty.ini")
	Fatal(t, IsNil(t, err, "read error"))
	// TODO: check c.Sections() is empty list (default section not included in that list).
}
