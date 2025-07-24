package convert

import (
	"reflect"
)

// CopyStructFields 通用结构体字段复制方法，将 src 的同名字段复制到 dst
func CopyStructFields(dst interface{}, src interface{}) {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	// dst 必须是指针
	if dstVal.Kind() != reflect.Ptr || dstVal.IsNil() {
		return
	}
	dstElem := dstVal.Elem()
	if dstElem.Kind() != reflect.Struct {
		return
	}

	// src 可以是指针或结构体
	if srcVal.Kind() == reflect.Ptr {
		if srcVal.IsNil() {
			return
		}
		srcVal = srcVal.Elem()
	}
	if srcVal.Kind() != reflect.Struct {
		return
	}

	dstType := dstElem.Type()
	for i := 0; i < dstElem.NumField(); i++ {
		field := dstType.Field(i)
		if !dstElem.Field(i).CanSet() {
			continue
		}
		srcField := srcVal.FieldByName(field.Name)
		if srcField.IsValid() && srcField.Type().AssignableTo(field.Type) {
			dstElem.Field(i).Set(srcField)
		}
	}
}

// CopyStructFieldsSlice 通用结构体切片字段复制方法，将 srcSlice 的每个元素复制到新建的 dstSlice 元素中
func CopyStructFieldsSlice(dstSlice interface{}, srcSlice interface{}) {
	dstVal := reflect.ValueOf(dstSlice)
	srcVal := reflect.ValueOf(srcSlice)

	// dstSlice 必须是指针
	if dstVal.Kind() != reflect.Ptr || dstVal.IsNil() {
		return
	}
	dstElem := dstVal.Elem()
	if dstElem.Kind() != reflect.Slice {
		return
	}

	// srcSlice 必须是切片
	if srcVal.Kind() == reflect.Ptr {
		if srcVal.IsNil() {
			return
		}
		srcVal = srcVal.Elem()
	}
	if srcVal.Kind() != reflect.Slice {
		return
	}

	dstElemType := dstElem.Type().Elem()
	for i := 0; i < srcVal.Len(); i++ {
		srcItem := srcVal.Index(i)
		// 创建新的目标元素
		var dstItem reflect.Value
		if dstElemType.Kind() == reflect.Ptr {
			dstItem = reflect.New(dstElemType.Elem())
			CopyStructFields(dstItem.Interface(), srcItem.Interface())
		} else {
			dstItem = reflect.New(dstElemType).Elem()
			CopyStructFields(dstItem.Addr().Interface(), srcItem.Interface())
		}
		dstElem.Set(reflect.Append(dstElem, dstItem))
	}
}
