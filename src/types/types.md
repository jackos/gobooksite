
---
```go
import "fmt"
```
---
# Types

## Terms
- Abstract Type - What a type should do, but not how it's done
- Concrete Type - Specifies what and how i.e how to store data and implement method
- Reciever - The type the method is being attached to, conventionally shorthand abbreviation of the types name

## Declarations

### Primitive

---
```go
type Score int
```
---
### Function

---
```go
type Converter func(string)Score
```
---
### Compound

---
```go
type TeamScores map[string]Score
```
---
### Struct

## Methods

---
```go
type Person struct {
	FirstName string
	LastName string
	Age int
}
```
---
---
```go
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```
---
---
```go
p := Person{
	"Jacky",
	"C",
	33,
}
p.String()
```
---
### Assigning a method value and calling it 

---
```go
type Adder struct {
	start int
}
```
---
---
```go
func (a *Adder) AddTo(val int) int {
	return a.start + val
}
```
---
---
```go
myAdder := Adder{40}
f := myAdder.AddTo
myAdder.start = f(20)
fmt.Println(f(20))
```
---
Method Expression

---
```go
f2 := Adder.AddTo
fmt.Println(f2(myAdder, 15))
```
---

---
```go
// assigning untyped constants is valid
var i int = 300
var s Score = 100
var hs HighScore = 200 
hs = s // compilation error!
s = i // compilation error!
s = Score(i) // ok
hs = HighScore(s) // ok
```
---

iota type (enumerations)

---
```go
type MailCategory int
```
---

---
```go
const (
			Uncategorized MailCategory =  iota
			Personal
			Spam
			Social
			Advertisements
)
```
---

---
```go
fmt.Println(Spam)
```
```output
2
```
---

---
```go
x := Personal
if x == Advertisements{
	fmt.Println("wow")
}
```
```output
wow
```
---
