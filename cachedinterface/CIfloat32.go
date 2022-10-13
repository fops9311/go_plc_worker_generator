package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output float32 data into PLC
type CIfloat32 struct {
	Value float32
	C     chan float32
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIfloat32) ReadChan() {
	c.Value = <-c.C
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIfloat32) WriteChan() {
	c.C <- c.Value
}