package expression

import "errors"

type Operation struct {
	ArgLimits                           // Arguments
	TrFn      func(arg []string) string // Translation function
}

var ArgsValidationError = errors.New("agrument number invalid")

func NewOperation(minArgs, maxArgs int, fn func([]string) string) Operation {
	return Operation{
		TrFn: fn,
		ArgLimits: ArgLimits{
			minArgs: minArgs,
			maxArgs: maxArgs,
		},
	}
}
func (op Operation) Translate(args []string) (res string, err error) {
	if e := op.ValidateArgs(args); e != nil {
		return "", e
	}
	return op.TrFn(args), nil
}

type ArgLimits struct {
	minArgs int
	maxArgs int
}

func (o ArgLimits) ValidateArgs(args []string) error {
	l := len(args)
	if l > o.maxArgs || l < o.minArgs {
		return ArgsValidationError
	}
	return nil
}
