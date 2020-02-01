// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package bench

import (
	"os"
	"testing"

	golog "log"

	"github.com/jrmsdev/gojc/log"
)

func benchSetup(b *testing.B) *os.File {
	fh, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	golog.SetOutput(fh)
	return fh
}

func BenchmarkOff(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("off")
	for i := 0; i < b.N; i++ {
		log.Debugf("%d", i)
	}
}

func BenchmarkDebug(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("debug")
	for i := 0; i < b.N; i++ {
		log.Debugf("%d", i)
	}
}
