package sqlstruct

import (
	"reflect"
)

func Marshall(a1 interface{}, b1 interface{}) error {
	if reflect.TypeOf(a1).Kind() == reflect.Slice {
		handleSlice(a1, b1)
		return nil
	}
	v := reflect.ValueOf(a1)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	f2 := reflect.Indirect(reflect.ValueOf(b1))

	for i := 0; i < v.NumField(); i++ {
		varName := v.Type().Field(i).Name
		varType := v.Type().Field(i).Type

		if varType.Kind() == reflect.Slice {

			//b1i := f2.FieldByName(varName).Interface()

			handleSlice(v.FieldByName(varName).Interface(), f2.FieldByName(varName).Addr().Interface())
			continue
		}

		if f2.Kind() == reflect.Ptr {
			f2 = f2.Elem()
		}

		e := f2.FieldByName(varName)

		varTypeOrig := varType.Name()

		if e.CanSet() {
			varTypeDest := e.Type().Name()
			switch  varTypeOrig {
			case "NullString":
				field := v.Field(i).Field(0)
				e.SetString(field.String())
				break
			case "NullInt64":
				field := v.Field(i).Field(0)
				e.SetInt(field.Int())
				break
			case "NullBool":
				field := v.Field(i).Field(0)
				e.SetBool(field.Bool())
				break
			case "NullFloat64":
				field := v.Field(i).Field(0)
				e.SetFloat(field.Float())
				break
			case "NullTime":
				e.Set(v.Field(i).Field(0))
				break
			default:
				if varTypeDest == "NullString" || varTypeDest == "NullInt64" ||
					varTypeDest == "NullBool" || varTypeDest == "NullFloat64" || varTypeDest == "NullTime" {
					e.Field(0).Set(v.FieldByName(varName))
					e.Field(1).Set(reflect.ValueOf(true))
				} else {
					if varTypeOrig == varTypeDest {
						e.Set(v.FieldByName(varName))
					}
				}

				break
			}
		}

	}

	return nil
}

func handleSlice(a1 interface{}, b1 interface{}) {
	var dest interface{}
	s := reflect.ValueOf(a1)
	m := reflect.ValueOf(b1)
	elemType := reflect.TypeOf(b1).Elem().Elem()

	isPtr := false
	if elemType.Kind() == reflect.Ptr {
		isPtr = true
		elemType = reflect.TypeOf(b1).Elem().Elem().Elem()
	}

	elemSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(b1).Elem().Elem()), 0, s.Len())

	for i := 0; i < s.Len(); i++ {
		dest = reflect.New(elemType).Interface()
		Marshall(s.Index(i).Interface(), dest)
		if isPtr {
			elemSlice = reflect.Append(elemSlice, reflect.ValueOf(dest))
		} else {
			elemSlice = reflect.Append(elemSlice, reflect.ValueOf(dest).Elem())
		}
	}

	m.Elem().Set(elemSlice)
}
