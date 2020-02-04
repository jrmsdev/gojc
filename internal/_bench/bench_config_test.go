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
	},
}

func BenchmarkConfigGetRaw(b *testing.B) {
	c := config.Map(bcfg)
	for i := 0; i < b.N; i++ {
		c.GetRaw("default", "testing")
	}
}
