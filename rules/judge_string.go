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
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _stringPrefixListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func _stringPrefixListString(v string, arr []string) bool {
	for i := range arr {
		if !strings.HasPrefix(v, arr[i]) {
			return false
		}
	}
	return true
}

func _stringSuffixListString(v string, arr []string) bool {
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
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _stringSuffixListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func _stringRegexListString(v string, arr []string) bool {
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
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _stringRegexListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
		return leftValue == right.Value().(string), nil
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _stringInListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func _stringInListString(v string, arr []string) bool {
	for i := range arr {
		if v == arr[i] {
			return true
		}
	}
	return false
}
