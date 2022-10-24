package internaldriver

import (
	"fmt"
	"log"
	"strconv"
)

// bool
func LinkInputbool(id string, def bool) func() bool {
	return func() bool {
		s, err := readVal("bool", id)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		b, err := strconv.ParseBool(s)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		return b
	}
}
func LinkOutputbool(id string) func(bool) {
	return func(v bool) {
		err := updateVal("bool", id, fmt.Sprintf("%v", v))
		if err != nil {
			log.Printf("[error] %v\n", err)
		}
	}
}

// int
func LinkInputint(id string, def int) func() int {
	return func() int {
		s, err := readVal("int", id)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		b, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		return int(b)
	}
}
func LinkOutputint(id string) func(int) {
	return func(v int) {
		err := updateVal("bool", id, fmt.Sprintf("%d", v))
		if err != nil {
			log.Printf("[error] %v\n", err)
		}
	}
}

// float64
func LinkInputfloat64(id string, def float64) func() float64 {
	return func() float64 {
		s, err := readVal("int", id)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		b, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		return (b)
	}
}
func LinkOutputfloat64(id string) func(float64) {
	return func(v float64) {
		err := updateVal("bool", id, fmt.Sprintf("%f", v))
		if err != nil {
			log.Printf("[error] %v\n", err)
		}
	}
}
