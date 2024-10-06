package utils

import (
	"fmt"
	"reflect"
)

func IsInsideArray(value interface{}, array interface{}) bool {
	if value == nil {
		panic("Value must existed")
	}

	if array == nil {
		panic("Array must existed")
	}

	if !IsArray(array) || !IsSlice(array) {
		panic("Argument #2 must be an Array or Slice")
	}

	var targetSlice []interface{}
	val := reflect.ValueOf(array)
	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		targetSlice[i] = ele
	}

	fmt.Println("new array:", targetSlice)

	return true
}

// func isValidArrayOrSlice(array interface{}) bool {
// 	isValid := true
// 	switch isValid {
// 		case
// 	}

// 	return isValid
// }

func IsSlice(slice interface{}) bool {
	v := reflect.TypeOf(slice)
	fmt.Println("reflect value:", reflect.ValueOf(slice))
	return v.Kind() != reflect.Slice
}

func IsArray(array interface{}) bool {
	v := reflect.TypeOf(array)
	fmt.Println("reflect value:", reflect.ValueOf(array))

	return v.Kind() != reflect.Array
}
