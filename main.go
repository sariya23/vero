package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

type S struct {
	Name string
}

func main() {
	var s S
	gofakeit.Struct(&s)
	fmt.Println(s)
	gofakeit.Struct(&s)
	fmt.Println(s)
}
