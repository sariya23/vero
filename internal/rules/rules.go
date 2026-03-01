package rules

import (
	rand2 "github.com/sariya23/vero/internal/rand"
)

var rand *rand2.Source

func init() {
	rand = rand2.NewRandSource()
}
