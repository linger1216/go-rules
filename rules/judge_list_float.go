package rules

type ListFloatJudge struct {
}

func _listFloatListIntegerEQ(left []float64, right []int64) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatListFloatEQ(left []float64, right []float64) bool {
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

func (s *ListFloatJudge) EQ(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
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
		return _listFloatListFloatEQ(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatListIntegerEQ(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListFloatJudge) NE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
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
		return !_listFloatListFloatEQ(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatListIntegerEQ(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListFloatJudge) GT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
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
		return !_listFloatListFloatGT(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatListIntegerGT(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListFloatJudge) GE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
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
		return !_listFloatListFloatGE(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatListIntegerGE(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListFloatJudge) LT(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
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
		return !_listFloatListFloatLT(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatListIntegerLT(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListFloatJudge) LE(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
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
		return _listFloatFloatLE(right.Value().(float64), leftValue), nil
	case TypeValueInteger:
		return _integerListFloatIn(right.Value().(int64), leftValue), nil
	case TypeValueListStrings:
		return false, ErrInvalidOperator
	case TypeValueListDoubles:
		return !_listFloatListFloatLE(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatListIntegerLE(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListFloatJudge) Contain(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
	if err != nil {
		return false, err
	}

	switch right.Type() {
	case TypeValueString:
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
		return _floatListFloatIn(right.Value().(float64), leftValue), nil
	case TypeValueInteger:
		return _integerListFloatIn(right.Value().(int64), leftValue), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return !_listFloatListFloatContain(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatListIntegerContain(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func (s *ListFloatJudge) Prefix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListFloatJudge) Suffix(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListFloatJudge) Regex(left interface{}, right IFiled) (bool, error) {
	return false, ErrInvalidType
}

func (s *ListFloatJudge) In(left interface{}, right IFiled) (bool, error) {
	leftValue, err := toListFloat(left)
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
	case TypeValueListDoubles:
		return !_listFloatListFloatIn(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatListIntegerIn(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _listFloatListIntegerGT(left []float64, right []int64) bool {
	for i := range left {
		if left[i] <= float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatListFloatGT(left []float64, right []float64) bool {
	for i := range left {
		if left[i] <= right[i] {
			return false
		}
	}
	return true
}

func _listFloatListIntegerGE(left []float64, right []int64) bool {
	for i := range left {
		if left[i] < float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatListFloatGE(left []float64, right []float64) bool {
	for i := range left {
		if left[i] < right[i] {
			return false
		}
	}
	return true
}

func _listFloatListIntegerLT(left []float64, right []int64) bool {
	for i := range left {
		if left[i] >= float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatListFloatLT(left []float64, right []float64) bool {
	for i := range left {
		if left[i] >= right[i] {
			return false
		}
	}
	return true
}

func _listFloatListIntegerLE(left []float64, right []int64) bool {
	for i := range left {
		if left[i] > float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatListFloatLE(left []float64, right []float64) bool {
	for i := range left {
		if left[i] > right[i] {
			return false
		}
	}
	return true
}

func _listFloatListIntegerContain(left []float64, right []int64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_integerListFloatIn(right[i], left) {
			return false
		}
	}
	return true
}

func _listFloatListFloatContain(left []float64, right []float64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_floatListFloatIn(right[i], left) {
			return false
		}
	}
	return true
}

func _listFloatListIntegerIn(left []float64, right []int64) bool {
	for i := range left {
		if !_floatListIntegerIn(left[i], right) {
			return false
		}
	}
	return true
}

func _listFloatListFloatIn(left []float64, right []float64) bool {
	for i := range left {
		if !_floatListFloatIn(left[i], right) {
			return false
		}
	}
	return true
}
