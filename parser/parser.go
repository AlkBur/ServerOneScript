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

func (p *Parser) expect(kind lexer.TokenType, values ...string) {
	if p.curToken.Is(kind, values...) {
		return
	}
	p.error("Неожиданный токен и значение %v", p.curToken)
}

func (p *Parser) expectTokenValues(values ...string) {
	if p.curToken.IsValue(values...) {
		return
	}
	p.error("Неожиданное значение %v", p.curToken)
}

func (p *Parser) expectTokens(kinds ...lexer.TokenType) {
	if p.curToken.IsToken(kinds...) {
		return
	}
	p.error("Неожиданный токен %v", p.curToken)
}

//NextToken - Получение следующего токена
func (p *Parser) NextToken() *lexer.Token {
	p.curToken = p.l.NextToken()
	return p.curToken
}

//Parse - Получение дерева AST
func (p *Parser) Parse() (tree *ast.Tree, err error) {
	tree = ast.New()

	p.NextToken()
	for p.curToken.Type != lexer.TokenEOF && p.err == nil {
		switch p.curToken.Type {
		case lexer.TokenEOF:
			break
		case lexer.TokenEnd:
			p.NextToken()
		default:
			node := p.parsePrimary()
			if p.err != nil {
				err = p.err
				return
			} else {
				tree.AddNode(node)
			}
		}
	}
	return
}

func (p *Parser) parsePrimary() (node ast.Node) {
	p.expect(lexer.TokenIdentifier)
	if p.err != nil {
		return nil
	}
	if p.curToken.IsValue("Перем", "Var") {
		p.NextToken()

		p.expect(lexer.TokenIdentifier)
		if p.err != nil {
			return nil
		}

		val := ast.NewVariableNode(p.curToken.Value, ast.NewBaseNode(ast.Undefined, p.curToken.Position))
		val.SetLocation(p.curToken.Position)

		p.NextToken()
		if p.curToken.Type == lexer.TokenIdentifier {
			p.expect(lexer.TokenIdentifier, "Экспорт", "Export")
			if p.err != nil {
				return nil
			}
			val.Export = true
			p.NextToken()
		}
		node = val
	} else {
		name := p.curToken.Value
		loc := p.curToken.Position
		p.NextToken()

		p.expect(lexer.TokenOperator, "=", "(")
		if p.err != nil {
			return nil
		}

		if p.curToken.Value == "=" {
			p.NextToken()
			val := p.parsePrimaryExpression()
			if p.err != nil {
				return nil
			}
			node = ast.NewVariableNode(name, val)
			node.SetLocation(loc)
		} else {
			//Это процедура или функция
			p.NextToken()

		}
	}

	if p.curToken.Type != lexer.TokenEOF {
		p.expect(lexer.TokenEnd, ";")
	}
	return node
}

func (p *Parser) parsePrimaryExpression() (node ast.Node) {
	var err error
	switch p.curToken.Type {

	case lexer.TokenIdentifier:
		if p.curToken.IsValue("Истина", "Ложь", "True", "False") {
			node, err = ast.NewBoolNode(p.curToken.Value)
			if err != nil {
				p.error(err.Error())
				return
			}
			node.SetLocation(p.curToken.Position)
			p.NextToken()
			return
		} else if p.curToken.IsValue("Неопределено", "Undefined, Null") {
			node = ast.NewBaseNode(ast.Undefined, p.curToken.Position)
			p.NextToken()
			return
		}
		//node = p.parseIdentifierExpression()
	case lexer.TokenNumber:
		node, err = ast.NewNumberNode(p.curToken.Value)
		if err != nil {
			p.error(err.Error())
			return
		}
		node.SetLocation(p.curToken.Position)
		p.NextToken()
		return node
	case lexer.TokenString:
		node = ast.NewStringNode(p.curToken.Value)
		node.SetLocation(p.curToken.Position)
		p.NextToken()
		return node
	case lexer.TokenDate:
		node, err = ast.NewDateNode(p.curToken.Value)
		if err != nil {
			p.error(err.Error())
			return
		}
		node.SetLocation(p.curToken.Position)
		p.NextToken()
		return node
	default:
		/*
			if token.Is(Bracket, "[") {
				node = p.parseArrayExpression(token)
			} else if token.Is(Bracket, "{") {
				node = p.parseMapExpression(token)
			} else {
				p.error("unexpected token %v", token)
			}
		*/
	}
	return nil
	//return p.parsePostfixExpression(node)
}
