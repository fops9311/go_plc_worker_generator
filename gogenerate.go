//go:generate go run ./generators/interfaces/main.go bool
//go:generate go run ./generators/interfaces/main.go float64
//go:generate go run ./generators/interfaces/main.go float32
//go:generate go run ./generators/interfaces/main.go int
//go:generate go run ./generators/interfaces/main.go string

//go:generate go run ./generators/plcs/main.go test_unit0
package main
