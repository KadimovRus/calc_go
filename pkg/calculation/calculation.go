package calculation

import (
	"regexp"
	"strconv"
	"strings"
)

func applyOperation(a, b float64, op string) (float64, error) {
	if op == "+" {
		return a + b, nil
	} else if op == "-" {
		return a - b, nil
	} else if op == "*" {
		return a * b, nil
	} else if op == "/" {
		if b == 0 {
			return 0, ErrInvalidExpression
		}
		return a / b, nil
	}
	return 0, ErrInvalidExpression
}

func volumeOperation(op string) int {
	if strings.Contains("+-", op) {
		return 1
	} else if strings.Contains("*/", op) {
		return 2
	}
	return 0
}

func Calc(expression string) (float64, error) {
	re := regexp.MustCompile(`[\d.]+|[+/*()-]`)
	symbols := re.FindAllString(expression, -1)

	var values []float64
	var ops []string

	for i, s := range symbols {
		switch s {
		case "(":
			ops = append(ops, s)
		case ")":
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				if len(values) < 2 {
					return 0, ErrInvalidExpression
				}
				val, err := applyOperation(values[len(values)-2], values[len(values)-1], ops[len(ops)-1])
				if err != nil {
					return 0, err
				}
				values = values[:len(values)-2]
				values = append(values, val)
				ops = ops[:len(ops)-1]
			}
			if len(ops) == 0 || ops[len(ops)-1] != "(" {
				return 0, ErrInvalidExpression
			}
			ops = ops[:len(ops)-1]

		case "+", "-", "*", "/":
			if i == 0 || i == len(symbols)-1 || (i > 0 && strings.Contains("+-*/", symbols[i-1])) {
				return 0, ErrInvalidExpression
			}
			for len(ops) > 0 && volumeOperation(ops[len(ops)-1]) >= volumeOperation(s) {
				if len(values) < 2 {
					return 0, ErrInvalidExpression
				}
				val, err := applyOperation(values[len(values)-2], values[len(values)-1], ops[len(ops)-1])
				if err != nil {
					return 0, err
				}
				values = values[:len(values)-2]
				values = append(values, val)
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, s)
		default:
			num, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return 0, ErrInvalidExpression
			}
			values = append(values, num)
		}
	}

	for len(ops) > 0 {
		if ops[len(ops)-1] == "(" {
			return 0, ErrInvalidExpression
		}
		if len(values) < 2 {
			return 0, ErrInvalidExpression
		}
		val, err := applyOperation(values[len(values)-2], values[len(values)-1], ops[len(ops)-1])
		if err != nil {
			return 0, err
		}
		values = values[:len(values)-2]
		values = append(values, val)
		ops = ops[:len(ops)-1]
	}

	if len(values) > 0 {
		return values[0], nil
	}
	return 0, ErrInvalidExpression
}
