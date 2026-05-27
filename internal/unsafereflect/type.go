/*
 * Copyright 2023 ByteDance Inc.
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
 *
 * Unsafe reflect package for mockey, copy most code from go/src/reflect/type.go,
 * allow to export the address of private member methods.
 */

package unsafereflect

import (
	"reflect"
	"unsafe"
)

// MethodByName returns the method with the given name.
// NOTE: This may fail, depending on whether the relevant function type is ignored during compilation
func MethodByName(r reflect.Type, name string) (typ reflect.Type, addr uintptr, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), 0, false
}

// copy from src/reflect/type.go
// rtype is the common implementation of most values.
// It is embedded in other struct types.
//
// rtype must be kept in sync with src/runtime/type.go:/^type._type.
type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      tflag   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C

	// In go 1.13 equal was replaced with "alg *typeAlg".
	// Since size(func) == size(ptr), the total size of rtype
	// and alignment of other field keeps the same, we do not
	// need to make an adaption for go1.13.
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte   // garbage collection data
	str       nameOff // string form
	ptrToThis typeOff // type for pointer to this type, may be zero
}

const kindMask = (1 << 5) - 1

func (t *rtype) Kind() reflect.Kind { _ = "STUB: not implemented"; return *new(reflect.Kind) }

type (
	tflag   uint8
	nameOff int32 // offset to a name
	typeOff int32 // offset to an *rtype
	textOff int32 // offset from top of text section
)

// resolveNameOff resolves a name offset from a base pointer.
// The (*rtype).nameOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
//
//go:linkname resolveNameOff reflect.resolveNameOff
func resolveNameOff(unsafe.Pointer, int32) unsafe.Pointer

func (t *rtype) nameOff(off nameOff) name { _ = "STUB: not implemented"; return *new(name) }

// resolveTypeOff resolves an *rtype offset from a base type.
// The (*rtype).typeOff method is a convenience wrapper for this function.
//
//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

func (t *rtype) typeOff(off typeOff) *rtype { _ = "STUB: not implemented"; return nil }

// toType convert rtype to reflect.Type
//
// The conversion is not guaranteed to be successful.
// If conversion failed, response will be nil
func toType(r *rtype) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// resolveTextOff resolves a function pointer offset from a base type.
// The (*rtype).textOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
//
//go:linkname resolveTextOff reflect.resolveTextOff
func resolveTextOff(unsafe.Pointer, int32) unsafe.Pointer

func (t *rtype) textOff(off textOff) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

const tflagUncommon tflag = 1 << 0

// uncommonType is present only for defined types or types with methods
type uncommonType struct {
	pkgPath nameOff // import path; empty for built-in types like int, string
	mcount  uint16  // number of methods
	xcount  uint16  // number of exported methods
	moff    uint32  // offset from this uncommontype to [mcount]method
	_       uint32  // unused
}

// ptrType represents a pointer type.
type ptrType struct {
	rtype
	elem *rtype // pointer element (pointed at) type
}

// funcType represents a function type.
type funcType struct {
	rtype
	inCount  uint16
	outCount uint16 // top bit is set if last input parameter is ...
}

func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// interfaceType represents an interface type.
type interfaceType struct {
	rtype
	pkgPath name      // import path
	methods []imethod // sorted by hash
}

type imethod struct {
	_ nameOff // unused name of method
	_ typeOff // unused .(*FuncType) underneath
}

func (t *rtype) methods() []method { _ = "STUB: not implemented"; return nil }

// Method on non-interface type
type method struct {
	name nameOff // name of method
	mtyp typeOff // method type (without receiver), not valid for private methods
	_    textOff // unused fn used in interface call (one-word receiver)
	tfn  textOff // fn used for normal method call
}

func (t *uncommonType) methods() []method { _ = "STUB: not implemented"; return nil }

// Struct field
type structField struct {
	_ name    // unused name is always non-empty
	_ *rtype  // unused type of field
	_ uintptr // unused byte offset of field
}

// structType
type structType struct {
	rtype
	pkgPath name
	fields  []structField // sorted by offset
}

type _String struct {
	Data unsafe.Pointer
	Len  int
}
