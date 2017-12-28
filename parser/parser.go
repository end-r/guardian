package parser

import (
	"fmt"
	"strconv"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// Parser ...
type Parser struct {
	scope         *ast.ScopeNode
	Expression    ast.ExpressionNode
	tokens        []token.Token
	modifiers     [][]string
	lastModifiers []string
	annotations   []*ast.Annotation
	index         int
	errs          util.Errors
	line          int
	simple        bool
}

func createParser(data string) *Parser {
	p := new(Parser)
	p.line = 1
	p.tokens, _ = lexer.LexString(data)
	p.scope = &ast.ScopeNode{
		ValidTypes: []ast.NodeType{
			ast.InterfaceDeclaration, ast.ClassDeclaration,
			ast.FuncDeclaration,
		},
	}
	return p
}

func (p *Parser) current() token.Token {
	return p.token(0)
}

func (p *Parser) next() {
	p.index++
}

func (p *Parser) token(offset int) token.Token {
	return p.tokens[p.index+offset]
}

// index 0, tknlength 1
// hasTokens(0) yes always
// hasTokens(1) yes
// hasTokens(2) no
func (p *Parser) hasTokens(offset int) bool {
	return p.index+offset <= len(p.tokens)
}

func (p *Parser) parseOptional(types ...token.Type) bool {
	if !p.hasTokens(1) {
		return false
	}
	for _, t := range types {
		if p.current().Type == t {
			p.next()
			return true
		}
	}
	return false
}

// TODO: clarify what this actually returns
func (p *Parser) parseRequired(types ...token.Type) token.Type {
	if !p.hasTokens(1) {
		p.addError(fmt.Sprintf("Required one of {%s}, found %s", listTypes(types), "nothing"))
		// TODO: what should be returned here
		return token.Invalid
	}
	for _, t := range types {
		if p.isNextToken(t) {
			p.next()
			return t
		}
	}
	p.addError(fmt.Sprintf("Required one of {%s}, found %s", listTypes(types), p.current().Name()))
	// correct return type
	return token.Invalid
}

func listTypes(types []token.Type) string {
	s := ""
	for _, t := range types {
		s += strconv.Itoa(int(t))
	}
	return s
}

func (p *Parser) parseIdentifier() string {
	if !p.hasTokens(1) {
		p.addError("Required identifier, nothing found")
		return ""
	}
	if p.current().Type != token.Identifier {
		p.addError(fmt.Sprintf("Required identifier, found %s", p.current().Name()))
		return ""
	}
	s := p.current().String()
	p.next()
	return s
}

func (p *Parser) validate(t ast.NodeType) {
	if p.scope != nil {
		if !p.scope.IsValid(t) {
			p.addError("Invalid declaration in scope")
		}
	}
}

func parseNewLine(p *Parser) {
	p.line++
	p.next()
}

func parseSingleLineComment(p *Parser) {
	p.parseRequired(token.LineComment)
	for p.hasTokens(1) {
		if p.current().Type != token.NewLine {
			p.next()
		} else {
			parseNewLine(p)
			return
		}

	}

}

func parseMultiLineComment(p *Parser) {
	p.parseRequired(token.CommentOpen)
	for p.hasTokens(1) && p.current().Type != token.CommentClose {
		p.next()
	}
	p.parseOptional(token.CommentClose)
}

func (p *Parser) parseModifierList() []string {
	var mods []string
	for p.hasTokens(1) && p.current().Type == token.Identifier {
		mods = append(mods, p.parseIdentifier())
	}
	return mods
}

func parseModifiers(p *Parser) {
	p.lastModifiers = p.parseModifierList()
}

func parseGroup(p *Parser) {
	if p.lastModifiers == nil {
		p.addError(errEmptyGroup)
		p.next()
	} else {
		p.modifiers = append(p.modifiers, p.lastModifiers)
		p.lastModifiers = nil
		p.parseEnclosedScope(token.OpenBracket, token.CloseBracket)
		p.modifiers = p.modifiers[:len(p.modifiers)-1]
	}
}

func (p *Parser) getModifiers() ast.Modifiers {
	var mods []string
	for _, m := range p.lastModifiers {
		mods = append(mods, m)
	}
	for _, ml := range p.modifiers {
		for _, m := range ml {
			mods = append(mods, m)
		}
	}
	return ast.Modifiers{
		Modifiers:   mods,
		Annotations: p.annotations,
	}
}

func (p *Parser) addError(message string) {
	err := util.Error{
		Message:    message,
		LineNumber: p.line,
	}
	p.errs = append(p.errs, err)
}

func (p *Parser) parseBracesScope(valids ...ast.NodeType) *ast.ScopeNode {
	return p.parseEnclosedScope(token.OpenBrace, token.CloseBrace, valids...)
}

func (p *Parser) parseEnclosedScope(opener, closer token.Type, valids ...ast.NodeType) *ast.ScopeNode {
	p.parseRequired(opener)
	scope := p.parseScope(closer, valids...)
	p.parseRequired(closer)
	return scope
}

func (p *Parser) parseScope(terminator token.Type, valids ...ast.NodeType) *ast.ScopeNode {
	scope := new(ast.ScopeNode)
	scope.Parent = p.scope
	p.scope = scope
	for p.hasTokens(1) {
		if p.current().Type == terminator {
			p.scope = scope.Parent
			return scope
		}
		found := false
		for _, c := range getPrimaryConstructs() {
			if c.is(p) {
				//fmt.Printf("FOUND: %s at index %d on line %d\n", c.name, p.index, p.line)
				c.parse(p)
				found = true
				break
			}
		}
		if !found {
			// try interpreting it as an expression
			saved := *p
			expr := p.parseExpression()
			if expr == nil {
				*p = saved
				//fmt.Printf("Unrecognised construct at index %d: %s\n", p.index, p.current().String())
				p.addError(fmt.Sprintf("Unrecognised construct: %s", p.current().String()))
				p.next()
			} else {
				switch expr.Type() {
				case ast.CallExpression:

					p.scope.AddSequential(expr)
					break
				default:
					p.addError(errDanglingExpression)
					p.next()
					break
				}

			}
		}
	}
	return scope
}

func parseAnnotation(p *Parser) {

	var a *ast.Annotation
	p.parseRequired(token.At)

	a.Name = p.parseIdentifier()

	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		a.Parameters = p.parseExpressionList()
		p.parseRequired(token.CloseBracket)
	}
	if p.annotations == nil {
		p.annotations = make([]*ast.Annotation, 0)
	}
	p.annotations = append(p.annotations, a)
}
