// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

const DEBUG string = "DEBUG"

type DebugLogger struct {
	*BaseLogger
}

func (l *DebugLogger) Error(args ...interface{}) {
	l.err(args...)
}

func (l *DebugLogger) Errorf(fmt string, args ...interface{}) {
	l.errf(fmt, args...)
}

func (l *DebugLogger) Warn(args ...interface{}) {
	l.wrn(args...)
}

func (l *DebugLogger) Warnf(fmt string, args ...interface{}) {
	l.wrnf(fmt, args...)
}

func (l *DebugLogger) Print(args ...interface{}) {
	l.prt(args...)
}

func (l *DebugLogger) Printf(fmt string, args ...interface{}) {
	l.prtf(fmt, args...)
}

func (l *DebugLogger) Info(args ...interface{}) {
	l.inf(args...)
}

func (l *DebugLogger) Infof(fmt string, args ...interface{}) {
	l.inff(fmt, args...)
}

func (l *DebugLogger) Debug(args ...interface{}) {
	l.dbg(args...)
}

func (l *DebugLogger) Debugf(fmt string, args ...interface{}) {
	l.dbgf(fmt, args...)
}
