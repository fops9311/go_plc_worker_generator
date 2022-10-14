package plc

import (
	"log"
	"time"

	cif "github.com/fops9311/go_plc_worker_generator/cachedinterface"
	driver "github.com/fops9311/go_plc_worker_generator/drivers/mqtt"
)

type TEST_UNIT0 struct {
	In cif.CIbool;
	
	Out cif.CIbool;
	State cif.CIint;
	
	StateStartTime time.Time
	tickInterval100ms int
}

func NewPLC_TEST_UNIT0(tickInterval100ms int) *TEST_UNIT0 {
	return &TEST_UNIT0{
		In: cif.CIbool{
			Value: false,
			Get: driver.LinkInputbool("metric/test_plc/plc/in/b",false),//func()bool{return false},
			},
		
		Out:  cif.CIbool{
			Value: false,
			Set: driver.LinkOutputbool("metric/test_plc/plc/out/b"),//func(v bool){log.Print("Set Out = ");fmt.Println(v)},
			},
		State:  cif.CIint{
			Value: 0,
			Set: driver.LinkOutputint("metric/test_plc/plc/state/pv"),//func(v int){log.Print("Set State = ");fmt.Println(v)},
			},
		
		StateStartTime: time.Now(),
		tickInterval100ms:tickInterval100ms,
	}
}
func (plc *TEST_UNIT0) readInputs() {
	plc.In.ReadChan()
	
}
func (plc *TEST_UNIT0) writeOutputs() {
	plc.Out.WriteChan()
	plc.State.WriteChan()
	
}
func (plc *TEST_UNIT0) tick() {
	plc.readInputs()
	plc.logic()
	plc.writeOutputs()
}
func (plc *TEST_UNIT0) logic() {
	switch plc.State.Value {
	
	case 0: // Out off, wait for In == true
		plc.Out.Value = false
		
		if (plc.In.Value){ plc.State.Value = 1; plc.resetStateTimer()} 
		
	
	case 1: // Out on, wait for In == false
		plc.Out.Value = true
		
		if (!plc.In.Value){ plc.State.Value = 2; plc.resetStateTimer()} 
		
	
	case 2: // Out on, wait for In == true
		plc.Out.Value = true
		
		if (plc.In.Value){ plc.State.Value = 3; plc.resetStateTimer()} 
		
	
	case 3: // Out off, wait for In == false
		plc.Out.Value = false
		
		if (!plc.In.Value){ plc.State.Value = 0; plc.resetStateTimer()} 
		
	
	default:
		plc.State.Value = 0
	}
}
func (plc *TEST_UNIT0) resetStateTimer() {
	log.Printf("%s passed. state changed to %d", time.Now().Sub(plc.StateStartTime),plc.State.Value)
	plc.StateStartTime = time.Now()
}
func (plc *TEST_UNIT0) Start() {
	for {
		plc.tick()
		<-time.NewTimer(time.Millisecond * time.Duration(plc.tickInterval100ms*100)).C
	}
}