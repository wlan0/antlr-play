// Code generated from Test.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Test

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTestListener is a complete listener for a parse tree produced by TestParser.
type BaseTestListener struct{}

var _ TestListener = &BaseTestListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTestListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTestListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTestListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTestListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTestListener) ExitExpression(ctx *ExpressionContext) {}
