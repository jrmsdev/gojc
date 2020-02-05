// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"

	. "github.com/jrmsdev/gojc/testing/check"
)

var evalCfg = Cfg{
	"default": Option{
		"testing": "ok",
	},
}

func TestConfigEval(t *testing.T) {
	c := New(evalCfg)
	IsEqual(t, c.Eval("default", "testing"), "testing", "eval")
	IsEqual(t, c.Eval("default", "    testing"), "    testing", "trim space")
	IsEqual(t, c.Eval("default", " \t testing"), " \t testing", "trim tab and space")
	IsEqual(t, c.Eval("default", "\ttesting\r\n"), "\ttesting\r\n", "trim tab and newline")
	IsEqual(t, c.Eval("default", "test\ting"), "test\ting", "trim tab in the middle")
	IsEqual(t, c.Eval("default", "test\n\ring"), "test\n\ring", "trim newline in the middle")

	IsEqual(t, c.Eval("default", ""), "", "empty expression")
	IsEqual(t, c.Eval("default", "\t"), "\t", "tab expression")

	IsEqual(t, c.Eval("default", "${testing}"), "ok", "expand")

	IsEqual(t, c.Eval("default", "${default:testing}:${testing}"), "ok:ok", "expand section")
}
