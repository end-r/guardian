package validator

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// NumericType... used to pass types into the s
type NumericType struct {
	size    int
	name    string
	signed  bool
	integer bool
}

type BooleanType struct {
}

func (v *Validator) validateBuiltinDeclarations(scope *ast.ScopeNode) {
	if scope.Declarations != nil {
		// order doesn't matter here
		for _, i := range scope.Declarations.Map() {
			// add in placeholders for all declarations
			v.validateDeclaration(i.(ast.Node))
		}
	}
}

func (v *Validator) validateBuiltinSequence(scope *ast.ScopeNode) {
	for _, node := range scope.Sequence {
		v.validate(node)
	}
}

type LiteralFunc func(*Validator, string) Type
type LiteralMap map[lexer.TokenType]LiteralFunc

func SimpleLiteral(typeName string) LiteralFunc {
	return func(v *Validator, data string) Type {
		return v.getNamedType(typeName)
	}
}

type OperatorFunc func(*Validator, ...Type) Type
type OperatorMap map[lexer.TokenType]OperatorFunc

func (m OperatorMap) Add(function OperatorFunc, types ...lexer.TokenType) {
	for _, t := range types {

		m[t] = function
	}
}

func SimpleOperator(typeName string) OperatorFunc {
	return func(v *Validator, types ...Type) Type {
		return v.getNamedType(typeName)
	}
}

func BinaryNumericOperator() OperatorFunc {
	return func(v *Validator, ts ...Type) Type {
		left := resolveUnderlying(ts[0])
		right := resolveUnderlying(ts[1])
		if na, ok := left.(NumericType); ok {
			if nb, ok := right.(NumericType); ok {
				if na.size > nb.size {
					return v.smallestNumericType(na.size, true)
				}
				return v.smallestNumericType(nb.size, true)
			}
		}
		return standards[Invalid]
	}
}

func BinaryIntegerOperator() OperatorFunc {
	return func(v *Validator, ts ...Type) Type {
		if na, ok := ts[0].(NumericType); ok && na.integer {
			if nb, ok := ts[1].(NumericType); ok && nb.integer {
				if na.size > nb.size {
					return v.smallestNumericType(na.size, false)
				}
				return v.smallestNumericType(nb.size, false)
			}
		}
		return standards[Invalid]
	}
}

func (v *Validator) smallestNumericType(bits int, allowFloat bool) Type {
	smallest := -1
	smallestType := Type(standards[Unknown])
	for _, typ := range v.primitives {
		n, ok := typ.(NumericType)
		if ok {
			if !n.integer && !allowFloat {
				continue
			}
			if smallest == -1 || n.size < smallest {
				if n.size >= bits {
					if n.signed {
						// TODO: only select signed?
						smallest = n.size
						smallestType = n
					}

				}
			}
		}
	}
	return smallestType
}

// can probs use logs here
func BitsNeeded(x int) int {
	if x == 0 {
		return 1
	}
	count := 0
	for x > 0 {
		count++
		x = x >> 1
	}
	return count
}