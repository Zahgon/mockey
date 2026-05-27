//go:build go1.20 && !go1.27
// +build go1.20,!go1.27

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

package fn

import (
	"reflect"

	"github.com/bytedance/mockey/internal/tool"
)

func NewAnalyzer(target interface{}, generic *bool, method *bool) Analyzer {
	_ = "STUB: not implemented"
	return *new(Analyzer)
}

type AnalyzerImpl struct {
	target    interface{}
	genericIn *bool
	methodIn  *bool

	targetValue        reflect.Value
	targetType         reflect.Type
	nameAnalyzer       *NameAnalyzer
	generic            bool
	method             bool // DO NOT use it unless absolutely necessary
	runtimeTargetType  reflect.Type
	runtimeTargetValue reflect.Value
	runtimeGenericInfo GenericInfo
}

// init initializes the AnalyzerImpl. If `a.genericIn` or `a.methodIn` is set, it will be used directly, else it will be
// determined by the NameAnalyzer.
func (a *AnalyzerImpl) init() *AnalyzerImpl {
	tool.DebugPrintf("[Analyzer.init] start analyze, genericIn: %v, methodIn: %v\n", a.genericIn, a.methodIn)
	a.targetValue, a.targetType = reflect.ValueOf(a.target), reflect.TypeOf(a.target)
	tool.DebugPrintf("[Analyzer.init] targetType: %v, targetValue: 0x%x\n", a.targetType, a.targetValue.Pointer())
	a.generic, a.method = a.isGeneric0(), a.isMethod0()
	a.runtimeTargetType = a.runtimeTargetType0()
	a.runtimeTargetValue, a.runtimeGenericInfo = a.runtimeTargetValueAndGenericInfo0()
	tool.DebugPrintf("[Analyzer.init] analyze finish, generic: %v, method: %v, runtimeTargetType: %v, runtimeTargetValue: 0x%x, runtimeGenericInfo: 0x%x\n", a.generic, a.method, a.runtimeTargetType, a.runtimeTargetValue.Pointer(), a.runtimeGenericInfo)
	return a
}

func (a *AnalyzerImpl) isGeneric0() bool { _ = "STUB: not implemented"; return false }

func (a *AnalyzerImpl) isMethod0() bool { _ = "STUB: not implemented"; return false }

// Analyze whether the function is a method.
// NOTE: For methods named 'func\d+', misjudgment may occur, and they will not be regarded as methods.
// ```
// type foo struct {}
// func (f *foo) func1() {}
// ```
// The fullname of 'f.func1' is 'main.foo.func1' which is regarded as an anonymous function in function 'main.foo'.

// A function without an argument or middleName is definitely not a method.

// If the receiver is a pointer type, it must be a method

// For exported functions, it is sufficient to iterate through all methods of the receiver and look for a method
// with the same type and pointer.

// For unexported functions, it is highly challenging to determine whether they are methods. We can rule out some
// cases and the rest are regarded as methods.

func (a *AnalyzerImpl) runtimeTargetType0() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

// for methods, generic information needs to be inserted at position 1 after go1.20

// for functions, generic information needs to be inserted at position 0

func (a *AnalyzerImpl) InputAdapter(inputName string, inputType reflect.Type) func([]reflect.Value) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

// check:
// 1. function:
//     a. non-generic function: func(inArgs) outArgs
//     b. generic function: func(info GenericInfo, inArgs) outArgs
// 2. method:
//     a. non-generic method: func(self *struct, inArgs) outArgs

// need to adapt the arguments.

func (a *AnalyzerImpl) ReversedInputAdapter(inputName string, inputType reflect.Type) func(inputArgs, extraArgs []reflect.Value) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

// check:
// 1. function:
//     a. non-generic function: func(inArgs) outArgs
//     b. generic function: func(info GenericInfo, inArgs) outArgs
// 2. method:
//     a. non-generic method: func(self *struct, inArgs) outArgs

// need to adapt the arguments.

type (
	fn         = func(targetArgs []reflect.Value) []reflect.Value
	reversedFn = func(inputArgs, extraArgs []reflect.Value) []reflect.Value
)

func (a *AnalyzerImpl) nonGenericAnalyzer(inputName string, inputType reflect.Type) (fn, reversedFn) {
	_ = "STUB: not implemented"
	// check:
	// 1. method:
	//     a. non-generic method: func(inArgs) outArgs
	return *new(fn), *new(reversedFn)
}

func (a *AnalyzerImpl) genericAnalyzer(inputName string, inputType reflect.Type) (fn, reversedFn) {
	_ = "STUB: not implemented"
	return *new(fn), *new(reversedFn)
}

// check function:
// 		a. generic function: func(inArgs) outArgs

// check method:
// 		a. generic method: func(GenericInfo, self *struct, inArgs)

// check method:
// 		a. generic method: func(self *struct, inArgs)

// check method:
// 		a. generic method: func(inArgs)
