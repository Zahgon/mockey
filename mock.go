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
	"sync"

	"github.com/bytedance/mockey/internal/fn"
	"github.com/bytedance/mockey/internal/monkey"
	"github.com/bytedance/mockey/internal/tool"
)

type FilterGoroutineType int64

const (
	Disable FilterGoroutineType = 0
	Include FilterGoroutineType = 1
	Exclude FilterGoroutineType = 2
)

type Mocker struct {
	target    reflect.Value // mock target value
	hook      reflect.Value // mock hook value
	proxy     reflect.Value // proxy pointer value
	times     int64
	mockTimes int64
	patch     *monkey.Patch
	lock      sync.Mutex
	isPatched bool
	builder   *MockBuilder

	outerCaller tool.CallerInfo
}

type MockBuilder struct {
	target          interface{}      // mock target
	originPtr       interface{}      // origin caller
	conditions      []*mockCondition // mock conditions
	filterGoroutine FilterGoroutineType
	gId             int64
	unsafe          bool
	analyzer        fn.Analyzer
}

// Mock mocks target function.
// From go1.20, Mock can automatically judge whether the target is generic or not. Before go1.20, you need to use
// MockGeneric to mock generic function.
func Mock(target interface{}, opt ...mockOptionFn) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

// MockUnsafe has the full ability of the Mock function and removes some security restrictions. This is an alternative
// when the Mock function fails. It may cause some unknown problems, so we recommend using Mock under normal conditions.
func MockUnsafe(target interface{}, opt ...mockOptionFn) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

// runtimeTargetType returns the type of the target function with generic type info if it's generic.
func (builder *MockBuilder) runtimeTargetType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (builder *MockBuilder) resetCondition() *MockBuilder { _ = "STUB: not implemented"; return nil }

// at least 1 condition is needed

// Origin add an origin hook which can be used to call un-mocked origin function
//
// For example:
//
//	 origin := Fun // only need the same type
//	 mock := func(p string) string {
//		 return origin(p + "mocked")
//	 }
//	 mock2 := Mock(Fun).To(mock).Origin(&origin).Build()
//
// Origin only works when call origin hook directly, target will still be mocked in recursive call
func (builder *MockBuilder) Origin(funcPtr interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) origin(funcPtr interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) lastCondition() *mockCondition { _ = "STUB: not implemented"; return nil }

func (builder *MockBuilder) newCondition() *mockCondition { _ = "STUB: not implemented"; return nil }

// When declares the condition hook that's called to determine whether the mock should be executed.
//
// The condition hook function must have the same parameters as the target function.
//
// The following example would execute the mock when input int is negative
//
//	func Fun(input int) string {
//		return strconv.Itoa(input)
//	}
//	Mock(Fun).When(func(input int) bool { return input < 0 }).Return("0").Build()
//
// Note that if the target function is a struct method, you may optionally include
// the receiver as the first argument of the condition hook function. For example,
//
//	type Foo struct {
//		Age int
//	}
//	func (f *Foo) GetAge(younger int) string {
//		return strconv.Itoa(f.Age - younger)
//	}
//	Mock((*Foo).GetAge).When(func(f *Foo, younger int) bool { return younger < 0 }).Return("0").Build()
func (builder *MockBuilder) When(when interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

// To declares the hook function that's called to replace the target function.
//
// The hook function must have the same signature as the target function.
//
// The following example would make Fun always return true
//
//	func Fun(input string) bool {
//		return input == "fun"
//	}
//
//	Mock(Fun).To(func(_ string) bool {return true}).Build()
//
// Note that if the target function is a struct method, you may optionally include
// the receiver as the first argument of the hook function. For example,
//
//	type Foo struct {
//		Name string
//	}
//	func (f *Foo) Bar(other string) bool {
//		return other == f.Name
//	}
//	Mock((*Foo).Bar).To(func(f *Foo, other string) bool {return true}).Build()
func (builder *MockBuilder) To(hook interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) Return(results ...interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) IncludeCurrentGoRoutine() *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) ExcludeCurrentGoRoutine() *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) FilterGoRoutine(filter FilterGoroutineType, gId int64) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) Build() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) build() { _ = "STUB: not implemented"; return }

// when condition is not set, just go into hook exec

// hook condition is not set, just go into original exec

// Origin call need extra args, which only can be obtained during the execution of mockerHook.

// Check if the currently called generic function instance matches the mocked generic function

func (mocker *Mocker) Patch() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) UnPatch() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) Release() *MockBuilder { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) ExcludeCurrentGoRoutine() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) FilterGoRoutine(filter FilterGoroutineType, gId int64) *Mocker {
	_ = "STUB: not implemented"
	return nil
}

func (mocker *Mocker) IncludeCurrentGoRoutine() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) When(when interface{}) *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) To(to interface{}) *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) Return(results ...interface{}) *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) Origin(funcPtr interface{}) *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) rePatch(do func()) *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) access() { _ = "STUB: not implemented"; return }

func (mocker *Mocker) mock() { _ = "STUB: not implemented"; return }

func (mocker *Mocker) Times() int { _ = "STUB: not implemented"; return 0 }

func (mocker *Mocker) MockTimes() int { _ = "STUB: not implemented"; return 0 }

func (mocker *Mocker) key() uintptr { _ = "STUB: not implemented"; return 0 }

func (mocker *Mocker) name() string { _ = "STUB: not implemented"; return "" }

func (mocker *Mocker) unPatch() { _ = "STUB: not implemented"; return }

func (mocker *Mocker) caller() tool.CallerInfo {
	_ = "STUB: not implemented"
	return *new(tool.CallerInfo)
}
