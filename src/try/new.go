package newCh

import (
	"fmt"
)

func newCh() {
	p := new(int)
	q := new(int)
	fmt.Println(*p)
	fmt.Println(p == q)
}
