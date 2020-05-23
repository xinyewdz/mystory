package util

import (
	"errors"
	"reflect"
	"story-api/global"
)

var log = global.MainLog

func CopyProperties(src interface{},dest interface{})error{
	sVal := reflect.ValueOf(src)
	sType := reflect.TypeOf(src)
	dVal := reflect.ValueOf(dest)
	dType := reflect.TypeOf(dest)
	// dst必须结构体指针类型
	if dType.Kind()!=reflect.Ptr||dType.Elem().Kind()!=reflect.Struct{
		return errors.New("destType error.type="+dType.Kind().String())
	}
	// src必须为结构体或者结构体指针
	if sType.Kind()==reflect.Ptr{
		sType,sVal = sType.Elem(),sVal.Elem()
	}
	if sType.Kind()!=reflect.Struct{
		return errors.New("srcType error.type="+sType.Kind().String())
	}
	dType = dType.Elem()
	dVal = dVal.Elem()
	fieldCount := dType.NumField()
	for idx:=0;idx<fieldCount;idx++{
		dField := dType.Field(idx)
		val := sVal.FieldByName(dField.Name)
		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !val.IsValid()||val.Type()!=dField.Type{
			continue
		}
		if dVal.Field(idx).CanSet() {
			dVal.Field(idx).Set(val)
		}
	}
	return nil
}
