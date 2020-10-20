package rules

import (
	"fmt"
	"testing"
)

func TestListStringJudge(t *testing.T) {
	tests := []struct {
		name   string
		left   interface{}
		right  IFiled
		expect bool
	}{
		// TODO: Add test cases.
		{
			"==",
			[]string{"hello"},
			NewValueListStrings("[hello]"),
			true,
		},

		{
			"==",
			[]string{},
			NewValueNull(),
			true,
		},

		{
			"!=",
			[]string{"hello"},
			NewValueListStrings("[world]"),
			true,
		},

		{
			"!=",
			[]string{"hello"},
			NewValueNull(),
			true,
		},

		{
			"contain",
			[]string{"ab", "bc"},
			NewValueString("ab"),
			true,
		},

		{
			"contain",
			[]string{"a", "b", "c"},
			NewValueListStrings("[a,b]"),
			true,
		},
		//
		{
			"prefix",
			[]string{"hello"},
			NewValueListStrings("[h,he]"),
			true,
		},

		{
			"suffix",
			[]string{"world"},
			NewValueListStrings("[d,ld,rld]"),
			true,
		},

		{
			"in",
			[]string{"a", "b"},
			NewValueListStrings("[a,b,c]"),
			true,
		},
	}
	s := &ListStringJudge{}
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
