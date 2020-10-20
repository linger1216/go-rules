// Code generated from Expr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompareExp is called when production compareExp is entered.
func (s *BaseExprListener) EnterCompareExp(ctx *CompareExpContext) {}

// ExitCompareExp is called when production compareExp is exited.
func (s *BaseExprListener) ExitCompareExp(ctx *CompareExpContext) {}

// EnterParenExp is called when production parenExp is entered.
func (s *BaseExprListener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production parenExp is exited.
func (s *BaseExprListener) ExitParenExp(ctx *ParenExpContext) {}

// EnterLogicalExp is called when production logicalExp is entered.
func (s *BaseExprListener) EnterLogicalExp(ctx *LogicalExpContext) {}

// ExitLogicalExp is called when production logicalExp is exited.
func (s *BaseExprListener) ExitLogicalExp(ctx *LogicalExpContext) {}

// EnterAttrPath is called when production attrPath is entered.
func (s *BaseExprListener) EnterAttrPath(ctx *AttrPathContext) {}

// ExitAttrPath is called when production attrPath is exited.
func (s *BaseExprListener) ExitAttrPath(ctx *AttrPathContext) {}

// EnterSubAttr is called when production subAttr is entered.
func (s *BaseExprListener) EnterSubAttr(ctx *SubAttrContext) {}

// ExitSubAttr is called when production subAttr is exited.
func (s *BaseExprListener) ExitSubAttr(ctx *SubAttrContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseExprListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseExprListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BaseExprListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseExprListener) ExitNull(ctx *NullContext) {}

// EnterString is called when production string is entered.
func (s *BaseExprListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseExprListener) ExitString(ctx *StringContext) {}

// EnterDouble is called when production double is entered.
func (s *BaseExprListener) EnterDouble(ctx *DoubleContext) {}

// ExitDouble is called when production double is exited.
func (s *BaseExprListener) ExitDouble(ctx *DoubleContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseExprListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseExprListener) ExitInteger(ctx *IntegerContext) {}

// EnterListOfIntegers is called when production listOfIntegers is entered.
func (s *BaseExprListener) EnterListOfIntegers(ctx *ListOfIntegersContext) {}

// ExitListOfIntegers is called when production listOfIntegers is exited.
func (s *BaseExprListener) ExitListOfIntegers(ctx *ListOfIntegersContext) {}

// EnterListOfDoubles is called when production listOfDoubles is entered.
func (s *BaseExprListener) EnterListOfDoubles(ctx *ListOfDoublesContext) {}

// ExitListOfDoubles is called when production listOfDoubles is exited.
func (s *BaseExprListener) ExitListOfDoubles(ctx *ListOfDoublesContext) {}

// EnterListOfStrings is called when production listOfStrings is entered.
func (s *BaseExprListener) EnterListOfStrings(ctx *ListOfStringsContext) {}

// ExitListOfStrings is called when production listOfStrings is exited.
func (s *BaseExprListener) ExitListOfStrings(ctx *ListOfStringsContext) {}

// EnterListStrings is called when production listStrings is entered.
func (s *BaseExprListener) EnterListStrings(ctx *ListStringsContext) {}

// ExitListStrings is called when production listStrings is exited.
func (s *BaseExprListener) ExitListStrings(ctx *ListStringsContext) {}

// EnterSubListOfStrings is called when production subListOfStrings is entered.
func (s *BaseExprListener) EnterSubListOfStrings(ctx *SubListOfStringsContext) {}

// ExitSubListOfStrings is called when production subListOfStrings is exited.
func (s *BaseExprListener) ExitSubListOfStrings(ctx *SubListOfStringsContext) {}

// EnterListDoubles is called when production listDoubles is entered.
func (s *BaseExprListener) EnterListDoubles(ctx *ListDoublesContext) {}

// ExitListDoubles is called when production listDoubles is exited.
func (s *BaseExprListener) ExitListDoubles(ctx *ListDoublesContext) {}

// EnterSubListOfDoubles is called when production subListOfDoubles is entered.
func (s *BaseExprListener) EnterSubListOfDoubles(ctx *SubListOfDoublesContext) {}

// ExitSubListOfDoubles is called when production subListOfDoubles is exited.
func (s *BaseExprListener) ExitSubListOfDoubles(ctx *SubListOfDoublesContext) {}

// EnterListIntegers is called when production listIntegers is entered.
func (s *BaseExprListener) EnterListIntegers(ctx *ListIntegersContext) {}

// ExitListIntegers is called when production listIntegers is exited.
func (s *BaseExprListener) ExitListIntegers(ctx *ListIntegersContext) {}

// EnterSubListOfIntegers is called when production subListOfIntegers is entered.
func (s *BaseExprListener) EnterSubListOfIntegers(ctx *SubListOfIntegersContext) {}

// ExitSubListOfIntegers is called when production subListOfIntegers is exited.
func (s *BaseExprListener) ExitSubListOfIntegers(ctx *SubListOfIntegersContext) {}
