// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

const MSG string = "MSG"

type MsgLogger struct {
	*BaseLogger
}

func (l *MsgLogger) Error(args ...interface{}) {
	l.err(args...)
}

func (l *MsgLogger) Errorf(fmt string, args ...interface{}) {
	l.errf(fmt, args...)
}

func (l *MsgLogger) Warn(args ...interface{}) {
	l.wrn(args...)
}

func (l *MsgLogger) Warnf(fmt string, args ...interface{}) {
	l.wrnf(fmt, args...)
}

func (l *MsgLogger) Print(args ...interface{}) {
	l.prt(args...)
}

func (l *MsgLogger) Printf(fmt string, args ...interface{}) {
	l.prtf(fmt, args...)
}

func (l *MsgLogger) Info(args ...interface{})              {}
func (l *MsgLogger) Infof(fmt string, args ...interface{}) {}

func (l *MsgLogger) Debug(args ...interface{})              {}
func (l *MsgLogger) Debugf(fmt string, args ...interface{}) {}
