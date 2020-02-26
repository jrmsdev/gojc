// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/jrmsdev/gojc/errors"
)

var (
	ErrParse     = errors.New("config parse: %s")
	ErrLineParse = errors.New("config parse, line %d: '%s'")
)

// MapJSON updates config data from src json string.
// Panics if there's any error.
func (c *Config) MapJSON(src []byte) {
	dst := make(Cfg)
	err := json.Unmarshal(src, &dst)
	if err != nil {
		panic(ErrParse.Format(err))
	}
	c.Map(dst)
}

// Read acts like ReadFile but using a filename instead.
func (c *Config) Read(filename string) error {
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	return c.ReadFile(fh)
}

// ReadFile parses file's content. Returns an error if any.
func (c *Config) ReadFile(file io.ReadSeeker) (err error) {
	err = nil
	s := make([]byte, 1)
	if _, err := file.Read(s); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	defer func() {
		if e := recover(); e != nil {
			var ok bool
			err, ok = e.(error)
			if !ok {
				panic(e)
			}
		}
	}()
	if string(s) == "{" {
		b, e := ioutil.ReadAll(file)
		if e != nil && e != io.EOF {
			return e
		}
		c.MapJSON(b)
		return err
	}
	return c.parse(file)
}

type optinfo struct {
	name string
	val  string
}

var parseSection = regexp.MustCompile(`^\[([0-9A-Za-z._-]+)\]$`)

func (c *Config) parse(file io.Reader) error {
	src := make(Cfg)
	sect := "default"
	s := bufio.NewScanner(file)
	lno := 0
	for s.Scan() {
		l := s.Text()
		lno += 1
		l2 := strings.TrimSpace(l)
		if l2 == "" || l2[0] == '#' || l2[0] == ';' {
			continue
		}
		m := parseSection.FindStringSubmatch(l)
		if m != nil {
			sect = m[1]
			_, ok := src[sect]
			if !ok {
				src[sect] = make(Option)
			}
		} else {
			opt, err := parseLine(lno, l)
			if err != nil {
				return err
			}
			if opt != nil {
				src[sect][opt.name] = opt.val
			}
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	c.Map(src)
	return nil
}

var parseOption = regexp.MustCompile(`^([0-9A-Za-z._-]+) *= *(.*)$`)

func parseLine(lno int, l string) (*optinfo, error) {
	m := parseOption.FindStringSubmatch(l)
	if m != nil {
		return &optinfo{m[1], m[2]}, nil
	}
	return nil, ErrLineParse.Format(lno, l)
}
