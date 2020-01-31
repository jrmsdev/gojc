// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package bench

import (
	"os"
	"testing"

	golog "log"

	"github.com/jrmsdev/gojc/log"
	log2 "github.com/jrmsdev/gojc/log2"
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

func BenchmarkLog2Off(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log2.SetLevel("off")
	for i := 0; i < b.N; i++ {
		log2.Debugf("%d", i)
	}
}

func BenchmarkIfOff(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("off")
	for i := 0; i < b.N; i++ {
		log.BenchIf("off", "%d", i)
	}
}

func BenchmarkSwitchOff(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("off")
	for i := 0; i < b.N; i++ {
		log.BenchSwitch("off", "%d", i)
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

func BenchmarkLog2Debug(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log2.SetLevel("debug")
	for i := 0; i < b.N; i++ {
		log2.Debugf("%d", i)
	}
}

func BenchmarkIfDebug(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("debug")
	for i := 0; i < b.N; i++ {
		log.BenchIf("debug", "%d", i)
	}
}

func BenchmarkSwitchDebug(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("debug")
	for i := 0; i < b.N; i++ {
		log.BenchSwitch("debug", "%d", i)
	}
}
