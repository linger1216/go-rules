package rules

import (
	"fmt"
	"strings"
	"testing"
)

func getFunc(funcName string, judge Judge) func(left interface{}, right IFiled) (bool, error) {
	switch strings.ToLower(funcName) {
	case "eq", "==":
		return judge.EQ
	case "ne", "!=":
		return judge.NE
	case "gt", ">":
		return judge.GT
	case "ge", ">=":
		return judge.GE
	case "lt", "<":
		return judge.LT
	case "le", "<=":
		return judge.LE
	case "contain":
		return judge.Contain
	case "prefix":
		return judge.Prefix
	case "suffix":
		return judge.Suffix
	case "regex":
		return judge.Regex
	case "in":
		return judge.In
	default:
		return nil
	}
}

func TestStringJudge(t *testing.T) {
	tests := []struct {
		name   string
		left   interface{}
		right  IFiled
		expect bool
	}{
		// TODO: Add test cases.
		{
			"==",
			"hello",
			NewValueString("hello"),
			true,
		},

		{
			"==",
			"",
			NewValueNull(),
			true,
		},

		{
			"!=",
			"hello",
			NewValueString("world"),
			true,
		},

		{
			"!=",
			"hello",
			NewValueNull(),
			true,
		},

		{
			">",
			"z",
			NewValueString("a"),
			true,
		},

		{
			">=",
			"z",
			NewValueString("a"),
			true,
		},

		{
			"<",
			"a",
			NewValueString("z"),
			true,
		},

		{
			"<=",
			"a",
			NewValueString("a"),
			true,
		},

		{
			"contain",
			"abc",
			NewValueString("ab"),
			true,
		},

		{
			"prefix",
			"hello",
			NewValueString("he"),
			true,
		},

		{
			"prefix",
			"hello",
			NewValueListStrings("[h,he,hel]"),
			true,
		},

		{
			"suffix",
			"world",
			NewValueString("ld"),
			true,
		},

		{
			"suffix",
			"world",
			NewValueListStrings("[d,ld,rld]"),
			true,
		},
		{
			"in",
			"a",
			NewValueString("a"),
			true,
		},
		{
			"in",
			"a",
			NewValueListStrings("[a,b,c]"),
			true,
		},
	}
	s := &StringJudge{}
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
