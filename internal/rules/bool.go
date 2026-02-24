package rules

import (
	"errors"
	"slices"
)

var (
	ErrUnknowBoolRuleName  = errors.New("unknown rule name")
	ErrUnknowBoolRuleValue = errors.New("unknown rule value")
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

type BoolRuleStructTag struct {
	Name  BoolRuleName
	Value BoolRuleValue
}

func NewBoolRuleStructTag(name, value string) (BoolRuleStructTag, error) {
	if !slices.Contains([]BoolRuleName{Only}, BoolRuleName(name)) {
		return BoolRuleStructTag{}, ErrUnknowBoolRuleName
	}
	if !slices.Contains([]BoolRuleValue{OnlyTrue, OnlyFalse}, BoolRuleValue(value)) {
		return BoolRuleStructTag{}, ErrUnknowBoolRuleValue
	}
	return BoolRuleStructTag{Name: BoolRuleName(name), Value: BoolRuleValue(value)}, nil
}

func GenerateBool(rule BoolRuleValue) bool {
	switch rule {
	case OnlyTrue:
		return true
	case OnlyFalse:
		return false
	default:
		return rand.Bool()
	}
}
