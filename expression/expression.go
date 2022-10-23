/*
Allows translation of string input expression into some other output expression

# Input expression has sequetial sintax, pipelike, functional style

Input expression consists of left hand side (LHS) and right hand side (RHS)

# LHS mast contain variable name defined by the RHS

RHS consists of operation sequences. Operations are divided with OperandDivider string, defined as package constant

# Opeation substring consists of operation name and args divided with OperandDivider string, defined as package constant

Operation carries previous operation result as the first argument. If operation is in the begining of pipe, it doesn't carry anything.

# Output expression has algebraic sintax, imperative style

Each operation result enclosed in parentheses so the amount of parentheses is not optimal, but order is guaranteed
*/
package expression

import (
	"errors"
	"strings"
)

const (
	OperationDivider  = "|>"
	OperandDivider    = " "
	ExpressionDivider = "="
)

/*
incorrect expression sintax. try something like

val{OperationDivider}{Operation}{Arg} {Arg}...{OperationDivider}{Operation}...
*/
var ExpressionError = errors.New("incorrect expression sintax")

// there must be one expression divider in expression
var ExpressionDividerNotFound = errors.New("expression divider " + ExpressionDivider + " not found in expression")

// one of the operations is unspesified
var OperationNotFound = errors.New("one of the operations not in operations")

/*
takes input string expression and translates it according to the operation defined in package
*/
func Translate(input string) (output string, err error) {
	//-- SPLIT
	l, r, err := splitInput(input)
	if err != nil {
		return "", err
	}
	//-- TRANSLATE
	r, err = translateExpr(r)
	if err != nil {
		return "", err
	}
	//-- JOIN
	output = l + "=" + r

	return output, nil
}

/*splits left and right parts of the input (variable and expression)*/
func splitInput(expr string) (r string, l string, err error) {
	s := strings.Split(expr, ExpressionDivider)
	if len(s) != 2 {
		return "", "", ExpressionDividerNotFound
	}
	l = s[0]
	r = s[1]
	return l, r, nil
}

/*translates the expression part*/
func translateExpr(expr string) (tExpr string, err error) {
	tExpr = ""
	var prev = ""
	s := strings.Split(expr, OperationDivider)
	for _, op := range s {
		prev, err = translateOperation(op, prev)
		if err != nil {
			return "", err
		}
		tExpr = prev
	}
	return tExpr, nil
}

/*translates individual operation*/
func translateOperation(op, prev string) (res string, err error) {
	p := strings.Split(op, OperandDivider)
	if len(p) < 1 {
		return "", ExpressionError
	}
	fn := strings.ToUpper(p[0])
	if prev != "" {
		p[0] = prev
	} else {
		p = p[1:]
	}
	if v, ok := Operations[fn]; ok {
		return v.Translate(p)
	}
	return "", OperationNotFound
}
