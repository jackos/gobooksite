
---
```go
import "fmt"
```
---

Recovering from a panic so divide by 0 doesn't crash the program. Not good practise, generally best to use error handling.

---
```go
func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Panic in div60: %+v\n", v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
```
```output
60
30
Panic in div60: runtime error: integer divide by zero
10
```
---