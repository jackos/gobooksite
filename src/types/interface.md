
---
```go
import "fmt"
```
---
## Terms
- Interface: list the methods that must be implemented by a concrete type
- Method set: the methods defined on an interface

Example of using a interface

---
```go
type GoodDog interface {
	roll()
	pat()
}

type BadDog interface {
	bark()
	howl()
	GoodDog
}

type Dog struct {
	color string
	size  int
}

type Wolf struct {
	Dog
}

func (Wolf) bark() {
	fmt.Println("Woof!")
}
func (Wolf) howl() {
	fmt.Println("Awwwwwoooooooooo!")
}

func (Dog) roll() {
	fmt.Println("He rolls over")
}

func (Dog) pat() {
	fmt.Println("He likes that!")
}

func main() {
	hendrix := Dog{"blonde", 2}
	wolfy := Wolf{Dog{"black", 10}}
	var hendrix_interface GoodDog = hendrix
	var wolfy_interface BadDog = wolfy

	hendrix_interface.pat()
	wolfy_interface.bark()
	wolfy_interface.roll()
}
```
---
Interfaces are implemented as two pointers, value and type. Having a type assigned makes it non-nil

---
```go
var s *string
```
---
---
```go
fmt.Println(s == nil)
```
---
---
```go
var i interface{}
```
---
---
```go
fmt.Println(i == nil)
```
---
---
```go
i = s
```
---
---
```go
fmt.Println(i == nil)
```
---
Empty interface, can represent any data type

---
```go
data := map[string]interface{}{}
```
---
