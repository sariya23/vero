package main

import (
	"fmt"

	"github.com/sariya23/vero/random"
)

type S struct {
	A int
	b bool
	T struct{ C int }
}

func main() {
	var s S

	random.Struct(&s)
	fmt.Println(s)

}
