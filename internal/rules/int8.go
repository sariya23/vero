package rules

import (
	"errors"
	"math"
	"slices"
	"strconv"
)

var (
	// ErrUnknowInt8RuleName is returned when the rule name is not "min" or "max".
	ErrUnknowInt8RuleName = errors.New("unknown int8 rule name")
	// ErrInvalidInt8RuleValue is returned when the rule value is not a valid int8.
	ErrInvalidInt8RuleValue = errors.New("invalid int8 rule value")
	// ErrInvalidInt8Range is returned when Min > Max.
	ErrInvalidInt8Range = errors.New("invalid int8 rule range")
)

// Int8RuleValue is the string form of an int8 rule value (parsed to int8).
type Int8RuleValue string

// Int8RuleName is the name of an int8 rule: "min" or "max".
type Int8RuleName string

const (
	// MinRuleName lower bound for generated int8.
	MinRuleName Int8RuleName = "min"

	// MaxRuleName upper bound for generated int8.
	MaxRuleName Int8RuleName = "max"
)

// Int8Rule is a single int8 constraint: name (min/max) and numeric value.
type Int8Rule struct {
	Name  Int8RuleName
	Value int8
}

// Int8Rules holds parsed int8 rules used by GenerateInt8.
type Int8Rules struct {
	Min *Int8Rule
	Max *Int8Rule
}

// NewInt8Rules builds an Int8Rules from a slice of Int8Rule.
// Duplicate rules overwrite each other; unknown rule names are ignored.
func NewInt8Rules(rules []Int8Rule) Int8Rules {
	var res Int8Rules
	for _, rule := range rules {
		if rule.Name == MinRuleName {
			res.Min = &rule
		} else if rule.Name == MaxRuleName {
			res.Max = &rule
		}
	}
	return res
}

// NewInt8Rule creates an Int8Rule from string name and value.
// name must be "min" or "max"; value must be a number in int8 range.
// Empty name and value return a zero rule with no error.
// Returns ErrUnknowInt8RuleName or ErrInvalidInt8RuleValue for invalid values.
func NewInt8Rule(name, value string) (Int8Rule, error) {
	if name == "" || value == "" {
		return Int8Rule{}, nil
	}
	if !slices.Contains([]Int8RuleName{MinRuleName, MaxRuleName}, Int8RuleName(name)) {
		return Int8Rule{}, ErrUnknowInt8RuleName
	}
	v, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return Int8Rule{}, ErrInvalidInt8RuleValue
	}
	return Int8Rule{Name: Int8RuleName(name), Value: int8(v)}, nil
}

// GenerateInt8 returns a random int8 in the range defined by rules.
// If both Min and Max are set, min must be <= max or ErrInvalidInt8Range is returned.
// With no rules, calls GenerateInt8WithoutRules().
func GenerateInt8(rules Int8Rules) (int8, error) {
	if rules.Min == nil && rules.Max == nil {
		return GenerateInt8WithoutRules(), nil
	}

	if rules.Min != nil && rules.Max != nil {
		if rules.Min.Value > rules.Max.Value {
			return 0, ErrInvalidInt8Range
		}
		return rand.Int8Range(rules.Min.Value, rules.Max.Value), nil
	}
	if rules.Min != nil && rules.Max == nil {
		return rand.Int8Range(rules.Min.Value, math.MaxInt8), nil
	}
	if rules.Max != nil && rules.Min == nil {
		return rand.Int8Range(math.MinInt8, rules.Max.Value), nil
	}
	return 0, ErrInvalidInt8Range
}

// GenerateInt8WithoutRules returns a random int8 with no rule constraints.
func GenerateInt8WithoutRules() int8 {
	return rand.Int8()
}
