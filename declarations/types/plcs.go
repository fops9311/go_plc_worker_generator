package plc

var PLCS map[string]PlcStartable = make(map[string]PlcStartable)

type PlcStartable interface {
	Start()
}

func init() {
	{{range $val := .Plcs}}PLCS["{{.Type | ToUpper}}_{{.Id | ToUpper}}"] = NewPLC_{{.Type | ToUpper}}_{{.Id | ToUpper}}({{.TickInterval100ms}})
	{{end}}
}
