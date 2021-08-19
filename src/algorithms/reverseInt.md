
---
```go
package main
```
---
---
```go
import "fmt"
```
---
---
```go
func reverse(x int) int {
	var result int
	for x != 0 {
		// multiply by 10 for placeholder for next digit
		result *= 10
		// Gets the last digit
		xModulo := x % 10
		// Add the last digit
		result += xModulo
		if result > 2147483647 || result < -2147483648 {
			return 0
		}
		// strip the last digit, no remainder because int
		x /= 10
	}
	return result
}
```
---
---
```go
func main(){
	x := reverse(-4232)
	fmt.Println(x)
	fmt.Prin
}
```
---
