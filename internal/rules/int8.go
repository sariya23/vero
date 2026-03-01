package rules

import (
	"errors"
	"math"
	"slices"
	"strconv"
)

var (
	ErrUnknowInt8RuleName   = errors.New("unknown int8 rule name")
	ErrInvalidInt8RuleValue = errors.New("invalid int8 rule value")
	ErrInvalidInt8Range     = errors.New("invalid int8 rule range")
)

type Int8RuleValue string
type Int8RuleName string

const (
	MinRuleName Int8RuleName = "min"
	MaxRuleName Int8RuleName = "max"
)

type Int8Rule struct {
	Name  Int8RuleName
	Value int8
}

type Int8Rules struct {
	Min *Int8Rule
	Max *Int8Rule
}

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

func GenerateInt8(rules Int8Rules) (int8, error) {
	if rules.Min == nil && rules.Max == nil {
		return GenerateInt8WithoutRules(), nil
	}

	if rules.Min != nil && rules.Max != nil {
		if rules.Min.Value < rules.Max.Value {
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

func GenerateInt8WithoutRules() int8 {
	return rand.Int8()
}
