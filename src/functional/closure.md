
## Define the Person struct

---
```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}
```
---
## Initialize a slice of 3 people

---
```go
people := []Person{
	{"Pat", "Patterson", 16},
	{"Amy", "Bobbart", 19},
	{"Bob", "Bobbart", 18},
}

sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
})

sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
})
```
```output
/tmp/main.go:4:1: syntax error: non-declaration statement outside function body
/tmp/main.go:9:24: missing function body
/tmp/main.go:9:39: method has multiple receivers
/tmp/main.go:9:44: syntax error: unexpected {, expecting (
/tmp/main.go:10:3: syntax error: unexpected return, expecting )
/tmp/main.go:13:1: syntax error: non-declaration statement outside function body
/tmp/main.go:13:24: missing function body
/tmp/main.go:13:39: method has multiple receivers
/tmp/main.go:14:3: syntax error: unexpected return, expecting )
```
---
Sort by age, then last name
People is captured by the anonymous function closure, 

---
```go
fmt.Println(people)
```
```output
/tmp/main.go:18:13: undefined: people
```
---
---
```go
// Returns a function that multiplies by the given factor
type twoFactor func(factor int) int
```
---
---
```go
func makeMult(base int) twoFactor {
	return func(factor int) int {
		return base *factor
	}
}
```
---
---
```go
twoBase := makeMult(2)
threeBase := makeMult(3)
fourBase := makeMult(4)
for i := 0; i < 10; i++ {
	fmt.Println(twoBase(i), threeBase(i), fourBase(i))
}
```
---
Remind caller to do something with a closure e.g. close a file, as it's returned to caller and Go doesn't allow unused values, they know that they should do something with it

---
```go
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, err
}
```
---
---
```go
f, closer, err := getFile("./out.txt")
if err != nil {
	log.Fatal(err)
}
defer closer()
```
---
