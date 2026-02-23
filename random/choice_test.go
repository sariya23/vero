package random

import (
	"slices"
	"testing"
)

const ANY_INT = -1

func TestChoice_EmptyCollection(t *testing.T) {
	t.Parallel()
	col := []int{}
	got := Choice(col)
	if got != *new(int) {
		t.Error("expected zero value for empty collection")
	}
}

func TestChoice(t *testing.T) {
	t.Parallel()
	col := []int{1, 2, 3, 4}
	got := Choice(col)

	if !slices.Contains(col, got) {
		t.Error("choice element must be in collection")
	}
}
