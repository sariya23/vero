package random

import (
	"reflect"
	"strings"

	"github.com/sariya23/vero/internal/rules"
)

const TagName = "rules"

func Struct(s any) any {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		panic("not struct")
	}
	structInstance := reflect.New(t).Elem()

	recursiveGenerateFillValue(structInstance)
	return structInstance.Interface()
}

func recursiveGenerateFillValue(structValue reflect.Value) {
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		if !field.CanSet() {
			continue
		}
		ruleTags := structValue.Type().Field(i).Tag.Get(TagName)
		switch field.Type().Kind() {
		case reflect.Bool:
			boolRules := buildBoolRules(ruleTags)
			generated := rules.GenerateBool(boolRules)
			field.Set(reflect.ValueOf(generated))
		case reflect.Int:
			field.Set(reflect.ValueOf(8))
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint16:
		case reflect.Uint32:
		case reflect.Uint64:
		case reflect.Uintptr:
		case reflect.Float32:
		case reflect.Float64:
		case reflect.Complex64:
		case reflect.Complex128:
		case reflect.Array:
		case reflect.Chan:
		case reflect.Func:
		case reflect.Interface:
		case reflect.Map:
		case reflect.Ptr:
		case reflect.Slice:
		case reflect.String:
		case reflect.Struct:
			recursiveGenerateFillValue(field)
		case reflect.UnsafePointer:
		}
	}
}

func buildBoolRules(tagString string) rules.BoolRules {
	if tagString == "" {
		return rules.BoolRules{}
	}
	allRules := strings.Split(tagString, ",")
	boolRules := make([]rules.BoolRule, 0, len(allRules))
	for _, r := range allRules {
		pair := strings.Split(r, "=")
		if !validRule(pair) {
			panic("invalid tag value")
		}
		boolRule, err := rules.NewBoolRule(pair[0], pair[1])
		if err != nil {
			panic(err)
		}
		boolRules = append(boolRules, boolRule)
	}

	return rules.NewBoolRules(boolRules)
}

func validRule(pair []string) bool {
	return len(pair) != 2 || pair[0] == "" || pair[1] == ""
}
