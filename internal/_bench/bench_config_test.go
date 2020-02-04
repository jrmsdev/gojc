// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package bench

import (
	"testing"

	"github.com/jrmsdev/gojc/config"
)

var bcfg = config.Cfg{
	"default": config.Option{
		"testing": "ok",
		"eval": "${testing}",
	},
}

func BenchmarkConfigMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		config.Map(bcfg)
	}
}

func BenchmarkConfigGetRaw(b *testing.B) {
	c := config.Map(bcfg)
	for i := 0; i < b.N; i++ {
		c.GetRaw("default", "testing")
	}
}

func BenchmarkConfigEval(b *testing.B) {
	c := config.Map(bcfg)
	for i := 0; i < b.N; i++ {
		c.Eval("default", "${testing}")
	}
}

func BenchmarkConfigGet(b *testing.B) {
	c := config.Map(bcfg)
	for i := 0; i < b.N; i++ {
		c.Get("default", "eval")
	}
}
