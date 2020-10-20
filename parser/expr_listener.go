// Code generated from Expr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterCompareExp is called when entering the compareExp production.
	EnterCompareExp(c *CompareExpContext)

	// EnterParenExp is called when entering the parenExp production.
	EnterParenExp(c *ParenExpContext)

	// EnterLogicalExp is called when entering the logicalExp production.
	EnterLogicalExp(c *LogicalExpContext)

	// EnterAttrPath is called when entering the attrPath production.
	EnterAttrPath(c *AttrPathContext)

	// EnterSubAttr is called when entering the subAttr production.
	EnterSubAttr(c *SubAttrContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterDouble is called when entering the double production.
	EnterDouble(c *DoubleContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterListOfIntegers is called when entering the listOfIntegers production.
	EnterListOfIntegers(c *ListOfIntegersContext)

	// EnterListOfDoubles is called when entering the listOfDoubles production.
	EnterListOfDoubles(c *ListOfDoublesContext)

	// EnterListOfStrings is called when entering the listOfStrings production.
	EnterListOfStrings(c *ListOfStringsContext)

	// EnterListStrings is called when entering the listStrings production.
	EnterListStrings(c *ListStringsContext)

	// EnterSubListOfStrings is called when entering the subListOfStrings production.
	EnterSubListOfStrings(c *SubListOfStringsContext)

	// EnterListDoubles is called when entering the listDoubles production.
	EnterListDoubles(c *ListDoublesContext)

	// EnterSubListOfDoubles is called when entering the subListOfDoubles production.
	EnterSubListOfDoubles(c *SubListOfDoublesContext)

	// EnterListIntegers is called when entering the listIntegers production.
	EnterListIntegers(c *ListIntegersContext)

	// EnterSubListOfIntegers is called when entering the subListOfIntegers production.
	EnterSubListOfIntegers(c *SubListOfIntegersContext)

	// ExitCompareExp is called when exiting the compareExp production.
	ExitCompareExp(c *CompareExpContext)

	// ExitParenExp is called when exiting the parenExp production.
	ExitParenExp(c *ParenExpContext)

	// ExitLogicalExp is called when exiting the logicalExp production.
	ExitLogicalExp(c *LogicalExpContext)

	// ExitAttrPath is called when exiting the attrPath production.
	ExitAttrPath(c *AttrPathContext)

	// ExitSubAttr is called when exiting the subAttr production.
	ExitSubAttr(c *SubAttrContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitDouble is called when exiting the double production.
	ExitDouble(c *DoubleContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitListOfIntegers is called when exiting the listOfIntegers production.
	ExitListOfIntegers(c *ListOfIntegersContext)

	// ExitListOfDoubles is called when exiting the listOfDoubles production.
	ExitListOfDoubles(c *ListOfDoublesContext)

	// ExitListOfStrings is called when exiting the listOfStrings production.
	ExitListOfStrings(c *ListOfStringsContext)

	// ExitListStrings is called when exiting the listStrings production.
	ExitListStrings(c *ListStringsContext)

	// ExitSubListOfStrings is called when exiting the subListOfStrings production.
	ExitSubListOfStrings(c *SubListOfStringsContext)

	// ExitListDoubles is called when exiting the listDoubles production.
	ExitListDoubles(c *ListDoublesContext)

	// ExitSubListOfDoubles is called when exiting the subListOfDoubles production.
	ExitSubListOfDoubles(c *SubListOfDoublesContext)

	// ExitListIntegers is called when exiting the listIntegers production.
	ExitListIntegers(c *ListIntegersContext)

	// ExitSubListOfIntegers is called when exiting the subListOfIntegers production.
	ExitSubListOfIntegers(c *SubListOfIntegersContext)
}
