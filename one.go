package main

import (
	"fmt"
	"reflect"
)

func SetKey(in interface{}, key string, value interface{}) error {
	structValue := reflect.ValueOf(in).Elem()
	fieldValue := structValue.FieldByName(key)
	if !fieldValue.IsValid() {
		return fmt.Errorf("No field %s in In", key)
	}
	if !fieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", key)
	}
	val := reflect.ValueOf(value)
	if fieldValue.Type() != val.Type() {
		if m, ok := value.(map[string]interface{}); ok {
			if fieldValue.Kind() == reflect.Struct {
				return NewStruct(m, fieldValue.Addr().Interface())
			}
		}
		return fmt.Errorf("Provided value type didn't match In field type")

	}
	fieldValue.Set(val)
	return nil
}

func NewStruct(m map[string]interface{}, s interface{}) error {
	for k, v := range m {
		err := SetKey(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

type In struct {
	key string
}

func main() {
	data := make(map[string]interface{})
	data["key"] = "value"
	result := &In{}
	err := NewStruct(data, result)
	fmt.Println(err)
	fmt.Println(result)
}
