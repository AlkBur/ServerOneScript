package lexer

import (
	"errors"
	"unicode"
)

func (l *Lexer) scanNumber() bool {
	if l.err != nil {
		return false
	}
	if !unicode.IsDigit(l.lastRune) {
		return false
	}

	l.position = l.stream.Position()

	countPoint := 0

	for unicode.IsDigit(l.lastRune) || l.lastRune == '.' {
		l.content.WriteRune(l.lastRune)
		l.lastRune, l.err = l.stream.Next()
		if l.err != nil {
			break
		}
		if l.lastRune == '.' {
			countPoint++
			if countPoint > 1 {
				l.err = errors.New("Некорректно указана десятичная точка в числе")
				break
			}
		}
	}
	if unicode.IsLetter(l.lastRune) {
		l.err = errors.New("Некорректный символ")
	}

	l.emit(TokenNumber)
	return true
}
