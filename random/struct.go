package random

import "reflect"

// Struct...
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
		fieldType := structValue.Field(i).Type().Kind()
		switch fieldType {
		case reflect.Bool:
			if !field.CanSet() {
				continue
			}
			field.Set(reflect.ValueOf(true))
		case reflect.Int:
			if !field.CanSet() {
				continue
			}
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
