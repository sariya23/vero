package random

import (
	"fmt"
	mrand "math/rand/v2"
)

// Sample returns k randomly selected elements from collection.
//
// Elements are selected without replacement. If collection contains
// duplicate values, they may appear in the result.
//
// If k is greater than len(collection), all elements are returned in random order.
// If k equals len(collection), the original slice is returned as-is.
// Sample panics if k <= 0.
func Sample[T any](collection []T, k int) []T {
	if k <= 0 {
		panic(fmt.Sprintf("'k' argument must be greater than zero"))
	}

	if len(collection) == k {
		return collection
	}

	if k > len(collection) {
		k = len(collection)
	}

	res := append([]T{}, collection...)
	mrand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	return res[:k:k]
}
