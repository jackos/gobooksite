
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

## Copying Slices
Taking a slice of another slice or array uses a pointer, so modifying or appending items will also effect the original. See below examples on how to avoid this

Once the new slice grows beyond the max it creates a new copy, now any modifications to any value won't effect the original

---
```go
a := []int{0, 1, 2}
b := a

b = append(b, 30)
b[2] = 60

fmt.Printf("a: %v len: %v cap: %v\n", a, len(a), cap(a))
fmt.Printf("b: %v len: %v cap: %v\n", b, len(b), cap(b))
```
```output
a: [0 1 2] len: 3 cap: 3
b: [0 1 60 30] len: 4 cap: 6
```
---
If the new slice didn't grow beyond the original, modifications will still effect the original

---
```go
c := []int{0, 1, 2}
d := c

d[2] = 60

fmt.Printf("a: %v len: %v cap: %v\n", c, len(c), cap(c))
fmt.Printf("b: %v len: %v cap: %v\n", d, len(d), cap(d))
```
```output
a: [0 1 60] len: 3 cap: 3
b: [0 1 60] len: 3 cap: 3
```
---
Copy to stop slice modifications effecting original

---
```go
e := []int{1, 2, 3, 4}
f := make([]int, 2)

copy(f, e[2:])
f[0] = 20

fmt.Printf("a: %v\n", e)
fmt.Printf("b: %v\n", f)
```
```output
a: [1 2 3 4]
b: [20 4]
```
---
Array to slice copy

---
```go
g := [5]int{1, 2, 3, 4, 5}
h := []int{99, 98, 97}
elementsCopied := copy(h[1:], g[3:])
fmt.Println("elements copied:",elementsCopied,"\ne:",g,"\nf:",h)
```
```output
elements copied: 2 
e: [1 2 3 4 5] 
f: [99 4 5]
```
---
