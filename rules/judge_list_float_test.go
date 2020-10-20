package rules

import (
	"fmt"
	"testing"
)

func TestListFloatJudge(t *testing.T) {
	tests := []struct {
		name   string
		left   interface{}
		right  IFiled
		expect bool
	}{
		// TODO: Add test cases.
		{
			"==",
			[]float64{1.0},
			NewValueListIntegers("[1]"),
			true,
		},

		{
			"==",
			[]float64{1.0},
			NewValueListFloats("[1.0]"),
			true,
		},

		{
			"!=",
			[]float64{1.0},
			NewValueListIntegers("[2]"),
			true,
		},

		{
			"!=",
			[]float64{1.0},
			NewValueListFloats("[2.0]"),
			true,
		},

		{
			">",
			[]float64{10},
			NewValueListIntegers("[1,2,3,4]"),
			true,
		},

		{
			">",
			[]float64{10},
			NewValueListFloats("[1,2,3,4]"),
			true,
		},

		{
			">=",
			[]float64{10},
			NewValueListIntegers("[1,2,3,4]"),
			true,
		},

		{
			">=",
			[]float64{10},
			NewValueListFloats("[1,2,3,4]"),
			true,
		},

		{
			"<",
			[]float64{1.0},
			NewValueListIntegers("[5,6,7,8]"),
			true,
		},

		{
			"<",
			[]float64{1.0},
			NewValueListFloats("[5,6,7,8]"),
			true,
		},

		{
			"<=",
			[]float64{1.0},
			NewValueListIntegers("[5,6,7,8]"),
			true,
		},

		{
			"<=",
			[]float64{1.0},
			NewValueListIntegers("[5,6,7,8]"),
			true,
		},

		{
			"contain",
			[]float64{1, 2, 3, 4},
			NewValueListIntegers("[1]"),
			true,
		},

		{
			"contain",
			[]float64{1, 2, 3, 4},
			NewValueListFloats("[1.0]"),
			true,
		},

		{
			"in",
			[]float64{1.0},
			NewValueListIntegers("[1,2,3]"),
			true,
		},

		{
			"in",
			[]float64{1.0},
			NewValueListFloats("[1.0,2.0,3]"),
			true,
		},
	}
	s := &ListFloatJudge{}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%s[%d]", tt.name, i), func(t *testing.T) {
			fn := getFunc(tt.name, s)
			got, err := fn(tt.left, tt.right)
			if err != nil {
				t.Errorf("%s error = %v", tt.name, err)
				return
			}
			if got != tt.expect {
				t.Errorf("%s got = %v, want %v", tt.name, got, tt.expect)
			}
		})
	}
}
