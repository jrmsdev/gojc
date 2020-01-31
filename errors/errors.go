// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package errors

import (
	"crypto/sha256"
	"time"

	gfmt "fmt"
)

type Error interface {
	Error() string
	Format(fmt string, args ...interface{}) error
}

type E struct {
	id   [sha256.Size]byte
	name string
	msg  string
}

func New(name string) *E {
	return &E{
		sha256.Sum256([]byte(time.Now().String())),
		name,
		"",
	}
}

func (e *E) Error() string {
	if e.msg != "" {
		return gfmt.Sprintf("%s: %s", e.name, e.msg)
	}
	return e.name
}

func (e *E) Format(fmt string, args ...interface{}) error {
	return Format(e, fmt, args...)
}

func (e *E) Is(err error) bool {
	return Is(e, err)
}

func Format(e *E, fmt string, args ...interface{}) error {
	return &E{e.id, e.name, gfmt.Sprintf(fmt, args...)}
}

func Is(e *E, err error) bool {
	x, ok := err.(*E)
	if !ok {
		return false
	}
	return e.id == x.id
}
