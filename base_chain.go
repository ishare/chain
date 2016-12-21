package chain

type BaseChain struct {
	handlers []*Handler // registered func executed when all processes come to barrier
}

func NewBaseChain() *BaseChain {
	res := new(BaseChain)
	return res
}

func (c *BaseChain) Append(f interface{}, args ...interface{}) *BaseChain {
	h := NewHandler(f, args...)
	c.handlers = append(c.handlers, h)
	return c
}

func (c *BaseChain) Prepend(f interface{}, args ...interface{}) *BaseChain {
	h := NewHandler(f, args...)
	c.handlers = append([]*Handler{h}, c.handlers...)
	return c
}

func (c *BaseChain) Run() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	for _, f := range c.handlers {
		f.Call()
	}
}

func (c *BaseChain) Call(f interface{}, args ...interface{}) *BaseChain {
	h := NewHandler(f, args...)
	c.handlers = append(c.handlers, h)
	h.Call()
	return c
}
