package rules

import (
	"math"
	"reflect"
	"testing"
)

func TestNewInt8Rule(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name                string
		ruleName, ruleValue string
		expected            Int8Rule
		expectedErr         error
	}{
		{
			name:      "success min",
			ruleName:  string(MinRuleName),
			ruleValue: "0",
			expected:  Int8Rule{Name: MinRuleName, Value: 0},
		},
		{
			name:      "success max",
			ruleName:  string(MaxRuleName),
			ruleValue: "100",
			expected:  Int8Rule{Name: MaxRuleName, Value: 100},
		},
		{
			name:     "empty rule name and rule value",
			expected: Int8Rule{},
		},
		{
			name:        "unknown rule name",
			ruleName:    "ABOBA",
			ruleValue:   "42",
			expected:    Int8Rule{},
			expectedErr: ErrUnknowInt8RuleName,
		},
		{
			name:        "invalid rule value not a number",
			ruleName:    string(MinRuleName),
			ruleValue:   "SHTA?",
			expected:    Int8Rule{},
			expectedErr: ErrInvalidInt8RuleValue,
		},
		{
			name:        "invalid rule value overflow",
			ruleName:    string(MaxRuleName),
			ruleValue:   "99999",
			expected:    Int8Rule{},
			expectedErr: ErrInvalidInt8RuleValue,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewInt8Rule(tc.ruleName, tc.ruleValue)
			if tc.expectedErr != err {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}
			if tc.expected != got {
				t.Errorf("expected rule: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestNewInt8Rules(t *testing.T) {
	minRule := Int8Rule{Name: MinRuleName, Value: 10}
	maxRule := Int8Rule{Name: MaxRuleName, Value: 100}
	t.Parallel()
	cases := []struct {
		name     string
		rules    []Int8Rule
		expected Int8Rules
	}{
		{
			name:     "only min rule",
			rules:    []Int8Rule{minRule},
			expected: Int8Rules{Min: &minRule},
		},
		{
			name:     "only max rule",
			rules:    []Int8Rule{maxRule},
			expected: Int8Rules{Max: &maxRule},
		},
		{
			name:     "min and max rules",
			rules:    []Int8Rule{minRule, maxRule},
			expected: Int8Rules{Min: &minRule, Max: &maxRule},
		},
		{
			name:     "duplicate rule",
			rules:    []Int8Rule{minRule, minRule},
			expected: Int8Rules{Min: &minRule},
		},
		{
			name:     "empty rules",
			rules:    nil,
			expected: Int8Rules{},
		},
		{
			name:     "some random rule",
			rules:    []Int8Rule{{Name: "ZXC", Value: 42}},
			expected: Int8Rules{},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := NewInt8Rules(tc.rules)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("expected rules: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestGenerateInt8_InvalidRange(t *testing.T) {
	t.Parallel()
	rules := Int8Rules{
		Min: &Int8Rule{Name: MinRuleName, Value: 50},
		Max: &Int8Rule{Name: MaxRuleName, Value: 10},
	}
	_, err := GenerateInt8(rules)
	if err != ErrInvalidInt8Range {
		t.Errorf("expected error: %v, got: %v", ErrInvalidInt8Range, err)
	}
}

func TestGenerateInt8_NoRules(t *testing.T) {
	t.Parallel()
	got, err := GenerateInt8(Int8Rules{})
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	_ = got
}

func TestGenerateInt8_WithRules_InRange(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		rules      Int8Rules
		minAllowed int8
		maxAllowed int8
	}{
		{
			name: "only min",
			rules: Int8Rules{
				Min: &Int8Rule{Name: MinRuleName, Value: 10},
			},
			minAllowed: 10,
			maxAllowed: math.MaxInt8,
		},
		{
			name: "only max",
			rules: Int8Rules{
				Max: &Int8Rule{Name: MaxRuleName, Value: 10},
			},
			minAllowed: math.MinInt8,
			maxAllowed: 10,
		},
		{
			name: "max and min",
			rules: Int8Rules{
				Min: &Int8Rule{Name: MinRuleName, Value: 10},
				Max: &Int8Rule{Name: MaxRuleName, Value: 100},
			},
			minAllowed: 10,
			maxAllowed: 100,
		},
		{
			name: "min and max equal",
			rules: Int8Rules{
				Min: &Int8Rule{Name: MinRuleName, Value: 42},
				Max: &Int8Rule{Name: MaxRuleName, Value: 42},
			},
			minAllowed: 42,
			maxAllowed: 42,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := GenerateInt8(tc.rules)
			if err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if got < tc.minAllowed || got > tc.maxAllowed {
				t.Errorf("expected value in [%d, %d], got: %d", tc.minAllowed, tc.maxAllowed, got)
			}
		})
	}
}
