// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

/*
Package errors implement some methods around Go error interface for easier
managing of errors.

Use an struct for error instances comparing and error message formating without
having to change error's identity/message.

Each new created error instance needs to have a name assigned, an unique id is
generated then so instances of the same object, as the ones returned by
e.Format, can be compared later using the e.Is method. Check the examples.
*/
package errors

import (
	gfmt "fmt"
)

// Error interface around the standard error interface.
type Error interface {
	error
	Format(fmt string, args ...interface{}) error
}

// E implements the Error interface.
type E struct {
	id   string
	name string
	msg  string
}

// New creates a new instance of E.
// Assigns a name to the error instance and generates its unique identifier.
func New(name string) *E {
	e := new(E)
	e.id = gfmt.Sprintf("%p:%q", e, name)
	e.name = name
	e.msg = ""
	return e
}

// Error returns the error's message string, as required by the error interface.
func (e *E) Error() string {
	if e.msg != "" {
		return gfmt.Sprintf("%s: %s", e.name, e.msg)
	}
	return e.name
}

// Format creates a new instance of E with the same identity but with a new
// error message.
//
// It returns an error interface implementation instead of an E pointer so it
// can be used in situations like:
//	var ErrExample = errors.New("example")
//	...
//	...
//	func DoSomething() error {
//		err := CheckAnotherStuff()
//		if err != nil {
//			return ErrExample.Format("%s", err)
//		}
//	}
func (e *E) Format(fmt string, args ...interface{}) error {
	err := new(E)
	err.id = e.id
	err.name = e.name
	err.msg = gfmt.Sprintf(fmt, args...)
	return err
}

// Is, shorcut method to call errors.Is(e, err).
func (e *E) Is(err error) bool {
	return Is(e, err)
}

// Is checks if err "is a" *E error. So, it checks if it was generated from the
// same E's instance (using Format).
func Is(e *E, err error) bool {
	x, ok := err.(*E)
	if !ok {
		return false
	}
	return e.id == x.id
}
