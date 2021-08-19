
# Looping 

Go has a single loop statement: for

### Standard for statement

---
```go
for i:= 0; i < 5; i++ {
	fmt.Println(i)
}
```
```output
0
1
2
3
4
```
---
### Condition only (while loop)


---
```go
i := 1
for i < 100 {
	fmt.Println(i)
	i = i * 2
}
```
```output
1
2
4
8
16
32
64
```
---
### Infinite Loop

---
```go
x := 0
for {
	x++
	fmt.Println(x)
	if x == 5 {
		break
	}
}
```
```output
1
2
3
4
5
```
---
### Looping on 
A for loop of a string converts UTF-8 to 32bit, so non standard characters are still treated as one

---
```go
m := "moon🌜y"
for i, v := range m {
	fmt.Printf("%v %8v %-6v\n", i, v, string(v))
}
```
```output
0      109 m     
1      111 o     
2      111 o     
3      110 n     
4   127772 🌜     
8      121 y
```
---
### Countinue to label

---
```go
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println("Outer")
	}
```
```output
0 104 h
1 101 e
2 108 l
0 97 a
1 112 p
2 112 p
3 108 l
```
---
### Goto

---
```go
func gotoer() {
	goto done
	for {
			fmt.Println("Skip the infinite loop")
	}
done:
	fmt.Println("Done")
}
```
---
---
```go
gotoer()
```
```output
Done
```
---
## Notes
- The value returned in a for-range loop is a copy, not a reference 

