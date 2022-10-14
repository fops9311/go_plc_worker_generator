package cachedinterface
//CI stands for cached interface\n
//This is a way to input or output bool data into PLC
type CIbool struct {
	Value bool
	Get     func()bool
	Set     func(bool)
}
//Use this method in every tick of plc logic to input data into PLC
func (c *CIbool) ReadChan() {
	c.Value = c.Get()
}
//Use this method in every tick of plc logic to output data into PLC
func (c *CIbool) WriteChan() {
	c.Set(c.Value)
}