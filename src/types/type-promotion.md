
---
```go
import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Descriptions() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

// Anything from Employee is promoted so Manager has direct access to it
type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	result := []Employee{}
	result = append(result, Employee{"Billy", "14560456"})
	result = append(result, Employee{"Donkey", "4662456"})
	return result
}

func main() {
	bob := Employee{"Bob", "42248465"}
	fred := Employee{"Fred", "45224865"}

	bossMan := Manager{bob, []Employee{fred}}

	bossMan.Reports = append(bossMan.Reports, bossMan.FindNewEmployees()...)

	fmt.Println("Boss name: ", bossMan.Name)
	fmt.Println("Boss id: ", bossMan.ID)
	fmt.Println("Boss reports:\n")
	for _, v := range bossMan.Reports {
		fmt.Printf("Name: %v\nid: %v\n\n", v.Name, v.ID)
	}
}

```
```output
Boss name:  Bob
Boss id:  42248465
Boss reports:

Name: Fred
id: 45224865

Name: Billy
id: 14560456

Name: Donkey
id: 4662456

```
---
