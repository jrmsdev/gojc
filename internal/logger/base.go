// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"fmt"
	"log"
)

type BaseLogger struct {
	depth int
}

func newBaseLogger() *BaseLogger {
	return &BaseLogger{4}
}

func (l *BaseLogger) SetFlags(flags string) {
	// TODO: implement SetFlags
}

func (l *BaseLogger) err(args ...interface{}) {
	log.Output(l.depth, "[E] "+fmt.Sprint(args...))
}

func (l *BaseLogger) errf(f string, args ...interface{}) {
	log.Output(l.depth, "[E] "+fmt.Sprintf(f, args...))
}

func (l *BaseLogger) wrn(args ...interface{}) {
	log.Output(l.depth, "[W] "+fmt.Sprint(args...))
}

func (l *BaseLogger) wrnf(f string, args ...interface{}) {
	log.Output(l.depth, "[W] "+fmt.Sprintf(f, args...))
}

func (l *BaseLogger) prt(args ...interface{}) {
	log.Output(l.depth, fmt.Sprint(args...))
}

func (l *BaseLogger) prtf(f string, args ...interface{}) {
	log.Output(l.depth, fmt.Sprintf(f, args...))
}

func (l *BaseLogger) inf(args ...interface{}) {
	log.Output(l.depth, "[I] "+fmt.Sprint(args...))
}

func (l *BaseLogger) inff(f string, args ...interface{}) {
	log.Output(l.depth, "[I] "+fmt.Sprintf(f, args...))
}

func (l *BaseLogger) dbg(args ...interface{}) {
	log.Output(l.depth, "[D] "+fmt.Sprint(args...))
}

func (l *BaseLogger) dbgf(f string, args ...interface{}) {
	log.Output(l.depth, "[D] "+fmt.Sprintf(f, args...))
}
