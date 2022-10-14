package plc

import (
	"fmt"
	"log"
	"time"

	cif "github.com/fops9311/go_plc_worker_generator/cachedinterface"
)

type TEST_UNIT0 struct {
	In cif.CIbool;
	
	Out cif.CIbool;
	State cif.CIint;
	
	StateStartTime time.Time
}

func NewPLC_TEST_UNIT0() *TEST_UNIT0 {
	return &TEST_UNIT0{
		In: cif.CIbool{
			Value: false,
			Get: func()bool{return false},
			},
		
		Out:  cif.CIbool{
			Value: false,
			Set: func(v bool){log.Print("Set Out = ");fmt.Println(v)},
			},
		State:  cif.CIint{
			Value: 0,
			Set: func(v int){log.Print("Set State = ");fmt.Println(v)},
			},
		
		StateStartTime: time.Now(),
	}
}
func (plc *TEST_UNIT0) ReadInputs() {
	plc.In.ReadChan()
	
}
func (plc *TEST_UNIT0) WriteOutputs() {
	plc.Out.WriteChan()
	plc.State.WriteChan()
	
}
func (plc *TEST_UNIT0) Tick() {
	plc.ReadInputs()
	plc.Logic()
	plc.WriteOutputs()
}
func (plc *TEST_UNIT0) Logic() {
	switch plc.State.Value {
	
	case 0:
		plc.Out.Value = false
		
		if (plc.In.Value){ plc.State.Value = 1; plc.ResetStateTimer()} 
		
	
	case 1:
		plc.Out.Value = true
		
		if (!plc.In.Value){ plc.State.Value = 2; plc.ResetStateTimer()} 
		
	
	case 2:
		plc.Out.Value = true
		
		if (plc.In.Value){ plc.State.Value = 3; plc.ResetStateTimer()} 
		
	
	case 3:
		plc.Out.Value = false
		
		if (!plc.In.Value){ plc.State.Value = 0; plc.ResetStateTimer()} 
		
	
	default:
		plc.State.Value = 0
	}
}
func (plc *TEST_UNIT0) ResetStateTimer() {
	log.Printf("%s passed", time.Now().Sub(plc.StateStartTime))
	plc.StateStartTime = time.Now()
}
