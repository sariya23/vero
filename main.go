package main

import (
	"fmt"

	"github.com/sariya23/vero/random"
)

type S struct {
	B bool `rules:"asd=qwe"`
}

func main() {
	var s S

	err := random.Struct(&s)

	fmt.Println(err, s)

}
