package rules

import (
	"fmt"
	"testing"
)

func TestFloatJudge(t *testing.T) {
	tests := []struct {
		name   string
		left   interface{}
		right  IFiled
		expect bool
	}{
		// TODO: Add test cases.
		{
			"==",
			1.0,
			NewValueInteger("1"),
			true,
		},

		{
			"==",
			1.0,
			NewValueFloat("1.0"),
			true,
		},

		{
			"!=",
			1.0,
			NewValueInteger("2"),
			true,
		},

		{
			"!=",
			1.0,
			NewValueFloat("2.0"),
			true,
		},

		{
			">",
			10.0,
			NewValueInteger("1"),
			true,
		},

		{
			">=",
			10.0,
			NewValueFloat("10"),
			true,
		},

		{
			"<",
			1.0,
			NewValueInteger("10"),
			true,
		},

		{
			"<=",
			1.0,
			NewValueFloat("10"),
			true,
		},

		{
			"contain",
			1.0,
			NewValueInteger("1"),
			true,
		},

		{
			"contain",
			1.0,
			NewValueFloat("1.0"),
			true,
		},

		{
			"in",
			1.0,
			NewValueInteger("1"),
			true,
		},
		{
			"in",
			1.0,
			NewValueListIntegers("[1,2,3]"),
			true,
		},

		{
			"in",
			1.0,
			NewValueFloat("1"),
			true,
		},
		{
			"in",
			1.0,
			NewValueListFloats("[1.0,2.0,3]"),
			true,
		},
	}
	s := &FloatJudge{}
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
