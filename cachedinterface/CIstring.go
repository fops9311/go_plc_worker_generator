package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output string data into PLC
type CIstring struct {
	Value string
	Get     func()string
	Set     func(string)
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIstring) ReadChan() {
	c.Value = c.Get()
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIstring) WriteChan() {
	c.Set(c.Value)
}