package rules

type ListIntegerJudge struct {
}

func _listIntegerEQListInteger(left []int64, right []int64) bool {
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

func _listIntegerEQListFloat(left []int64, right []float64) bool {
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listIntegerEQListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listIntegerEQListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return !_listIntegerEQListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerEQListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listIntegerGTListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerGTListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listIntegerGEListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return !_listIntegerGEListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listIntegerLTListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listIntegerLTListInteger(leftValue, right.Value().([]int64)), nil
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listIntegerLEListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listIntegerLEListInteger(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _listIntegerContainListInteger(left []int64, right []int64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_integerInListInteger(right[i], left) {
			return false
		}
	}
	return true
}

func _listIntegerContainListFloat(left []int64, right []float64) bool {
	// right的每一个元素都在left里
	for i := range right {
		if !_integerInListInteger(int64(right[i]), left) {
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listIntegerContainListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listIntegerContainListInteger(leftValue, right.Value().([]int64)), nil
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

func _listIntegerInListInteger(left, right []int64) bool {
	for i := range left {
		if !_integerInListInteger(left[i], right) {
			return false
		}
	}
	return true
}

func _listIntegerInListFloat(left []int64, right []float64) bool {
	for i := range left {
		if !_integerInListFloat(left[i], right) {
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
	case TypeValueBoolean:
	case TypeValueNull:
	case TypeValueDouble:
	case TypeValueInteger:
	case TypeValueListStrings:
	case TypeValueListDoubles:
		return _listIntegerInListFloat(leftValue, right.Value().([]float64)), nil
	case TypeValueListIntegers:
		return _listIntegerInListInteger(leftValue, right.Value().([]int64)), nil
	}
	return false, ErrInvalidType
}

func _listIntegerGTListInteger(left []int64, right []int64) bool {
	for i := range left {
		if left[i] <= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerGTListFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) <= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerGEListInteger(left []int64, right []int64) bool {
	for i := range left {
		if left[i] < right[i] {
			return false
		}
	}
	return true
}

func _listIntegerGEListFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) < right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLTListInteger(left []int64, right []int64) bool {
	for i := range left {
		if left[i] >= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLTListFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) >= right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLEListInteger(left []int64, right []int64) bool {
	for i := range left {
		if left[i] > right[i] {
			return false
		}
	}
	return true
}

func _listIntegerLEListFloat(left []int64, right []float64) bool {
	for i := range left {
		if float64(left[i]) > right[i] {
			return false
		}
	}
	return true
}
