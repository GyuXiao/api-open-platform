package tools

import "reflect"

func StructToMap(obj interface{}) map[string]interface{} {
	// 通过反射获取结构体字段和值
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	result := make(map[string]interface{})
	// for 循环遍历每个字段
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)
		// 判断字段是否是指针类型
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				result[field.Name] = nil
			} else {
				result[field.Name] = fieldValue.Elem().Interface()
			}
			continue
		}
		// 非指针类型直接赋值
		result[field.Name] = fieldValue.Interface()
	}
	return result
}
