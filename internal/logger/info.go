// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

const INFO string = "INFO"

type InfoLogger struct {
	*BaseLogger
}

func (l *InfoLogger) Error(args ...interface{}) {
	l.err(args...)
}

func (l *InfoLogger) Errorf(fmt string, args ...interface{}) {
	l.errf(fmt, args...)
}

func (l *InfoLogger) Warn(args ...interface{}) {
	l.wrn(args...)
}

func (l *InfoLogger) Warnf(fmt string, args ...interface{}) {
	l.wrnf(fmt, args...)
}

func (l *InfoLogger) Print(args ...interface{}) {
	l.prt(args...)
}

func (l *InfoLogger) Printf(fmt string, args ...interface{}) {
	l.prtf(fmt, args...)
}

func (l *InfoLogger) Info(args ...interface{}) {
	l.inf(args...)
}

func (l *InfoLogger) Infof(fmt string, args ...interface{}) {
	l.inff(fmt, args...)
}

func (l *InfoLogger) Debug(args ...interface{})              {}
func (l *InfoLogger) Debugf(fmt string, args ...interface{}) {}
