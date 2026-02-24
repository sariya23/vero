package main

import (
	"fmt"

	"github.com/sariya23/vero/random"
)

type S struct {
	B bool `rules:"only=true"`
	C bool
}

func main() {
	var s S

	q := random.Struct(s).(S)

	fmt.Println(q)

}
