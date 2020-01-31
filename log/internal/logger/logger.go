// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"errors"
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

func New(lvl string) (Logger, error) {
	if lvl == OFF {
		return &OffLogger{base}, nil
	} else if lvl == ERROR {
		return &ErrorLogger{base}, nil
	} else if lvl == WARN {
		return &WarnLogger{base}, nil
	} else if lvl == MSG {
		return &MsgLogger{base}, nil
	} else if lvl == INFO {
		return &InfoLogger{base}, nil
	} else if lvl == DEBUG {
		return &DebugLogger{base}, nil
	}
	return nil, errors.New("invalid log level: " + lvl)
}
