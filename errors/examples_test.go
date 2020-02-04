// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package errors_test

import (
	"fmt"

	"github.com/jrmsdev/gojc/errors"
)

func ExampleError() {
	// An E struct's instance is created.
	var ErrExample = errors.New("example %d")

	// Format returns an error interface implementation.
	err := ErrExample.Format(1)

	// Check if err `is an` ErrExample error.
	fmt.Println(ErrExample.Is(err))

	// err message is printed using Error() method, as required by the error
	// interface.
	fmt.Println(err)

	// Even if Format was run against ErrExample instance, it wasn't changed.
	// A new instance of an error with the same identity was generated instead,
	// so err `is an` ErrExample error but they both have different messages.
	fmt.Println(ErrExample.Error())

	// That's useful in example when you have a base error but you need to
	// generate a different error message with runtime information from the
	// moment when the error happened.
	flag := 1
	if flag == 1 {
		fmt.Println("ok")
	} else {
		err := ErrExample.Format(flag)
		fmt.Println(err)

		// And later in your code you can check what the err means and do
		// something about it.
		if errors.Is(ErrExample, err) {
			fmt.Println("found")
		} else {
			fmt.Println("not found")
		}
	}

	// Output:
	// true
	// example 1
	// example %d
	// ok
}

func ExampleIs() {
	// An E struct's instance is created.
	var ErrExample = errors.New("example")

	// Format returns an error interface implementation.
	err := ErrExample.Format("%d", 1)

	// Check if err `is an` ErrExample error.
	fmt.Println(ErrExample.Is(err))

	// Output:
	// true
}

func Example_format() {
	// An E struct's instance is created.
	var ErrExample = errors.New("example: %d")

	// Format returns an error interface implementation.
	err := ErrExample.Format(1)

	// Check if err `is an` ErrExample error.
	fmt.Println(err)

	// Output:
	// example: 1
}
