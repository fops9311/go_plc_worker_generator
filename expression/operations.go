package expression

import "strings"

var Operations = map[string]Operation{
	"SET": NewOperation(1, 1,
		func(args []string) string {
			return "(" + strings.Join(args, "") + ")"
		},
	),
	"ADD": NewOperation(2, 50,
		func(args []string) string {
			return "(" + strings.Join(args, "+") + ")"
		},
	),
	"SUB": NewOperation(2, 50,
		func(args []string) string {
			return "(" + strings.Join(args, "-") + ")"
		},
	),
	"DIV": NewOperation(2, 50,
		func(args []string) string {
			return "(" + strings.Join(args, "/") + ")"
		},
	),
	"MUL": NewOperation(2, 50,
		func(args []string) string {
			return "(" + strings.Join(args, "*") + ")"
		},
	),
	"INTTOFLOAT": NewOperation(1, 1,
		func(args []string) string {
			return "INTTOFLOAT(" + strings.Join(args, "") + ")"
		},
	),
	"FLOATTOINT": NewOperation(1, 1,
		func(args []string) string {
			return "FLOATTOINT(" + strings.Join(args, "") + ")"
		},
	),
	"NOT": NewOperation(1, 1,
		func(args []string) string {
			return "!(" + strings.Join(args, "") + ")"
		},
	),
	"GT": NewOperation(2, 2,
		func(args []string) string {
			return "(" + strings.Join(args, ">") + ")"
		},
	),
	"LT": NewOperation(2, 2,
		func(args []string) string {
			return "(" + strings.Join(args, "<") + ")"
		},
	),
	"GE": NewOperation(2, 2,
		func(args []string) string {
			return "(" + strings.Join(args, ">=") + ")"
		},
	),
	"LE": NewOperation(2, 2,
		func(args []string) string {
			return "(" + strings.Join(args, "<=") + ")"
		},
	),
	"OR": NewOperation(2, 2,
		func(args []string) string {
			return "(" + strings.Join(args, "||") + ")"
		},
	),
	"AND": NewOperation(2, 2,
		func(args []string) string {
			return "(" + strings.Join(args, "&&") + ")"
		},
	),
}
