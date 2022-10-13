package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output int data into PLC
type CIint struct {
	Value int
	C     chan int
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIint) ReadChan() {
	c.Value = <-c.C
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIint) WriteChan() {
	c.C <- c.Value
}