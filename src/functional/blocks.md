
## Shadowing

### Shadow variable x

---
```go
func block() {
x := 10
if x > 5 {
	fmt.Println("Inner block", x)
	x:= 5 
	fmt.Println("Reinit inner block:", x)
	}
	fmt.Println("Outer block", x)
}
```
---
---
```go
block()
```
```output
Inner block 10
Reinit inner block: 5
Outer block 10
```
---
## Universe block


Redeclaring a variable in the universe block isn't detected by any linters or the shadow tool, be careful not to shadow anything in the universe block.

---
```go
fmt.Println(true)
true := 10
fmt.Println(true)
```
```output
true
10
```
---
## Initializer in if/else statement scope

---
```go
if n := rand.Intn(10); n == 0 {
	fmt.Println("That's too low")
} else if n > 5 {
	fmt.Println("That's too big:", n)
} else {
	fmt.Println("That's a good number:", n)
}
```
```output
That's a good number: 1
```
---
## Notes
### Scope
- Outside any function = package block
- Imports = file block
- Inside function / parameters = block
- Anything inside {} is a another block

### Access
- Can access any outer scope identifiers
