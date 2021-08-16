
---
```go
import "fmt"
```
---
---
```go
type IntTree struct {
	val         int
	left, right *IntTree
}
```
---
---
```go
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}
```
---
---
```go
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		fmt.Println("No Match on", val)
		return false
	case val < it.val:
		fmt.Println("Going left", val, "<", it.val)
		return it.left.Contains(val)
	case val > it.val:
		fmt.Println("Going right", val, ">", it.val)
		return it.right.Contains(val)
	default:
		fmt.Println("Matched", val)
		return true
	}
}
```
---
---
```go
func main() {
	t := IntTree{}
	t.Insert(10)
	t.Insert(10)
	t.Insert(9)
	t.Insert(5)
	fmt.Println(t.Contains(6))
	fmt.Println(t.Contains(5))
}
```
---
