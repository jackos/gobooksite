
## Shadowing

---
```go
import "fmt"
```
```output
Imports are done automatically. Please remove import statement
```
---
Shadow variable (x) in inner block, dissallowing access to outer block

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
```output
/tmp/main.go:28:14: undefined: c

```
---
---
```go
block()
```
---
## Universe block


Redeclaring a variable in the universe block isn't detected by any linters or the shadow tool! Be careful not to shadow anything in the universe block!

---
```go
fmt.Println(true)
true := 10
fmt.Println(true)
```
---
If initializer used in the scope of any if/else statements

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
---
## Notes
### Scope
- Outside any function = package block
- Imports = file block
- Inside function / parameters = block
- Anything inside {} is a another block

### Access
- Can access any outer scope identifiers
