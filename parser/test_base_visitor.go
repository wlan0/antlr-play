// Code generated from Test.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Test

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTestVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTestVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
