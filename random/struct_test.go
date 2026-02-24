package random

import (
	"reflect"
	"testing"
)

type TestBoolNoRules struct {
	Field bool
}
type TestBoolOnlyTrue struct {
	Field bool `rules:"only=true"`
}

type TestBoolOnlyFalse struct {
	Field bool `rules:"only=false"`
}

type TestBoolUnknownRuleValue struct {
	Field bool `rules:"only=aboba"`
}

type TestBoolUnknownRuleName struct {
	Field bool `rules:"aboba=true"`
}

type TestBoolInvalidTag struct {
	Field bool `rules:"aboba="`
}

func TestStructBool(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name      string
		model     any
		expected  any
		mustPanic bool
	}{
		{
			name:      "argument not struct",
			model:     8,
			mustPanic: true,
		},
		{
			name:     "BOOL: no rules",
			model:    TestBoolNoRules{},
			expected: nil,
		},
		{
			name:     "BOOL: only true",
			model:    TestBoolOnlyTrue{},
			expected: TestBoolOnlyTrue{Field: true},
		},
		{
			name:     "BOOL: only false",
			model:    TestBoolOnlyFalse{},
			expected: TestBoolOnlyFalse{Field: false},
		},
		{
			name:      "BOOL: unknown rule value",
			model:     TestBoolUnknownRuleValue{},
			mustPanic: true,
		},
		{
			name:      "BOOL: unknown rule name",
			model:     TestBoolUnknownRuleName{},
			mustPanic: true,
		},
		{
			name:      "BOOL: invalid tag",
			model:     TestBoolInvalidTag{},
			mustPanic: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer func() {
				didPanic := recover() != nil
				if didPanic != tc.mustPanic {
					t.Errorf("expected that panic `%v`, got `%v`", tc.mustPanic, didPanic)
				}
			}()
			got := Struct(tc.model)

			if tc.expected != nil {
				if !reflect.DeepEqual(got, tc.expected) {
					t.Errorf("got %v, want %v", tc.model, tc.expected)
				}
			}
		})
	}
}
