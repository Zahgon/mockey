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
)

type Analyzer interface {
	// TargetType returns the type of the target function or method.
	TargetType() reflect.Type

	// TargetValue returns the value of the target function or method.
	TargetValue() reflect.Value

	// RuntimeTargetType returns the actual type of the target function at runtime. It contains generic type info if the
	// target is generic:
	//  1. function:
	//     a. non-generic function: func(inArgs) outArgs
	//     b. generic function: func(GenericInfo, inArgs) outArgs
	//  2. method:
	//     a. non-generic method: func(self *struct, inArgs) outArgs
	//     b. generic method:
	//     - BEFORE go1.20: func(GenericInfo, self *struct, inArgs) outArgs
	//     - AFTER go1.20: func(self *struct, GenericInfo, inArgs) outArgs
	RuntimeTargetType() reflect.Type

	// RuntimeTargetValue returns the actual value of the target function or method. If the target is generic, the returned
	// value is the gcshape function. Otherwise, it is the same as the target.
	RuntimeTargetValue() reflect.Value

	// IsGeneric returns true if the target is a generic function or method.
	IsGeneric() bool

	// GenericInfo returns the type info of the generic target.
	GenericInfo() GenericInfo

	// InputAdapter generates an adapter function to adapt the input arguments of the RuntimeTargetType() to the inputType.
	// These inputTypes are valid:
	//  1. function:
	//     a. non-generic function: func(inArgs) outArgs
	//     b. generic function: func(info GenericInfo, inArgs) outArgs OR func(inArgs) outArgs
	//  2. method:
	//     a. non-generic method: func(self *struct, inArgs) outArgs OR func(inArgs) outArgs
	//     b. generic method: func(GenericInfo, self *struct, inArgs) outArgs OR func(self *struct, inArgs) outArgs OR func(inArgs) outArgs
	InputAdapter(inputName string, inputType reflect.Type) func([]reflect.Value) []reflect.Value

	// ReversedInputAdapter generates an adapter function to adapt the input arguments of the inputType to the RuntimeTargetType().
	// These inputTypes are valid:
	//  1. function:
	//     a. non-generic function: func(inArgs) outArgs
	//     b. generic function: func(info GenericInfo, inArgs) outArgs OR func(inArgs) outArgs
	//  2. method:
	//     a. non-generic method: func(self *struct, inArgs) outArgs OR func(inArgs) outArgs
	//     b. generic method: func(GenericInfo, self *struct, inArgs) outArgs OR func(self *struct, inArgs) outArgs OR func(inArgs) outArgs
	ReversedInputAdapter(inputName string, inputType reflect.Type) func(inputArgs, extraArgs []reflect.Value) []reflect.Value
}

var (
	genericAnalyzedCount int64
	genericFallbackCount int64
)

func (a *AnalyzerImpl) TargetValue() reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func (a *AnalyzerImpl) TargetType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (a *AnalyzerImpl) RuntimeTargetType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (a *AnalyzerImpl) RuntimeTargetValue() reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func (a *AnalyzerImpl) IsGeneric() bool { _ = "STUB: not implemented"; return false }

func (a *AnalyzerImpl) GenericInfo() GenericInfo {
	_ = "STUB: not implemented"
	return *new(GenericInfo)
}

// runtimeTargetValueAndGenericInfo0 obtains the runtime value of the target and the generic information.
func (a *AnalyzerImpl) runtimeTargetValueAndGenericInfo0() (reflect.Value, GenericInfo) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(GenericInfo)
}

// Obtain the jump address and generic information address of the generic function through instruction analysis

// Create a function value based on the runtime type and the obtained jump address

// Fallback genericInfo: obtains generic information by means of actual execution

// extract genericInfo from args
