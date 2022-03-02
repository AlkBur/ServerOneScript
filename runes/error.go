package runes

import (
	"fmt"
)

type Error struct {
	Location
	Message string
	Snippet string
	Module  string
}

func (e *Error) Error() string {
	return e.format()
}

func (e *Error) format() string {
	if e.Location.Empty() {
		return e.Message
	}
	module := e.Module
	if module == "" {
		module = "main"
	}
	return fmt.Sprintf(
		"%s: %s (%d:%d)%s",
		e.Module,
		e.Message,
		e.line,
		e.column,
		e.Snippet,
	)
}

func (e *Error) Bind(stream *Source) *Error {
	e.Snippet = stream.StringLine(e.Location.line)
	e.Module = stream.module
	return e
}
