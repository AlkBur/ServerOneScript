package parser

import (
	"github.com/AlkBur/ServerOneScript/ast"
	"github.com/AlkBur/ServerOneScript/lexer"
)

//parseExpression - Получение выражения
func (p *Parser) parseExpressionStatement() (node ast.Node) {
	if !p.expectTokens(lexer.TokenIdentifier) {
		return
	}
	name := p.curToken.Value
	loc := p.curToken.Position
	p.NextToken()
	switch p.curToken.Value {
	case "=":
		p.NextToken()
		expr := p.parseExpression()
		if p.err == nil {
			node = ast.NewVariableNode(name, expr)
			node.SetLocation(loc)
		}
	case "(":
		node = p.parseFunctionCall(name)
	default:
		p.error("Неожиданный токен для выражения %v", p.curToken)
		return
	}
	if p.curToken.Type != lexer.TokenEOF {
		p.expect(lexer.TokenSyntax, ";")
	}
	return
}

//parseExpression - Получение выражения
func (p *Parser) parseExpression() (node ast.Node) {
	node = p.ParsePrimaryExpression()
	if p.err != nil {
		return
	}
	if p.curToken.Type == lexer.TokenEOF || p.curToken.Is(lexer.TokenSyntax, ";", ")", ",") {
		return
	}
	node = p.parseBinaryOperation(0, node)
	return
}

func (p *Parser) ParsePrimaryExpression() (node ast.Node) {
	if p.curToken.IsToken(lexer.TokenIdentifier) {
		node = p.parseIdentifierExpression()
		return
	} else if p.curToken.IsToken(lexer.TokenNumber) {
		node = p.parseNumberExpression()
		return
	} else if p.curToken.IsToken(lexer.TokenString) {
		node = p.parseStringExpression()
		return
	} else if p.curToken.IsToken(lexer.TokenDate) {
		node = p.parseDateExpression()
		return
	} else if p.curToken.Is(lexer.TokenSyntax, "(") {
		p.NextToken()
		return p.parseParentExpression()
	} else if p.curToken.Is(lexer.TokenOperator, "-", "+") {
		BinOp := p.curToken.Value
		location := p.curToken.Position
		p.NextToken()

		if p.curToken.Type == lexer.TokenNumber {
			var err error
			node, err = ast.NewNumberNode(BinOp + p.curToken.Value)
			node.SetLocation(location)
			if err != nil {
				p.error("Ошибка при получинии числа %v", err)
				return
			}
			p.NextToken()
		} else {
			LHS, err := ast.NewNumberNode(BinOp + "1")
			LHS.SetLocation(location)
			if err != nil {
				p.error("Ошибка при получинии числа %v", err)
				return
			}

			TokPrec := GetPrecedence("*")
			RHS := p.ParsePrimaryExpression()
			if p.err != nil {
				return nil
			}

			// Если BinOp связан с RHS меньшим приоритетом, чем оператор после RHS,
			// то берём часть вместе с RHS как LHS.
			NextPrec := GetPrecedence(p.curToken.Value)
			if TokPrec < NextPrec {
				RHS = p.parseBinaryOperation(TokPrec+1, RHS)
				if p.err != nil {
					return nil
				}
			}

			// Собираем LHS/RHS.
			node = ast.NewBinaryNode(BinOp, LHS, RHS)
			p.parseBinaryOperation(NextPrec, node)
		}
		return
	}
	p.error("unknown token when expecting an expression")
	return
}

func (p *Parser) parseIdentifierExpression() (node ast.Node) {
	if p.curToken.IsValue("TRUE", "ИСТИНА", "FALSE", "ЛОЖЬ") {
		node = p.parseBooleanExpression()
		return
	} else if p.curToken.IsValue("UNDEFINED", "НЕОПРЕДЕЛЕНО") {
		node = p.parseUndefinedExpression()
		return
	} else if p.curToken.IsValue("NULL") {
		node = p.parseNullExpression()
		return
	}

	IdName := p.curToken.Value
	loc := p.curToken.Position
	p.NextToken()
	if p.curToken.Value != "(" {
		node = ast.NewLiteralNode(IdName)
		node.SetLocation(loc)
		return
	}
	p.NextToken()
	node = p.parseFunctionCall(IdName)
	if p.err != nil {
		return
	}
	node.SetLocation(loc)
	return
}

func (p *Parser) parseNumberExpression() (node ast.Node) {
	var err error
	node, err = ast.NewNumberNode(p.curToken.Value)
	node.SetLocation(p.curToken.Position)
	if err != nil {
		p.error("Ошибка при получинии числа %v", err)
	}
	p.NextToken()
	return
}

func (p *Parser) parseDateExpression() (node ast.Node) {
	var err error
	node, err = ast.NewDateNode(p.curToken.Value)
	node.SetLocation(p.curToken.Position)
	if err != nil {
		p.error("Ошибка при получинии даты %v", err)
	}
	p.NextToken()
	return
}

func (p *Parser) parseBooleanExpression() (node ast.Node) {
	var err error
	node, err = ast.NewBoolNode(p.curToken.Value)
	node.SetLocation(p.curToken.Position)
	if err != nil {
		p.error("Ошибка при получинии boolean %v", err)
	}
	p.NextToken()
	return
}

func (p *Parser) parseParentExpression() (node ast.Node) {
	node = p.parseExpression()
	if p.err != nil {
		return
	}

	if p.curToken.Value != ")" {
		p.error("expected ')'")
		return
	}
	p.NextToken()
	return
}

func (p *Parser) parseStringExpression() (node ast.Node) {
	node = ast.NewStringNode(p.curToken.Value)
	node.SetLocation(p.curToken.Position)
	p.NextToken()
	return
}

func (p *Parser) parseNullExpression() (node ast.Node) {
	node = ast.NewBaseNode(ast.Null, p.curToken.Position)
	p.NextToken()
	return
}

func (p *Parser) parseUndefinedExpression() (node ast.Node) {
	node = ast.NewBaseNode(ast.Undefined, p.curToken.Position)
	p.NextToken()
	return
}
