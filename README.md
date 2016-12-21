# chain
call chain, bool chain, to be extend

Usage:
-----------------
basic chain call:
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

bool chain call: (stop call when there is a func return false)
```go
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