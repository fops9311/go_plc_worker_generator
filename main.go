package main

import (
	"fmt"
	"os"
	"os/signal"

	internaldb "github.com/fops9311/go_plc_worker_generator/drivers/internaldb"
	"github.com/fops9311/go_plc_worker_generator/plc"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	for k := range plc.PLCS {
		go plc.PLCS[k].Start()
	}

	select {
	case sig := <-c:
		fmt.Printf("Got %s signal. Aborting...\n", sig)
		var end chan bool = make(chan bool)
		internaldb.Close <- end
		<-end
	}

}
