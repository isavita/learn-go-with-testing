package reflection

import "reflect"

func Walk(x interface{}, fn func(string)) {
	val := getValue(x)
	walkValue := func(v reflect.Value) {
		Walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		fnType := val.Type()
		var args []reflect.Value
		for i := 0; i < fnType.NumIn(); i++ {
			argType := fnType.In(i)
			args = append(args, reflect.New(argType).Elem())
		}
		valFnResult := val.Call(args)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
