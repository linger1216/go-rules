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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) != right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue != right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) > right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue > right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) >= right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue >= right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) < right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue < right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) <= right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue <= right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
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
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return false, ErrInvalidOperator
	case TypeValueDouble:
		return float64(leftValue) == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == right.Value().(int64), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _integerListFloatIn(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _integerIntegerIn(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _integerListFloatIn(v int64, arr []float64) bool {
	for i := range arr {
		if float64(v) == arr[i] {
			return true
		}
	}
	return false
}

func _integerIntegerIn(v int64, arr []int64) bool {
	for i := range arr {
		if v == arr[i] {
			return true
		}
	}
	return false
}
