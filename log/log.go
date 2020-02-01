// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"os"

	"github.com/jrmsdev/gojc/errors"
	"github.com/jrmsdev/gojc/log/internal/logger"
)

var ErrInvalidLevel = errors.New("invalid log level")

var l logger.Logger

var level = map[string]int{
	"default": logger.MSG,
	"quiet": logger.WARN,
	"off": logger.OFF,
	"error": logger.ERROR,
	"warn": logger.WARN,
	"msg": logger.MSG,
	"info": logger.INFO,
	"debug": logger.DEBUG,
}

func istty(fh *os.File) bool {
	if st, err := fh.Stat(); err == nil {
		m := st.Mode()
		if m&os.ModeDevice != 0 && m&os.ModeCharDevice != 0 {
			return true
		}
	}
	return false
}

func init() {
	colored := istty(os.Stderr)
	l = logger.New(logger.OFF, colored)
}

func SetLevel(lvl string) error {
	if n, ok := level[lvl]; !ok {
		return ErrInvalidLevel.Format("%s (%d)", lvl, n)
	} else {
		return l.SetLevel(n)
	}
}

func Error(args ...interface{}) {
	l.Print(logger.ERROR, args...)
}

func Errorf(fmt string, args ...interface{}) {
	l.Printf(logger.ERROR, fmt, args...)
}

func Warn(args ...interface{}) {
	l.Print(logger.WARN, args...)
}

func Warnf(fmt string, args ...interface{}) {
	l.Printf(logger.WARN, fmt, args...)
}

func Print(args ...interface{}) {
	l.Print(logger.MSG, args...)
}

func Printf(fmt string, args ...interface{}) {
	l.Printf(logger.MSG, fmt, args...)
}

func Info(args ...interface{}) {
	l.Print(logger.INFO, args...)
}

func Infof(fmt string, args ...interface{}) {
	l.Printf(logger.INFO, fmt, args...)
}

func Debug(args ...interface{}) {
	l.Print(logger.DEBUG, args...)
}

func Debugf(fmt string, args ...interface{}) {
	l.Printf(logger.DEBUG, fmt, args...)
}
