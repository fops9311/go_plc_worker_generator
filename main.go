package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fops9311/go_plc_worker_generator/example"
	"github.com/fops9311/go_plc_worker_generator/mqttpubsub"
)

func init() {
	go mqttpubsub.Init()
}
func main() {

	plc := example.NewPLC()
	plc_in := false
	plc_out := MQTTCache{
		c: &mqttpubsub.Publish,
	}
	go func() {
		for message := range mqttpubsub.Subsribe {
			switch message.Topic {
			case "metric/test_plc/plc/in/b":
				if b, err := strconv.ParseBool(message.Body); err == nil {
					plc_in = b
				} else {
					plc_in = false
				}
			}
			//log.Printf("got from mqtt:%s\n", message)
		}
	}()
	go func() {
		for {
			out := <-plc.Out.C
			plc_out.update(mqttpubsub.Message{Topic: "metric/test_plc/plc/out/b", Body: fmt.Sprintf("%v", out)})
		}
	}()
	go func() {
		for {
			//log.Println("tick...")
			go ChanAssignValue(&plc.In.C, plc_in)
			go plc.Tick()
			//log.Println("active goroutines: ", runtime.NumGoroutine())
			<-time.NewTimer(time.Microsecond * 300).C
		}
	}()
	for {
	}
}

var fliper bool

func ChanWriter(c *chan bool) {
	log.Printf("SEND: %v\n", fliper)
	*c <- fliper
	fliper = !fliper
}

func ChanAssignValue(c *chan bool, value bool) {
	*c <- value
}
func ChanReader(c *chan bool) {
	log.Printf("GOT: %v\n", <-*c)
}

type MQTTCache struct {
	mem mqttpubsub.Message
	c   *chan mqttpubsub.Message
}

func (m *MQTTCache) update(new mqttpubsub.Message) {
	if m.mem.Topic != new.Topic || m.mem.Body != new.Body {
		m.mem = new
		*(m.c) <- m.mem
	}
}
