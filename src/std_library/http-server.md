
---
```go
import "fmt"
```
---
---
```go
func logOutput(message string) {
	fmt.Println(message)
}
```
---
---
```go
type SimpleDataStore struct {
	userData map[string]string
}
```
---
---
```go
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}
jack
```
---
---
```go
SimpleDataStore:
func NewSimpleDataStore() SimpleDataStore {
return SimpleDataStore{
userData: map[string]string{
"1": "Fred",
"2": "Mary",
```
---
