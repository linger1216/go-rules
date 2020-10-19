package rules

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linger1216/go-rules/parser"
)

type ExpressionListener struct {
	stack *objStack
	obj   map[string]interface{}
}

func NewExpressionListener(obj map[string]interface{}) *ExpressionListener {
	ret := &ExpressionListener{&objStack{}, obj}
	return ret
}

func (e *ExpressionListener) push(i IFiled) {
	if !e.stack.empty() && i.Type() == TypeAttr {
		switch e.stack.peek().Type() {
		case TypeAttr:
			e.stack.pop()
		}
	}
	e.stack.push(i)
}

func (e *ExpressionListener) Result() bool {
	e.stack.pop()
	return false
}

func (e *ExpressionListener) VisitTerminal(node antlr.TerminalNode) {
}

func (e *ExpressionListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (e *ExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (e *ExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (e *ExpressionListener) EnterCompareExp(c *parser.CompareExpContext) {
}

func (e *ExpressionListener) EnterParenExp(c *parser.ParenExpContext) {
}

func (e *ExpressionListener) EnterLogicalExp(c *parser.LogicalExpContext) {
}

func (e *ExpressionListener) EnterAttrPath(c *parser.AttrPathContext) {
}

func (e *ExpressionListener) EnterSubAttr(c *parser.SubAttrContext) {
}

func (e *ExpressionListener) EnterBoolean(c *parser.BooleanContext) {
}

func (e *ExpressionListener) EnterNull(c *parser.NullContext) {
}

func (e *ExpressionListener) EnterString(c *parser.StringContext) {
}

func (e *ExpressionListener) EnterDouble(c *parser.DoubleContext) {
}

func (e *ExpressionListener) EnterInteger(c *parser.IntegerContext) {
}

func (e *ExpressionListener) EnterListOfIntegers(c *parser.ListOfIntegersContext) {
}

func (e *ExpressionListener) EnterListOfDoubles(c *parser.ListOfDoublesContext) {
}

func (e *ExpressionListener) EnterListOfStrings(c *parser.ListOfStringsContext) {
}

func (e *ExpressionListener) EnterListStrings(c *parser.ListStringsContext) {
}

func (e *ExpressionListener) EnterSubListOfStrings(c *parser.SubListOfStringsContext) {
}

func (e *ExpressionListener) EnterListDoubles(c *parser.ListDoublesContext) {
}

func (e *ExpressionListener) EnterSubListOfDoubles(c *parser.SubListOfDoublesContext) {
}

func (e *ExpressionListener) EnterListIntegers(c *parser.ListIntegersContext) {
}

func (e *ExpressionListener) EnterSubListOfIntegers(c *parser.SubListOfIntegersContext) {
}

func (e *ExpressionListener) ExitCompareExp(c *parser.CompareExpContext) {
	right, left := e.stack.pop(), e.stack.pop()

	var err error
	leftValue := left.Value()
	if left.Type() == TypeAttr {
		leftValue, err = GetProperty(e.obj, left.Value().(string))
		if err != nil {
			panic(err)
		}
	}

	_ = leftValue
	_ = right

	//
	//switch c.GetOp().GetTokenType() {
	//case parser.ExprLexerEQ:
	//	fmt.Println("==")
	//	res, err := judgeEQ(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerNE:
	//	fmt.Println("!=")
	//	res, err := judgeNE(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerGT:
	//	fmt.Println(">")
	//	res, err := judgeGT(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerLT:
	//	fmt.Println("<")
	//	res, err := judgeLT(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerGE:
	//	fmt.Println(">=")
	//	res, err := judgeGE(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerLE:
	//	fmt.Println("<=")
	//	res, err := judgeLE(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerCONTAIN:
	//	fmt.Println("contain")
	//	res, err := judgeContain(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerPREFIX:
	//	fmt.Println("prefix")
	//	res, err := judgePrefix(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerREGEX:
	//	fmt.Println("regex")
	//	fmt.Println("prefix")
	//	res, err := judgeRegex(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerIN:
	//	fmt.Println("in")
	//	fmt.Println("prefix")
	//	res, err := judgeIn(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//case parser.ExprLexerSUFFIX:
	//	fmt.Println("suffix")
	//	fmt.Println("prefix")
	//	res, err := judgeSuffix(leftValue, right)
	//	if err != nil {
	//		panic(err)
	//	}
	//	e.push(NewResultBoolean(res))
	//}
}

func (e *ExpressionListener) ExitParenExp(c *parser.ParenExpContext) {
}

func (e *ExpressionListener) ExitLogicalExp(c *parser.LogicalExpContext) {
	//left, right
	fmt.Println(c.LOGICAL_OPERATOR().GetText())
	e.push(NewLogical(c.LOGICAL_OPERATOR().GetText()))
}

func (e *ExpressionListener) ExitAttrPath(c *parser.AttrPathContext) {
	fmt.Println("ExitAttrPath:", c.GetText())
	e.push(NewAttr(c.GetText()))
}

func (e *ExpressionListener) ExitSubAttr(c *parser.SubAttrContext) {
}

func (e *ExpressionListener) ExitBoolean(c *parser.BooleanContext) {
	fmt.Println("ExitBoolean:", c.GetText())
	e.push(NewValueBoolean(c.GetText()))
}

func (e *ExpressionListener) ExitNull(c *parser.NullContext) {
	fmt.Println("ExitNull:", c.GetText())
	e.push(NewValueNull(c.GetText()))
}

func (e *ExpressionListener) ExitString(c *parser.StringContext) {
	fmt.Println("ExitString:", c.GetText())
	e.push(NewValueString(c.GetText()))
}

func (e *ExpressionListener) ExitDouble(c *parser.DoubleContext) {
	fmt.Println("ExitDouble:", c.GetText())
	e.push(NewValueDouble(c.GetText()))
}

func (e *ExpressionListener) ExitInteger(c *parser.IntegerContext) {
	fmt.Println("ExitInteger:", c.GetText())
	e.push(NewValueInteger(c.GetText()))
}

func (e *ExpressionListener) ExitListStrings(c *parser.ListStringsContext) {
	fmt.Println("ExitListStrings:", c.GetText())
	e.push(NewValueListStrings(c.GetText()))
}

func (e *ExpressionListener) ExitListDoubles(c *parser.ListDoublesContext) {
	fmt.Println("ExitListDoubles:", c.GetText())
	e.push(NewValueListDoubles(c.GetText()))
}

func (e *ExpressionListener) ExitListIntegers(c *parser.ListIntegersContext) {
	fmt.Println("ExitListIntegers:", c.GetText())
	e.push(NewValueListIntegers(c.GetText()))
}

func (e *ExpressionListener) ExitListOfIntegers(c *parser.ListOfIntegersContext) {
}

func (e *ExpressionListener) ExitListOfDoubles(c *parser.ListOfDoublesContext) {
}

func (e *ExpressionListener) ExitListOfStrings(c *parser.ListOfStringsContext) {
}

func (e *ExpressionListener) ExitSubListOfStrings(c *parser.SubListOfStringsContext) {
}

func (e *ExpressionListener) ExitSubListOfDoubles(c *parser.SubListOfDoublesContext) {
}

func (e *ExpressionListener) ExitSubListOfIntegers(c *parser.SubListOfIntegersContext) {
}
