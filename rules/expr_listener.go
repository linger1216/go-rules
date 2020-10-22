package rules

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linger1216/go-rules/parser"
	"strings"
)

type ExprListener struct {
	stack  *objStack
	obj    map[string]interface{}
	judges map[string]Judge
	err    error
}

func NewExprListener(obj map[string]interface{}) *ExprListener {
	ret := &ExprListener{&objStack{}, obj, make(map[string]Judge), nil}
	ret.judges["string"] = &StringJudge{}
	ret.judges["list_string"] = &ListStringJudge{}
	ret.judges["float"] = &FloatJudge{}
	ret.judges["list_float"] = &ListFloatJudge{}
	ret.judges["integer"] = &IntegerJudge{}
	ret.judges["list_integer"] = &ListIntegerJudge{}
	ret.judges["boolean"] = &BooleanJudge{}
	return ret
}

func (e *ExprListener) push(i IFiled) {
	if !e.stack.empty() && i.Type() == TypeAttr {
		switch e.stack.peek().Type() {
		case TypeAttr:
			e.stack.pop()
		}
	}
	e.stack.push(i)
}

func (e *ExprListener) Result() (bool, error) {

	if e.err != nil {
		return false, e.err
	}

	if !e.stack.empty() {
		res := e.stack.pop()
		if res.Type() != TypeResultBoolean {
			return false, ErrInvalidType
		}
		return res.Value().(bool), nil
	}
	return false, ErrInvalidType
}

func (e *ExprListener) VisitTerminal(node antlr.TerminalNode) {
}

func (e *ExprListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (e *ExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (e *ExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (e *ExprListener) EnterCompareExp(c *parser.CompareExpContext) {
}

func (e *ExprListener) EnterParenExp(c *parser.ParenExpContext) {
}

func (e *ExprListener) EnterLogicalExp(c *parser.LogicalExpContext) {
}

func (e *ExprListener) EnterAttrPath(c *parser.AttrPathContext) {
}

func (e *ExprListener) EnterSubAttr(c *parser.SubAttrContext) {
}

func (e *ExprListener) EnterBoolean(c *parser.BooleanContext) {
}

func (e *ExprListener) EnterNull(c *parser.NullContext) {
}

func (e *ExprListener) EnterString(c *parser.StringContext) {
}

func (e *ExprListener) EnterDouble(c *parser.DoubleContext) {
}

func (e *ExprListener) EnterInteger(c *parser.IntegerContext) {
}

func (e *ExprListener) EnterListOfIntegers(c *parser.ListOfIntegersContext) {
}

func (e *ExprListener) EnterListOfDoubles(c *parser.ListOfDoublesContext) {
}

func (e *ExprListener) EnterListOfStrings(c *parser.ListOfStringsContext) {
}

func (e *ExprListener) EnterListStrings(c *parser.ListStringsContext) {
}

func (e *ExprListener) EnterSubListOfStrings(c *parser.SubListOfStringsContext) {
}

func (e *ExprListener) EnterListDoubles(c *parser.ListDoublesContext) {
}

func (e *ExprListener) EnterSubListOfDoubles(c *parser.SubListOfDoublesContext) {
}

func (e *ExprListener) EnterListIntegers(c *parser.ListIntegersContext) {
}

func (e *ExprListener) EnterSubListOfIntegers(c *parser.SubListOfIntegersContext) {
}

func (e *ExprListener) ExitCompareExp(c *parser.CompareExpContext) {
	if e.stack.empty() {
		e.err = ErrInvalidRule
	}

	if e.err != nil {
		return
	}

	right, left := e.stack.pop(), e.stack.pop()

	var err error
	leftValue := left.Value()
	if left.Type() == TypeAttr {
		leftValue, err = GetProperty(e.obj, left.Value().(string))
		if err != nil {
			e.err = err
			return
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
		e.err = ErrInvalidOperator
		return
	}

	var res bool
	switch c.GetOp().GetTokenType() {
	case parser.ExprLexerEQ:
		//fmt.Println("==")
		res, err = judge.EQ(leftValue, right)
	case parser.ExprLexerNE:
		//fmt.Println("!=")
		res, err = judge.NE(leftValue, right)
	case parser.ExprLexerGT:
		//fmt.Println(">")
		res, err = judge.GT(leftValue, right)
	case parser.ExprLexerLT:
		//fmt.Println("<")
		res, err = judge.LT(leftValue, right)
	case parser.ExprLexerGE:
		//fmt.Println(">=")
		res, err = judge.GE(leftValue, right)
	case parser.ExprLexerLE:
		//fmt.Println("<=")
		res, err = judge.LE(leftValue, right)
	case parser.ExprLexerCONTAIN:
		//fmt.Println("contain")
		res, err = judge.Contain(leftValue, right)
	case parser.ExprLexerPREFIX:
		//fmt.Println("prefix")
		res, err = judge.Prefix(leftValue, right)
	case parser.ExprLexerREGEX:
		//fmt.Println("regex")
		res, err = judge.Regex(leftValue, right)
	case parser.ExprLexerIN:
		//fmt.Println("in")
		res, err = judge.In(leftValue, right)
	case parser.ExprLexerSUFFIX:
		//fmt.Println("suffix")
		res, err = judge.Suffix(leftValue, right)
	}

	if err != nil {
		e.err = err
		return
	}

	e.push(NewResultBoolean(res))
}

func (e *ExprListener) ExitParenExp(c *parser.ParenExpContext) {
	fmt.Println("ExitParenExp:", c.NOT().GetText())

	if e.stack.empty() {
		e.err = ErrInvalidRule
	}

	if e.err != nil {
		return
	}

	ele := e.stack.pop()
	boolean := !ele.Value().(bool)
	e.stack.push(NewResultBoolean(boolean))
}

func (e *ExprListener) ExitLogicalExp(c *parser.LogicalExpContext) {

	if e.stack.empty() {
		e.err = ErrInvalidRule
	}

	if e.err != nil {
		return
	}

	//fmt.Println(c.LOGICAL_OPERATOR().GetText())
	var res bool
	right, left := e.stack.pop(), e.stack.pop()
	if right.Type() == left.Type() && left.Type() == TypeResultBoolean {
		switch strings.ToLower(c.LOGICAL_OPERATOR().GetText()) {
		case "and":
			res = left.Value().(bool) && right.Value().(bool)
		case "or":
			res = left.Value().(bool) || right.Value().(bool)
		}
	}
	e.push(NewResultBoolean(res))
}

func (e *ExprListener) ExitAttrPath(c *parser.AttrPathContext) {
	//fmt.Println("ExitAttrPath:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewAttr(c.GetText()))
}

func (e *ExprListener) ExitSubAttr(c *parser.SubAttrContext) {
}

func (e *ExprListener) ExitBoolean(c *parser.BooleanContext) {
	//fmt.Println("ExitBoolean:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueBoolean(c.GetText()))
}

func (e *ExprListener) ExitNull(c *parser.NullContext) {
	//fmt.Println("ExitNull:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueNull())
}

func (e *ExprListener) ExitString(c *parser.StringContext) {
	//fmt.Println("ExitString:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueString(c.GetText()))
}

func (e *ExprListener) ExitDouble(c *parser.DoubleContext) {
	//fmt.Println("ExitDouble:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueFloat(c.GetText()))
}

func (e *ExprListener) ExitInteger(c *parser.IntegerContext) {
	//fmt.Println("ExitInteger:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueInteger(c.GetText()))
}

func (e *ExprListener) ExitListStrings(c *parser.ListStringsContext) {
	//fmt.Println("ExitListStrings:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueListStrings(c.GetText()))
}

func (e *ExprListener) ExitListDoubles(c *parser.ListDoublesContext) {
	//fmt.Println("ExitListDoubles:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueListFloats(c.GetText()))
}

func (e *ExprListener) ExitListIntegers(c *parser.ListIntegersContext) {
	//fmt.Println("ExitListIntegers:", c.GetText())
	if e.err != nil {
		return
	}
	e.push(NewValueListIntegers(c.GetText()))
}

func (e *ExprListener) ExitListOfIntegers(c *parser.ListOfIntegersContext) {
}

func (e *ExprListener) ExitListOfDoubles(c *parser.ListOfDoublesContext) {
}

func (e *ExprListener) ExitListOfStrings(c *parser.ListOfStringsContext) {
}

func (e *ExprListener) ExitSubListOfStrings(c *parser.SubListOfStringsContext) {
}

func (e *ExprListener) ExitSubListOfDoubles(c *parser.SubListOfDoublesContext) {
}

func (e *ExprListener) ExitSubListOfIntegers(c *parser.SubListOfIntegersContext) {
}
