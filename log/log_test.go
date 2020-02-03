// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log_test

import (
	"bytes"
	"strings"
	"testing"

	glog "log"

	"github.com/jrmsdev/gojc/log"

	. "github.com/jrmsdev/gojc/testing/check"
)

var buf *bytes.Buffer

func init() {
	buf = new(bytes.Buffer)
}

func check(t *testing.T, match string) {
	t.Helper()
	got := strings.TrimSpace(buf.String())
	buf.Reset()
	//~ t.Log(got)
	Match(t, match, got, "check")
}

func TestLog(t *testing.T) {
	oldw := glog.Writer()
	glog.SetOutput(buf)
	defer glog.SetOutput(oldw)

	log.Print("t0")
	check(t, "")

	log.SetLevel("debug")

	log.Error("e1")
	check(t, "e1")

	log.Errorf("%s%d", "e", 2)
	check(t, "e2")

	log.Warn("w1")
	check(t, "w1")

	log.Warnf("%s%d", "w", 2)
	check(t, "w2")

	log.Print("t1")
	check(t, "t1")

	log.Printf("%s", "t2")
	check(t, "t2")

	log.Info("i1")
	check(t, "i1")

	log.Infof("%s%d", "i", 2)
	check(t, "\\[I] i2")

	log.Debug("d1")
	check(t, ` \[D] d1$`)

	log.Debugf("%s%d", "d", 2)
	check(t, ` \[D] d2$`)

	log.Debug("testing")
	check(t, `^\d\d\d\d/\d\d/\d\d \d\d:\d\d:\d\d\.\d\d\d\d\d\d .*log_test\.go:\d\d: \[D] testing$`)
}

func TestInit(t *testing.T) {
	err := log.SetLevel("debug")
	if err != nil {
		t.Error(err)
	}
}

func TestInitFail(t *testing.T) {
	err := log.SetLevel("fail")
	if err == nil {
		t.Error("no error")
	}
}

func TestPanic(t *testing.T) {
	oldw := glog.Writer()
	glog.SetOutput(buf)
	defer glog.SetOutput(oldw)
	defer func() {
		if err := recover(); err != nil {
			check(t, ` \[PANIC] testing$`)
			IsEqual(t, err, "testing", "error message")
		} else {
			t.Error("panic was not called")
		}
	}()
	log.Panic("testing")
}

func TestPanicf(t *testing.T) {
	oldw := glog.Writer()
	glog.SetOutput(buf)
	defer glog.SetOutput(oldw)
	defer func() {
		if err := recover(); err != nil {
			check(t, ` \[PANIC] testing$`)
			IsEqual(t, err, "testing", "error message")
		} else {
			t.Error("panic was not called")
		}
	}()
	log.Panicf("testing")
}
