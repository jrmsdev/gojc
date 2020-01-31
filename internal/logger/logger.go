// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"log"
)

type Logger interface {
	Error(args ...interface{})
	Errorf(fmt string, args ...interface{})

	Warn(args ...interface{})
	Warnf(fmt string, args ...interface{})

	Print(args ...interface{})
	Printf(fmt string, args ...interface{})

	Info(args ...interface{})
	Infof(fmt string, args ...interface{})

	Debug(args ...interface{})
	Debugf(fmt string, args ...interface{})

	SetFlags(flags string)
}

var base *BaseLogger
var flagsDefault int = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile

func init() {
	base = newBaseLogger()
	log.SetFlags(flagsDefault)
}

func New(lvl string) Logger {
	if lvl == OFF {
		return &OffLogger{base}
	} else if lvl == ERROR {
		return &ErrorLogger{base}
	} else if lvl == WARN {
		return &WarnLogger{base}
	} else if lvl == MSG {
		return &MsgLogger{base}
	} else if lvl == INFO {
		return &InfoLogger{base}
	} else if lvl == DEBUG {
		return &DebugLogger{base}
	}
	panic("unkown logger: " + lvl)
}
