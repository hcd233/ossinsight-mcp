package ossinsight

import (
	"fmt"
	"net/url"
	"reflect"
)

// toQuery 将查询参数转换为 url.Values
//
//	@param query
//	@return url.Values
//	@author centonhuang
//	@update 2025-06-09 20:49:33
func toQuery(query any) url.Values {
	values := url.Values{}

	// 使用反射获取结构体字段
	v := reflect.ValueOf(query)
	t := v.Type()

	// 如果是指针类型,获取其指向的元素
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 获取json tag
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}

		// 如果字段值不为空,则添加到查询参数中
		if !value.IsZero() {
			values.Set(tag, fmt.Sprint(value.Interface()))
		}
	}

	return values
}
