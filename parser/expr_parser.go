// Code generated from Expr.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Expr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 105,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 5, 2, 25,
	10, 2, 3, 2, 5, 2, 28, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 5, 2, 40, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2,
	47, 10, 2, 12, 2, 14, 2, 50, 11, 2, 3, 3, 3, 3, 5, 3, 54, 10, 3, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 64, 10, 5, 3, 5, 3, 5,
	5, 5, 68, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 83, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 5, 9, 93, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 103, 10, 11, 3, 11, 2, 3, 2, 12, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 2, 3, 3, 2, 13, 23, 2, 111, 2, 39, 3, 2, 2, 2,
	4, 51, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2, 8, 72, 3, 2, 2, 2, 10, 74, 3, 2,
	2, 2, 12, 82, 3, 2, 2, 2, 14, 84, 3, 2, 2, 2, 16, 92, 3, 2, 2, 2, 18, 94,
	3, 2, 2, 2, 20, 102, 3, 2, 2, 2, 22, 24, 8, 2, 1, 2, 23, 25, 7, 9, 2, 2,
	24, 23, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 28, 7,
	32, 2, 2, 27, 26, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29,
	30, 7, 3, 2, 2, 30, 31, 5, 2, 2, 2, 31, 32, 7, 4, 2, 2, 32, 40, 3, 2, 2,
	2, 33, 34, 5, 4, 3, 2, 34, 35, 7, 32, 2, 2, 35, 36, 9, 2, 2, 2, 36, 37,
	7, 32, 2, 2, 37, 38, 5, 8, 5, 2, 38, 40, 3, 2, 2, 2, 39, 22, 3, 2, 2, 2,
	39, 33, 3, 2, 2, 2, 40, 48, 3, 2, 2, 2, 41, 42, 12, 4, 2, 2, 42, 43, 7,
	32, 2, 2, 43, 44, 7, 10, 2, 2, 44, 45, 7, 32, 2, 2, 45, 47, 5, 2, 2, 5,
	46, 41, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3,
	2, 2, 2, 49, 3, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 53, 7, 24, 2, 2, 52,
	54, 5, 6, 4, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 5, 3, 2, 2,
	2, 55, 56, 7, 5, 2, 2, 56, 57, 5, 4, 3, 2, 57, 7, 3, 2, 2, 2, 58, 73, 7,
	11, 2, 2, 59, 73, 7, 12, 2, 2, 60, 73, 7, 26, 2, 2, 61, 73, 7, 27, 2, 2,
	62, 64, 7, 6, 2, 2, 63, 62, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 3,
	2, 2, 2, 65, 67, 7, 28, 2, 2, 66, 68, 7, 29, 2, 2, 67, 66, 3, 2, 2, 2,
	67, 68, 3, 2, 2, 2, 68, 73, 3, 2, 2, 2, 69, 73, 5, 18, 10, 2, 70, 73, 5,
	14, 8, 2, 71, 73, 5, 10, 6, 2, 72, 58, 3, 2, 2, 2, 72, 59, 3, 2, 2, 2,
	72, 60, 3, 2, 2, 2, 72, 61, 3, 2, 2, 2, 72, 63, 3, 2, 2, 2, 72, 69, 3,
	2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 71, 3, 2, 2, 2, 73, 9, 3, 2, 2, 2, 74,
	75, 7, 7, 2, 2, 75, 76, 5, 12, 7, 2, 76, 11, 3, 2, 2, 2, 77, 78, 7, 26,
	2, 2, 78, 79, 7, 31, 2, 2, 79, 83, 5, 12, 7, 2, 80, 81, 7, 26, 2, 2, 81,
	83, 7, 8, 2, 2, 82, 77, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 13, 3, 2, 2,
	2, 84, 85, 7, 7, 2, 2, 85, 86, 5, 16, 9, 2, 86, 15, 3, 2, 2, 2, 87, 88,
	7, 27, 2, 2, 88, 89, 7, 31, 2, 2, 89, 93, 5, 16, 9, 2, 90, 91, 7, 27, 2,
	2, 91, 93, 7, 8, 2, 2, 92, 87, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 17,
	3, 2, 2, 2, 94, 95, 7, 7, 2, 2, 95, 96, 5, 20, 11, 2, 96, 19, 3, 2, 2,
	2, 97, 98, 7, 28, 2, 2, 98, 99, 7, 31, 2, 2, 99, 103, 5, 20, 11, 2, 100,
	101, 7, 28, 2, 2, 101, 103, 7, 8, 2, 2, 102, 97, 3, 2, 2, 2, 102, 100,
	3, 2, 2, 2, 103, 21, 3, 2, 2, 2, 13, 24, 27, 39, 48, 53, 63, 67, 72, 82,
	92, 102,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'.'", "'-'", "'['", "']'", "", "", "", "", "", "'=='",
	"'!='", "'>'", "'<'", "'>='", "'<='", "'contain'", "'prefix'", "'suffix'",
	"'regex'", "", "", "", "", "", "", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "NOT", "LOGICAL_OPERATOR", "BOOLEAN", "NULL",
	"IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CONTAIN", "PREFIX", "SUFFIX",
	"REGEX", "ATTRNAME", "INDEX", "STRING", "DOUBLE", "INT", "EXP", "NEWLINE",
	"COMMA", "SP",
}

var ruleNames = []string{
	"query", "attrPath", "subAttr", "value", "listStrings", "subListOfStrings",
	"listDoubles", "subListOfDoubles", "listIntegers", "subListOfIntegers",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ExprParser struct {
	*antlr.BaseParser
}

func NewExprParser(input antlr.TokenStream) *ExprParser {
	this := new(ExprParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expr.g4"

	return this
}

// ExprParser tokens.
const (
	ExprParserEOF              = antlr.TokenEOF
	ExprParserT__0             = 1
	ExprParserT__1             = 2
	ExprParserT__2             = 3
	ExprParserT__3             = 4
	ExprParserT__4             = 5
	ExprParserT__5             = 6
	ExprParserNOT              = 7
	ExprParserLOGICAL_OPERATOR = 8
	ExprParserBOOLEAN          = 9
	ExprParserNULL             = 10
	ExprParserIN               = 11
	ExprParserEQ               = 12
	ExprParserNE               = 13
	ExprParserGT               = 14
	ExprParserLT               = 15
	ExprParserGE               = 16
	ExprParserLE               = 17
	ExprParserCONTAIN          = 18
	ExprParserPREFIX           = 19
	ExprParserSUFFIX           = 20
	ExprParserREGEX            = 21
	ExprParserATTRNAME         = 22
	ExprParserINDEX            = 23
	ExprParserSTRING           = 24
	ExprParserDOUBLE           = 25
	ExprParserINT              = 26
	ExprParserEXP              = 27
	ExprParserNEWLINE          = 28
	ExprParserCOMMA            = 29
	ExprParserSP               = 30
)

// ExprParser rules.
const (
	ExprParserRULE_query             = 0
	ExprParserRULE_attrPath          = 1
	ExprParserRULE_subAttr           = 2
	ExprParserRULE_value             = 3
	ExprParserRULE_listStrings       = 4
	ExprParserRULE_subListOfStrings  = 5
	ExprParserRULE_listDoubles       = 6
	ExprParserRULE_subListOfDoubles  = 7
	ExprParserRULE_listIntegers      = 8
	ExprParserRULE_subListOfIntegers = 9
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyFrom(ctx *QueryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(ExprParserSP)
}

func (s *CompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserSP, i)
}

func (s *CompareExpContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExprParserEQ, 0)
}

func (s *CompareExpContext) NE() antlr.TerminalNode {
	return s.GetToken(ExprParserNE, 0)
}

func (s *CompareExpContext) GT() antlr.TerminalNode {
	return s.GetToken(ExprParserGT, 0)
}

func (s *CompareExpContext) LT() antlr.TerminalNode {
	return s.GetToken(ExprParserLT, 0)
}

func (s *CompareExpContext) GE() antlr.TerminalNode {
	return s.GetToken(ExprParserGE, 0)
}

func (s *CompareExpContext) LE() antlr.TerminalNode {
	return s.GetToken(ExprParserLE, 0)
}

func (s *CompareExpContext) CONTAIN() antlr.TerminalNode {
	return s.GetToken(ExprParserCONTAIN, 0)
}

func (s *CompareExpContext) PREFIX() antlr.TerminalNode {
	return s.GetToken(ExprParserPREFIX, 0)
}

func (s *CompareExpContext) SUFFIX() antlr.TerminalNode {
	return s.GetToken(ExprParserSUFFIX, 0)
}

func (s *CompareExpContext) IN() antlr.TerminalNode {
	return s.GetToken(ExprParserIN, 0)
}

func (s *CompareExpContext) REGEX() antlr.TerminalNode {
	return s.GetToken(ExprParserREGEX, 0)
}

func (s *CompareExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterCompareExp(s)
	}
}

func (s *CompareExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitCompareExp(s)
	}
}

type ParenExpContext struct {
	*QueryContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *ParenExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExprParserNOT, 0)
}

func (s *ParenExpContext) SP() antlr.TerminalNode {
	return s.GetToken(ExprParserSP, 0)
}

func (s *ParenExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterParenExp(s)
	}
}

func (s *ParenExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitParenExp(s)
	}
}

type LogicalExpContext struct {
	*QueryContext
}

func NewLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpContext {
	var p = new(LogicalExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *LogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *LogicalExpContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *LogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(ExprParserSP)
}

func (s *LogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserSP, i)
}

func (s *LogicalExpContext) LOGICAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(ExprParserLOGICAL_OPERATOR, 0)
}

func (s *LogicalExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterLogicalExp(s)
	}
}

func (s *LogicalExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitLogicalExp(s)
	}
}

func (p *ExprParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *ExprParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ExprParserRULE_query, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserT__0, ExprParserNOT, ExprParserSP:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ExprParserNOT {
			{
				p.SetState(21)
				p.Match(ExprParserNOT)
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ExprParserSP {
			{
				p.SetState(24)
				p.Match(ExprParserSP)
			}

		}
		{
			p.SetState(27)
			p.Match(ExprParserT__0)
		}
		{
			p.SetState(28)
			p.query(0)
		}
		{
			p.SetState(29)
			p.Match(ExprParserT__1)
		}

	case ExprParserATTRNAME:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.AttrPath()
		}
		{
			p.SetState(32)
			p.Match(ExprParserSP)
		}
		{
			p.SetState(33)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExprParserIN)|(1<<ExprParserEQ)|(1<<ExprParserNE)|(1<<ExprParserGT)|(1<<ExprParserLT)|(1<<ExprParserGE)|(1<<ExprParserLE)|(1<<ExprParserCONTAIN)|(1<<ExprParserPREFIX)|(1<<ExprParserSUFFIX)|(1<<ExprParserREGEX))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(34)
			p.Match(ExprParserSP)
		}
		{
			p.SetState(35)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_query)
			p.SetState(39)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(40)
				p.Match(ExprParserSP)
			}
			{
				p.SetState(41)
				p.Match(ExprParserLOGICAL_OPERATOR)
			}
			{
				p.SetState(42)
				p.Match(ExprParserSP)
			}
			{
				p.SetState(43)
				p.query(3)
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_attrPath
	return p
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(ExprParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterAttrPath(s)
	}
}

func (s *AttrPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitAttrPath(s)
	}
}

func (p *ExprParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExprParserRULE_attrPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(ExprParserATTRNAME)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ExprParserT__2 {
		{
			p.SetState(50)
			p.SubAttr()
		}

	}

	return localctx
}

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_subAttr
	return p
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSubAttr(s)
	}
}

func (s *SubAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSubAttr(s)
	}
}

func (p *ExprParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExprParserRULE_subAttr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(ExprParserT__2)
	}
	{
		p.SetState(54)
		p.AttrPath()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListOfDoublesContext struct {
	*ValueContext
}

func NewListOfDoublesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfDoublesContext {
	var p = new(ListOfDoublesContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfDoublesContext) ListDoubles() IListDoublesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListDoublesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListDoublesContext)
}

func (s *ListOfDoublesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterListOfDoubles(s)
	}
}

func (s *ListOfDoublesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitListOfDoubles(s)
	}
}

type ListOfStringsContext struct {
	*ValueContext
}

func NewListOfStringsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfStringsContext {
	var p = new(ListOfStringsContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfStringsContext) ListStrings() IListStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *ListOfStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterListOfStrings(s)
	}
}

func (s *ListOfStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitListOfStrings(s)
	}
}

type BooleanContext struct {
	*ValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ExprParserBOOLEAN, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitBoolean(s)
	}
}

type NullContext struct {
	*ValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(ExprParserNULL, 0)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitNull(s)
	}
}

type StringContext struct {
	*ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExprParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitString(s)
	}
}

type DoubleContext struct {
	*ValueContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ExprParserDOUBLE, 0)
}

func (s *DoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterDouble(s)
	}
}

func (s *DoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitDouble(s)
	}
}

type ListOfIntegersContext struct {
	*ValueContext
}

func NewListOfIntegersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfIntegersContext {
	var p = new(ListOfIntegersContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ListOfIntegersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfIntegersContext) ListIntegers() IListIntegersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListIntegersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListIntegersContext)
}

func (s *ListOfIntegersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterListOfIntegers(s)
	}
}

func (s *ListOfIntegersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitListOfIntegers(s)
	}
}

type IntegerContext struct {
	*ValueContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(ExprParserINT, 0)
}

func (s *IntegerContext) EXP() antlr.TerminalNode {
	return s.GetToken(ExprParserEXP, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *ExprParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExprParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(ExprParserBOOLEAN)
		}

	case 2:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(ExprParserNULL)
		}

	case 3:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.Match(ExprParserSTRING)
		}

	case 4:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)
			p.Match(ExprParserDOUBLE)
		}

	case 5:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ExprParserT__3 {
			{
				p.SetState(60)
				p.Match(ExprParserT__3)
			}

		}
		{
			p.SetState(63)
			p.Match(ExprParserINT)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(64)
				p.Match(ExprParserEXP)
			}

		}

	case 6:
		localctx = NewListOfIntegersContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(67)
			p.ListIntegers()
		}

	case 7:
		localctx = NewListOfDoublesContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(68)
			p.ListDoubles()
		}

	case 8:
		localctx = NewListOfStringsContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(69)
			p.ListStrings()
		}

	}

	return localctx
}

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_listStrings
	return p
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterListStrings(s)
	}
}

func (s *ListStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitListStrings(s)
	}
}

func (p *ExprParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExprParserRULE_listStrings)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(ExprParserT__4)
	}
	{
		p.SetState(73)
		p.SubListOfStrings()
	}

	return localctx
}

// ISubListOfStringsContext is an interface to support dynamic dispatch.
type ISubListOfStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubListOfStringsContext differentiates from other interfaces.
	IsSubListOfStringsContext()
}

type SubListOfStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfStringsContext() *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_subListOfStrings
	return p
}

func (*SubListOfStringsContext) IsSubListOfStringsContext() {}

func NewSubListOfStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_subListOfStrings

	return p
}

func (s *SubListOfStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfStringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExprParserSTRING, 0)
}

func (s *SubListOfStringsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, 0)
}

func (s *SubListOfStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *SubListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSubListOfStrings(s)
	}
}

func (s *SubListOfStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSubListOfStrings(s)
	}
}

func (p *ExprParser) SubListOfStrings() (localctx ISubListOfStringsContext) {
	localctx = NewSubListOfStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExprParserRULE_subListOfStrings)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(ExprParserSTRING)
		}
		{
			p.SetState(76)
			p.Match(ExprParserCOMMA)
		}
		{
			p.SetState(77)
			p.SubListOfStrings()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Match(ExprParserSTRING)
		}
		{
			p.SetState(79)
			p.Match(ExprParserT__5)
		}

	}

	return localctx
}

// IListDoublesContext is an interface to support dynamic dispatch.
type IListDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListDoublesContext differentiates from other interfaces.
	IsListDoublesContext()
}

type ListDoublesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListDoublesContext() *ListDoublesContext {
	var p = new(ListDoublesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_listDoubles
	return p
}

func (*ListDoublesContext) IsListDoublesContext() {}

func NewListDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDoublesContext {
	var p = new(ListDoublesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_listDoubles

	return p
}

func (s *ListDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfDoublesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *ListDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDoublesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterListDoubles(s)
	}
}

func (s *ListDoublesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitListDoubles(s)
	}
}

func (p *ExprParser) ListDoubles() (localctx IListDoublesContext) {
	localctx = NewListDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExprParserRULE_listDoubles)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(ExprParserT__4)
	}
	{
		p.SetState(83)
		p.SubListOfDoubles()
	}

	return localctx
}

// ISubListOfDoublesContext is an interface to support dynamic dispatch.
type ISubListOfDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubListOfDoublesContext differentiates from other interfaces.
	IsSubListOfDoublesContext()
}

type SubListOfDoublesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfDoublesContext() *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_subListOfDoubles
	return p
}

func (*SubListOfDoublesContext) IsSubListOfDoublesContext() {}

func NewSubListOfDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_subListOfDoubles

	return p
}

func (s *SubListOfDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfDoublesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ExprParserDOUBLE, 0)
}

func (s *SubListOfDoublesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, 0)
}

func (s *SubListOfDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfDoublesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *SubListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfDoublesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSubListOfDoubles(s)
	}
}

func (s *SubListOfDoublesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSubListOfDoubles(s)
	}
}

func (p *ExprParser) SubListOfDoubles() (localctx ISubListOfDoublesContext) {
	localctx = NewSubListOfDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExprParserRULE_subListOfDoubles)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Match(ExprParserDOUBLE)
		}
		{
			p.SetState(86)
			p.Match(ExprParserCOMMA)
		}
		{
			p.SetState(87)
			p.SubListOfDoubles()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(ExprParserDOUBLE)
		}
		{
			p.SetState(89)
			p.Match(ExprParserT__5)
		}

	}

	return localctx
}

// IListIntegersContext is an interface to support dynamic dispatch.
type IListIntegersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListIntegersContext differentiates from other interfaces.
	IsListIntegersContext()
}

type ListIntegersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIntegersContext() *ListIntegersContext {
	var p = new(ListIntegersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_listIntegers
	return p
}

func (*ListIntegersContext) IsListIntegersContext() {}

func NewListIntegersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIntegersContext {
	var p = new(ListIntegersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_listIntegers

	return p
}

func (s *ListIntegersContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIntegersContext) SubListOfIntegers() ISubListOfIntegersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfIntegersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntegersContext)
}

func (s *ListIntegersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIntegersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIntegersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterListIntegers(s)
	}
}

func (s *ListIntegersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitListIntegers(s)
	}
}

func (p *ExprParser) ListIntegers() (localctx IListIntegersContext) {
	localctx = NewListIntegersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExprParserRULE_listIntegers)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(ExprParserT__4)
	}
	{
		p.SetState(93)
		p.SubListOfIntegers()
	}

	return localctx
}

// ISubListOfIntegersContext is an interface to support dynamic dispatch.
type ISubListOfIntegersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubListOfIntegersContext differentiates from other interfaces.
	IsSubListOfIntegersContext()
}

type SubListOfIntegersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfIntegersContext() *SubListOfIntegersContext {
	var p = new(SubListOfIntegersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_subListOfIntegers
	return p
}

func (*SubListOfIntegersContext) IsSubListOfIntegersContext() {}

func NewSubListOfIntegersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfIntegersContext {
	var p = new(SubListOfIntegersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_subListOfIntegers

	return p
}

func (s *SubListOfIntegersContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfIntegersContext) INT() antlr.TerminalNode {
	return s.GetToken(ExprParserINT, 0)
}

func (s *SubListOfIntegersContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ExprParserCOMMA, 0)
}

func (s *SubListOfIntegersContext) SubListOfIntegers() ISubListOfIntegersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfIntegersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntegersContext)
}

func (s *SubListOfIntegersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfIntegersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfIntegersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSubListOfIntegers(s)
	}
}

func (s *SubListOfIntegersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSubListOfIntegers(s)
	}
}

func (p *ExprParser) SubListOfIntegers() (localctx ISubListOfIntegersContext) {
	localctx = NewSubListOfIntegersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ExprParserRULE_subListOfIntegers)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(ExprParserINT)
		}
		{
			p.SetState(96)
			p.Match(ExprParserCOMMA)
		}
		{
			p.SetState(97)
			p.SubListOfIntegers()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(ExprParserINT)
		}
		{
			p.SetState(99)
			p.Match(ExprParserT__5)
		}

	}

	return localctx
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExprParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
