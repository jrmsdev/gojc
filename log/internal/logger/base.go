// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
	"log"
)

type BaseLogger struct {}

func newBaseLogger() *BaseLogger {
	return &BaseLogger{}
}

func (l *BaseLogger) SetFlags(flags string) {
	// TODO: implement SetFlags
}

func (l *BaseLogger) out(prefix, msg string) {
	log.Output(5, fmt.Sprintf("%s%s", prefix, msg))
}

func (l *BaseLogger) err(args ...interface{}) {
	l.out("[E] ", fmt.Sprint(args...))
}

func (l *BaseLogger) errf(f string, args ...interface{}) {
	l.out("[E] ", fmt.Sprintf(f, args...))
}

func (l *BaseLogger) wrn(args ...interface{}) {
	l.out("[W] ", fmt.Sprint(args...))
}

func (l *BaseLogger) wrnf(f string, args ...interface{}) {
	l.out("[W] ", fmt.Sprintf(f, args...))
}

func (l *BaseLogger) prt(args ...interface{}) {
	l.out("", fmt.Sprint(args...))
}

func (l *BaseLogger) prtf(f string, args ...interface{}) {
	l.out("", fmt.Sprintf(f, args...))
}

func (l *BaseLogger) inf(args ...interface{}) {
	l.out("[I] ", fmt.Sprint(args...))
}

func (l *BaseLogger) inff(f string, args ...interface{}) {
	l.out("[I] ", fmt.Sprintf(f, args...))
}

func (l *BaseLogger) dbg(args ...interface{}) {
	l.out("[D] ", fmt.Sprint(args...))
}

func (l *BaseLogger) dbgf(f string, args ...interface{}) {
	l.out("[D] ", fmt.Sprintf(f, args...))
}
