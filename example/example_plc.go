package example

import (
	"log"
	"time"

	cif "github.com/fops9311/go_plc_worker_generator/cachedinterface"
)

type plc struct {
	In             cif.CIbool
	Out            cif.CIbool
	State          cif.CIint
	StateStartTime time.Time
}

func NewPLC() *plc {
	return &plc{
		In:             cif.CIbool{Value: false, C: make(chan bool)},
		Out:            cif.CIbool{Value: false, C: make(chan bool)},
		State:          cif.CIint{Value: 0, C: make(chan int)},
		StateStartTime: time.Now(),
	}
}
func (p *plc) NewTestBinaryPLC() {

}
func (p *plc) ReadInputs() {
	p.In.ReadChan()
}
func (p *plc) WriteOutputs() {
	p.Out.WriteChan()
}
func (p *plc) Tick() {
	p.ReadInputs()
	p.Logic()
	p.WriteOutputs()
}
func (p *plc) Logic() {
	switch p.State.Value {
	case 0:
		p.Out.Value = false
		if p.In.Value {
			p.State.Value = 1
			p.debugStateTime()
		}
	case 1:
		p.Out.Value = true
		if !p.In.Value {
			p.State.Value = 2
			p.debugStateTime()
		}
	case 2:
		p.Out.Value = true
		if p.In.Value {
			p.State.Value = 3
			p.debugStateTime()
		}
	case 3:
		p.Out.Value = false
		if !p.In.Value {
			p.State.Value = 0
			p.debugStateTime()
		}
	}
}
func (p *plc) debugStateTime() {
	log.Printf("%s passed", time.Now().Sub(p.StateStartTime))
	p.StateStartTime = time.Now()
}
