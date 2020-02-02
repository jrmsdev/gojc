// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"bytes"
	"strings"
	"testing"

	glog "log"

	x "github.com/jrmsdev/gojc/testing/check"
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
	x.Match(t, match, got, "check")
}

var exitStatus int
func mockExit(status int) {
	exitStatus = status
}

func TestFatal(t *testing.T) {
	oldw := glog.Writer()
	glog.SetOutput(buf)
	defer glog.SetOutput(oldw)
	oldx := osExit
	osExit = mockExit
	defer func() { osExit = oldx }()
	exitStatus = 0
	Fatal(9, "testing")
	check(t, `: \[FATAL] testing`)
	x.IsEqual(t, exitStatus, 9, "exit status")
}

func TestFatalf(t *testing.T) {
	oldw := glog.Writer()
	glog.SetOutput(buf)
	defer glog.SetOutput(oldw)
	oldx := osExit
	osExit = mockExit
	defer func() { osExit = oldx }()
	exitStatus = 0
	Fatalf(8, "testing")
	check(t, `: \[FATAL] testing`)
	x.IsEqual(t, exitStatus, 8, "exit status")
}
