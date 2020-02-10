// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/jrmsdev/gojc/errors"
)

var (
	ErrParse = errors.New("config parse: %s")
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
		return err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	if string(s) == "{" {
		b, e := ioutil.ReadAll(file)
		if e != nil {
			return e
		}
		c.MapJSON(b)
		return err
	}
	return c.parse(file)
}

func (c *Config) parse(file io.Reader) error {
	s := bufio.NewScanner(file)
	for s.Scan() {
		l := s.Text()
		println(l)
	}
	return s.Err()
}
