
---
```go
import (
	"errors"
	"fmt"
	"log"
	"os"
)
```
---

---
```go
type Status int
```
---

---
```go
type CustomError struct {
	description string
	category    string
	status      Status
	childError  error
	codes       []int
}
```
---

---
```go
const (
	TypeErr Status = iota + 1
	StatusErr
	UnlawfulErr
)
```
---

---
```go
func (ce CustomError) Error() string {
	return fmt.Sprintf("|Error: %v| |Status: %v| |Description: %v|", ce.category, ce.status, ce.description)
}
```
---

---
```go
func (ce CustomError) Unwrap() error {
	return ce.childError
}
```
---

---
```go
func readFileWrapError(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}
```
---

---
```go
func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	fmt.Println()
	if err != nil {
		return nil, CustomError{
			"Could not read file correctly",
			"FileReading",
			3,
			err,
			[]int{1, 2, 4, 6},
		}
	} else {
		return data, nil
	}
}
```
---

---
```go
func main() {
	err := readFileWrapError("nofile.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Custom Error:\n%v\n", err)
		} else {
			fmt.Println(err)
		}
	}
	var ce CustomError
	data, err := readFile("nofile.txt")
	if err != nil {
		if errors.As(err, &ce) {
			fmt.Printf("Error type: %T\n", &ce)
			fmt.Println("Error codes:", ce.codes)
			// Handle business logic here
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(data)
}
```
```output
Custom Error:
in fileChecker: open nofile.txt: no such file or directory

```
---
