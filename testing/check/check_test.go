// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package check_test

import (
	"testing"

	. "github.com/jrmsdev/gojc/testing/check"
)

func TestCheck(t *testing.T) {
	Fatal(t, NotNil(t, true, "Fatal"))
	IsNil(t, nil, "IsNil")
	NotNil(t, true, "NotNil")
	IsTrue(t, true, "IsTrue")
	IsFalse(t, false, "IsFalse")
	IsEqual(t, 1, 1, "IsEqual")
	NotEqual(t, 1, 2, "NotEqual")
	Match(t, "^t", "testing", "Match")
	NotMatch(t, "^a", "testing", "NotMatch")
}
