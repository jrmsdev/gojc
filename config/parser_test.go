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

func TestMapJSON(t *testing.T) {
	c := New(nil)
	c.MapJSON(jsonCfg)
	IsTrue(t, c.HasSection("default"), "default section")
	IsFalse(t, c.HasSection("invalid"), "invalid section")
	IsTrue(t, c.HasOption("default", "testing"), "default option")
	IsFalse(t, c.HasOption("default", "invalid"), "invalid option")
	IsEqual(t, c.Get("default", "testing"), "ok", "json map")
}

func TestRead(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/config.ini")
	IsNil(t, err, "read error")
	//~ IsEqual(t, c.GetRaw("default", "testing"), "ok", "get raw")
}

func TestReadJSON(t *testing.T) {
	c := New(nil)
	err := c.Read("testdata/config.json")
	IsNil(t, err, "json read error")
	IsEqual(t, c.GetRaw("default", "testing"), "ok", "json get raw")
}
