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

// InjectInto injects the raw codes into the target to make a new function. The target is the target function pointer.
func InjectInto(target reflect.Value, fnCode []byte) { _ = "STUB: not implemented"; return }

// ensure the code is executable

// make a new function to receive the code

// set the target with the new made function

func MakeFunc(typ reflect.Type, addr uintptr) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
