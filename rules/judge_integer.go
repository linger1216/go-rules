package rules

type IntegerJudge struct {
}

func (s *IntegerJudge) EQ(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *IntegerJudge) NE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) != right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue != right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *IntegerJudge) GT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}
	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) > right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue > right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *IntegerJudge) GE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) >= right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue >= right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *IntegerJudge) LT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) < right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue < right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *IntegerJudge) LE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) <= right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue <= right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *IntegerJudge) Contain(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}
	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *IntegerJudge) Prefix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *IntegerJudge) Suffix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *IntegerJudge) Regex(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *IntegerJudge) In(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return float64(leftValue) == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == right.Value().(int64), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _integerInListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _integerInListInteger(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _integerInListFloat(v int64, arr []float64) bool {
	for i := range arr {
		if float64(v) == arr[i] {
			return true
		}
	}
	return false
}

func _integerInListInteger(v int64, arr []int64) bool {
	for i := range arr {
		if v == arr[i] {
			return true
		}
	}
	return false
}
