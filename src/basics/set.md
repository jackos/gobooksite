
# Sets

## Using a map as a set

### Map keys don't duplicate

---
```go
intSet := map[int]bool{}
vals := []int{1, 1, 2, 2, 3}
for _, v := range vals { 
	intSet[v] = true
}
fmt.Println(intSet)
```
```output
map[1:true 2:true 3:true]
```
---
### Trying to retrieve a key that doesn't exist gives false

---
```go
fmt.Println(intSet[5])
```
```output
false
```
---
## Use a struct as a set

---
```go
structSet := map[int]struct{}{}
var exists = struct{}{}
ints := []int{1, 1, 2, 2, 3}
for _, v := range ints {
	structSet[v] = exists
}

for i, v := range structSet {
	fmt.Println(i, v)
}

if _, ok := structSet[4]; ok {
	fmt.Println("number is in the set")
} else {
	fmt.Println("not in the set")
}
```
```output
1 {}
2 {}
3 {}
not in the set
```
---
## Make a set type

---
```go
var exists = struct{}{}
type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}
```
---

---
```go
s := NewSet()
s.Add("Peter")
s.Add("David")
fmt.Println(s.Contains("Peter"))  // True
fmt.Println(s.Contains("George")) // False
s.Remove("David")
fmt.Println(s.Contains("David")) // False
```
```output
true
false
false
```
---
