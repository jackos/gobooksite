






# Using a map as a set

---
```go
import "fmt"
```
---
Inserting same value into map doesn't result in duplicates

---
```go
intSet := map[int]bool{}
vals := []int{1, 1, 2, 2, 3}
for _, v := range vals { 
	intSet[v] = true
}
fmt.Println(intSet)
```
---





Trying to retrieve a key that doesn't exist gives false

---
```go
fmt.Println(intSet[5])
```
---
Use a struct as a set

---
```go
intSet := map[int]struct{}{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
	intSet[v] = struct{}{}
}
fmt.Println(intSet)
```
```output
map[1:{} 2:{} 3:{} 5:{} 7:{} 8:{} 9:{} 10:{}]
```
---
