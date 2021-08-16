
---
```go
import (
	"time"
	"fmt"
)
```
---

Use multiply to convert units to a time.Duration

---
```go
d := 2 * time.Hour + 30 * time.Minute
fmt.Println(d)
```
```output
2h30m0s
```
---

Using Equal method corrects for timezones

---
```go
x := time.Now()
y := x
if(x.Equal(y)){
	fmt.Println("Equal!")
} else { 
	fmt.Println("Not equal!")
}
```
```output
Equal!
```
---

Get current time with timezone

---
```go
fmt.Println(time.Now())
```
```output
2021-07-09 20:50:46.11267622 +0800 AWST m=+69.426266556
```
---

Very cool way of specifying date format, 1 represents seconds, 6 represents year, and 7 is timezone. Everything else is intuitive

---
```go
tEarly := time.Parse("06-05-04 03:02:01", "20-01-01 01:01:01")
tLate := time.Parse("06-05-04 03:02:01", "20-01-01 12:01:01")
fmt.Println(tEarly)
```
```output
2020-01-01 01:01:01 +0000 UTC
```
---

Print a portion of time

---
```go
fmt.Println(tEarly.Year())
```
```output
2020
```
---

Compare two times (before, after, equal)

---
```go
fmt.Println(tEarly.Before(tLate))
```
```output
true
```
---

Return the time elapsed between two dates

---
```go
fmt.Println()
```
---
