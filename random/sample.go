package random

import (
	"fmt"
	"math/rand/v2"
)

// Sample returns elements...
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
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	return res[:k:k]
}
