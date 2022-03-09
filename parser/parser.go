package parser

import (
	"fmt"
	"github.com/AlkBur/ServerOneScript/ast"
	"github.com/AlkBur/ServerOneScript/lexer"
	"github.com/AlkBur/ServerOneScript/runes"
)

type Parser struct {
	l        *lexer.Lexer
	debug    bool
	curToken *lexer.Token
	err      *runes.Error
	depth    int // closure call depth
}

//New - Инициализация парсера
func New(b []byte, module string) *Parser {
	return &Parser{
		l: lexer.Parse(runes.New(b, module)),
	}
}

//NewFromString - Инициализация парсера
func NewFromString(str string, module string) *Parser {
	return &Parser{
		l: lexer.ParseString(str, module),
	}
}

func (p *Parser) error(format string, args ...interface{}) {
	if p.err == nil { // show first error
		p.err = &runes.Error{
			Location: p.curToken.Position,
			Message:  fmt.Sprintf(format, args...),
		}
	}
}

func (p *Parser) expect(kind lexer.TokenType, values ...string) bool {
	if p.curToken.Is(kind, values...) {
		return true
	}
	p.error("Неожиданный токен и значение %v", p.curToken)
	return false
}

func (p *Parser) expectTokenValues(values ...string) bool {
	if p.curToken.IsValue(values...) {
		return true
	}
	p.error("Неожиданное значение %v", p.curToken)
	return false
}

func (p *Parser) expectTokens(kinds ...lexer.TokenType) bool {
	if p.curToken.IsToken(kinds...) {
		return true
	}
	p.error("Неожиданный токен %v", p.curToken)
	return false
}

//NextToken - Получение следующего токена
func (p *Parser) NextToken() *lexer.Token {
	if p.curToken == nil || p.curToken.Type != lexer.TokenEOF {
		p.curToken = p.l.NextToken()
	}
	return p.curToken
}

//Parse - Получение дерева AST
func (p *Parser) Parse() (tree *ast.Tree, err error) {
	tree = ast.New()

	p.NextToken()
	for p.curToken.Type != lexer.TokenEOF && p.err == nil {
		node := p.parsePrimary()
		if p.err == nil && !isNil(node) {
			tree.AddNode(node)
			p.NextToken()
		}
	}
	if p.err != nil {
		err = p.err
	}
	return
}

func (p *Parser) parsePrimary() (node ast.Node) {
	if !p.expectTokens(lexer.TokenIdentifier) {
		return
	}

	if p.curToken.Is(lexer.TokenIdentifier, "IF", "ЕСЛИ") {

	} else if p.curToken.Is(lexer.TokenIdentifier, "FUNCTION", "ФУНКЦИЯ") {

	} else if p.curToken.Is(lexer.TokenIdentifier, "PROCEDURE", "ПРОЦЕДУРА") {

	} else if p.curToken.Is(lexer.TokenIdentifier, "VAR", "ПЕРЕМ") {
		p.NextToken()
		if !p.expectTokens(lexer.TokenIdentifier) {
			return
		}
		node = ast.NewVariableNode(p.curToken.Value, ast.NewBaseNode(ast.Undefined, p.curToken.Position))
		node.SetLocation(p.curToken.Position)
		p.NextToken()
		if p.curToken.Type != lexer.TokenEOF {
			p.expect(lexer.TokenSyntax, ";")
		}
		return
	} else if p.curToken.Is(lexer.TokenIdentifier, "RETURN", "ВОЗВРАТ") {
		p.NextToken()
		if p.curToken.Type != lexer.TokenEOF || p.curToken.Is(lexer.TokenSyntax, ";") {
			node = ast.NewReturnNode(ast.NewBaseNode(ast.Undefined, p.curToken.Position))
			return
		}
		node = ast.NewReturnNode(p.parseExpression())
		return
	} else if p.curToken.Is(lexer.TokenIdentifier, "CONTINUE", "ПРОДОЛЖИТЬ") {
		node = ast.NewBaseNode(ast.Continue, p.curToken.Position)
		p.NextToken()
		if !p.expect(lexer.TokenSyntax, ";") {
			node = nil
			return
		}
		return
	}

	node = p.parseExpressionStatement()
	return
}
