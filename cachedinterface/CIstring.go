package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output string data into PLC
type CIstring struct {
	Value string
	C     chan string
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIstring) ReadChan() {
	c.Value = <-c.C
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIstring) WriteChan() {
	c.C <- c.Value
}