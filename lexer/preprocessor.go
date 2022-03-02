package lexer

import (
	"unicode"
)

func (l *Lexer) scanPreprocessor() bool {
	if l.err != nil {
		return false
	}
	if l.lastRune != '#' {
		return false
	}

	l.position = l.stream.Position()

	for {
		l.lastRune, l.err = l.stream.Next()
		if l.err != nil {
			break
		}
		if !unicode.IsLetter(l.lastRune) {
			break
		}
		l.content.WriteRune(l.lastRune)
	}

	l.emit(TokenPreprocessor)
	return true

}
