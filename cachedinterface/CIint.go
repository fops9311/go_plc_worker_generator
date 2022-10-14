package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output int data into PLC
type CIint struct {
	Value int
	Get     func()int
	Set     func(int)
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIint) ReadChan() {
	c.Value = c.Get()
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIint) WriteChan() {
	c.Set(c.Value)
}