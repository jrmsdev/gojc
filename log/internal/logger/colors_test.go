// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	//~ "fmt"
	"log"
	"strings"
	"testing"

	. "github.com/jrmsdev/gojc/testing/check"
)

func checkColor(t *testing.T, lvl, col string) {
	t.Helper()
	got := strings.TrimSpace(buf.String())
	buf.Reset()
	var s string = ": " + col + "testing" + reset
	if !strings.HasSuffix(got, s) {
		t.Errorf("%s: color not found", lvl)
	}
}

func TestLoggerColors(t *testing.T) {
	oldw := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(oldw)
	IsEqual(t, levelColor[PANIC], red, "color")
	IsEqual(t, levelColor[FATAL], red, "color")
	IsEqual(t, levelColor[OFF], white, "color")
	IsEqual(t, levelColor[ERROR], red, "color")
	IsEqual(t, levelColor[WARN], yellow, "color")
	IsEqual(t, levelColor[MSG], white, "color")
	IsEqual(t, levelColor[INFO], cyan, "color")
	IsEqual(t, levelColor[DEBUG], green, "color")
	l := New(MSG, true)
	l.Print(MSG, "testing")
	checkColor(t, "MSG", levelColor[MSG])
}
