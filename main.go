// Copyright 2018, Brian Wiborg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package compare provides functions for comparing types by value.
package compare

import (
	"reflect"
	"strconv"
)

// Comparable defines the interface for a comparable type.
type Comparable interface {
	Int64() int64
}

func getValues(a, b interface{}) (int64, int64) {
	var (
		left, right int64
		av          = reflect.ValueOf(a)
		bv          = reflect.ValueOf(b)
		at          = reflect.TypeOf(a)
		bt          = reflect.TypeOf(b)
	)

	switch av.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		left = int64(av.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		left = av.Int()
	case reflect.String:
		left, _ = strconv.ParseInt(av.String(), 10, 64)
	default:
		if at.Implements(reflect.TypeOf((*Comparable)(nil)).Elem()) {
			left = a.(Comparable).Int64()
		}
		panic("can not compare unsupported type")
	}

	switch bv.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		right = int64(bv.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		right = bv.Int()
	case reflect.String:
		right, _ = strconv.ParseInt(bv.String(), 10, 64)
	default:
		if bt.Implements(reflect.TypeOf((*Comparable)(nil)).Elem()) {
			right = b.(Comparable).Int64()
		}
		panic("can not compare unsupported type")
	}

	return left, right
}

// Gt takes two types and returns true if the left type is greather then the rigth type.
// If the type is an arrays, chan, map or slice, this function will consider its length as value.
// If the type is a string, it will be parsed as a base-10 integer.
// If the type is a struct and satisfies Combarable, its Int64() function will be used.
func Gt(a, b interface{}) bool {
	left, right := getValues(a, b)
	return left > right
}

// Ge takes two types and returns true if the left type is greater than or equal to teh right type.
// If the types is an array, chan, map or slice, this function will consider its length as value.
// If the type is a string, it will be parsed as a base-10 integer.
// If the type is a struct and satisfies Combarable, its Int64() function will be used.
func Ge(a, b interface{}) bool {
	left, right := getValues(a, b)
	return left >= right
}

// Lt takes two types and returns true if the left type is less than the right type.
// If the type is an array, chan, map or slice, this function will consider its length as value.
// If the type is a string, it will be parsed as a base-10 integer.
// If the type is a struct and satisfies Combarable, its Int64() function will be used.
func Lt(a, b interface{}) bool {
	left, right := getValues(a, b)
	return left < right
}

// Le takes two types and returns true if the left type is less than or equal to the right type.
// If the type is an array, chan, map or slice, this function will consider its length as value.
// If the type is a string, it will be parsed as a base-10 integer.
// If the type is a struct and satisfies Combarable, its Int64() function will be used.
func Le(a, b interface{}) bool {
	left, right := getValues(a, b)
	return left <= right
}

// Eq takes two types and returns true if the left type is equal to the right type.
// If the type is an array, chan, map or slice, this function will consider its length as value.
// If the type is a string, it will be parsed as a base-10 integer.
// If the type is a struct and satisfies Combarable, its Int64() function will be used.
func Eq(a, b interface{}) bool {
	left, right := getValues(a, b)
	return left == right
}
