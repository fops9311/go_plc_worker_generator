//go:generate go run ./generators/interfaces/main.go bool
//go:generate go run ./generators/interfaces/main.go float64
//go:generate go run ./generators/interfaces/main.go float32
//go:generate go run ./generators/interfaces/main.go int
//go:generate go run ./generators/interfaces/main.go string

//go:generate go build -o decl ./declarations/
//go:generate ./decl ./declarations/plcs.json ./declarations/types/ ./declarations/instances/ ./declarations/default.template ./plc/

package main
