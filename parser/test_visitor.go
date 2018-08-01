// Code generated from Test.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Test

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TestParser.
type TestVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TestParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}
}
