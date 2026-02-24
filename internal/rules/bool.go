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

func NewBoolRule(name, value string) (BoolRule, error) {
	if !slices.Contains([]BoolRuleName{Only}, BoolRuleName(name)) {
		return BoolRule{}, ErrUnknowBoolRuleName
	}
	if !slices.Contains([]BoolRuleValue{OnlyTrue, OnlyFalse}, BoolRuleValue(value)) {
		return BoolRule{}, ErrUnknowBoolRuleValue
	}
	return BoolRule{Name: BoolRuleName(name), Value: BoolRuleValue(value)}, nil
}

func GenerateBool(rules []BoolRule) (bool, error) {
	return false, nil
}

func GenerateBoolWithoutRules() bool {
	return rand.Bool()
}
