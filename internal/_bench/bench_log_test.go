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

func BenchmarkLogOff(b *testing.B) {
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

func BenchmarkLogError(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.Init()
	for i := 0; i < b.N; i++ {
		log.Errorf("%d", i)
	}
}

func BenchmarkLogWarn(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("warn")
	for i := 0; i < b.N; i++ {
		log.Warnf("%d", i)
	}
}

func BenchmarkLogPrint(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.Init()
	for i := 0; i < b.N; i++ {
		log.Printf("%d", i)
	}
}

func BenchmarkLogInfo(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("info")
	for i := 0; i < b.N; i++ {
		log.Infof("%d", i)
	}
}

func BenchmarkLogInfoPrint(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("info")
	for i := 0; i < b.N; i++ {
		log.Printf("%d", i)
	}
}

func BenchmarkLogDebug(b *testing.B) {
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

func BenchmarkLogColors(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("msg")
	for i := 0; i < b.N; i++ {
		log.Printf("%d", i)
	}
}

func BenchmarkLogColorsDebug(b *testing.B) {
	oldw := golog.Writer()
	devnull := benchSetup(b)
	defer func() {
		devnull.Close()
		golog.SetOutput(oldw)
	}()
	log.SetLevel("debug")
	for i := 0; i < b.N; i++ {
		log.Printf("%d", i)
	}
}
