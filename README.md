# chain
basic call chain, bool chain, to be extend..

Usage:
-----------------
**basic chain call:**
```go
func TestBaseChainCall(t *testing.T) {
	NewBaseChain().Call(test1, 1).Call(test1, 2)
}

func TestBaseChain(t *testing.T) {
	c := NewBaseChain()
	c.Append(test1, 1)
	c.Append(test1, 2)
	c.Run()
}

func test1(a int) {
	fmt.Println(a)
}
```

**bool chain call:** 
Use `Must`, `Should` to restrict how many conditions it's gonna to match.
If any `Must` func returns false, it breaks execution immediately.
```go
func TestBoolChainMust(t *testing.T) {
	NewBoolChain().Must(less, 1, 2).MinShouldMatch(1).Should(less, 5, 2).Should(2, 3)
}

func TestBoolChainCall(t *testing.T) {
	NewBoolChain().Call(less, 1, 2).Call(less, 5, 2).Call(2, 3)
}

func TestBoolChain(t *testing.T) {
	c := NewBoolChain()
	c.Append(less, 1, 2)
	c.Append(less, 5, 2)
	c.Append(less, 2, 3)
	fmt.Println(c.Run())
}

func less(a, b int) bool {
	fmt.Printf("compare %d, %d\n", a, b)
	return a < b
}
```

