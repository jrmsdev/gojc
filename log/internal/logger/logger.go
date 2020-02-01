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
	Print(lvl int, args ...interface{})
	Printf(lvl int, fmt string, args ...interface{})
}

type L struct {
	lvl   int
	depth int
	print func(l *L, lvl int, msg string)
}

func New(lvl int, colored bool) Logger {
	log.SetFlags(flagsDefault)
	fn := print
	if colored {
		fn = color
	}
	return &L{lvl, 4, fn}
}

func print(l *L, lvl int, msg string) {
	tag := levelTag[lvl]
	log.Output(l.depth, tag+msg)
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
		l.print(l, lvl, gfmt.Sprint(args...))
	}
}

func (l *L) Printf(lvl int, fmt string, args ...interface{}) {
	if lvl <= l.lvl {
		l.print(l, lvl, gfmt.Sprintf(fmt, args...))
	}
}

// colors

var (
	white = "\033[0;0m"
	cyan  = "\033[0;36m"
	//~ blue = "\033[0;34m"
	red    = "\033[0;31m"
	yellow = "\033[0;33m"
	green  = "\033[0;32m"
	grey   = "\033[1;30m"
	reset  = "\033[0m"
)

var levelColor = map[int]string{
	PANIC: red,
	FATAL: red,
	OFF:   white,
	ERROR: red,
	WARN:  yellow,
	MSG:   white,
	INFO:  cyan,
	DEBUG: green,
}

func color(l *L, lvl int, msg string) {
	clr := levelColor[lvl]
	log.Output(l.depth, clr+msg+reset)
}
