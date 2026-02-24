package random

import (
	"reflect"

	"github.com/sariya23/vero/internal"
)

// Struct...
// TODO: добавить генерацию всех полей
// TODO: вынести генерацию в internal
func Struct(objPtr any) {
	value := reflect.ValueOf(objPtr)
	if value.Kind() != reflect.Ptr {
		panic("objPtr must be a pointer to struct")
	}

	if value.Elem().Kind() != reflect.Struct {
		panic("objPtr must be a pointer to struct")
	}

	recursiveGenerateFillValue(value.Elem())
}

func recursiveGenerateFillValue(structValue reflect.Value) {
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		if !field.CanSet() {
			continue
		}
		switch field.Type().Kind() {
		case reflect.Bool:
			structTag := structValue.Type().Field(i).Tag
			v := internal.GenerateBool(string(structTag))
			field.Set(reflect.ValueOf(v))
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
