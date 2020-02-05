// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"encoding/json"
	"io"

	"github.com/jrmsdev/gojc/errors"
)

var (
	ErrParse = errors.New("config parse: %s")
)

// MapJSON updates config data from src json string.
// Panics if there's any error.
func (c *Config) MapJSON(src string) {
	dst := make(Cfg)
	err := json.Unmarshal([]byte(src), &dst)
	if err != nil {
		panic(ErrParse.Format(err))
	}
	c.Map(dst)
}

// Read parses filename content and returns a new config instance. Or an error,
// if any.
func (c *Config) Read(filename string) (*Config, error) {
	//~ fh, err := os.Open(filename)
	//~ if err != nil {
		//~ return nil, err
	//~ }
	//~ defer fh.Close()
	//~ return ReadFile(fh)
	return nil, nil
}

// ReadFile acts like Read but using a file pointer instead of a filename.
func (c *Config) ReadFile(file io.Reader) (*Config, error) {
	//~ c, err := parseFile(file)
	//~ if err != nil {
		//~ return nil, err
	//~ }
	//~ return &Config{c}, nil
	return nil, nil
}
