package helpers

import (
	"encoding/json"
)

func StructToJSONString(value interface{}) (str string, err error) {
	obj, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(obj), nil
}

func StructToSafeJSONString(value interface{}) string {
	result, err := StructToJSONString(value)
	if err != nil {
		return err.Error()
	} else {
		return result
	}
}

func JSONStringToStruct(str string, value interface{}) (err error) {
	err = json.Unmarshal([]byte(str), value)
	if err != nil {
		return err
	}
	return nil
}
