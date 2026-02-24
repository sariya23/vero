package rules

import (
	"reflect"
	"testing"
)

func TestNewBoolRule(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name                string
		ruleName, ruleValue string
		expected            BoolRule
		expectedErr         error
	}{
		{
			name:      "success",
			ruleName:  string(Only),
			ruleValue: string(OnlyTrue),
			expected:  BoolRule{Name: Only, Value: OnlyTrue},
		},
		{
			name:        "empty rule name and rule value",
			expected:    BoolRule{},
			expectedErr: nil,
		},
		{
			name:        "unknown rule name",
			ruleName:    "ABOBA",
			ruleValue:   string(OnlyFalse),
			expected:    BoolRule{},
			expectedErr: ErrUnknowBoolRuleName,
		},
		{
			name:        "unknown rule value",
			ruleName:    string(Only),
			ruleValue:   "SHTA?",
			expected:    BoolRule{},
			expectedErr: ErrUnknowBoolRuleValue,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewBoolRule(tc.ruleName, tc.ruleValue)
			if tc.expectedErr != err {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}
			if tc.expected != got {
				t.Errorf("expected rule: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestNewBoolRules(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		rules    []BoolRule
		expected BoolRules
	}{
		{
			name:     "not empty rules",
			rules:    []BoolRule{{Name: Only, Value: OnlyTrue}},
			expected: BoolRules{OnlyTrue: &BoolRule{Name: Only, Value: OnlyTrue}},
		},
		{
			name:     "duplicate rule",
			rules:    []BoolRule{{Name: Only, Value: OnlyTrue}, {Name: Only, Value: OnlyTrue}},
			expected: BoolRules{OnlyTrue: &BoolRule{Name: Only, Value: OnlyTrue}},
		},
		{
			name:     "empty rules",
			rules:    nil,
			expected: BoolRules{},
		},
		{
			name:     "some random rule",
			rules:    []BoolRule{{Name: "ZXC", Value: "QWE"}},
			expected: BoolRules{},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := NewBoolRules(tc.rules)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("expected rule: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestGenerateBool_DefinedResult(t *testing.T) {
	cases := []struct {
		name     string
		rules    BoolRules
		expected bool
	}{
		{
			name: "only true",
			rules: BoolRules{
				OnlyTrue: &BoolRule{Name: Only, Value: OnlyTrue},
			},
			expected: true,
		},
		{
			name: "only false",
			rules: BoolRules{
				OnlyFalse: &BoolRule{Name: Only, Value: OnlyFalse},
			},
			expected: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := GenerateBool(tc.rules)
			if tc.expected != got {
				t.Errorf("expected result: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestGenerateBool_NotDefinedResult(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		rules BoolRules
	}{
		{
			name: "specified both rules",
			rules: BoolRules{
				OnlyTrue:  &BoolRule{Name: Only, Value: OnlyTrue},
				OnlyFalse: &BoolRule{Name: Only, Value: OnlyFalse},
			},
		},
		{
			name:  "no rules",
			rules: BoolRules{},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			GenerateBool(tc.rules)
		})
	}
}
