package rules

import (
	"errors"
	"slices"
)

var (
	// ErrUnknowBoolRuleName is returned when the rule name is not recognized.
	ErrUnknowBoolRuleName = errors.New("unknown bool rule name")
	// ErrUnknowBoolRuleValue is returned when the rule value is not recognized.
	ErrUnknowBoolRuleValue = errors.New("unknown bool rule value")
)

// BoolRuleValue is the value of a bool rule ("true" or "false").
type BoolRuleValue string

// BoolRuleName is the name of a bool rule (e.g. Only).
type BoolRuleName string

const (
	// OnlyTrue restrict generation to true only.
	OnlyTrue BoolRuleValue = "true"

	// OnlyFalse restrict generation to false only.
	OnlyFalse BoolRuleValue = "false"
)

const (
	// Only rule name for bool: fix value to true or false.
	Only BoolRuleName = "only"
)

// BoolRule is a single bool constraint: name (e.g. Only) and value (true/false).
type BoolRule struct {
	Name  BoolRuleName
	Value BoolRuleValue
}

// BoolRules holds parsed bool rules used by GenerateBool.
type BoolRules struct {
	OnlyTrue  *BoolRule
	OnlyFalse *BoolRule
}

// NewBoolRules builds a BoolRules from a slice of BoolRule.
// Duplicate rules overwrite each other; unknown rule names are ignored.
func NewBoolRules(rules []BoolRule) BoolRules {
	var res BoolRules
	for _, rule := range rules {
		if rule.Name == Only {
			if rule.Value == OnlyFalse {
				res.OnlyFalse = &rule
			} else if rule.Value == OnlyTrue {
				res.OnlyTrue = &rule
			}
		}
	}
	return res
}

// NewBoolRule creates a BoolRule from string name and value.
// Empty name and value return a zero rule with no error.
// Returns ErrUnknowBoolRuleName or ErrUnknowBoolRuleValue for invalid values.
func NewBoolRule(name, value string) (BoolRule, error) {
	if name == "" || value == "" {
		return BoolRule{}, nil
	}
	if !slices.Contains([]BoolRuleName{Only}, BoolRuleName(name)) {
		return BoolRule{}, ErrUnknowBoolRuleName
	}
	if !slices.Contains([]BoolRuleValue{OnlyTrue, OnlyFalse}, BoolRuleValue(value)) {
		return BoolRule{}, ErrUnknowBoolRuleValue
	}
	return BoolRule{Name: BoolRuleName(name), Value: BoolRuleValue(value)}, nil
}

// GenerateBool returns a random bool according to rules.
// If only OnlyTrue or only OnlyFalse is set, returns that value.
// If both or neither are set, returns the result of GenerateBoolWithoutRules().
func GenerateBool(rules BoolRules) bool {
	if rules.OnlyTrue != nil && rules.OnlyFalse != nil {
		return GenerateBoolWithoutRules()
	}
	if rules.OnlyTrue != nil {
		return true
	}
	if rules.OnlyFalse != nil {
		return false
	}
	return GenerateBoolWithoutRules()
}

// GenerateBoolWithoutRules returns a random bool with no rule constraints.
func GenerateBoolWithoutRules() bool {
	return rand.Bool()
}
