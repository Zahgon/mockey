//go:build !go1.20
// +build !go1.20

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

	targetValue        reflect.Value
	targetType         reflect.Type
	generic            bool
	runtimeTargetType  reflect.Type
	runtimeTargetValue reflect.Value
	runtimeGenericInfo GenericInfo
}

func (a *AnalyzerImpl) init() *AnalyzerImpl {
	tool.DebugPrintf("[Analyzer.init] start analyze, genericIn: %v\n", a.genericIn)
	a.targetValue, a.targetType = reflect.ValueOf(a.target), reflect.TypeOf(a.target)
	tool.DebugPrintf("[Analyzer.init] targetType: %v, targetValue: 0x%x\n", a.targetType, a.targetValue.Pointer())
	if a.genericIn == nil {
		a.generic = false
	} else {
		a.generic = *a.genericIn
	}
	a.runtimeTargetType = a.runtimeTargetType0()
	a.runtimeTargetValue, a.runtimeGenericInfo = a.runtimeTargetValueAndGenericInfo0()
	tool.DebugPrintf("[Analyzer.init] analyze finish, generic: %v, runtimeTargetType: %v, runtimeTargetValue: 0x%x, runtimeGenericInfo: 0x%x\n", a.generic, a.runtimeTargetType, a.runtimeTargetValue.Pointer(), a.runtimeGenericInfo)
	return a
}

func (a *AnalyzerImpl) runtimeTargetType0() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

// generic information needs to be inserted at position 0

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
//     b. generic method: func(GenericInfo, self *struct, inArgs) outArgs

// check:
// 1. function:
//     a. generic function: func(inArgs) outArgs
// 2. method:
//     a. non-generic method: func(inArgs) outArgs
//     b. generic method: func(self *struct, inArgs) outArgs

// check:
// 1. method:
//     a. generic method: func(inArgs) outArgs

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
//     b. generic method: func(GenericInfo, self *struct, inArgs) outArgs

// check:
// 1. function:
//     a. generic function: func(inArgs) outArgs
// 2. method:
//     a. non-generic method: func(inArgs) outArgs
//     b. generic method: func(self *struct, inArgs) outArgs

// check:
// 1. method:
//     a. generic method: func(inArgs) outArgs
