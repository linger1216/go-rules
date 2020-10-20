package rules

func toString(op interface{}) (string, error) {
	switch val := op.(type) {
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	}
	return "", ErrInvalidType
}

func toListString(op interface{}) ([]string, error) {
	switch val := op.(type) {
	case []string:
		return val, nil
	}
	return nil, ErrInvalidType
}

func toBoolean(op interface{}) (bool, error) {
	switch val := op.(type) {
	case bool:
		return val, nil
	}
	return false, ErrInvalidType
}

func toListBoolean(op interface{}) ([]bool, error) {
	switch val := op.(type) {
	case []bool:
		return val, nil
	}
	return nil, ErrInvalidType
}

func toFloat(op interface{}) (float64, error) {
	switch val := op.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	}
	return 0, ErrInvalidType
}

func toListFloat(op interface{}) ([]float64, error) {
	switch val := op.(type) {
	case []float64:
		return val, nil
	}
	return nil, ErrInvalidType
}

func toInt64(op interface{}) (int64, error) {
	switch val := op.(type) {
	case int:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	}
	return 0, ErrInvalidType
}

func toListInt64(op interface{}) ([]int64, error) {
	switch val := op.(type) {
	case []int64:
		return val, nil
	}
	return nil, ErrInvalidType
}

type Judge interface {
	EQ(left interface{}, right IFiled) (bool, error)
	NE(left interface{}, right IFiled) (bool, error)
	GT(left interface{}, right IFiled) (bool, error)
	GE(left interface{}, right IFiled) (bool, error)
	LT(left interface{}, right IFiled) (bool, error)
	LE(left interface{}, right IFiled) (bool, error)
	Contain(left interface{}, right IFiled) (bool, error)
	Prefix(left interface{}, right IFiled) (bool, error)
	Suffix(left interface{}, right IFiled) (bool, error)
	Regex(left interface{}, right IFiled) (bool, error)
	In(left interface{}, right IFiled) (bool, error)
}
