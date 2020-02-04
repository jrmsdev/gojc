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
	c := Map(evalCfg)
	IsEqual(t, c.Eval("testing"), "testing", "eval")
	IsEqual(t, c.Eval("    testing"), "testing", "trim space")
	IsEqual(t, c.Eval(" \t testing"), "testing", "trim tab and space")
	IsEqual(t, c.Eval("\ttesting\r\n"), "testing", "trim tab and newline")
	IsEqual(t, c.Eval("test\ting"), "test\ting", "no trim tab in the middle")
	IsEqual(t, c.Eval("test\n\ring"), "testing", "trim newline in the middle")

	//~ IsEqual(t, c.Eval("${testing}"), "ok", "expand")
}
