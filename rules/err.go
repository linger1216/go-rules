package rules

import "fmt"

var (
	ErrInvalidType     = fmt.Errorf("invalid type")
	ErrInvalidOperator = fmt.Errorf("invalid operator")
	ErrInvalidRule     = fmt.Errorf("invalid rule")
)
