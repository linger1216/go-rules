package rules

func _listStringEQ(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range right {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

// [a,b,c] [b,c]
func _listStringContain(left, right []string) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_stringIn(right[i], left) {
			return false
		}
	}
	return true
}

// [a,b,c] [b,c]
func _listStringPrefix(left, right []string) bool {
	for i := range left {
		if !_stringPrefix(left[i], right) {
			return false
		}
	}
	return true
}

func _listStringSuffix(left, right []string) bool {
	for i := range left {
		if !_stringSuffix(left[i], right) {
			return false
		}
	}
	return true
}

func _listStringRegex(left, right []string) bool {
	for i := range left {
		if !_stringRegex(left[i], right) {
			return false
		}
	}
	return true
}

func _listStringIn(left, right []string) bool {
	for i := range left {
		if !_stringIn(left[i], right) {
			return false
		}
	}
	return true
}

type ListStringJudge struct {
}

func (s *ListStringJudge) EQ(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return _listStringEQ(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *ListStringJudge) NE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return !_listStringEQ(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *ListStringJudge) GT(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListStringJudge) GE(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListStringJudge) LT(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListStringJudge) LE(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListStringJudge) Contain(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return !_listStringContain(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *ListStringJudge) Prefix(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return !_listStringPrefix(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *ListStringJudge) Suffix(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return !_listStringSuffix(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *ListStringJudge) Regex(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return !_listStringRegex(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *ListStringJudge) In(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListString(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return !_listStringIn(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}
