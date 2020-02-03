// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"log"

	gfmt "fmt"

	"github.com/jrmsdev/gojc/errors"
)

var ErrInvalidLevel = errors.New("invalid logger level")
var flagsDefault int = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile

const (
	PANIC = iota
	FATAL
	OFF
	ERROR
	WARN
	MSG
	INFO
	DEBUG
)

var levelTag = map[int]string{
	PANIC: "",
	FATAL: "",
	OFF:   "",
	ERROR: "[E] ",
	WARN:  "[W] ",
	MSG:   "",
	INFO:  "[I] ",
	DEBUG: "[D] ",
}

type Logger interface {
	SetLevel(lvl int) error
	Colors() bool
	SetColors(enable bool)
	Print(lvl int, args ...interface{})
	Printf(lvl int, fmt string, args ...interface{})
}

type L struct {
	depth   int
	lvl     int
	colored bool
}

func New(lvl int, colored bool) Logger {
	log.SetFlags(flagsDefault)
	return &L{3, lvl, colored}
}

func (l *L) SetLevel(lvl int) error {
	if lvl < OFF || lvl > DEBUG {
		return ErrInvalidLevel.Format("%d", lvl)
	}
	l.lvl = lvl
	return nil
}

func (l *L) Print(lvl int, args ...interface{}) {
	if lvl <= l.lvl {
		msg := gfmt.Sprint(args...)
		if l.colored {
			clr := levelColor[lvl]
			log.Output(l.depth, clr+msg+reset)
		} else {
			tag := levelTag[lvl]
			log.Output(l.depth, tag+msg)
		}
	}
}

func (l *L) Printf(lvl int, fmt string, args ...interface{}) {
	if lvl <= l.lvl {
		msg := gfmt.Sprintf(fmt, args...)
		if l.colored {
			clr := levelColor[lvl]
			log.Output(l.depth, clr+msg+reset)
		} else {
			tag := levelTag[lvl]
			log.Output(l.depth, tag+msg)
		}
	}
}
