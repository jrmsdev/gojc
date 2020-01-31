// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

const OFF string = "OFF"

type OffLogger struct {
	*BaseLogger
}

func (l *OffLogger) Error(args ...interface{})              {}
func (l *OffLogger) Errorf(fmt string, args ...interface{}) {}

func (l *OffLogger) Warn(args ...interface{})              {}
func (l *OffLogger) Warnf(fmt string, args ...interface{}) {}

func (l *OffLogger) Print(args ...interface{})              {}
func (l *OffLogger) Printf(fmt string, args ...interface{}) {}

func (l *OffLogger) Info(args ...interface{})              {}
func (l *OffLogger) Infof(fmt string, args ...interface{}) {}

func (l *OffLogger) Debug(args ...interface{})              {}
func (l *OffLogger) Debugf(fmt string, args ...interface{}) {}

func (l *OffLogger) SetFlags(flags string) {}
