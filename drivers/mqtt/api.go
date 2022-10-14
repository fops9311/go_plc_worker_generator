package mqttdriver

import (
	"fmt"
	"log"
	"strconv"
)

// bool
func LinkInputbool(id string, def bool) func() bool {
	inputM.Lock()
	defer inputM.Unlock()
	if _, ok := inputBuffer[id]; !ok {
		inputBuffer[id] = fmt.Sprintf("%v", def)

	}
	return func() bool {
		b, err := strconv.ParseBool(inputBuffer[id])
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		return b
	}
}
func LinkOutputbool(id string) func(bool) {
	return func(v bool) {
		Publish <- Message{
			Topic: id,
			Body:  fmt.Sprintf("%v", v),
		}
	}
}

// int
func LinkInputint(id string, def int) func() int {
	inputM.Lock()
	defer inputM.Unlock()
	if _, ok := inputBuffer[id]; !ok {
		inputBuffer[id] = fmt.Sprintf("%v", def)

	}
	return func() int {
		b, err := strconv.ParseInt(inputBuffer[id], 10, 64)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		return int(b)
	}
}
func LinkOutputint(id string) func(int) {
	return func(v int) {
		Publish <- Message{
			Topic: id,
			Body:  fmt.Sprintf("%d", v),
		}
	}
}

// float64
func LinkInputfloat64(id string, def float64) func() float64 {
	inputM.Lock()
	defer inputM.Unlock()
	if _, ok := inputBuffer[id]; !ok {
		inputBuffer[id] = fmt.Sprintf("%v", def)

	}
	return func() float64 {
		b, err := strconv.ParseFloat(inputBuffer[id], 64)
		if err != nil {
			log.Printf("[error] %v\n", err)
			return def
		}
		return (b)
	}
}
func LinkOutputfloat64(id string) func(float64) {
	return func(v float64) {
		Publish <- Message{
			Topic: id,
			Body:  fmt.Sprintf("%f", v),
		}
	}
}
