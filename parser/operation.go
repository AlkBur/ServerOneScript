package parser

import (
	"github.com/AlkBur/ServerOneScript/ast"
	"github.com/AlkBur/ServerOneScript/lexer"
)

func (p *Parser) parseBinaryOperation(ExprPrec int, LHS ast.Node) (node ast.Node) {
	// Если это бинарный оператор, получаем его приоритет
	for {
		p.expectTokens(lexer.TokenOperator)
		if p.err != nil {
			return
		}

		TokPrec := GetPrecedence(p.curToken.Value)
		if TokPrec < 0 {
			p.error("Не указан приоритет для операции %v", p.curToken.Value)
			return
		}

		// Если этот бинарный оператор связывает выражения по крайней мере так же,
		// как текущий, то используем его
		if TokPrec < ExprPrec {
			node = LHS
			return
		}

		// Отлично, мы знаем, что это бинарный оператор.
		BinOp := p.curToken.Value
		loc := p.curToken.Position
		p.NextToken() // получаем бинарный оператор

		// Разобрать первичное выражение после бинарного оператора
		RHS := p.ParsePrimaryExpression()
		if p.err != nil {
			return
		}

		if p.curToken.Type != lexer.TokenEOF && !p.curToken.Is(lexer.TokenSyntax, ";", ")") {
			// Если BinOp связан с RHS меньшим приоритетом, чем оператор после RHS,
			// то берём часть вместе с RHS как LHS.
			NextPrec := GetPrecedence(p.curToken.Value)
			if TokPrec < NextPrec {
				RHS = p.parseBinaryOperation(TokPrec+1, RHS)
				if p.err != nil {
					return
				}
			}
		} else {
			node = ast.NewBinaryNode(BinOp, LHS, RHS)
			node.SetLocation(loc)
			return
		}

		// Собираем LHS/RHS.
		LHS = ast.NewBinaryNode(BinOp, LHS, RHS)
		LHS.SetLocation(loc)
	}
	p.error("Ошибка в Binary Operation")
	return
}
