package random

import (
	"reflect"
	"strings"

	"github.com/sariya23/vero/internal/rules"
)

const TagName = "rules"

// Struct...
// TODO: добавить генерацию всех полей
// TODO: вынести генерацию в internal
func Struct(objPtr any) error {
	value := reflect.ValueOf(objPtr)
	if value.Kind() != reflect.Ptr {
		panic("objPtr must be a pointer to struct")
	}

	if value.Elem().Kind() != reflect.Struct {
		panic("objPtr must be a pointer to struct")
	}

	return recursiveGenerateFillValue(value.Elem())
}

func recursiveGenerateFillValue(structValue reflect.Value) error {
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		if !field.CanSet() {
			continue
		}
		switch field.Type().Kind() {
		case reflect.Bool:
			tags := strings.Split(structValue.Type().Field(i).Tag.Get(TagName), ",")
			boolRules := make([]rules.BoolRule, 0, len(tags))
			for _, structTag := range tags {
				pair := strings.Split(structTag, "=")
				boolRule, err := rules.NewBoolRule(pair[0], pair[1])
				if err != nil {
					return err
				}
				boolRules = append(boolRules, boolRule)
			}
			generated, err := rules.GenerateBool(boolRules)
			if err != nil {
				return err
			}
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
			err := recursiveGenerateFillValue(field)
			if err != nil {
				return err
			}
		case reflect.UnsafePointer:
		}
	}
	return nil
}
