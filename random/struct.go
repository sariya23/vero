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

	for i := 0; i < value.Elem().NumField(); i++ {
		field := value.Elem().Field(i)
		fieldType := value.Elem().Type().Field(i)
		generatedValue := recursiveGenerateFillValue(fieldType.Type.Kind())
		if !field.CanSet() {
			continue
		}
		field.Set(reflect.ValueOf(generatedValue))
	}
}

func recursiveGenerateFillValue(structValue reflect.Value) any {
	for i := 0; i < structValue.Elem().NumField(); i++ {
		field := structValue.Elem().Field(i)
		fieldType := structValue.Elem().Type().Field(i)
		generatedValue := recursiveGenerateFillValue(fieldType.Type.Kind())
		if !field.CanSet() {
			continue
		}
		field.Set(reflect.ValueOf(generatedValue))
	}
	switch fieldTypeKind {
	case reflect.Bool:
		return true
	case reflect.Int:
		return 8
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
		recursiveGenerateFillValue(fieldTypeKind)
	case reflect.UnsafePointer:
	}
	return nil
}
