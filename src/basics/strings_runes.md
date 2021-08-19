
# Strings

Converting string to runes retains emoji's as a single item

---
```go
s := "Hello, 🌞"
bs := []byte(s)
rs := []rune(s)
fmt.Println("Bytes:", bs)
fmt.Println("Runes:", rs)
```
```output
Bytes: [72 101 108 108 111 44 32 240 159 140 158]
Runes: [72 101 108 108 111 44 32 127774]
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
```output
Bytes: Hello, 🌞
Runes: Hello, 🌞
```
---
### String slice

---
```go
h := "Hello there"
for _, v := range(h[6:9]) {
	fmt.Printf("%c", v)
}
```
```output
the
```
---
Converting an int to a sting with give the UTF-8 representation

---
```go
x := 50
fmt.Println(string(x))
xString := fmt.Sprintf("%v", x)
fmt.Println(xString)
```
```output
2
50
```
---
