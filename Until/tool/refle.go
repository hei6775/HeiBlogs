package tool

import "reflect"

//获取接口中存放的实例的类型名
func GetTypeName(i interface{}) string {

	if i == nil {
		return ""
	}

	rt := reflect.TypeOf(i)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return rt.Name()
}
