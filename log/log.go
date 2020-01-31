// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"github.com/jrmsdev/gojc/log/internal/logger"
)

var l *logger.Logger

func init() {
	l, _ = logger.New("off")
}

func SetLevel(level string) error {
	return l.SetLevel(level)
}

func Error(args ...interface{}) {
	l.Error(l.Format("error", args...))
}

func Errorf(format string, args ...interface{}) {
	l.Error(l.Formatf("error", format, args...))
}

func Warn(args ...interface{}) {
	l.Warn(l.Format("warn", args...))
}

func Warnf(format string, args ...interface{}) {
	l.Warn(l.Formatf("warn", format, args...))
}

func Print(args ...interface{}) {
	l.Print(l.Format("msg", args...))
}

func Printf(format string, args ...interface{}) {
	l.Print(l.Formatf("msg", format, args...))
}

func Info(args ...interface{}) {
	l.Info(l.Format("info", args...))
}

func Infof(format string, args ...interface{}) {
	l.Info(l.Formatf("info", format, args...))
}

func Debug(args ...interface{}) {
	l.Debug(l.Format("debug", args...))
}

func Debugf(format string, args ...interface{}) {
	l.Debug(l.Formatf("debug", format, args...))
}

func BenchIf(lvl string, args ...interface{}) {
	if lvl == "error" {
	} else if lvl == "info" {
	} else if lvl == "debug" {
		l.Debug(l.Format("debug", args...))
	}
}

func BenchSwitch(lvl string, args ...interface{}) {
	switch lvl {
		case "error":
			l.Error(l.Format("error", args...))

		case "info":
			l.Info(l.Format("info", args...))

		case "debug":
			l.Debug(l.Format("debug", args...))

		default:
	}
}
