package utils

import (
	"reflect"
)

func IsInsideArray(value interface{}, array interface{}) bool {
	// REFACTOR: Need to apply another algorthm for searching
	if value == nil {
		panic("Value must existed")
	}

	if array == nil {
		panic("Array must existed")
	}

	if !isValidArrayOrSlice(array) {
		panic("Argument #2 must be an Array or Slice")
	}

	var targetSlice []interface{}
	val := reflect.ValueOf(array)
	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		targetSlice = append(targetSlice, ele.Interface())
	}

	for _, ele := range targetSlice {
		if ele == value {
			return true
		}
	}

	return false
}

func isValidArrayOrSlice(array interface{}) bool {
	if IsArray(array) {
		return true
	}

	if IsSlice(array) {
		return true
	}

	return false
}

func IsSlice(slice interface{}) bool {
	v := reflect.TypeOf(slice)
	return v.Kind() != reflect.Slice
}

func IsArray(array interface{}) bool {
	v := reflect.TypeOf(array)
	return v.Kind() != reflect.Array
}
