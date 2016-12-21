package chain

type BoolChain struct {
	BaseChain
	ret bool
}

func NewBoolChain() *BoolChain {
	res := new(BoolChain)
	return res
}

func (c *BoolChain) Run() bool {
	for _, f := range c.handlers {
		if !f.BoolCall() {
			return false
		}
	}
	return true
}

func (c *BoolChain) Call(f interface{}, args ...interface{}) *BoolChain {
	if !c.ret {
		return c
	}
	h := NewHandler(f, args...)
	c.handlers = append(c.handlers, h)
	c.ret = h.BoolCall()
	return c
}
