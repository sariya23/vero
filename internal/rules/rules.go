package rules

import "github.com/sariya23/vero/internal"

var rand *internal.RandSource

func init() {
	rand = internal.NewRandSource()
}
