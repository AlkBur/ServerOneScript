package runes

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Source struct {
	module   string
	data     []rune
	lines    []Line
	location Location
	offset   int
}

func NewFromFile(filename string) (*Source, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(bytes, strings.TrimSuffix(filename, filepath.Ext(filename))), nil
}

func NewFromString(data, module string) *Source {
	return New([]byte(data), module)
}

func New(data []byte, module string) *Source {
	s := &Source{
		module: module,
		data:   bytes.Runes(data),
		lines:  make([]Line, 0),
		location: Location{
			line:   1,
			column: 0,
		},
		offset: -1,
	}

	start := 0
	for i, r := range s.data {
		if r == '\n' {
			s.lines = append(s.lines, Line{start, i})
			start = i + 1
		}
	}
	s.lines = append(s.lines, Line{start, len(s.data)})
	return s
}

func (s *Source) Next() (r rune, err error) {
	nextID := s.offset + 1
	if nextID >= len(s.data) {
		err = io.EOF
		return
	}
	r = s.data[nextID]
	if r == '\n' {
		s.location.line++
		s.location.column = -1
	}
	s.offset = nextID
	s.location.column++
	return
}

func (s *Source) Prev() (r rune, err error) {
	nextID := s.offset - 1
	if nextID < 0 {
		err = io.EOF
		return
	}
	r = s.data[nextID]
	s.offset = nextID
	if r == '\n' {
		s.location.column = 0
		s.location.line--
		for i := nextID; i < 0; i-- {
			if r == '\n' {
				break
			}
			s.location.column++
		}
		s.location.column--
	}
	return
}

func (s *Source) Line() int {
	return s.location.line
}

func (s *Source) Column() int {
	return s.location.column
}

func (s *Source) Position() Location {
	return s.location
}

func (s *Source) StringLine(line int) string {
	if line > len(s.lines) || line < 1 {
		return ""
	}
	return string(s.data[s.lines[line-1].start:s.lines[line-1].end])
}
