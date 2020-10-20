package rules

func _listStringEQListString(left, right []string) bool {
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

func _listStringContainListString(left, right []string) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_stringInListString(right[i], left) {
			return false
		}
	}
	return true
}

// [a,b,c] [b,c]
func _listStringPrefixListString(left, right []string) bool {
	for i := range left {
		if !_stringPrefixListString(left[i], right) {
			return false
		}
	}
	return true
}

func _listStringSuffixListString(left, right []string) bool {
	for i := range left {
		if !_stringSuffixListString(left[i], right) {
			return false
		}
	}
	return true
}

func _listStringRegexListString(left, right []string) bool {
	for i := range left {
		if !_stringRegexListString(left[i], right) {
			return false
		}
	}
	return true
}

func _listStringInListString(left, right []string) bool {
	for i := range left {
		if !_stringInListString(left[i], right) {
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
	case TypeValueBoolean:
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _listStringEQListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueBoolean:
	case TypeValueNull:
		return len(leftValue) > 0, nil
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return !_listStringEQListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
		return _stringInListString(right.Value().(string), leftValue), nil
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _listStringContainListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _listStringPrefixListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _listStringSuffixListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return !_listStringRegexListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
		return _listStringInListString(leftValue, right.Value().([]string)), nil
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}
