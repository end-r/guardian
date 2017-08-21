package ast

import "github.com/end-r/vmgen"

type TypeDeclarationNode struct {
	Identifier string
	Value      ReferenceNode
}

// Type ...
func (n TypeDeclarationNode) Type() NodeType { return TypeDeclaration }

func (n TypeDeclarationNode) Validate(t NodeType) bool {
	return true
}

func (n TypeDeclarationNode) Declare(key string, node Node) {

}

func (n TypeDeclarationNode) Traverse(vm *vmgen.VM) {

}

type FuncDeclarationNode struct {
	Identifier string
	Parameters []Node
	Results    []Node
	IsAbstract bool
}

func (n FuncDeclarationNode) Type() NodeType { return FuncDeclaration }

func (n FuncDeclarationNode) Validate(t NodeType) bool {
	return true
}

func (n FuncDeclarationNode) Declare(key string, node Node) {

}

func (n FuncDeclarationNode) Traverse(vm *vmgen.VM) {

}

type ClassDeclarationNode struct {
	Identifier   string
	IsAbstract   bool
	Supers       []ReferenceNode
	Interfaces   []ReferenceNode
	declarations map[string][]Node
}

func (n ClassDeclarationNode) Type() NodeType { return ClassDeclaration }

func (n ClassDeclarationNode) Validate(t NodeType) bool {
	switch t {
	case ClassDeclaration, InterfaceDeclaration, ContractDeclaration:
		return true
	}
	return false
}

func (n ClassDeclarationNode) Declare(key string, node Node) {

}

func (n ClassDeclarationNode) Traverse(vm *vmgen.VM) {

}

type InterfaceDeclarationNode struct {
	Identifier   string
	IsAbstract   bool
	Declarations []Node
	Supers       []ReferenceNode
}

func (n InterfaceDeclarationNode) Type() NodeType { return InterfaceDeclaration }

func (n InterfaceDeclarationNode) Validate(t NodeType) bool {
	switch t {
	case FuncDeclaration:
		return true
	}
	return false
}

func (n InterfaceDeclarationNode) Declare(key string, node Node) {
	n.Declarations = append(n.Declarations, node)
}

func (n InterfaceDeclarationNode) Traverse(vm *vmgen.VM) {

}

type ContractDeclarationNode struct {
	Identifier   string
	IsAbstract   bool
	Supers       []ReferenceNode
	Declarations map[string][]Node
}

func (n ContractDeclarationNode) Type() NodeType { return ContractDeclaration }

func (n ContractDeclarationNode) Validate(t NodeType) bool {
	switch t {
	// allowable declarations
	case InterfaceDeclaration, ClassDeclaration, FuncDeclaration:
		return true
	}
	return false
}
func (n ContractDeclarationNode) Declare(key string, node Node) {
	if n.Declarations == nil {
		n.Declarations = make(map[string][]Node)
	}
	n.Declarations[key] = append(n.Declarations[key], node)
}

func (n ContractDeclarationNode) Traverse(vm *vmgen.VM) {
	// find our variable declarations
	for _, decl := range n.Declarations["field"] {
		decl.Type()
		// set key
		bytes := []byte("decl")
		params := make([]byte, 0)
		params = append(params, byte(len(bytes)))
		params = append(params, bytes...)
		vm.AddBytecode("PUSH", params...)
		// set value (default value)
		vm.AddBytecode("PUSH")
		vm.AddBytecode("STORE")
	}
}

type MapTypeNode struct {
	Key   Node
	Value Node
}

func (n MapTypeNode) Type() NodeType { return MapType }

func (n MapTypeNode) Validate(t NodeType) bool {
	return true
}

func (n MapTypeNode) Declare(key string, node Node) {

}

func (n MapTypeNode) Traverse(vm *vmgen.VM) {

}

type ArrayTypeNode struct {
	Value Node
}

func (n ArrayTypeNode) Type() NodeType { return ArrayType }

func (n ArrayTypeNode) Validate(t NodeType) bool {
	return true
}

func (n ArrayTypeNode) Declare(key string, node Node) {

}

func (n ArrayTypeNode) Traverse(vm *vmgen.VM) {

}

type ParameterNode struct {
}
