package rules

import (
	"errors"
	"slices"
)

var (
	ErrUnknowBoolRuleName  = errors.New("unknown bool rule name")
	ErrUnknowBoolRuleValue = errors.New("unknown bool rule value")
)

type BoolRuleValue string
type BoolRuleName string

const (
	OnlyTrue  BoolRuleValue = "true"
	OnlyFalse BoolRuleValue = "false"
)

const (
	Only BoolRuleName = "only"
)

type BoolRule struct {
	Name  BoolRuleName
	Value BoolRuleValue
}

type BoolRules struct {
	OnlyTrue  *BoolRule
	OnlyFalse *BoolRule
}

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

func GenerateBool(rules BoolRules) (bool, error) {
	if rules.OnlyTrue != nil && rules.OnlyFalse != nil {
		return GenerateBoolWithoutRules(), nil
	}
	if rules.OnlyTrue != nil {
		return true, nil
	}
	if rules.OnlyFalse != nil {
		return false, nil
	}
	return GenerateBoolWithoutRules(), nil
}

func GenerateBoolWithoutRules() bool {
	return rand.Bool()
}
