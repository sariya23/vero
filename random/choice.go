package random

import "math/rand/v2"

// Choice returns a randomly selected element from the given slice.
// If the slice is empty, it returns the zero value of T.
func Choice[T any](collection []T) T {
	if len(collection) == 0 {
		return *new(T)
	}

	pickIndex := rand.IntN(len(collection))
	return collection[pickIndex]
}
