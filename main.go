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
		"obj": map[string]interface{}{
			"bool":   false,
			"number": 4,
			"string": "value",
			"float":  3.1415,
			"arr":    []string{"a", "b", "c"},
		},
		"arr": []string{"a", "b", "c"},
	}

	rule := `not(number == 4 and bool == false) and string == "value" and float >= 3.1415 and obj.arr[0] == "a"`
	res, err := rules.Evaluate(rule, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", res)
}
