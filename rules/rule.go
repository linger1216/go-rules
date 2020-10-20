package rules

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linger1216/go-rules/parser"
)

func Evaluate(rule string, items map[string]interface{}) (bool, error) {
	is := antlr.NewInputStream(rule)
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(stream)

	exprListener := NewExprListener(items)
	antlr.ParseTreeWalkerDefault.Walk(exprListener, p.Query())
	if exprListener.err != nil {
		return false, exprListener.err
	}
	return exprListener.Result()
}

func PrepareExprParser(rule string) *parser.ExprParser {
	is := antlr.NewInputStream(rule)
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(stream)
	return p
}

func EvaluateByParser(p *parser.ExprParser, items map[string]interface{}) (bool, error) {
	exprListener := NewExprListener(items)
	antlr.ParseTreeWalkerDefault.Walk(exprListener, p.Query())
	if exprListener.err != nil {
		return false, exprListener.err
	}
	return exprListener.Result()
}
