
---
```go
import (
	"errors"
	"fmt"
	"log"
)
```
---

Example of using an interface and type assertion to run a different method depending on the type

---
```go
type treeNode struct {
	val    Val
	lchild *treeNode
	rchild *treeNode
}
type Val interface {
	process(int, int) int
}

type number int
type operator byte

func (n number) process(int, int) int {
	fmt.Println("Hey! Can't process this!")
	return int(n)
}

func (o operator) process(left int, right int) int {
	switch o {
	case '+':
		fmt.Printf("%v%v%v = %v\n", left, string(o), right, left+right)
		return left + right
	case '-':
		fmt.Printf("%v%v%v = %v\n", left, string(o), right, left-right)
		return left - right
	case '*':
		fmt.Printf("%v%v%v = %v\n", left, o, right, left*right)
		return left * right
	case '/':
		fmt.Printf("%v%v%v = %v\n", left, o, right, left/right)
		return left / right
	default:
		return -1
	}
}

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("invalid expression")
	case number:
		// we know that t.val is of type number, so return the
		// int value
		return int(val), nil
	case operator:
		// we know that t.val is of type operator, so
		// find the values of the left and right children, then
		// call the process() method on operator to return the
		// result of processing their values.
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		// if a new treeVal type is defined, but walkTree wasn't updated
		// to process it, this detects it
		return 0, errors.New("unknown node type")
	}
}

func main() {
	t := treeNode{
		val: number(10),
	}
	t2 := treeNode{
		val: number(9),
	}
	t3 := treeNode{
		val:    operator('-'),
		lchild: &t,
		rchild: &t2,
	}

	x, err := walkTree(&t3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
}
```
---
