package lexer

func (l *Lexer) scanOperator() bool {
	if l.err != nil {
		return false
	}

	l.position = l.stream.Position()

	switch l.lastRune {
	case '+', '-', '*', '%', ',', '.', '(', ')', '[', ']', '=':
		l.content.WriteRune(l.lastRune)
		l.lastRune, l.err = l.stream.Next()
		l.emit(TokenOperator)
		return true
	case '<', '>':
		last := l.lastRune
		l.content.WriteRune(l.lastRune)
		l.lastRune, l.err = l.stream.Next()
		if (l.lastRune == '=') || (last == '<' && l.lastRune == '>') {
			l.content.WriteRune(l.lastRune)
			l.lastRune, l.err = l.stream.Next()
		}
		l.emit(TokenOperator)
		return true
	case ';':
		last := l.lastRune
		l.content.WriteRune(l.lastRune)
		l.lastRune, l.err = l.stream.Next()
		if (l.lastRune == '=') || (last == '<' && l.lastRune == '>') {
			l.content.WriteRune(l.lastRune)
			l.lastRune, l.err = l.stream.Next()
		}
		l.emit(TokenEnd)
		return true
	}
	return false
}
