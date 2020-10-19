package rules

type FloatJudge struct {
}

func (s *FloatJudge) EQ(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return leftValue == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == float64(right.Value().(int64)), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *FloatJudge) NE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return leftValue != right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue != float64(right.Value().(int64)), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *FloatJudge) GT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
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
		return leftValue > right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue > float64(right.Value().(int64)), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *FloatJudge) GE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
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
		return leftValue >= right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue >= float64(right.Value().(int64)), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *FloatJudge) LT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
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
		return leftValue < right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue < float64(right.Value().(int64)), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *FloatJudge) LE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
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
		return leftValue <= right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue <= float64(right.Value().(int64)), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *FloatJudge) Contain(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
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
		return leftValue == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == float64(right.Value().(int64)), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return false, ErrInvalidOperator
	case TypeValueListIntegers:
		return false, ErrInvalidOperator
	}
	return false, ErrInvalidType
}

func (s *FloatJudge) Prefix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *FloatJudge) Suffix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *FloatJudge) Regex(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *FloatJudge) In(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toFloat(left)
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
		return leftValue == right.Value().(float64), nil
	case TypeValueInteger:
		return leftValue == float64(right.Value().(int64)), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _floatListFloatIn(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _floatListIntegerIn(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _floatListFloatIn(v float64, arr []float64) bool {
	for i := range arr {
		if v == arr[i] {
			return true
		}
	}
	return false
}

func _floatListIntegerIn(v float64, arr []int64) bool {
	for i := range arr {
		if v == float64(arr[i]) {
			return true
		}
	}
	return false
}
