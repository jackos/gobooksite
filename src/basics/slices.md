
### Create from an array

---
```go
array := [4]int{1,2,3,4}
slice := array[:3]
fmt.Println("array:", array)
fmt.Println("slice:", slice)
```
```output
array: [1 2 3 4]
slice: [1 2 3]
```
---
### Shorthand initializer

---
```go
a := []int{10, 0, 0}
```
```output
[10 0 0]
```
---
### Append two slices together

---
```go
c := []int{10, 20, 30}
d := []int{1, 2: 3, 4, 5: 10}
e := append(c, d...)
fmt.Println(e)
```
```output
[10 20 30 1 0 3 4 0 10]
```
---
### Initialize with make

---
```go
f := make([]int, 5)
fmt.Printf("%v len: %v cap %v\n", f, len(f), cap(f))
```
```output
[0 0 0 0 0] len: 5 cap 5
```
---
### Make with capacity

---
```go
g := make([]int, 5, 32)
fmt.Printf("%v len: %v cap %v", g, len(g), cap(g))
fmt.Println()
```
```output
[0 0 0 0 0] len: 5 cap 32
```
---

This sets asside 32 ints in memory to save the runtime doing extra work moving things around in memory, appending to a slice would otherwise have to create a new slice.
