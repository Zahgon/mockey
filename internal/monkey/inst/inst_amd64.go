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

// MOVABS RDX, to
// JMP RDX

func BranchInto(to uintptr) (res []byte) { _ = "STUB: not implemented"; return nil }

// MOVABS RDX, to
// JMP [RDX]

// rdxMOV moves the 64bit value to rdx register, using the following instruction:
// MOVABS RDX, val
func rdxMOV(val uintptr) []byte { _ = "STUB: not implemented"; return nil }
