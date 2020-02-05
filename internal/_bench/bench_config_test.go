// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package bench

import (
	"strconv"
	"testing"

	"github.com/jrmsdev/gojc/config"
)

var bcfg = config.Cfg{
	"default": config.Option{
		"testing": "ok",
		"eval":    "${testing}",
	},
}

func BenchmarkConfigMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		config.New(bcfg)
	}
}

func BenchmarkConfigGetRaw(b *testing.B) {
	c := config.New(bcfg)
	for i := 0; i < b.N; i++ {
		c.GetRaw("default", "testing")
	}
}

func BenchmarkConfigEval(b *testing.B) {
	c := config.New(bcfg)
	for i := 0; i < b.N; i++ {
		c.Eval("default", "${testing}")
	}
}

func BenchmarkConfigGet(b *testing.B) {
	c := config.New(bcfg)
	for i := 0; i < b.N; i++ {
		c.Get("default", "eval")
	}
}

func BenchmarkConfigSet(b *testing.B) {
	c := config.New(bcfg)
	for i := 0; i < b.N; i++ {
		c.Set("default", strconv.Itoa(i), "x")
	}
}

func BenchmarkConfigUpdate(b *testing.B) {
	c := config.New(bcfg)
	c.Set("default", "x", "")
	for i := 0; i < b.N; i++ {
		c.Update("default", "x", strconv.Itoa(i))
	}
}
