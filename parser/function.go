package parser

import (
	"github.com/AlkBur/ServerOneScript/ast"
	"github.com/AlkBur/ServerOneScript/lexer"
)

//parseExpression - Получение выражения
func (p *Parser) parseFunctionCall(name string) (node ast.Node) {
	if p.curToken.Is(lexer.TokenSyntax, ")") {
		node = ast.NewFunctionCallNode(name, nil)
		p.NextToken()
		return
	}

	args := make([]ast.Node, 0)

	for {
		arg := p.parseExpression()
		if p.err != nil {
			return
		}
		args = append(args, arg)

		if p.curToken.Is(lexer.TokenSyntax, ",") {
			p.NextToken()
		} else {
			break
		}
	}

	if !p.expect(lexer.TokenSyntax, ")") {
		return
	}

	node = ast.NewFunctionCallNode(name, args)
	p.NextToken()
	return
}
