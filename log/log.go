// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"errors"
	"strings"
	"sync"

	"github.com/jrmsdev/gojc/internal/logger"
)

var l logger.Logger
var m *sync.Mutex

func init() {
	m = &sync.Mutex{}
	setLogger(logger.OFF)
}

func setLogger(lvl string) {
	m.Lock()
	defer m.Unlock()
	l = logger.New(strings.ToUpper(lvl))
}

func Init(lvl string) (err error) {
	// TODO: defer recover from possible logger.New panic
	defer func() {
		if e := recover(); e != nil {
			if s, ok := e.(string); ok {
				err = errors.New(s)
			} else {
				panic(e)
			}
		}
	}()
	setLogger(lvl)
	return err
}

func SetFlags(flags string) {
	l.SetFlags(flags)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func Errorf(fmt string, args ...interface{}) {
	l.Errorf(fmt, args...)
}

func Warn(args ...interface{}) {
	l.Warn(args...)
}

func Warnf(fmt string, args ...interface{}) {
	l.Warnf(fmt, args...)
}

func Print(args ...interface{}) {
	l.Print(args...)
}

func Printf(fmt string, args ...interface{}) {
	l.Printf(fmt, args...)
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Infof(fmt string, args ...interface{}) {
	l.Infof(fmt, args...)
}

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Debugf(fmt string, args ...interface{}) {
	l.Debugf(fmt, args...)
}
