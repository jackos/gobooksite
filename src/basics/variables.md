
### Declare and initialize

---
```go
var a int = 10
```
---
### Inferred type

---
```go
var b = 10
fmt.Println(reflect.TypeOf(b))
```
```output
int
```
---
### Initialized to 0

---
```go
var c int
fmt.Println(c)
```
```output
0
```
---
### Multi initializer

---
```go
var d, e int = 10, 20
```
---
### Inferred multi

---
```go
var f, g = 10, "wow"
fmt.Printf("%v\n%v",reflect.TypeOf(f), reflect.TypeOf(g))
```
```output
int
string
```
---
### Multiline initializer

---
```go
var (
	h    int
	i        = 20
	j    int = 30
	k, l     = 40, "hello"
	m, n string
)
fmt.Println(h, i, j, k, l, m, n)
```
```output
0 20 30 40 hello
```
---
### Shorthand initializer

---
```go
o := 10
p, q := 30, "hello"
fmt.Println(o, p, q)
```
```output
10 30 hello
```
---
### Const declaration

---
```go
// Like C this swaps out the consts with literals during compile time
const (
	idKey   string = "id"
	nameKey        = "name"
)
fmt.Println(idKey, nameKey)
```
```output
id name
```
---
