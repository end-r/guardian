package parser

import (
	"axia/guardian/go/compiler/ast"
	"testing"

	"github.com/end-r/goutil"
)

func TestParseReturnStatementSingleConstant(t *testing.T) {
	p := createParser("return 6")
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
}

func TestParseAssignmentStatementSingleConstant(t *testing.T) {
	p := createParser("x = 6")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
}

func TestParseIfStatement(t *testing.T) {
	p := createParser(`if x > 4 {

	} elif x < 4 {

	} else {

	}`)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
}

func TestParseForStatementCondition(t *testing.T) {
	p := createParser(`for x < 5 {}`)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
}

func TestParseForStatementInitCondition(t *testing.T) {
	p := createParser(`for x := 0; x < 5 {}`)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
}

func TestParseForStatementInitConditionStatement(t *testing.T) {
	p := createParser(`for x := 0; x < 5; x++ {}`)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
}

func TestParseSwitchStatement(t *testing.T) {
	p := createParser(`switch x {}`)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementSingleCase(t *testing.T) {
	p := createParser(`switch x { case 5{}}`)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementMultiCase(t *testing.T) {
	p := createParser(`switch x {
		case 5 {
			x += 2
			break
		}
		case 4{
			x *= 2
			break
		}
	}`)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementExclusive(t *testing.T) {
	p := createParser(`exclusive switch x {}
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseCaseStatementSingle(t *testing.T) {
	p := createParser(`case 5 {}`)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)
}

func TestParseCaseStatementMultiple(t *testing.T) {
	p := createParser(`case 5, 8, 9 {}`)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)
}

func TestEmptyReturnStatement(t *testing.T) {
	p := createParser("return")
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
}

func TestSingleLiteralReturnStatement(t *testing.T) {
	p := createParser(`return "twenty"`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 1, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.Literal, "wrong literal type")
}

func TestMultipleLiteralReturnStatement(t *testing.T) {
	p := createParser(`return "twenty", "thirty"`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 2, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.Literal, "wrong result 0 type")
	goutil.AssertNow(t, u.Results[1].Type() == ast.Literal, "wrong result 1 type")
}

func TestSingleReferenceReturnStatement(t *testing.T) {
	p := createParser(`return twenty`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 1, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.Reference, "wrong result 0 type")
}

func TestMultipleReferenceReturnStatement(t *testing.T) {
	p := createParser(`return twenty, thirty`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 2, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.Reference, "wrong result 0 type")
	goutil.AssertNow(t, u.Results[1].Type() == ast.Reference, "wrong result 1 type")
}

func TestSingleCallReturnStatement(t *testing.T) {
	p := createParser(`return param()`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 1, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.CallExpression, "wrong result 0 type")
}

func TestMultipleCallReturnStatement(t *testing.T) {
	p := createParser(`return a(param, "param"), b()`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 2, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.CallExpression, "wrong result 0 type")
	goutil.AssertNow(t, u.Results[1].Type() == ast.CallExpression, "wrong result 1 type")
}

func TestSingleArrayLiterallReturnStatement(t *testing.T) {
	p := createParser(`return [int]{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 1, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
}

func TestMultipleArrayLiteralReturnStatement(t *testing.T) {
	p := createParser(`return [string]{"one", "two"}, [Dog]{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 2, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
	goutil.AssertNow(t, u.Results[1].Type() == ast.ArrayLiteral, "wrong result 1 type")
}

func TestSingleMapLiterallReturnStatement(t *testing.T) {
	p := createParser(`return map[string]int{"one":2, "two":3}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 1, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
}

func TestMultipleMapLiteralReturnStatement(t *testing.T) {
	p := createParser(`return map[string]int{"one":2, "two":3}, map[int]Dog{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 2, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
	goutil.AssertNow(t, u.Results[1].Type() == ast.ArrayLiteral, "wrong result 1 type")
}

func TestSingleCompositeLiterallReturnStatement(t *testing.T) {
	p := createParser(`return Dog{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 1, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
}

func TestMultipleCompositeLiteralReturnStatement(t *testing.T) {
	p := createParser(`return Cat{name:"Doggo"}, Dog{name:"Katter"}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	u := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	u.(ast.ReturnStatementNode)
	goutil.AssertNow(t, len(u.Results) == 2, "wrong result length")
	goutil.AssertNow(t, u.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
	goutil.AssertNow(t, u.Results[1].Type() == ast.ArrayLiteral, "wrong result 1 type")
}

func TestSimpleLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x = 5")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(n.Left) == 1, "should be one left value")
	goutil.AssertNow(t, n.Left[0].Type() == ast.Reference, "wrong left type")
	l := p.scope.(ast.ReferenceNode)
	goutil.Assert(t, len(l.Names) == 1 && l.Names[0] == "x", "wrong left name")

}

func TestMultiToSingleLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x, y = 5")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be two left values")
	goutil.AssertNow(t, a.Left[0].Type() == ast.Reference, "wrong left type")
}

func TestMultiLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x, y = 5, 3")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(n.Left) == 1, "should be two left values")
}

func TestSimpleReferenceAssignmentStatement(t *testing.T) {
	p := createParser("x = a")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiToSingleReferenceAssignmentStatement(t *testing.T) {
	p := createParser("x, y = a")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiReferenceAssignmentStatement(t *testing.T) {
	p := createParser("x, y = a, b")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestSimpleCallAssignmentStatement(t *testing.T) {
	p := createParser("x = a()")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiToSingleCallAssignmentStatement(t *testing.T) {
	p := createParser("x, y = ab()")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiCallAssignmentStatement(t *testing.T) {
	p := createParser("x, y = a(), b()")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestSimpleCompositeLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x = Dog{}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiToSingleCompositeLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x, y = Dog{}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiCompositeLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x, y = Dog{}, Cat{}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestSimpleArrayLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x = [int]{3, 5}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiToSingleArrayLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x, y = [int]{3, 5}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiArrayLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x, y = [int]{1, 2}, [int]{}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestSimpleMapLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x = [int]{3, 5}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiToSingleMapLiteralAssignmentStatement(t *testing.T) {
	p := createParser("x, y = [int]{3, 5}")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
}

func TestMultiMapLiteralAssignmentStatement(t *testing.T) {
	p := createParser(`x, y = map[int]string{1:"A", 2:"B"}, map[string]int{"A":3, "B": 4}`)
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	n := p.Scope.Nodes("flow")[0]
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	n.(ast.AssignmentStatementNode)
	s
}
