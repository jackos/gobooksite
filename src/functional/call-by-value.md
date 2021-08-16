
---
```go
import "fmt"
```
---
---
```go
type person struct{
	age int
	name string
}
```
---
---
```go
func changeValues(i int, s string, p person){
	i = i * 2
	s = "Changed string"
	p.age = 600
	p.name = "Changed Name"
}
```
---
Modifying the values passed in does not effect the values in the outer scope

---
```go
p := person{20, "Original Name"}
i := 2
s := "Original string"

changeValues(i, s, p)
fmt.Println(i, s, p)
```
---
Modifying a map does effect the original, as it's a pointer

---
```go
func modMap(m map[int]string){
	m[0] = "Changed value"
	m[1] = "This value will be deleted"
	m[2] = "Added value"
	delete(m, 1)
}
```
---
---
```go
m := map[int]string{
	0: "first",
	1: "second",
}
modMap(m)
fmt.Println(m)
```
---
Modifying a slice will only effect items that have already been initialized, can't add new items.

---
```go
func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}
```
---
---
```go
s := []int{1,2,3}
modSlice(s)
fmt.Println(s)
```
---
Experimenting
