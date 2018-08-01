package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wlan0/antlr-play/parser"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Please specify a file to parse\n")
		os.Exit(1)
	}

	file := os.Args[1]

	input, err := antlr.NewFileStream(file)
	if err != nil {
		fmt.Printf("Error reading file:%s", err)
	}

	lexer := parser.NewTestLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewTestParser(stream)
	p.BuildParseTrees = true
	el := &errorListener{file: file}
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	tree := p.Expression()
	if el.Error() != "" {
		fmt.Printf("%s\n", el)
	}
	fmt.Printf("%s", antlr.TreesIndentedStringTree(tree, "", nil, p))
}

type errorListener struct {
	*antlr.DefaultErrorListener

	file string
	Err  error
}

func (el *errorListener) Error() string {
	if el.Err == nil {
		return ""
	}
	return el.Err.Error()
}

func (el *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.Err = fmt.Errorf("[ERR %s:%d:%d] %s", el.file, line, column, msg)
}
