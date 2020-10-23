# go-rules



### usage
```
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

	rule := `number == 4 and bool == false and string == "value" and float >= 3.1415 and obj.arr[0] == "a"`
	res, err := rules.Evaluate(rule, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", res)
}
```





### Logical Operator

- and
- or



### Compare Operator

- `== or eq`
- `!= or ne`
- `> or gt`
- `>= or ge`
- `< or lt`
- `<= or le`
- `contain`
- `prefix`
- `suffix`
- `in`
- `regex`



### Value

| Type         | example value |
| ------------ | ------------- |
| boolean      | true or false |
| null         | null          |
| string       | "hello"       |
| list string  | [s1, s2, s3]  |
| float        | 3.1415        |
| list float   | [1.0,2.0]     |
| integer      | 1             |
| list integer | [1,2,3]       |