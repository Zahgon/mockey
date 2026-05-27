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

package tool

import (
	"reflect"
)

func ReflectCall(f reflect.Value, args []reflect.Value) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func NewFuncTypeByOut(ft reflect.Type, newOutTypes ...reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func NewFuncTypeByInsertIn(ft reflect.Type, newInTypes ...reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func NewFuncTypeByReplaceIn(ft reflect.Type, newInType reflect.Type, newInIndex int) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func MakeEmptyInArgs(ft reflect.Type) []reflect.Value { _ = "STUB: not implemented"; return nil }

func MakeEmptyOutArgs(ft reflect.Type) []reflect.Value { _ = "STUB: not implemented"; return nil }

func MakeEmtpy(typ reflect.Type) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func MakeReturnValues(ft reflect.Type, results ...interface{}) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}
