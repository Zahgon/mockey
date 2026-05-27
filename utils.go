/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mockey

import (
	"reflect"
)

// GetMethod resolves a method with the specified name from the given instance.
// Supports finding:
// - Exported and unexported methods
// - Methods for value types and pointer types
// - Methods in nested anonymous fields
// - Method fields of structs
// Parameters:
// - instance: The instance to find the method on
// - methodName: The name of the method to find
// Return value:
// - The interface of the found method, triggers assertion failure if not found
func GetMethod(instance interface{}, methodName string, opt ...methodOptionFn) (res interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func getMethod(val reflect.Value, methodName string, opts *methodOption) (method reflect.Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// if is nil pointer but with type, try get method from pointer type

// check nested if it has anonymous fields

// check elem type for exported method

// check elem type for field method

// check elem type for unexported method

// check ptr type for exported or unexported method

// getFieldMethod gets a functional field's value as an instance
// The return instance is not original field but a new function object points to
// the same function.
// for example:
//
//	  type Fn func()
//	  type Foo struct {
//			privateField Fn
//	  }
//	  func NewFoo() Foo { return Foo{ privateField: func() { /*do nothing*/ } }}
//
// getFieldMethod(NewFoo(),"privateField") will return a function object which
// points to the anonymous function in NewFoo
func getFieldMethod(v reflect.Value, fieldName string) (res reflect.Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// GetPrivateMethod resolve a certain public method from an instance.
// Deprecated, this is an old API in mockito. Please use GetMethod instead.
func GetPrivateMethod(instance interface{}, methodName string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// GetNestedMethod resolves a certain public method in anonymous structs, it will
// look for the specific method in every anonymous struct field recursively.
// Deprecated, this is an old API in mockito. Please use GetMethod instead.
func GetNestedMethod(instance interface{}, methodName string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func getNestedMethod(val reflect.Value, methodName string) (reflect.Method, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Method), false
}

// there is no need to acquire non-anonymous method

// a struct receiver is prior to the corresponding pointer receiver

// unexportedMethodByName resolve an unexported method from an instance
func unexportedMethodByName(instanceType reflect.Type, methodName string, opts *methodOption) (res reflect.Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// GetGoroutineId gets the current goroutine ID
func GetGoroutineId() int64 { _ = "STUB: not implemented"; return 0 }
