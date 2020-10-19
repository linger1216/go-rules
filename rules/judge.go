package rules

import "strings"

func toString(op interface{}) (string, error) {
	switch val := op.(type) {
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	}
	return "", ErrInvalidType
}

func toListString(op interface{}) ([]string, error) {
	switch val := op.(type) {
	case []string:
		return val, nil
	}
	return nil, ErrInvalidType
}

func toBoolean(op interface{}) (bool, error) {
	switch val := op.(type) {
	case bool:
		return val, nil
	}
	return false, ErrInvalidType
}

func toListBoolean(op interface{}) ([]bool, error) {
	switch val := op.(type) {
	case []bool:
		return val, nil
	}
	return nil, ErrInvalidType
}

func toFloat(op interface{}) (float64, error) {
	switch val := op.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	}
	return 0, ErrInvalidType
}

func toListFloat(op interface{}) ([]float64, error) {
	switch val := op.(type) {
	case []float64:
		return val, nil
	}
	return nil, ErrInvalidType
}

func toInt64(op interface{}) (int64, error) {
	switch val := op.(type) {
	case int:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	}
	return 0, ErrInvalidType
}

func toListInt64(op interface{}) ([]int64, error) {
	switch val := op.(type) {
	case []int64:
		return val, nil
	}
	return nil, ErrInvalidType
}

func EQListIntegers(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EQListDoubles(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ContainString(arr []string, v string) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func ContainBoolean(arr []bool, v bool) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func ContainInteger(arr []int64, v int64) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func ContainFloat(arr []float64, v float64) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func ContainListString(arr, v []string) bool {
	for i := range v {
		if !ContainString(arr, v[i]) {
			return false
		}
	}
	return true
}

func ContainListInteger(arr, v []int64) bool {
	for i := range v {
		if !ContainInteger(arr, v[i]) {
			return false
		}
	}
	return true
}

func ContainListFloat(arr, v []float64) bool {
	for i := range v {
		if !ContainFloat(arr, v[i]) {
			return false
		}
	}
	return true
}

func PrefixString(left, right string) bool {
	return strings.HasPrefix(left, right)
}

func PrefixListString(left string, arr []string) bool {
	for i := range arr {
		if !PrefixString(left, arr[i]) {
			return false
		}
	}
	return true
}

func judgeEQ(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return v == right.Value().(string), nil
	case TypeValueBoolean:
		v, err := toBoolean(left)
		if err != nil {
			return false, err
		}
		return v == right.Value().(bool), nil
	case TypeValueNull:
		return left == nil, nil
	case TypeValueDouble:
		v, err := toFloat(left)
		if err != nil {
			return false, err
		}
		return v == right.Value().(float64), nil
	case TypeValueInteger:
		v, err := toInt64(left)
		if err != nil {
			return false, err
		}
		return v == right.Value().(int64), nil
	case TypeValueListStrings:
		v, err := toListString(left)
		if err != nil {
			return false, err
		}
		return _listStringEQListString(v, right.Value().([]string)), nil
	case TypeValueListDoubles:
		v, err := toListFloat(left)
		if err != nil {
			return false, err
		}
		return EQListDoubles(v, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		v, err := toListInt64(left)
		if err != nil {
			return false, err
		}
		return EQListIntegers(v, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}
func judgeNE(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return v != right.Value().(string), nil
	case TypeValueBoolean:
		v, err := toBoolean(left)
		if err != nil {
			return false, err
		}
		return v != right.Value().(bool), nil
	case TypeValueNull:
		return left != nil, nil
	case TypeValueDouble:
		v, err := toFloat(left)
		if err != nil {
			return false, err
		}
		return v != right.Value().(float64), nil
	case TypeValueInteger:
		v, err := toInt64(left)
		if err != nil {
			return false, err
		}
		return v != right.Value().(int64), nil
	case TypeValueListStrings:
		v, err := toListString(left)
		if err != nil {
			return false, err
		}
		return !_listStringEQListString(v, right.Value().([]string)), nil
	case TypeValueListDoubles:
		v, err := toListFloat(left)
		if err != nil {
			return false, err
		}
		return !EQListDoubles(v, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		v, err := toListInt64(left)
		if err != nil {
			return false, err
		}
		return !EQListIntegers(v, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}
func judgeGT(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return v > right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		v, err := toFloat(left)
		if err != nil {
			return false, err
		}
		return v > right.Value().(float64), nil
	case TypeValueInteger:
		v, err := toInt64(left)
		if err != nil {
			return false, err
		}
		return v > right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}
func judgeGE(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return v >= right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		v, err := toFloat(left)
		if err != nil {
			return false, err
		}
		return v >= right.Value().(float64), nil
	case TypeValueInteger:
		v, err := toInt64(left)
		if err != nil {
			return false, err
		}
		return v >= right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}
func judgeLT(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return v < right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		v, err := toFloat(left)
		if err != nil {
			return false, err
		}
		return v < right.Value().(float64), nil
	case TypeValueInteger:
		v, err := toInt64(left)
		if err != nil {
			return false, err
		}
		return v < right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}
func judgeLE(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return v <= right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		v, err := toFloat(left)
		if err != nil {
			return false, err
		}
		return v <= right.Value().(float64), nil
	case TypeValueInteger:
		v, err := toInt64(left)
		if err != nil {
			return false, err
		}
		return v <= right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}
func judgeContain(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toListString(left)
		if err != nil {
			return false, err
		}
		return ContainString(v, right.Value().(string)), nil
	case TypeValueBoolean:
		v, err := toListBoolean(left)
		if err != nil {
			return false, err
		}
		return ContainBoolean(v, right.Value().(bool)), nil
	case TypeValueNull:
		return left == nil, nil
	case TypeValueDouble:
		v, err := toListFloat(left)
		if err != nil {
			return false, err
		}
		return ContainFloat(v, right.Value().(float64)), nil
	case TypeValueInteger:
		v, err := toListInt64(left)
		if err != nil {
			return false, err
		}
		return ContainInteger(v, right.Value().(int64)), nil
	case TypeValueListStrings:
		v, err := toListString(left)
		if err != nil {
			return false, err
		}
		return ContainListString(v, right.Value().([]string)), nil
	case TypeValueListDoubles:
		v, err := toListFloat(left)
		if err != nil {
			return false, err
		}
		return ContainListFloat(v, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		v, err := toListInt64(left)
		if err != nil {
			return false, err
		}
		return ContainListInteger(v, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}
func judgePrefix(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixString(v, right.Value().(string)), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixListString(v, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}
func judgeSuffix(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixString(v, right.Value().(string)), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixListString(v, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}
func judgeRegex(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixString(v, right.Value().(string)), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixListString(v, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func judgeIn(left interface{}, right IFiled) (bool, error) {
	switch right.Type() {
	case TypeValueString:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixString(v, right.Value().(string)), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		v, err := toString(left)
		if err != nil {
			return false, err
		}
		return PrefixListString(v, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

type Judge interface {
	EQ(left, right IFiled) (bool, error)
	NE(left, right IFiled) (bool, error)
	GT(left, right IFiled) (bool, error)
	GE(left, right IFiled) (bool, error)
	LT(left, right IFiled) (bool, error)
	LE(left, right IFiled) (bool, error)
	Contain(left, right IFiled) (bool, error)
	Prefix(left, right IFiled) (bool, error)
	Suffix(left, right IFiled) (bool, error)
	Regex(left, right IFiled) (bool, error)
	In(left, right IFiled) (bool, error)
}
