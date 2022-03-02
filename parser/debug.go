package parser

import (
	"errors"
	"fmt"
	"github.com/AlkBur/ServerOneScript/lexer"
	"os"
)

func errorft(tok *lexer.Token, format string, v ...interface{}) error {

	var tokString string
	if tok != nil {
		tokString = tok.Value
	}
	gs := format + "\n " + tokString
	return errorf(gs, v...)
}

func errorf(format string, v ...interface{}) error {
	return errors.New(fmt.Sprintf(format+"\n", v...))
}

func (p *Parser) traceIn(funcname string) int {
	if !p.debug {
		return 0
	}
	//if GENERATION == 1 {
	//	funcname = getCallerName(2)
	//}
	debugf("func %s is gonna read %s", funcname, p.l.Content())
	return 0
}

func (p *Parser) traceOut(funcname string) {
	if !p.debug {
		return
	}
	if r := recover(); r != nil {
		os.Exit(1)
	}
	/*
		if GENERATION == 1 {
			funcname = getCallerName(2)
		}
	*/
	debugf("func %s end after %s", funcname, p.l.Content())
}

func debugf(format string, v ...interface{}) {
	s2 := fmt.Sprintf(format, v...)
	var b []byte = []byte(s2)
	b = append(b, '\n')
	os.Stderr.Write(b)
}
