// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log_test

import (
	"github.com/jrmsdev/gojc/log"
)

func setupOutput() {
}

func ExampleSetLevel() {
	log.SetLevel("off")
	log.Print(1) // nothing printed
	log.SetLevel("default")
	log.Print(2) // 2 is printed
}
