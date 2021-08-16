
---
```go
import (
	"fmt"
	"errors"
	"os"
	"strconv"
)
```
---
## Simulate optional parameters
---
```go
type MyFuncOpts struct {
	FirstName string
	LastName string
	Age int
}
```
---
---
```go
func printNameAge(opts MyFuncOpts) error {
	if opts.LastName != "" {
		fmt.Println(opts.LastName)
	}
	if opts.Age != 0 {
		fmt.Println("Age =", opts.Age)
		} else { 
			fmt.Println("No age provided")
	}
}
```
---
---
```go
func test() {
	printNameAge(MyFuncOpts {
		LastName: "Clayton",
	})
}
```
---
---
```go
test()
```
```output
Clayton
```
---
## Variadic input parameters and slices

---
```go
func addTo(base int, values ...int) []int {
	out := make([]int, 0, len(values))
	for _, v := range values {
		out = append(out, base+v)
	}
	return out
}
```
```output
No age provided
```
---
### Standard call

---
```go
sum := addTo(29,123,14,13)
fmt.Println(sum)
```
```output
[152 43 42]
```
---
### Spread an array into the parameters

---
```go
x := []int{20,30,40}
sum := addTo(10, x...)
fmt.Println(sum)
```
```output
[30 40 50]
```
---
## Named return values

---
```go
func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}
```
---
---
```go
x, y, err := divAndRemainder(5, 0)
if err != nil {
	fmt.Println("Oh no an error:", err)
	os.Exit(1)
}
fmt.Println(x, y, err)
```
```output
Oh no an error: Cannot divide by zero
```
---
Blank return (Go community says this is bad idea as it's unclear!)

---
```go
func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}
```
---
---
```go
x, y, z := divAndRemainder(10, 2)
fmt.Println(x, y, z)
```
```output
5 0 <nil>
```
---
## Anonymous function
---
```go
	for i := 0; i < 3; i++ {
		x := func(j int) int {
			fmt.Println("Printing", j, "from inside of an anonymous function")
			return i
		}(i)
		fmt.Println("Printing", x, "from the value returned")
	}
```
```output
Printing 0 from inside of an anonymous function
Printing 0 from the value returned
Printing 1 from inside of an anonymous function
Printing 1 from the value returned
Printing 2 from inside of an anonymous function
```
---
## First Class Functions

---
```go
func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }
```
---
---
```go
// Documentation for type
type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
```
---
---
```go
func run() {
	expressions := [][]string{
		{"2", "-", "3"},
		{"2", "+", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "plus", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression", expression, "(needs three parameters)")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
```
---
---
```go
run()
```
```output
-1
5
6
0
unsupported operator: %
strconv.Atoi: parsing "two": invalid syntax
invalid expression [5] (needs three parameters)
```
---
