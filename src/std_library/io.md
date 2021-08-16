
---
```go
import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)
```
---

Standard read file through buffer

---
```go
file, err := os.Open("./out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 100)
	for {
		count, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		if count == 0 {
			return
		}
		fmt.Print(string(data[:count]))
	}
```
---

Read through zip and use buffer for more operations

---
```go
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
```
---

---
```go
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}
```
---

---
```go
	r, closer, err := buildGZipReader("./test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counts
```
---
