package random

import (
	"slices"
	"testing"
)

func TestSample(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		collection  []any
		k           int
		expectedLen int
		mustPanic   bool
	}{
		{
			name:        "k less then len(collection)",
			collection:  []any{"qwe", "asd", "zxc"},
			k:           1,
			expectedLen: 1,
		},
		{
			name:        "k greater then len(collection)",
			collection:  []any{"qwe", "asd", "zxc"},
			k:           10,
			expectedLen: 3,
		},
		{
			name:        "k is negative",
			collection:  []any{"qwe", "asd", "zxc"},
			k:           -10,
			expectedLen: 3,
			mustPanic:   true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer func() {
				if r := recover(); r != nil {
					if !tc.mustPanic {
						t.Error("unexpected panic")
					}
				}
			}()
			got := Sample(tc.collection, tc.k)

			if len(got) != tc.expectedLen {
				t.Errorf("expected len %d, got %d", tc.expectedLen, len(got))
			}

			for _, v := range got {
				if !slices.Contains(tc.collection, v) {
					t.Errorf("initial collection not contains elevent '%v'", v)
				}
			}
		})
	}
}
