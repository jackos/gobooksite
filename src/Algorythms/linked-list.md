
Slow implementation

---
```go
import "fmt"
```
---

---
```go
type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}
```
---

---
```go
func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}
```
---

---
```go
l := LinkedList{
	Value: 200,
}
m := LinkedList{
	Value: 300,
	Next:  &l,
}

fmt.Println(m.Value, m.Next.Value)
```
```output
300 200
{300 0xc0002ec0e0} 0xc0002ec0e0
```
---
