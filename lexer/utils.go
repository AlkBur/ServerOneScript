package lexer

import "unicode"

type HashSet map[string]struct{}

func NewHashSet() HashSet {
	return make(HashSet)
}

func (s HashSet) Add(val string) {
	s[val] = struct{}{}
}

func (s HashSet) Remove(val string) {
	delete(s, val)
}

func (s HashSet) Has(val string) bool {
	_, ok := s[val]
	return ok
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}
