package m5_download_reflect

import "reflect"

func InvokeReflection() {
	var target MyType
	var methodName string = "SafeMethod"
	methodName = "UnsafeMethod"
	v := reflect.ValueOf(target)
	m := v.MethodByName(methodName)
	m.Call(nil)
}
