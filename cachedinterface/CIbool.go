package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output bool data into PLC
type CIbool struct {
	Value bool
	C     chan bool
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIbool) ReadChan() {
	c.Value = <-c.C
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIbool) WriteChan() {
	c.C <- c.Value
}