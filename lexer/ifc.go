package lexer

import (
	"io/ioutil"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/util"
)

// Lexer ...
type Lexer struct {
	buffer      []byte
	byteOffset  uint
	line        int
	column      int
	tokens      []token.Token
	tokenOffset int
	errors      util.Errors
}

// Lex ...
func Lex(bytes []byte) (tokens []token.Token, errs util.Errors) {
	l := new(Lexer)
	l.byteOffset = 0
	l.buffer = bytes
	l.next()
	return l.tokens, l.errors
}

// LexFile ...
func LexFile(path string) (tokens []token.Token, errs util.Errors) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, append(errs, util.Error{
			Message: "File does not exist",
		})
	}
	return Lex(bytes)
}

// LexString lexes a string
func LexString(str string) (tokens []token.Token, errs util.Errors) {
	return Lex([]byte(str))
}
