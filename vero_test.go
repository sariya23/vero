package main

import "testing"

func TestValidPair(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		pair     []string
		expected bool
	}{
		{
			name:     "valid",
			pair:     []string{"a", "b"},
			expected: true,
		},
		{
			name:     "no left part",
			pair:     []string{"", "b"},
			expected: false,
		},
		{
			name:     "no right part",
			pair:     []string{"a", ""},
			expected: false,
		},
		{
			name:     "no pair",
			pair:     []string{"", ""},
			expected: false,
		},
		{
			name:     "nil",
			pair:     nil,
			expected: false,
		},
		{
			name:     "to many value",
			pair:     []string{"a", "b", "c"},
			expected: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := validRule(tc.pair)
			if got != tc.expected {
				t.Errorf("got %v; expected %v", got, tc.expected)
			}
		})
	}
}
