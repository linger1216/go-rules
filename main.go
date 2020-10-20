package main

import (
	"fmt"
	"github.com/linger1216/go-rules/rules"
)

func main() {
	m := map[string]interface{}{
		"bool":   false,
		"number": 4,
		"string": "value",
		"float":  3.1415,
	}

	rule := `number == 4 and bool == false and string == "value" and float >= 3.1415`
	res, err := rules.Evaluate(rule, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("result1:", res)
}
