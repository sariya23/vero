package rules

import "testing"

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
