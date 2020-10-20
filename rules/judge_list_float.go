package rules

type ListFloatJudge struct {
}

func _listFloatEQListInteger(left []float64, right []int64) bool {
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

func _listFloatEQListFloat(left []float64, right []float64) bool {
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
	case TypeValueBoolean:
	case TypeValueNull:
		return len(leftValue) > 0, nil
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listFloatEQListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatEQListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
		return len(leftValue) == 0, nil
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return !_listFloatEQListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listFloatEQListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listFloatGTListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatGTListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listFloatGEListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatGEListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listFloatLTListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatLTListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listFloatLEListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatLEListInteger(leftValue, right.Value().([]int64)), nil
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
		return _floatInListFloat(right.Value().(float64), leftValue), nil
	case TypeValueInteger:
		return _integerInListFloat(right.Value().(int64), leftValue), nil
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listFloatContainListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatContainListInteger(leftValue, right.Value().([]int64)), nil
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
		return _listFloatInListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listFloatInListInteger(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _listFloatGTListInteger(left []float64, right []int64) bool {
	for i := range left {
		if left[i] <= float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatGTListFloat(left []float64, right []float64) bool {
	for i := range left {
		if left[i] <= right[i] {
			return false
		}
	}
	return true
}

func _listFloatGEListInteger(left []float64, right []int64) bool {
	for i := range left {
		if left[i] < float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatGEListFloat(left []float64, right []float64) bool {
	for i := range left {
		if left[i] < right[i] {
			return false
		}
	}
	return true
}

func _listFloatLTListInteger(left []float64, right []int64) bool {
	for i := range left {
		if left[i] >= float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatLTListFloat(left []float64, right []float64) bool {
	for i := range left {
		if left[i] >= right[i] {
			return false
		}
	}
	return true
}

func _listFloatLEListInteger(left []float64, right []int64) bool {
	for i := range left {
		if left[i] > float64(right[i]) {
			return false
		}
	}
	return true
}

func _listFloatLEListFloat(left []float64, right []float64) bool {
	for i := range left {
		if left[i] > right[i] {
			return false
		}
	}
	return true
}

func _listFloatContainListInteger(left []float64, right []int64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_integerInListFloat(right[i], left) {
			return false
		}
	}
	return true
}

func _listFloatContainListFloat(left []float64, right []float64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_floatInListFloat(right[i], left) {
			return false
		}
	}
	return true
}

func _listFloatInListInteger(left []float64, right []int64) bool {
	for i := range left {
		if !_floatInListInteger(left[i], right) {
			return false
		}
	}
	return true
}

func _listFloatInListFloat(left []float64, right []float64) bool {
	for i := range left {
		if !_floatInListFloat(left[i], right) {
			return false
		}
	}
	return true
}
