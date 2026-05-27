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

package inst

func BranchTo(to uintptr) (res []byte) { _ = "STUB: not implemented"; return nil }

// MOV x26, to // fake
// BR x26

// BranchInto create a branch into command
//
// Go supports passing function arguments from go 1.17 (see https://go.dev/doc/go1.17).
// We could not use x0~x18 register. As an alternative, we use R19 register (see https://go.googlesource.com/go/+/refs/heads/master/src/cmd/compile/abi-internal.md).
func BranchInto(to uintptr) (res []byte) { _ = "STUB: not implemented"; return nil }

// MOV x26, to // fake
// LDR x19, [x26]
// BR x19

const x26 uint32 = 0b11010

// x26MOV moves the 64bit value to x26 register, using the following four instructions:
// MOVZ x26, val[0:16]
// MOVK x26, val[16:32]
// MOVK x26, val[32:48]
// MOVK x26, val[48:64]
func x26MOV(val uintptr) (res []byte) { _ = "STUB: not implemented"; return nil }

// x26MOVZ see https://developer.arm.com/documentation/ddi0596/2021-12/Base-Instructions/MOVZ--Move-wide-with-zero-
func x26MOVZ(val uintptr) []byte { _ = "STUB: not implemented"; return nil }

// x26MOVK see https://developer.arm.com/documentation/ddi0596/2021-12/Base-Instructions/MOVK--Move-wide-with-keep-
func x26MOVK(val uintptr, shift int) []byte { _ = "STUB: not implemented"; return nil }
