package chain

type BoolChain struct {
	BaseChain
	result         bool
	shouldHandlers []*Handler
	minShouldMatch int
}

func NewBoolChain() *BoolChain {
	res := new(BoolChain)
	return res
}

func (c *BoolChain) MinShouldMatch(minShouldMatch int) *BoolChain {
	c.minShouldMatch = minShouldMatch
	return c
}

func (c *BoolChain) Result() bool {
	return c.result
}

func (c *BoolChain) Must(f interface{}, args ...interface{}) *BoolChain {
	h := NewHandler(f, args...)
	c.handlers = append(c.handlers, h)
	return c
}

func (c *BoolChain) Should(f interface{}, args ...interface{}) *BoolChain {
	h := NewHandler(f, args...)
	c.shouldHandlers = append(c.shouldHandlers, h)
	return c
}

func (c *BoolChain) MustMatch(handlers ...*Handler) *BoolChain {
	c.handlers = append(c.handlers, handlers...)
	return c
}

func (c *BoolChain) ShouldMatch(handlers ...*Handler) *BoolChain {
	c.shouldHandlers = append(c.shouldHandlers, handlers...)
	return c
}

func (c *BoolChain) Run() bool {
	for _, f := range c.handlers {
		if !f.BoolCall() {
			return false
		}
	}
	shouldMatched := 0
	for _, f := range c.shouldHandlers {
		if f.BoolCall() {
			shouldMatched++
		}
	}
	return shouldMatched > c.minShouldMatch
}

func (c *BoolChain) Call(f interface{}, args ...interface{}) *BoolChain {
	if !c.result {
		return c
	}
	h := NewHandler(f, args...)
	c.handlers = append(c.handlers, h)
	c.result = h.BoolCall()
	return c
}
