
Keep your code in GOPATH which is typically ~/go/src/
Make a subdirectory for the extra package you want:

---
```markdown

├── go.mod
├── go.sum
├── main.go
└── mypackage
    └── mypackage.go
```
---
In mypackage (make sure to capitalize anything you want to export):
---
```go
package mypackage

func Double(a int) int {
	return a * 2
}
```
---
Now in main.go you can import the package:

---
```go
package main

import (
	"fmt"

	"github.com/[user]/[repo]/package_example/mypackage"
)

func main() {
	num := mypackage.Double(2)
	output := print.Format(num)

	fmt.Println(output)
}
```
---
run:
---
```fish
go mod init
go mod tidy
```
---
When you upload this to github, others will be able to import mypackage via the same import line