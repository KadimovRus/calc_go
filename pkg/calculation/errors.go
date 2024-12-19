package calculation

import "errors"

var (
	ErrInvalidExpression      = errors.New("expression is not valid")
	ErrUnbalancedParentheses  = errors.New("unbalanced parentheses")
	ErrDivisionByZero         = errors.New("division by zero")
	ErrInvalidTypeOfOperation = errors.New("invalid type of operation")
)
