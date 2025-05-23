package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
)

func IntToString(intVal int) string {
	return strconv.Itoa(intVal)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IsContain(target string, list []string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}
	return false
}

func InterfaceArrayToStringArray(data []interface{}) (i []string) {
	for _, param := range data {
		i = append(i, param.(string))
	}
	return i
}

func StructToJsonString(param interface{}) (string, error) {
	dataType, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	dataString := string(dataType)
	return dataString, nil
}

// 传入指针,修改指针
func JsonStringToStruct(s string, args interface{}) error {
	err := json.Unmarshal([]byte(s), args)
	return err
}

// 使用泛型,返回结构体
func JsonStringToStruct2[T any](s string) (T, error) {
	var obj T
	err := json.Unmarshal([]byte(s), &obj)
	return obj, err
}

// MD5(时间戳+sendId+随机数) 偏业务的方法,可以分离出去
func GetMsgID(sendId string) string {
	t := int64ToString(GetCurrentTimestampByNano())
	return Md5(t + sendId + int64ToString(rand.Int63n(GetCurrentTimestampByNano())))
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
