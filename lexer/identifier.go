package lexer

import (
	"unicode"
)

func (l *Lexer) scanIdentifier() bool {
	if l.err != nil {
		return false
	}
	if !isLetter(l.lastRune) {
		return false
	}

	l.position = l.stream.Position()

	l.content.WriteRune(l.lastRune)
	for {
		l.lastRune, l.err = l.stream.Next()
		if l.err != nil {
			break
		}
		if isLetter(l.lastRune) || unicode.IsNumber(l.lastRune) {
			l.content.WriteRune(l.lastRune)
		} else {
			break
		}
	}

	l.emit(TokenIdentifier)
	return true
}
