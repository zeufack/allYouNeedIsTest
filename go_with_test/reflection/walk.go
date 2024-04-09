package reflection

import "reflect"

func walk(x interface{}, fun func(input string)) {
	val := getValue(x)

	numberOfValue := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numberOfValue = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numberOfValue = val.Len()
		getField = val.Index
	case reflect.String:
		fun(val.String())
	}

	for i := 0; i < numberOfValue; i++ {
		walk(getField(i).Interface(), fun)
	}
	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		walk(val.Index(i).Interface(), fun)
	// 	}
	// 	return
	// }
	// if val.Kind() == reflect.Pointer {
	// 	val = val.Elem()
	// // }
	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)

	// 	switch field.Kind() {
	// 	case reflect.String:
	// 		fun(field.String())
	// 	case reflect.Struct:
	// 		walk(field.Interface(), fun)
	// 	}
	// }

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
