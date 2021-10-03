# JSON
---
```go
import (
	"encoding/json"
	"fmt"
)
```
---

---
```go
f := struct {
Name string `json:"name"`
Age int `json:"age"`
}
err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
```
---

## json pretty print
```go
fmt.Println(json.MarshalIndent(data, "", "    "))
```
## Example
```go
import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
]`

	var unmarshalled []Person

	//
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Println(unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Billy"
	m2.LastName = "Bob"
	m2.HairColor = "green"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.Marshal(mySlice)
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
```