// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package errors

import (
	"testing"

	gerr "errors"

	. "github.com/jrmsdev/gojc/testing/check"
)

func TestError(t *testing.T) {
	e := New("testing")
	NotNil(t, e.id, "e.id")
	IsEqual(t, e.name, "testing", "e.name")
	IsEqual(t, e.msg, "", "e.msg")
	IsEqual(t, e, e, "e")
	x := New("testing")
	NotEqual(t, x, e, "x")
	IsEqual(t, x.Error(), "testing", "x.Error")
}

func TestFormat(t *testing.T) {
	e := New("t0")
	IsEqual(t, e.Error(), "t0", "e.Error")
	r := e.Format("%s", "t1")
	IsEqual(t, e.Error(), "t0", "e.Error")
	IsEqual(t, r.Error(), "t0: t1", "r.Error")
}

func TestIsError(t *testing.T) {
	e := New("t0")
	IsEqual(t, e, e, "e")
	r := e.Format("%s", "t1")
	NotEqual(t, r, e, "r")
	IsTrue(t, e.Is(r), "e.Is(r)")
	x := gerr.New("testing")
	IsFalse(t, e.Is(x), "e.Is(x)")
}
