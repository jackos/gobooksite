
# Closures

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
- initialize a slice of 3 people
- Sort by age, then last name
- People is captured by the anonymous function closure, 

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
fmt.Println(people)
```
```output
[{Bob Bobbart 18} {Amy Bobbart 19} {Pat Patterson 16}]
```
---

## Defining a closure type

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
```output
0 0 0
2 3 4
4 6 8
6 9 12
8 12 16
10 15 20
12 18 24
14 21 28
16 24 32
18 27 36
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
_, closer, err := getFile("./out.txt")
if err != nil {
	log.Fatal(err)
}
defer closer()
```
```output
2 3 4
4 6 8
6 9 12
8 12 16
10 15 20
12 18 24
14 21 28
16 24 32
18 27 36
gobook-output-start
2021/08/18 10:27:47 open ./out.txt: no such file or directory
exit status 1
```
---
