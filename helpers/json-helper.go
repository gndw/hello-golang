package helpers

import (
	"container/list"
	"encoding/json"
	gerrors "errors"
	"hello-golang/errors"
	"hello-golang/requests"
)

func StructToJSONString(value interface{}) (string, *errors.Error) {
	obj, err := json.Marshal(value)
	if err != nil {
		e := errors.CreateError(errors.FailedToParseToJSON)
		e.AddDetail(errors.CreateErrorFromPrimitiveError(err))
		return "", e
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

func JSONStringToStruct(str string, value interface{}, validations *[]requests.Validation) *errors.Error {
	unmarshal_error := json.Unmarshal([]byte(str), value)
	if unmarshal_error != nil {
		return errors.CreateErrorFromPrimitiveError(unmarshal_error)
	}

	li := list.New()
	for _, v := range *validations {
		e := v()
		if e != nil {
			li.PushBack(e)
		}
	}

	if li.Len() > 0 {
		validation_error := errors.CreateErrorFromPrimitiveError(gerrors.New("object not contains specific data"))
		for ve := li.Front(); ve != nil; ve = ve.Next() {
			validation_error.AddDetail(ve.Value.(*errors.Error))
		}
		return validation_error
	} else {
		return nil
	}
}
