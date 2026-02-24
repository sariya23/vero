package main

import (
	"fmt"

	"github.com/sariya23/vero/random"
)

type S struct {
	B bool `rules:"aboba="`
}

func main() {
	var s S

	q := random.Struct(s).(S)

	fmt.Println(q)

}
