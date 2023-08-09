package object

import (
	"reflect"
)

func GetObjectInfo(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj).Elem()
	m := map[string]string{}
	m["type_name"] = t.Name()
	return m
}
