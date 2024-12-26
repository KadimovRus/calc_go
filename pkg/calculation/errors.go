package calculation

import "errors"

var (
	ErrInvalidExpression   = errors.New("expression is not valid")
	ErrInternalServerError = errors.New("internal server error")
)
