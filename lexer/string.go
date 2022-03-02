package lexer

import (
	"errors"
	"io"
)

func (l *Lexer) scanString() bool {
	if l.err != nil {
		return false
	}
	if l.lastRune != '"' {
		return false
	}

	l.position = l.stream.Position()

	bEnd := false
	bNextLine := false
	for {
		l.lastRune, l.err = l.stream.Next()
		if l.err != nil {
			break
		}
		if bNextLine {
			l.deleteSpace()
		}

		if l.lastRune == '\n' || l.lastRune == '\r' {
			bNextLine = true
			l.content.WriteRune(l.lastRune)
			continue
		}

		if bNextLine {
			if l.lastRune == '/' {
				if !l.deleteComment() {
					l.err = errors.New("Incorrect comment")
					break
				} else {
					continue
				}
			}

			if l.lastRune == '|' {
				bNextLine = false
				continue
			} else {
				break
			}
		}

		if l.lastRune == '"' {
			l.lastRune, l.err = l.stream.Next()
			if l.err != nil {
				if errors.Is(l.err, io.EOF) {
					bEnd = true
				}
				break
			}
			if l.lastRune != '"' {
				bEnd = true
				break
			}
		}

		l.content.WriteRune(l.lastRune)

	}
	if !bEnd {
		l.err = errors.New("Incorrect string")
	}
	l.emit(TokenString)
	return true
}
