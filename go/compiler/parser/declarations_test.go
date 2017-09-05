package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"

	"github.com/end-r/goutil"
)

func TestParseInterfaceDeclarationEmpty(t *testing.T) {
	p := createParser(`interface Wagable {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.AssertNow(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, i.Supers == nil, "wrong supers")
}

func TestParseInterfaceDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
	goutil.AssertNow(t, i.Supers[0].Names[0] == "Visible", "wrong supers 0 name")
}

func TestParseInterfaceDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
	goutil.AssertNow(t, i.Supers[0].Names[0] == "Visible", "wrong supers 0 name")
	goutil.AssertNow(t, i.Supers[1].Names[0] == "Movable", "wrong supers 1 name")
}

func TestParseInterfaceDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract interface Wagable {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.IsAbstract, "wrong abstract")
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationEmpty(t *testing.T) {
	p := createParser(`contract Wagable {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable is Visible inherits Object {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits A,B is Visible, Movable  {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Object is Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationMultipleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A,B {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseContractDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract contract Wagable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable is Visible inherits Object {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits A,B is Visible, Movable  {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Object is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A,B {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract class Wagable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseTypeDeclaration(t *testing.T) {
	p := createParser(`type Wagable int`)
	fmt.Println(p.lexer.Tokens)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	goutil.Assert(t, isTypeDeclaration(p), "should detect type decl")
	parseTypeDeclaration(p)
}

func TestParseExplicitVarDeclaration(t *testing.T) {
	p := createParser(`a string`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 2, "wrong token length")
	goutil.Assert(t, isExplicitVarDeclaration(p), "should detect expvar decl")
	parseExplicitVarDeclaration(p)
}

func TestParseEventDeclarationEmpty(t *testing.T) {
	p := createParser(`event Notification()`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(eventKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(eventKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 0, "wrong param length")
}

func TestParseEventDeclarationSingle(t *testing.T) {
	p := createParser(`event Notification(string)`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(eventKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(eventKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 1, "wrong param length")
}

func TestParseEventDeclarationMultiple(t *testing.T) {
	p := createParser(`event Notification(string, string)`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 7, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(eventKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(eventKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 2, "wrong param length")
}

func TestParseEnum(t *testing.T) {
	p := createParser(`enum Weekday {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(enumKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(enumKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Weekday", "wrong identifier")
}

func TestParseEnumInheritsSingle(t *testing.T) {
	p := createParser(`enum Day inherits Weekday {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(enumKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(enumKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseEnumInheritsMultiple(t *testing.T) {
	p := createParser(`enum Day inherits Weekday, Weekend {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 8, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(enumKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(enumKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseFuncNoParameters(t *testing.T) {
	p := createParser(`func foo(){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 0, "wrong param length")
}

func TestParseFuncOneParameter(t *testing.T) {
	p := createParser(`func foo(a int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 1, "wrong param length")
}

func TestParseFuncParameters(t *testing.T) {
	p := createParser(`func foo(a int, b string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 2, "wrong param length")
}

func TestParseFuncMultiplePerType(t *testing.T) {
	p := createParser(`func foo(a, b int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 1, "wrong param length")
}

func TestParseFuncMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`func foo(a, b int, c string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 2, "wrong param length")
}

func TestParseConstructorNoParameters(t *testing.T) {
	p := createParser(`constructor(){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 0, "wrong param length")
}

func TestParseConstructorOneParameter(t *testing.T) {
	p := createParser(`constructor(a int){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorParameters(t *testing.T) {
	p := createParser(`constructor(a int, b string){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
}

func TestParseConstructorMultiplePerType(t *testing.T) {
	p := createParser(`constructor(a, b int){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`constructor(a, b int, c string){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
}
