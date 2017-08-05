package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

func (p *Parser) parseExpression() ast.Node {
	// Guardian expressions can be arbitrarily chained
	// e.g. array[expr]
	// the expr could be 5 + 4 + 3, or 5 + 4 + getNumber()
	// this is all stored in one expression Node
	// however, ORDER IS IMPORTANT
	// (5 + 4) * 3 vs 5 + 4 * 3
	// these expression are not evaluated at compile time
	// actually, maybe evaluate constants fully

	// (dog() - 5) + 6
	// !((dog() - 5) + 6 > 10)
	switch p.current().Type {
	case lexer.TknMap:
		return p.parseMapLiteral()
	case lexer.TknOpenSquare:
		return p.parseArrayLiteral()
	case lexer.TknString, lexer.TknNumber, lexer.TknCharacter:
		return p.parseLiteral()
	case lexer.TknIdentifier:
		// TODO: check in range
		// TODO: won't work for package.name()
		switch p.token(1).Type {
		case lexer.TknOpenBracket:
			return p.parseCallExpression()
		case lexer.TknOpenBrace:
			return p.parseCompositeLiteral()
		}
	}
	return p.parseReference()
}

func (p *Parser) parseUnaryExpression() (n ast.UnaryExpressionNode) {
	return n
}

func (p *Parser) parseBinaryExpression() (n ast.BinaryExpressionNode) {
	return n
}

func (p *Parser) parseCallExpression() (n ast.CallExpressionNode) {
	n.Name = p.parseIdentifier()
	p.parseRequired(lexer.TknOpenBracket)
	if !p.parseOptional(lexer.TknCloseBracket) {
		n.Arguments = p.parseExpressionList()
		p.parseRequired(lexer.TknCloseBracket)
	}
	return n
}

func (p *Parser) parseArrayLiteral() (n ast.ArrayLiteralNode) {
	// [string:3]{"Dog", "Cat", ""}
	p.parseRequired(lexer.TknOpenSquare)
	n.Key = p.parseType()
	if !p.parseOptional(lexer.TknCloseSquare) {
		n.Data = append(n.Data, p.parseExpression())
		for p.parseOptional(lexer.TknComma) {
			n.Data = append(n.Data, p.parseExpression())
		}
	}
	return n
}

func (p *Parser) parseMapLiteral() (n ast.MapLiteralNode) {
	p.parseRequired(lexer.TknMap)
	p.parseRequired(lexer.TknOpenSquare)
	n.Key = p.parseType()
	p.parseRequired(lexer.TknCloseSquare)
	n.Value = p.parseType()
	p.parseRequired(lexer.TknOpenBrace)
	if !p.parseOptional(lexer.TknCloseBrace) {
		firstKey := p.parseExpression()
		p.parseRequired(lexer.TknColon)
		firstValue := p.parseExpression()
		n.Data[firstKey] = firstValue
		for p.parseOptional(lexer.TknComma) {
			key := p.parseExpression()
			p.parseRequired(lexer.TknColon)
			value := p.parseExpression()
			n.Data[key] = value
		}
	}
	return n
}

func (p *Parser) parseExpressionList() (list []ast.Node) {
	list = append(list, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		list = append(list, p.parseExpression())
	}
	return list
}

func (p *Parser) parseIndexExpression() (n ast.IndexExpressionNode) {
	n.Expression = p.parseExpression()
	n.Index = p.parseExpression()
	return n
}

func (p *Parser) parseSliceExpression() (n ast.SliceExpressionNode) {
	n.Expression = p.parseExpression()
	p.parseRequired(lexer.TknOpenSquare)
	n.Low = p.parseExpression()
	p.parseRequired(lexer.TknColon)
	if !p.parseOptional(lexer.TknCloseSquare) {
		n.High = p.parseExpression()
		p.parseRequired(lexer.TknCloseSquare)
	}
	return n
}

func (p *Parser) parseReference() (n ast.ReferenceNode) {
	n.Names = make([]string, 0)
	n.Names = append(n.Names, p.parseIdentifier())
	for p.parseOptional(lexer.TknDot) {
		n.Names = append(n.Names, p.parseIdentifier())
	}
	return n
}

func (p *Parser) parseLiteral() (n ast.LiteralNode) {
	n.LiteralType = p.current().Type
	return n
}

func (p *Parser) parseCompositeLiteral() (n ast.CompositeLiteralNode) {
	return n
}
