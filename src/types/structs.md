
---
```go
import "fmt"
```
---
## Define a struct

---
```go
type person struct {
	name string
	age int
}
```
---
---
```go
fred := person {
	"Fred",
	20,
}
fmt.Println(fred)
```
---
Anonymous struct

---
```go
var person struct {
	name string
	age int
	pet string 
}
```
---


---
```go
person.name = "bob"
person.age = 50
person.pet = "dog"
fmt.Println(person)
```
---
Shorthand

---
```go
pet := struct {
	name string
	kind string
}{
	name: "Fido",
	kind: "dog",
}
fmt.Println(pet)
```
---
## Notes

### Initialization
- No difference between assigning empty struct literal and not assigning a value, both initialize all fields to 0

### Conversions
- Structs can be type converted if fields are in the same order, with same naming and types

### Comparisons
- Structs are comparable if of the same type, and no field contains a slice, map, channel or function
- Anonymous structs can be compared to a typed struct if contains same order, names and types

---
```go
type firstPerson struct {
	name string
	age int
}
```
---
---
```go
f := firstPerson{
	name: "Bob",
	age: 50,
}

var g struct {
	name string
	age int
}

g = f
fmt.Println(g == f)
```
---
