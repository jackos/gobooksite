
---
```go
import (
	"io"
	"log"
	"os"
	"fmt"
)
```
---
---
```go
f, err := os.Open("/home/jacko/vimwiki/go/out.txt")
if err != nil {
	log.Fatal(err)
}
```
---
---
```go
// defer f.Close()
// ## defer keyword not working in gobook
```
---
---
```go
buffer := make([]byte, 1028)
```
---
---
```go
count := 0
```
---
---
```go
for {
	count, err = f.Read(buffer)
	// os.Stdout.Write(buffer[:count])
	// ## os.Stdout not working in gobook
	fmt.Println(string(buffer[:count]))
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
		break
	}
}
```
```output


```
---
