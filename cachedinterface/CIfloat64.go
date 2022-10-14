package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output float64 data into PLC
type CIfloat64 struct {
	Value float64
	Get     func()float64
	Set     func(float64)
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIfloat64) ReadChan() {
	c.Value = c.Get()
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIfloat64) WriteChan() {
	c.Set(c.Value)
}