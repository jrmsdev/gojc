// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

/*
Package log implementes higher-level features for an application logger.

Messages are printed using standard's Go log package so the lib can nicely work
with other's application code.

By default only error, warning and feedback messages are printed, but other
levels of messages can be set as follows:

	* off: no messages at all, not even error messages.
	* error: error messages only.
	* warn: warning + error messages.
	* msg: warning + error + feedback (Print and Printf) messages.
	* info: warning + error + msg + extra info messages.
	* debug: all previous messages plus debug info.

Other useful levels are set:

	* default: error, warning and msg.
	* quiet: errors and warnings only.

If the application is running on a tty (to be honest, if os.Stderr is a tty)
messages are colored. Otherwise messages are time stamped plain lines of text.
*/
package log

import (
	"os"

	"github.com/jrmsdev/gojc/errors"
	"github.com/jrmsdev/gojc/log/internal/logger"
)

// ErrInvalidLevel is returned by SetLevel if an invalid level name is provided.
var ErrInvalidLevel = errors.New("invalid log level")

var l logger.Logger

var level = map[string]int{
	"default": logger.MSG,
	"quiet":   logger.WARN,
	"off":     logger.OFF,
	"error":   logger.ERROR,
	"warn":    logger.WARN,
	"msg":     logger.MSG,
	"info":    logger.INFO,
	"debug":   logger.DEBUG,
}

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
	colored := istty(os.Stderr)
	l = logger.New(logger.OFF, colored)
}

// SetLevel sets the logger level of messages.
func SetLevel(lvl string) error {
	if n, ok := level[lvl]; !ok {
		return ErrInvalidLevel.Format("%s (%d)", lvl, n)
	} else {
		return l.SetLevel(n)
	}
}

// Error prints an error messages.
func Error(args ...interface{}) {
	l.Print(logger.ERROR, args...)
}

// Errorf prints a formatted error messages.
func Errorf(fmt string, args ...interface{}) {
	l.Printf(logger.ERROR, fmt, args...)
}

// Warn prints a warning message.
func Warn(args ...interface{}) {
	l.Print(logger.WARN, args...)
}

// Warnf prints a formatted warning message.
func Warnf(fmt string, args ...interface{}) {
	l.Printf(logger.WARN, fmt, args...)
}

// Print prints a "feedback" message.
func Print(args ...interface{}) {
	l.Print(logger.MSG, args...)
}

// Printf prints a formatted "feedback" message.
func Printf(fmt string, args ...interface{}) {
	l.Printf(logger.MSG, fmt, args...)
}

// Info prints an info message.
func Info(args ...interface{}) {
	l.Print(logger.INFO, args...)
}

// Infof prints a formatted info message.
func Infof(fmt string, args ...interface{}) {
	l.Printf(logger.INFO, fmt, args...)
}

// Debug prints a debug message.
func Debug(args ...interface{}) {
	l.Print(logger.DEBUG, args...)
}

// Debugf prints a formatted debug message.
func Debugf(fmt string, args ...interface{}) {
	l.Printf(logger.DEBUG, fmt, args...)
}
