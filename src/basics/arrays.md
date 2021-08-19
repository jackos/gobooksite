
# Arrays

### Initialize array to 0 values

---
```go
a := [3]int{}
fmt.Println(a)
```
```output
[0 0 0]
```
---
### Set to width of provided literal

---
```go
b := [...]int{5 + 5, 20, 30}
fmt.Println(b)
```
```output
[10 20 30]
```
---
### Compare arrays

---
```go
// Arrays can be compared in go, but Slices cannot
println(a == b)
```
```output
false
```
---
### Set by index, blank are initialized to 0

---
```go
d := [6]int{1, 2: 4, 5, 5: 100}
fmt.Println(d)
```
```output
[1 0 4 5 0 100]
```
---
### Multidimensional array

---
```go
e := [2][4]int{}
fmt.Println(e)
```
```output
[[0 0 0 0] [0 0 0 0]]
```
---
_Go has a simple multidimensional array implementation which incurs performance penalities_
