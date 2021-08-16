
---
```go
import (
	"log"
	"fmt"
)
```
---
type conversion: change
type assertion: reveal

Conversion example

---
```go
type MyInt int
```
---
---
```go
func main() {
	var i MyInt = 20
	i2 := int(i)
	fmt.Println(i2)
}
```
```output
20
```
---
Assertion example (happens at runtime), always check type assertions!

---
```go
func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2, ok := i.(MyInt)
	if !ok {
		log.Fatalf("unexpected type for %v", i)
	}
	fmt.Println(i2)
}
```
```output
20
```
---
Type switch to check type and do appropriate action
i is shadowing here, one of few examples where shadowing is idiomatic

---
```go
func doThings(i interface{}) {
	switch i := i.(type) {
	case nil:
		fmt.Println(i, "is of type interface{}")
	case int:
		fmt.Println(i, "is of type int")
	case MyInt:
		fmt.Println(i, "is of type MyInt")
	case string:
		fmt.Println(i, "is of type string")

	default:
		fmt.Println("I don't know what type this value is:", j)
	}
}

func main() {
	var j MyInt = 10
	doThings(j)
}
```
---

Example in the standard library using type assertion from the standard librar

---
```go
// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	// function continues...
}
```
---

Example of using fallback code when implementing a newer version of an API
This comes from database/sql/driver in the standard library

---
```go
func ctxDriverStmtExec(ctx context.Context, si driver.Stmt,
	nvdargs []driver.NamedValue) (driver.Result, error) {
if siCtx, is := si.(driver.StmtExecContext); is {
	return siCtx.ExecContext(ctx, nvdargs)
}
// fallback code is here
}
```
---

---
```go

```
---
