package mqttdriver

import (
	"sync"
)

var inputBuffer = make(map[string]string)
var inputM sync.Mutex
var init_flag bool = false

func init() {
	inputM.Lock()
	defer inputM.Unlock()
	if !init_flag {
		go Init()
		go func() {
			for message := range Subsribe {
				inputM.Lock()
				inputBuffer[message.Topic] = message.Body
				inputM.Unlock()
			}
		}()
		init_flag = true
	}
}
