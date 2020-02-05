// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"testing"

	"github.com/jrmsdev/gojc/config"

	. "github.com/jrmsdev/gojc/testing/check"
)

var tcfg = config.Cfg{
	"default": config.Option{
		"testing": "ok",
	},
}

func TestConfigParse(t *testing.T) {
	Map(tcfg)
	IsEqual(t, Get("default", "testing"), "ok", "get map value")

	Update("x", "y", "true")
	IsTrue(t, GetBool("x", "y"), "get true bool")

	Update("x", "y", "false")
	IsFalse(t, GetBool("x", "y"), "get false bool")

	Update("x", "y", "1.2")
	IsEqual(t, GetFloat32("x", "y"), float32(1.2), "get float32")

	Update("x", "y", "1.2")
	IsEqual(t, GetFloat64("x", "y"), float64(1.2), "get float64")

	Update("x", "y", "-9")
	IsEqual(t, GetInt("x", "y"), -9, "get int")

	Update("x", "y", "-9")
	IsEqual(t, GetInt8("x", "y"), int8(-9), "get int8")

	Update("x", "y", "-9")
	IsEqual(t, GetInt16("x", "y"), int16(-9), "get int16")

	Update("x", "y", "-9")
	IsEqual(t, GetInt32("x", "y"), int32(-9), "get int32")

	Update("x", "y", "-9")
	IsEqual(t, GetInt64("x", "y"), int64(-9), "get int64")

	Update("x", "y", "9")
	IsEqual(t, GetUint("x", "y"), uint(9), "get uint")

	Update("x", "y", "9")
	IsEqual(t, GetUint8("x", "y"), uint8(9), "get uint8")

	Update("x", "y", "9")
	IsEqual(t, GetUint16("x", "y"), uint16(9), "get uint16")

	Update("x", "y", "9")
	IsEqual(t, GetUint32("x", "y"), uint32(9), "get uint32")

	Update("x", "y", "9")
	IsEqual(t, GetUint64("x", "y"), uint64(9), "get uint64")
}
