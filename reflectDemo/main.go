package main

import (
	"fmt"
	"reflect"
	"strings"
)

// ProjectSession
type ProjectSession struct {
	ID    int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null"`
	Name  string `gorm:"column:name;type:varchar(50)"`
	Title string `gorm:"column:title;type:varchar(100)"`
}

func main() {
	var ProjectSessionDataArray []ProjectSession
	var single ProjectSession

	single.ID = 1
	single.Name = "测试"
	single.Title = "测试标题"

	ProjectSessionDataArray = append(ProjectSessionDataArray, single)

	ProjectSessionDataMap := Struct2Map(ProjectSessionDataArray)
	fmt.Println("转换后结果", ProjectSessionDataMap)
	Fields := column(ProjectSessionDataMap, "ID")
	fmt.Println("转换后结果", Fields)
	IDs := implode(Fields, ",")
	fmt.Println("转换后结果", IDs)
}

/**
获取二维map指定字段为新数组
*/
func column(arrayData []map[string]interface{}, field string) (returnData []interface{}) {
	for _, value := range arrayData {
		returnData = append(returnData, value[field])
	}
	return
}

/**
数组转字符 按指定field隔开
*/
func implode(arrayData []interface{}, field string) (returnString string) {
	for _, value := range arrayData {
		returnString = returnString + fmt.Sprint(value) + field
	}

	if returnString != "" {
		returnString = strings.TrimRight(returnString, field)
	}

	return
}

/**
结构体切片转map 反射
*/
func Struct2Map(obj interface{}) (returnData []map[string]interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// 判断类型
	switch t.Kind() {
	case reflect.Slice:
		// 切片
		for i := 0; i < v.Len(); i++ {
			data := ReflectValue2Map(v.Index(i))
			returnData = append(returnData, data)
		}
	case reflect.Struct:
		// 结构体
		data := ReflectValue2Map(v)
		returnData = append(returnData, data)
	default:
		panic("数据结构异常")
	}

	return returnData
}

/**
结构体转map 反射
*/
func ReflectValue2Map(obj reflect.Value) map[string]interface{} {
	var data = make(map[string]interface{})
	for ii := 0; ii < obj.NumField(); ii++ {
		data[obj.Type().Field(ii).Name] = obj.Field(ii).Interface()
	}

	return data
}
