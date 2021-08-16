
## Intro
Taking a slice of another slice or array uses a pointer, so modifying or appending items will also effect the original. See below examples on how to avoid this

---
```go
import "fmt"
```
---
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
a := []int{0, 1, 2}
b := a

b[2] = 60

fmt.Printf("a: %v len: %v cap: %v\n", a, len(a), cap(a))
fmt.Printf("b: %v len: %v cap: %v\n", b, len(b), cap(b))
```
```output
a: [0 1 60] len: 3 cap: 3
b: [0 1 60] len: 3 cap: 3
```
---
Copy to stop slice modifications effecting original

---
```go
a := []int{1, 2, 3, 4}
b := make([]int, 2)

copy(b, a[2:])
b[0] = 20

fmt.Printf("a: %v\n", a)
fmt.Printf("b: %v\n", b)
```
```output
a: [1 2 3 4]
b: [20 4]
```
---
Array to slice copy

---
```go
e := [5]int{1, 2, 3, 4, 5}
f := []int{99, 98, 97}
elementsCopied := copy(f[1:], e[3:])
fmt.Println("elements copied:",elementsCopied,"\ne:",e,"\nf:",f)
```
```output
elements copied: 2 
e: [1 2 3 4 5] 
f: [99 4 5]
```
---
