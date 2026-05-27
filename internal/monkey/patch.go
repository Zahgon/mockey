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

package monkey

import (
	"reflect"
)

// Patch is a context that holds the address and original codes of the patched function.
type Patch struct {
	size int
	code []byte
	base uintptr
}

// Base returns the address of the patched function.
func (p *Patch) Base() uintptr {
	_ = "STUB: not implemented"

	// Unpatch restores the patched function to the original function.
	return 0
}

func (p *Patch) Unpatch() { _ = "STUB: not implemented"; return }

// PatchValue replace the target function with a hook function, and stores the target function in the proxy function
// for future restore. Target and hook are values of function. Proxy is a value of proxy function pointer.
func PatchValue(target, hook, proxy reflect.Value, unsafe bool) *Patch {
	_ = "STUB: not implemented"
	return nil
}

// The first few bytes of the target function code

// construct the branch instruction, i.e. jump to the hook function

// construct the proxy code

// search the cutting point of the target code, i.e. the minimum length of full instructions that is longer than the hookCode

// save the original code before the cutting point

// construct the branch instruction, i.e. jump to the cutting point

// inject the proxy code to the proxy function

// replace target function codes before the cutting point

func PatchFunc(fn, hook, proxy interface{}, unsafe bool) *Patch {
	_ = "STUB: not implemented"
	return nil
}
