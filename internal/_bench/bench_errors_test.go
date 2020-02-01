// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package bench

import (
	"testing"

	"github.com/jrmsdev/gojc/errors"
)

func BenchmarkErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		errors.New("testing" + string(i))
	}
}
