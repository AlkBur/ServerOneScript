package runes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	snippetNotFound = "line not found, wanted '%v'"
)

func TestStringSource_SnippetMultiLine(t *testing.T) {
	assert := assert.New(t)

	source := NewFromString("hello\nworld\nmy\nbub\n", "")

	str := source.StringLine(1)
	assert.Equal(str, "hello", snippetNotFound, 1)

	str = source.StringLine(2)
	assert.Equal(str, "world", snippetNotFound, 2)

	str = source.StringLine(3)
	assert.Equal(str, "my", snippetNotFound, 3)

	str = source.StringLine(4)
	assert.Equal(str, "bub", snippetNotFound, 4)

	str = source.StringLine(5)
	assert.Equal(str, "", snippetNotFound, 5)

}

func TestStringSource_SnippetSingleLine(t *testing.T) {
	assert := assert.New(t)

	source := NewFromString("hello, world", "")

	str := source.StringLine(1)
	assert.Equal(str, "hello, world", snippetNotFound, 1)

	str = source.StringLine(2)
	assert.Equal(str, "", snippetNotFound, 2)

}
