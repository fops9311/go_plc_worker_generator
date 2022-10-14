package main

import (
	"github.com/fops9311/go_plc_worker_generator/plc"
)

func main() {
	plc.NewPLC_TEST_UNIT0(3).Start()
}
