package expression

import (
	"errors"
	"strings"
)

const (
	operationDivider  = "|>"
	operandDivider    = " "
	expressionDivider = "="
)

var ExpressionError = errors.New("incorrect expression sintax")

func Translate(input string) (output string, err error) {
	//-- SPLIT
	r, l, err := splitExpr(input)
	if err != nil {
		return "", err
	}
	//-- TRANSLATE
	l, err = translateLeft(l)
	if err != nil {
		return "", err
	}
	//-- JOIN
	output = r + "=" + l

	return output, nil
}

func splitExpr(expr string) (r string, l string, err error) {
	s := strings.Split(expr, expressionDivider)
	if len(s) != 2 {
		return "", "", ExpressionError
	}
	r = s[0]
	l = s[1]
	return r, l, nil
}
func translateLeft(l string) (tExpr string, err error) {
	tExpr = ""
	var prev = ""
	s := strings.Split(l, operationDivider)
	for _, op := range s {
		prev, err = translateOperation(op, prev)
		if err != nil {
			return "", ExpressionError
		}
		tExpr = prev
	}
	return tExpr, nil
}
func translateOperation(op, prev string) (res string, err error) {
	p := strings.Split(op, operandDivider)
	if len(p) < 1 {
		return "", ExpressionError
	}
	fn := strings.ToUpper(p[0])
	if prev != "" {
		p[0] = prev
	} else {
		p = p[1:]
	}
	switch fn {
	case "ADD":
		res = "(" + strings.Join(p, "+") + ")"
	case "SUB":
		res = "(" + strings.Join(p, "-") + ")"
	case "MUL":
		res = "(" + strings.Join(p, "*") + ")"
	case "DIV":
		res = "(" + strings.Join(p, "/") + ")"
	default:
		return "", ExpressionError
	}
	return res, nil
}

/*

Out1=ADD In1 1|>MULT 2 -> Out1 += ((In1+1) * 2)

*/
