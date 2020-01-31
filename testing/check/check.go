// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package check

import (
	//~ "errors"
	"regexp"
	"testing"
)

func Fatal(t *testing.T, ok bool) {
	if !ok {
		t.FailNow()
	}
}

func IsNil(t *testing.T, got interface{}, errmsg string) bool {
	t.Helper()
	if got != nil {
		t.Errorf("%s is not nil: (%T)%v", errmsg, got, got)
		return false
	}
	return true
}

func NotNil(t *testing.T, got interface{}, errmsg string) bool {
	t.Helper()
	if got == nil {
		t.Errorf("%s is nil: (%T)%v", errmsg, got, got)
		return false
	}
	return true
}

func IsTrue(t *testing.T, got bool, errmsg string) bool {
	t.Helper()
	if !got {
		t.Errorf("%s: is %v", errmsg, got)
		return false
	}
	return true
}

func IsFalse(t *testing.T, got bool, errmsg string) bool {
	t.Helper()
	if got {
		t.Errorf("%s: is %v", errmsg, got)
		return false
	}
	return true
}

func IsEqual(t *testing.T, got, expect interface{}, errmsg string) bool {
	t.Helper()
	if got != expect {
		t.Errorf("%s: '(%T)%v' != '(%T)%v'", errmsg, got, got, expect, expect)
		return false
	}
	return true
}

func NotEqual(t *testing.T, got, expect interface{}, errmsg string) bool {
	t.Helper()
	if got == expect {
		t.Errorf("%s: '(%T)%v' == '(%T)%v'", errmsg, got, got, expect, expect)
		return false
	}
	return true
}

func Match(t *testing.T, pat, s, errmsg string) bool {
	t.Helper()
	m, err := regexp.MatchString(pat, s)
	if err != nil {
		t.Fatalf("ERROR %s: %s", errmsg, err)
		return false
	}
	if !m {
		t.Errorf("%s: '%s' mismatch '%s'", errmsg, pat, s)
		return false
	}
	return true
}

func NotMatch(t *testing.T, pat, s, errmsg string) bool {
	t.Helper()
	m, err := regexp.MatchString(pat, s)
	if err != nil {
		t.Fatalf("ERROR %s: %s", errmsg, err)
		return false
	}
	if m {
		t.Errorf("%s: '%s' match '%s'", errmsg, pat, s)
		return false
	}
	return true
}
