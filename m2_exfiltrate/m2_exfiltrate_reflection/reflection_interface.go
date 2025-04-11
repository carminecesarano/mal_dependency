package m2_exfiltrate_reflection

import "reflect"

func InvokeReflection() {
	var target MyType
	var methodName string = "SafeMethod"
	methodName = "UnsafeMethod"
	v := reflect.ValueOf(target)
	m := v.MethodByName(methodName)
	m.Call(nil)
}
