package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IteratorReturnsCorrectLine(t *testing.T) {
	assert := assert.New(t)
	code := "\r\nF = 1;\r\nD = 2;\r\nX = 3;\r\n"
	lex := ParseString(code, "")
	loc := lex.Location()

	assert.Equal(1, loc.Line())
	assert.Equal(0, loc.Column())

	lex.NextToken()
	loc = lex.Location()

	assert.Equal(2, loc.Line())
	assert.Equal(1, loc.Column())
	assert.Equal("F = 1;\r", lex.GetCodeLine(2))

	lex.NextToken()
	loc = lex.Location()
	assert.Equal(3, loc.Column())
	assert.Equal("=", lex.Content())

	lex.NextToken()
	loc = lex.Location()
	assert.Equal(5, loc.Column())
	assert.Equal("1", lex.Content())

	for lex.Content() != "D" && lex.Token != TokenEOF {
		lex.NextToken()
	}

	loc = lex.Location()
	assert.Equal(1, loc.Column())
	assert.Equal(3, loc.Line())
}
