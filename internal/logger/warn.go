// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

const WARN string = "WARN"

type WarnLogger struct {
	*BaseLogger
}

func (l *WarnLogger) Error(args ...interface{}) {
	l.err(args...)
}

func (l *WarnLogger) Errorf(fmt string, args ...interface{}) {
	l.errf(fmt, args...)
}

func (l *WarnLogger) Warn(args ...interface{}) {
	l.wrn(args...)
}

func (l *WarnLogger) Warnf(fmt string, args ...interface{}) {
	l.wrnf(fmt, args...)
}

func (l *WarnLogger) Print(args ...interface{})              {}
func (l *WarnLogger) Printf(fmt string, args ...interface{}) {}

func (l *WarnLogger) Info(args ...interface{})              {}
func (l *WarnLogger) Infof(fmt string, args ...interface{}) {}

func (l *WarnLogger) Debug(args ...interface{})              {}
func (l *WarnLogger) Debugf(fmt string, args ...interface{}) {}
