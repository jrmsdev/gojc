// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

const ERROR string = "ERROR"

type ErrorLogger struct {
	*BaseLogger
}

func (l *ErrorLogger) Error(args ...interface{}) {
	l.err(args...)
}

func (l *ErrorLogger) Errorf(fmt string, args ...interface{}) {
	l.errf(fmt, args...)
}

func (l *ErrorLogger) Warn(args ...interface{})              {}
func (l *ErrorLogger) Warnf(fmt string, args ...interface{}) {}

func (l *ErrorLogger) Print(args ...interface{})              {}
func (l *ErrorLogger) Printf(fmt string, args ...interface{}) {}

func (l *ErrorLogger) Info(args ...interface{})              {}
func (l *ErrorLogger) Infof(fmt string, args ...interface{}) {}

func (l *ErrorLogger) Debug(args ...interface{})              {}
func (l *ErrorLogger) Debugf(fmt string, args ...interface{}) {}
