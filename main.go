package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/linger1216/go-rules/parser"
	"github.com/linger1216/go-rules/rules"
)

func main() {

	//sourceStr := `my email is gerrylon@163.com`
	//matched, _ := regexp.MatchString(`[\w-]+@[\w]+(?:\.[\w]+)+`, sourceStr)
	//fmt.Printf("%v", matched) // true
	// one.number == 4 and a.b.c regex "[1-9]?\\+" or z > 6 and (x IN [1,2,3])

	// `number == 4`
	// `list_number == 4`
	// `string prefix "val"`
	is := antlr.NewInputStream(`bool == false and number >= 4 and string prefix "val"`)
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	exprListener := rules.NewExpressionListener(map[string]interface{}{
		"bool":   false,
		"number": 4,
		"list_number": []int64{
			1, 2, 3,
		},
		"string": "value",
		"list_string": []string{
			"aaa", "bbb", "ccc",
		},
		"float": 3.1415,
		"list_float": []float64{
			3.1, 3.2, 3.3,
		},
		"two": map[string]interface{}{
			"three": []string{
				"aaa", "bbb", "ccc",
			},
		},
		"four": map[string]interface{}{
			"five": []int{
				11, 22, 33,
			},
		},
	})
	p := parser.NewExprParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(exprListener, p.Query())
	//if exprListener.err != nil {
	//	panic(error())
	//}
	fmt.Println("*********************")
	fmt.Println("result:", exprListener.Result())
	fmt.Println("*********************")
}
