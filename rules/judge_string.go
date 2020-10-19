package rules

import (
	"regexp"
	"strings"
)

type StringJudge struct {
}

func (s *StringJudge) EQ(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return leftValue == right.Value().(string), nil
	case TypeValueBoolean:
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *StringJudge) NE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return leftValue != right.Value().(string), nil
	case TypeValueBoolean:
	case TypeValueNull:
		return len(leftValue) > 0, nil
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *StringJudge) GT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return leftValue > right.Value().(string), nil
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *StringJudge) GE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return leftValue >= right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *StringJudge) LT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return leftValue < right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *StringJudge) LE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return leftValue <= right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *StringJudge) Contain(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return leftValue == right.Value().(string), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *StringJudge) Prefix(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}
	switch right.Type() {
	case TypeValueString:
		return strings.HasPrefix(leftValue, right.Value().(string)), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _stringPrefix(leftValue, right.Value().([]string)), nil
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func _stringPrefix(v string, arr []string) bool {
	for i := range arr {
		if !strings.HasPrefix(v, arr[i]) {
			return false
		}
	}
	return true
}

func _stringSuffix(v string, arr []string) bool {
	for i := range arr {
		if !strings.HasSuffix(v, arr[i]) {
			return false
		}
	}
	return true
}

func (s *StringJudge) Suffix(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}
	switch right.Type() {
	case TypeValueString:
		return strings.HasSuffix(leftValue, right.Value().(string)), nil
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return _stringSuffix(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func _stringRegex(v string, arr []string) bool {
	for i := range arr {
		b, _ := regexp.MatchString(v, arr[i])
		if !b {
			return false
		}
	}
	return true
}

func (s *StringJudge) Regex(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}
	switch right.Type() {
	case TypeValueString:
		// a.b.c regex "[1-9]?\\+"
		// right is pattern
		return regexp.MatchString(right.Value().(string), leftValue)
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return _stringRegex(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *StringJudge) In(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toString(left)
	if err != nil {
		return false, err
	}
	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return _stringIn(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func _stringIn(v string, arr []string) bool {
	for i := range arr {
		if v == arr[i] {
			return true
		}
	}
	return false
}
