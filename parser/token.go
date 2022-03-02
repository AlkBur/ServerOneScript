package parser

import (
	"github.com/AlkBur/ServerOneScript/lexer"
	"io"
	"sync"
)

var tokensPool = sync.Pool{
	New: func() interface{} { return []*lexer.Token{} },
}

func getTokens() []*lexer.Token {
	return tokensPool.Get().([]*lexer.Token)
}

func putTokens(tokens []*lexer.Token) {
	tokens = tokens[:0] // сброс
	tokensPool.Put(tokens)
}

type TokenStream struct {
	tokens    []*lexer.Token
	nextToken int
}

func NewTokenStream(t []*lexer.Token) *TokenStream {
	return &TokenStream{
		tokens: t,
	}
}

func (st *TokenStream) Close() {
	putTokens(st.tokens)
}

func (st *TokenStream) NextToken() (t *lexer.Token, err error) {
	if st.nextToken >= len(st.tokens) {
		err = io.EOF
		return
	}
	t = st.tokens[st.nextToken]
	st.nextToken++
	return
}
