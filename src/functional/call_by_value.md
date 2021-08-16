
# Call by value

Like C, Go is pass by value

---
```go
type person struct{
	age int
	name string
}

func changeValues(i int, s string, p person) {
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
```output
2 Original string {20 Original Name}
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
```output
map[0:Changed value 2:Added value]
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
sl := []int{1,2,3}
modSlice(sl)
fmt.Println(sl)
```
```output
[2 4 6]
```
---
