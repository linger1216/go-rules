package rules

type BooleanJudge struct {
}

func (s *BooleanJudge) EQ(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toBoolean(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
		return leftValue == right.Value().(bool), nil
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *BooleanJudge) NE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toBoolean(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
		return leftValue != right.Value().(bool), nil
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
	case TypeValueListIntegers:
	}
	return false, ErrInvalidType
}

func (s *BooleanJudge) GT(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) GE(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) LT(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) LE(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) Contain(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) Prefix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) Suffix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) Regex(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *BooleanJudge) In(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}
