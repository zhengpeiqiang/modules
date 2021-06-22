package mapping

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

/*
	data 传入的源map
	obj 要转换的最终目标结构体
	isHump 是否开启驼峰转换
*/
func SlStruct(data map[string]interface{}, obj interface{}, isHump int) error {
	for k, v := range data {
		err := SetField(obj, k, v, isHump)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	obj 要转换的最终目标结构体
	name for循环中的 键名
	value for循环中的 键值
	isHump 是否开启驼峰转换 group_id ==》 GroupId
*/
func SetField(obj interface{}, name string, value interface{}, isHump int) error {
	structValue := reflect.ValueOf(obj).Elem() //结构体属性值
	if isHump == 1 {
		name = Marshal(name)
	}
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		return errors.New(fmt.Sprintf("No such field: %s in obj", name))
	}

	if !structFieldValue.CanSet() {
		return errors.New(fmt.Sprintf("Cannot set %s field value", name))
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}

/*
	value 值
	ntype 类型
*/
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}
	//else if .......增加其他一些类型的转换
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
