// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

var colored  bool
var flagsDefault int = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile

func istty(fh *os.File) bool {
	if st, err := fh.Stat(); err == nil {
		m := st.Mode()
		if m&os.ModeDevice != 0 && m&os.ModeCharDevice != 0 {
			return true
		}
	}
	return false
}

func init() {
	log.SetFlags(flagsDefault)
	colored = false
	if istty(os.Stderr) {
		colored = true
	}
}

func off(msg string) {}

func out(msg string) {
	log.Output(3, msg)
}

type Logger struct {
	Format  func(lvl string, args ...interface{}) string
	Formatf func(lvl string, fmt string, args ...interface{}) string
	Error   func(msg string)
	Warn    func(msg string)
	Print   func(msg string)
	Info    func(msg string)
	Debug   func(msg string)

	m *sync.Mutex
}

func New(level string) (*Logger, error) {
	l := new(Logger)
	l.m = new(sync.Mutex)
	l.Format = mfmt
	l.Formatf = mfmtf
	l.Error = off
	l.Warn = off
	l.Print = off
	l.Info = off
	l.Debug = off
	return l, l.SetLevel(level)
}

func (l *Logger) SetLevel(level string) error {
	l.m.Lock()
	defer l.m.Unlock()
	if level == "off" {
		l.Error = off
		l.Warn = off
		l.Print = off
		l.Info = off
		l.Debug = off
	} else if level == "error" {
		l.Error = out
		l.Warn = off
		l.Print = off
		l.Info = off
		l.Debug = off
	} else if level == "warn" {
		l.Error = out
		l.Warn = out
		l.Print = off
		l.Info = off
		l.Debug = off
	} else if (level == "msg" || level == "default") {
		l.Error = out
		l.Warn = out
		l.Print = out
		l.Info = off
		l.Debug = off
	} else if level == "info" {
		l.Error = out
		l.Warn = out
		l.Print = out
		l.Info = out
		l.Debug = off
	} else if level == "debug" {
		l.Error = out
		l.Warn = out
		l.Print = out
		l.Info = out
		l.Debug = out
	} else {
		return errors.New("invalid log level: " + level)
	}
	if colored {
		l.Format = color
		l.Formatf = colorf
	} else {
		l.Format = mfmt
		l.Formatf = mfmtf
	}
	return nil
}

var (
	white = "\033[0;0m"
	cyan = "\033[0;36m"
	//~ blue = "\033[0;34m"
	red    = "\033[0;31m"
	yellow = "\033[0;33m"
	green  = "\033[0;32m"
	grey   = "\033[1;30m"
	reset  = "\033[0m"
)

var lvlTag = map[string]string{
	"error": "[E] ",
	"warn":  "[W] ",
	"msg":   "",
	"info":   "[I] ",
	"debug": "[D] ",
}

var lvlColor = map[string]string{
	"error": red,
	"warn":  yellow,
	"msg":   white,
	"info":  cyan,
	"debug": green,
}

func mfmt(lvl string, args ...interface{}) string {
	tag := lvlTag[lvl]
	return fmt.Sprintf("%s%s", tag, fmt.Sprint(args...))
}

func mfmtf(lvl string, format string, args ...interface{}) string {
	tag := lvlTag[lvl]
	return fmt.Sprintf("%s%s", tag, fmt.Sprintf(format, args...))
}

func color(lvl string, args ...interface{}) string {
	clr := lvlColor[lvl]
	return fmt.Sprint(clr, fmt.Sprint(args...), reset)
}

func colorf(lvl string, format string, args ...interface{}) string {
	clr := lvlColor[lvl]
	return fmt.Sprint(clr, fmt.Sprintf(format, args...), reset)
}
