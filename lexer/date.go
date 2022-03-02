package lexer

import (
	"errors"
	"unicode"
)

func (l *Lexer) scanDate() bool {
	if l.err != nil {
		return false
	}
	if l.lastRune != '\'' {
		return false
	}

	l.position = l.stream.Position()

	l.lastRune, l.err = l.stream.Next()
	for l.err == nil && unicode.IsNumber(l.lastRune) {
		l.content.WriteRune(l.lastRune)
		l.lastRune, l.err = l.stream.Next()
	}
	if l.lastRune != '\'' || (l.content.Len() != 8 && l.content.Len() != 14) {
		l.err = errors.New("Incorrect date: " + l.content.String())
	} else {
		l.lastRune, l.err = l.stream.Next()
	}
	l.emit(TokenDate)
	return true
}
