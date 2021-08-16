
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
