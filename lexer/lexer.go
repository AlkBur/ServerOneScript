package lexer

import (
	"errors"
	"github.com/AlkBur/ServerOneScript/runes"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	stream   *runes.Source
	lastRune rune
	content  strings.Builder

	Token      TokenType
	directives HashSet

	position runes.Location

	err error
}

func Parse(s *runes.Source) *Lexer {
	l := &Lexer{
		stream:     s,
		lastRune:   ' ',
		Token:      TokenUnknown,
		directives: NewHashSet(),
	}
	l.position.SetLine(1)
	return l
}

func ParseString(s string, module string) *Lexer {
	return Parse(runes.NewFromString(s, module))
}

func ParseBytes(b []byte, module string) *Lexer {
	return Parse(runes.New(b, module))
}

func (l *Lexer) NextToken() *Token {
	l.content.Reset()
	l.Token = TokenUnknown
	l.err = nil

	l.deleteSpace()
	if l.err == nil && l.lastRune == '/' {
		if l.deleteComment() {
			return l.NextToken()
		} else {
			l.content.WriteRune('/')
			l.emit(TokenOperator)
			return l.NewToken()
		}
	}
	if l.scanOperator() {
		return l.NewToken()
	}

	if l.scanIdentifier() { // идентификатор
		return l.NewToken()
	}

	if l.scanNumber() { // Число: [0-9.]+
		return l.NewToken()
	}

	if l.scanString() {
		return l.NewToken()
	}

	if l.scanPreprocessor() {
		return l.NewToken()
	}

	if l.scanDate() {
		return l.NewToken()
	}

	if l.err == nil {
		l.lastRune, l.err = l.stream.Next()
	}

	// Проверка конца файла.
	if l.err != nil {
		if errors.Is(l.err, io.EOF) {
			l.Token = TokenEOF
		} else {
			l.Token = TokenError
		}
	} else {
		l.Token = TokenUnknown
	}

	return l.NewToken()
}

func (l *Lexer) deleteSpace() {
	if l.err != nil {
		return
	}
	for unicode.IsSpace(l.lastRune) {
		l.lastRune, l.err = l.stream.Next()
		if l.err != nil {
			break
		}
	}
}

func (l *Lexer) deleteComment() bool {
	if l.err != nil {
		return false
	}
	if l.lastRune == '/' {
		l.lastRune, l.err = l.stream.Next()
		if l.err != nil {
			return false
		}
		if l.lastRune == '/' {
			// Комментарий до конца строки
			for l.err == nil && l.lastRune != '\n' && l.lastRune != '\r' {
				l.lastRune, l.err = l.stream.Next()
				if l.err != nil {
					break
				}
			}
			return true
		} else {
			return false
		}
	}
	return false
}

func (l *Lexer) emit(token TokenType) {

	if l.err != nil {
		if errors.Is(l.err, io.EOF) {
			l.Token = token
			l.err = nil
		} else {
			l.Token = TokenError
		}
	} else {
		l.Token = token
	}
	/*
		if l.err == nil && l.Token != TokenEOF && l.Token != TokenDate &&
			l.Token != TokenNumber && l.Token != TokenPreprocessor {

			token, ok := TokenStringID[strings.ToUpper(l.content.String())]
			if ok {
				l.Token = token
			}
		}
	*/
}

func (l *Lexer) Content() string {
	return l.content.String()
}

func (l *Lexer) Error() error {
	return l.err
}

func (l *Lexer) Location() runes.Location {
	return l.position
}

func (l *Lexer) BindError(err *runes.Error) error {
	return errors.New(err.Bind(l.stream).Error())
}

func (l *Lexer) GetCodeLine(num int) string {
	return l.stream.StringLine(num)
}
