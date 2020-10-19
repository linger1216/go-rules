package rules

type ListIntegerJudge struct {
}

func _listIntegerEQ(left []int64, right []int64) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func _listIntegerEQFloat(left []int64, right []float64) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if float64(left[i]) != right[i] {
			return false
		}
	}
	return true
}

func (s *ListIntegerJudge) EQ(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
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
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _listIntegerEQFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listIntegerEQ(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) NE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) != 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return !_listIntegerEQFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerEQ(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) GT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) != 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _listIntegerGTFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerGT(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) GE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) != 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _listIntegerGEFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerGE(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) LT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) != 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _listIntegerLTFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerLT(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) LE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) != 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _listIntegerLEFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerLE(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _listIntegerContain(left []int64, right []int64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_integerIntegerIn(right[i], left) {
			return false
		}
	}
	return true
}

func _listIntegerContainFloat(left []int64, right []float64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_integerIntegerIn(int64(right[i]), left) {
			return false
		}
	}
	return true
}

func (s *ListIntegerJudge) Contain(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) != 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _listIntegerContainFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerContain(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) Prefix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) Suffix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListIntegerJudge) Regex(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func _listIntegerIn(left, right []int64) bool {
	for i := range left {
		if !_integerIntegerIn(left[i], right) {
			return false
		}
	}
	return true
}

func _listIntegerInFloat(left []int64, right []float64) bool {
	for i := range left {
		if !_integerListFloatIn(left[i], right) {
			return false
		}
	}
	return true
}

func (s *ListIntegerJudge) In(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListInt64(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
		return false, ErrInvalidOperator
	case TypeValueBoolean:
		return false, ErrInvalidOperator
	case TypeValueNull:
		return len(leftValue) != 0, nil
	case TypeValueDouble:
		return false, ErrInvalidOperator
	case TypeValueInteger:
		return false, ErrInvalidOperator
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return _listIntegerInFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerIn(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _listIntegerGT(left []int64, right []int64) bool {
	for i := range left {
		if left[i] <= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerGTFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) <= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerGE(left []int64, right []int64) bool {
	for i := range left {
		if left[i] < right[i] {
			return false
		}
	}
	return true
}

func _listIntegerGEFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) < right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLT(left []int64, right []int64) bool {
	for i := range left {
		if left[i] >= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLTFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) >= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLE(left []int64, right []int64) bool {
	for i := range left {
		if left[i] > right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLEFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) > right[i] {
			return false
		}
	}
	return true
}
