/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govalidator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ToString convert the input to a string.
func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
	return res
}

// ToJSON convert the input to a valid JSON string
func ToJSON(obj interface{}) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		res = []byte("")
	}
	return string(res), err
}

// ToFloat convert the input string to a float, or 0.0 if the input is not a float.
func ToFloat(value interface{}) (res float64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = float64(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = float64(val.Uint())
	case float32, float64:
		res = val.Float()
	case string:
		res, err = strconv.ParseFloat(val.String(), 64)
		if err != nil {
			res = 0
		}
	default:
		err = fmt.Errorf("ToInt: unknown interface type %T", value)
		res = 0
	}

	return
}

// ToInt convert the input string or any int type to an integer type 64, or 0 if the input is not an integer.
func ToInt(value interface{}) (res int64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		res = int64(val.Uint())
	case float32, float64:
		res = int64(val.Float())
	case string:
		if IsInt(val.String()) {
			res, err = strconv.ParseInt(val.String(), 0, 64)
			if err != nil {
				res = 0
			}
		} else {
			err = fmt.Errorf("ToInt: invalid numeric format %g", value)
			res = 0
		}
	default:
		err = fmt.Errorf("ToInt: unknown interface type %T", value)
		res = 0
	}

	return
}

// ToBoolean convert the input string to a boolean.
func ToBoolean(str string) (bool, error) {
	return strconv.ParseBool(str)
}
