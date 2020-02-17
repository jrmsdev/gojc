// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"

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
	if match != "" {
		Match(t, match, got, "check")
	} else {
		IsEqual(t, got, match, "check")
	}
}

func TestLogger(t *testing.T) {
	l := New(OFF, false)
	IsFalse(t, l.colored, "l.colored")
	IsEqual(t, levelTag[PANIC], "", "tag")
	IsEqual(t, levelTag[FATAL], "", "tag")
}

func TestLoggerOFF(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	IsEqual(t, levelTag[OFF], "", "tag")
	l := New(OFF, false)
	l.Print(MSG, "testing")
	check(t, "")
	l.Printf(MSG, "testing")
	check(t, "")
}

func TestLoggerERROR(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	IsEqual(t, levelTag[ERROR], "[E] ", "tag")
	l := New(ERROR, false)
	l.Print(PANIC, "testing")
	check(t, ` testing$`)
	l.Print(FATAL, "testing")
	check(t, ` testing$`)
	l.Print(ERROR, "testing")
	check(t, ` \[E] testing$`)
	l.Printf(WARN, "testing")
	check(t, "")
	l.Printf(MSG, "testing")
	check(t, "")
	l.Printf(INFO, "testing")
	check(t, "")
	l.Printf(DEBUG, "testing")
	check(t, "")
}

func TestLoggerWARN(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	IsEqual(t, levelTag[WARN], "[W] ", "tag")
	l := New(WARN, false)
	l.Print(PANIC, "testing")
	check(t, ` testing$`)
	l.Print(FATAL, "testing")
	check(t, ` testing$`)
	l.Print(ERROR, "testing")
	check(t, ` \[E] testing$`)
	l.Printf(WARN, "testing")
	check(t, ` \[W] testing$`)
	l.Printf(MSG, "testing")
	check(t, "")
	l.Printf(INFO, "testing")
	check(t, "")
	l.Printf(DEBUG, "testing")
	check(t, "")
}

func TestLoggerMSG(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	IsEqual(t, levelTag[MSG], "", "tag")
	l := New(MSG, false)
	l.Print(PANIC, "testing")
	check(t, ` testing$`)
	l.Print(FATAL, "testing")
	check(t, ` testing$`)
	l.Print(ERROR, "testing")
	check(t, ` \[E] testing$`)
	l.Printf(WARN, "testing")
	check(t, ` \[W] testing$`)
	l.Printf(MSG, "testing")
	check(t, " testing$")
	l.Printf(INFO, "testing")
	check(t, "")
	l.Printf(DEBUG, "testing")
	check(t, "")
}

func TestLoggerINFO(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	IsEqual(t, levelTag[INFO], "[I] ", "tag")
	l := New(INFO, false)
	l.Print(PANIC, "testing")
	check(t, ` testing$`)
	l.Print(FATAL, "testing")
	check(t, ` testing$`)
	l.Print(ERROR, "testing")
	check(t, ` \[E] testing$`)
	l.Printf(WARN, "testing")
	check(t, ` \[W] testing$`)
	l.Printf(MSG, "testing")
	check(t, " testing$")
	l.Printf(INFO, "testing")
	check(t, ` \[I] testing$`)
	l.Printf(DEBUG, "testing")
	check(t, "")
}

func TestLoggerDEBUG(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	IsEqual(t, levelTag[DEBUG], "[D] ", "tag")
	l := New(DEBUG, false)
	l.Print(PANIC, "testing")
	check(t, ` testing$`)
	l.Print(FATAL, "testing")
	check(t, ` testing$`)
	l.Print(ERROR, "testing")
	check(t, ` \[E] testing$`)
	l.Printf(WARN, "testing")
	check(t, ` \[W] testing$`)
	l.Printf(MSG, "testing")
	check(t, " testing$")
	l.Printf(INFO, "testing")
	check(t, ` \[I] testing$`)
	l.Printf(DEBUG, "testing")
	check(t, ` \[D] testing$`)
}

func TestLoggerSetLevel(t *testing.T) {
	l := New(OFF, false)
	IsEqual(t, l.lvl, OFF, "off level")
	l.SetLevel(DEBUG)
	IsEqual(t, l.lvl, DEBUG, "debug level")
	if err := l.SetLevel(PANIC); err == nil {
		t.Error("set level panic did not failed")
	} else {
		IsTrue(t, ErrInvalidLevel.Is(err), "invalid error")
	}
	l.SetLevel(MSG)
	IsEqual(t, l.lvl, MSG, "msg level")
}

func TestLoggerFlags(t *testing.T) {
	l := New(OFF, false)
	IsEqual(t, log.Flags(), flagsDefault, "off flags")
	l.SetLevel(MSG)
	IsEqual(t, log.Flags(), flagsDefault, "msg flags")
	l.SetLevel(DEBUG)
	IsEqual(t, log.Flags(), flagsDebug, "debug flags")
	l.SetColors(true)
	l.SetLevel(MSG)
	IsEqual(t, log.Flags(), flagsColored, "colored flags")
	l.SetLevel(DEBUG)
	IsEqual(t, log.Flags(), flagsColoredDebug, "colored debug flags")
}

func TestLoggerLevelColorFlags(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	l := New(MSG, true)
	IsEqual(t, log.Flags(), flagsDefault, "initial flags")
	l.SetColors(true)
	IsEqual(t, log.Flags(), flagsColored, "colored flags")
	l.SetLevel(DEBUG)
	IsEqual(t, log.Flags(), flagsColoredDebug, "colored debug flags")
	l.SetColors(false)
	IsEqual(t, log.Flags(), flagsDebug, "no color debug flags")
	l.SetLevel(MSG)
	IsEqual(t, log.Flags(), flagsDefault, "default flags")
	l.SetColors(true)
	IsEqual(t, log.Flags(), flagsColored, "colored flags")
}
