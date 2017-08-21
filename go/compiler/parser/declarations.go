package parser

import (
	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
)

func parseInterfaceDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknInterface)
	identifier := p.parseIdentifier()

	var inherits []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
	}

	p.parseRequired(lexer.TknOpenBrace)

	n := ast.InterfaceDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		IsAbstract: abstract,
	}

	p.Scope.Declare("interface", n)

	p.parent = p.Scope
	p.Scope = n
}

// like any list parser, but enforces that each node must be a reference
func (p *Parser) parseReferenceList() []ast.ReferenceNode {
	var refs []ast.ReferenceNode
	first := p.parseReference()
	refs = append(refs, first)
	for p.parseOptional(lexer.TknComma) {
		refs = append(refs, p.parseReference())
	}
	return refs
}

func parseClassDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknClass)
	identifier := p.parseIdentifier()

	// is and inherits can be in any order

	var inherits, interfaces []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
		if p.parseOptional(lexer.TknIs) {
			interfaces = p.parseReferenceList()
		}
	} else if p.parseOptional(lexer.TknIs) {
		interfaces = p.parseReferenceList()
		if p.parseOptional(lexer.TknInherits) {
			inherits = p.parseReferenceList()
		}
	}

	p.parseRequired(lexer.TknOpenBrace)

	n := ast.ClassDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		Interfaces: interfaces,
		IsAbstract: abstract,
	}

	p.Scope.Declare("class", n)

	p.parent = p.Scope
	p.Scope = n
}

func parseContractDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)
	p.parseRequired(lexer.TknContract)
	identifier := p.parseIdentifier()

	// is and inherits can be in any order

	var inherits []ast.ReferenceNode

	if p.parseOptional(lexer.TknInherits) {
		inherits = p.parseReferenceList()
	}

	p.parseRequired(lexer.TknOpenBrace)

	n := ast.ContractDeclarationNode{
		Identifier: identifier,
		Supers:     inherits,
		IsAbstract: abstract,
	}

	p.Scope.Declare("contract", n)

	p.parent = p.Scope
	p.Scope = n
}

func (p *Parser) parseParameters() []ast.Node {
	var params []ast.Node
	p.parseRequired(lexer.TknOpenBracket)

	p.parseRequired(lexer.TknCloseBracket)
	return params
}

func (p *Parser) parseResults() []ast.Node {
	return nil
}

func parseFuncDeclaration(p *Parser) {

	abstract := p.parseOptional(lexer.TknAbstract)

	identifier := p.parseIdentifier()

	params := p.parseParameters()

	results := p.parseResults()

	p.parseRequired(lexer.TknOpenBrace)

	p.validate(ast.FuncDeclaration)

	n := ast.FuncDeclarationNode{
		Identifier: identifier,
		Parameters: params,
		Results:    results,
		IsAbstract: abstract,
	}

	p.Scope.Declare("func", n)

	p.parent = p.Scope
	p.Scope = n
}

func parseTypeDeclaration(p *Parser) {
	p.parseRequired(lexer.TknType)
	identifier := p.parseIdentifier()

	value := p.parseReference()

	p.validate(ast.TypeDeclaration)

	n := ast.TypeDeclarationNode{
		Identifier: identifier,
		Value:      value,
	}

	p.Scope.Declare("type", n)
}

func parseMapType(p *Parser) {

	p.parseRequired(lexer.TknMap)
	p.parseRequired(lexer.TknOpenSquare)

	key := p.parseType()

	p.parseRequired(lexer.TknCloseSquare)

	value := p.parseType()

	p.validate(ast.MapType)

	p.Scope.Declare("", ast.MapTypeNode{
		Key:   key,
		Value: value,
	})
}

func parseArrayType(p *Parser) {
	p.parseRequired(lexer.TknOpenSquare)

	//typ := p.parseExpression()
	//var max ast.Node

	if p.parseOptional(lexer.TknColon) {
		//	max = p.parseExpression()
	}
	p.validate(ast.ArrayType)

	p.Scope.Declare("", ast.ArrayTypeNode{
	//Value: typ,
	})
}
