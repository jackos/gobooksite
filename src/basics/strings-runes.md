
---
```go
import "fmt"
```
---
Converting string to runes retains emoji's as a single item

---
```go
s := "Hello, 🌞"
bs := []byte(s)
rs := []rune(s)
fmt.Println("Bytes:", bs)
fmt.Println("Runes:", rs)
```
---
Converting back to string from either still formats correctly

---
```go
bytes := string(bs)
runes := string(rs)
fmt.Println("Bytes:",bytes)
fmt.Println("Runes:",runes)
```
---
---
```go
s := "Hello there"
for i, v := range(s[6:9]) {
	fmt.Printf("%c", v)
}
```
---
---
```go
x := 50
xString := string(x)
fmt.Println(xString)
```
```output
2
```
---

