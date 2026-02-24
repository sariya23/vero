package random

import (
	"reflect"
	"testing"

	"github.com/sariya23/vero/internal/rules"
)

type Any struct {
}

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
	intPtr := 8
	cases := []struct {
		name        string
		structPtr   any
		expected    any
		expectedErr error
		mustPanic   bool
	}{
		{
			name:      "argument not struct",
			structPtr: &intPtr,
			mustPanic: true,
		},
		{
			name: "argument not pointer",
			structPtr: struct {
			}{},
			mustPanic: true,
		},
		{
			name:      "BOOL: no rules",
			structPtr: &TestBoolNoRules{},
			expected:  Any{},
		},
		{
			name:      "BOOL: only true",
			structPtr: &TestBoolOnlyTrue{},
			expected:  &TestBoolOnlyTrue{Field: true},
		},
		{
			name:      "BOOL: only false",
			structPtr: &TestBoolOnlyFalse{},
			expected:  &TestBoolOnlyFalse{Field: false},
		},
		{
			name:        "BOOL: unknown rule value",
			structPtr:   &TestBoolUnknownRuleValue{},
			expectedErr: rules.ErrUnknowBoolRuleValue,
			expected:    &TestBoolUnknownRuleValue{},
		},
		{
			name:        "BOOL: unknown rule name",
			structPtr:   &TestBoolUnknownRuleName{},
			expectedErr: rules.ErrUnknowBoolRuleName,
			expected:    &TestBoolUnknownRuleName{},
		},
		{
			name:      "BOOL: invalid tag",
			structPtr: &TestBoolInvalidTag{},
			mustPanic: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer func() {
				if r := recover(); r != nil {
					if !tc.mustPanic {
						t.Errorf("unexpected panic")
					}
				}
			}()
			err := Struct(tc.structPtr)
			if tc.expectedErr != err {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(tc.expected, Any{}) {
				if !reflect.DeepEqual(tc.expected, tc.structPtr) {
					t.Errorf("got %v, want %v", tc.structPtr, tc.expected)
				}
			}
		})
	}
}
