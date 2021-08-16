
---
```go
import "fmt"
```
---
---
```go
x := 10
pointer := &x
fmt.Println("memory address:", pointer) 
fmt.Println("value:", *pointer)
```
---
Nil pointer behaviour

---
```go
var x *int
```
---
---
```go
fmt.Println(x == nil)
```
---
---
```go
fmt.Println(*x)
```
---
Check for a nil pointer

---
```go
	var x *int
	z := 20
	x = &z
	if x != nil {
		fmt.Println(*x)
	}
```
---
Initialize and provide a pointer to the given primitive type with the 'new' keyword

---
```go
var x = new(int)
```
---
---
```go
fmt.Println(*x)
```
---
---
```go
Create a pointer instance with a stuct
```
---
---
```go
type Foo struct {
	name string
}
```
---
---
```go
x := &Foo{"billy"}
fmt.Println(x.name)
```
---
Workaround struct with a field that's a pointer to a primitive type

---
```go
type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}
```
---
Method 1: Place value in a variable and use its address

---
```go
	x := "Perry"
	p := person{
		FirstName:  "Pat",
		MiddleName: &x,
		LastName:   "Peterson",
	}
	fmt.Println(*p.MiddleName)
```
---
Method 2: Use a helper function that returns the address to the parameter which is initialized into memory on the function call

---
```go
// Helper function to return pointer to a string constant, allowing more concise syntax
func p(s string) *string {
	return &s
}
```
---
---
```go
// ## Doesn't work in gobook, says too many arguments
p := person{
	FirstName: "Pat",
	MiddleName: p("Perry"),
	LastName: "Peterson",
}
```
---


