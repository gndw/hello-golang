package helpers

import (
	"encoding/json"
	"hello-golang/errors"
)

func StructToJSON(value interface{}) (string, *errors.Error) {
	obj, err := json.Marshal(value)
	if err != nil {
		e := errors.CreateError(errors.FailedToParseToJSON)
		e.AddDetail(errors.CreateErrorFromPrimitiveError(err))
		return "", e
	}

	return string(obj), nil
}

func StructToSafeJSONString(value interface{}) string {
	result, err := StructToJSON(value)
	if err != nil {
		return err.Error()
	} else {
		return result
	}
}
