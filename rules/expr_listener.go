package rules

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linger1216/go-rules/parser"
	"strings"
)

type ExpressionListener struct {
	stack  *objStack
	obj    map[string]interface{}
	judges map[string]Judge
}

func NewExpressionListener(obj map[string]interface{}) *ExpressionListener {
	ret := &ExpressionListener{&objStack{}, obj, make(map[string]Judge)}
	ret.judges["string"] = &StringJudge{}
	ret.judges["list_string"] = &ListStringJudge{}
	ret.judges["float"] = &FloatJudge{}
	ret.judges["list_float"] = &ListFloatJudge{}
	ret.judges["integer"] = &IntegerJudge{}
	ret.judges["list_integer"] = &ListIntegerJudge{}
	ret.judges["boolean"] = &BooleanJudge{}
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
	res := e.stack.pop()
	if res.Type() != TypeBoolean {
		return false
	}
	return res.Value().(bool)
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

	typeName := ""
	switch leftValue.(type) {
	case bool:
		typeName = "boolean"
	case string:
		typeName = "string"
	case []string:
		typeName = "list_string"
	case int64, int:
		typeName = "integer"
	case []int64, []int:
		typeName = "list_integer"
	case float64:
		typeName = "float"
	case []float64:
		typeName = "list_float"
	}

	judge, ok := e.judges[typeName]
	if !ok {
		panic("not implement")
	}

	var res bool
	switch c.GetOp().GetTokenType() {
	case parser.ExprLexerEQ:
		fmt.Println("==")
		res, err = judge.EQ(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerNE:
		fmt.Println("!=")
		res, err = judge.NE(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerGT:
		fmt.Println(">")
		res, err = judge.GT(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerLT:
		fmt.Println("<")
		res, err = judge.LT(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerGE:
		fmt.Println(">=")
		res, err = judge.GE(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerLE:
		fmt.Println("<=")
		res, err = judge.LE(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerCONTAIN:
		fmt.Println("contain")
		res, err = judge.Contain(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerPREFIX:
		fmt.Println("prefix")
		res, err = judge.Prefix(leftValue, right)
		if err != nil {
			panic(err)
		}
		e.push(NewResultBoolean(res))
	case parser.ExprLexerREGEX:
		fmt.Println("regex")
		res, err = judge.Regex(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerIN:
		fmt.Println("in")
		res, err = judge.In(leftValue, right)
		if err != nil {
			panic(err)
		}
	case parser.ExprLexerSUFFIX:
		fmt.Println("suffix")
		res, err = judge.Suffix(leftValue, right)
		if err != nil {
			panic(err)
		}
	}
	e.push(NewResultBoolean(res))
}

func (e *ExpressionListener) ExitParenExp(c *parser.ParenExpContext) {
}

func (e *ExpressionListener) ExitLogicalExp(c *parser.LogicalExpContext) {
	//left, right
	fmt.Println(c.LOGICAL_OPERATOR().GetText())
	var res bool
	right, left := e.stack.pop(), e.stack.pop()
	if right.Type() == left.Type() && left.Type() == TypeBoolean {
		switch strings.ToLower(c.LOGICAL_OPERATOR().GetText()) {
		case "and":
			res = left.Value().(bool) && right.Value().(bool)
		case "or":
			res = left.Value().(bool) || right.Value().(bool)
		}
	}
	e.push(NewResultBoolean(res))
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
