package ast

// Node interface for storage in AST
type Node interface {
	Type() NodeType
	Validate(NodeType) bool
	Declare(string, Node)
}

type FileNode struct {
	name string
}

func (n FileNode) Type() NodeType { return File }

func (n FileNode) Validate(t NodeType) bool {
	return true
}
func (n FileNode) Declare(key string, node Node) {

}

type PackageNode struct {
	name string
}

func (n PackageNode) Type() NodeType { return File }

func (n PackageNode) Validate(t NodeType) bool {
	return true
}
func (n PackageNode) Declare(key string, node Node) {

}

type ProgramNode struct {
}

func (n ProgramNode) Type() NodeType { return File }

func (n ProgramNode) Validate(t NodeType) bool {
	return true
}
func (n ProgramNode) Declare(key string, node Node) {

}