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

import (
	"golang.org/x/arch/arm64/arm64asm"
)

const (
	instLen = 4 // arm64 instruction length is 4 bytes
)

func calcFnAddrRange(name string, fn func()) (uintptr, uintptr) {
	_ = "STUB: not implemented"
	return 0, 0
}

func Disassemble(code []byte, required int, checkLen bool) int { _ = "STUB: not implemented"; return 0 }

func GetGenericAddr(addr uintptr, maxScan int) (jumpAddr, genericInfoAddr uintptr) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Find the only jumpInst and filter the extra call

// Find the latest infoInst before the jumpInst

// In some cases, genericInfoAddr needs to be calculated and cannot be directly obtained by analyzing instructions.

type posInst struct {
	pos  int
	addr uintptr
	inst arm64asm.Inst
}

func (pi *posInst) String() string { _ = "STUB: not implemented"; return "" }

func newGenericJmpInst(base uintptr, pos int, inst arm64asm.Inst) *genericJmpInst {
	_ = "STUB: not implemented"
	return nil
}

type genericJmpInst struct {
	*posInst
	jumpAddr      uintptr
	isExtraCall   bool
	extraCallName string
}

func (g *genericJmpInst) String() string { _ = "STUB: not implemented"; return "" }

func (g *genericJmpInst) init() *genericJmpInst {
	g.jumpAddr = g.calcJumpAddr()
	g.isExtraCall, g.extraCallName = isGenericProxyCallExtra(g.jumpAddr)
	return g
}

func (g *genericJmpInst) calcJumpAddr() uintptr { _ = "STUB: not implemented"; return 0 }

func newGenericInfoInst(base uintptr, pos int, inst arm64asm.Inst) *genericInfoInst {
	_ = "STUB: not implemented"
	return nil
}

type genericInfoInst struct {
	adrp *posInst
	add  *posInst
}

func (g *genericInfoInst) String() string { _ = "STUB: not implemented"; return "" }

func (g *genericInfoInst) matchJumpInst(jumpInst *genericJmpInst) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *genericInfoInst) putAddInst(pos int, inst arm64asm.Inst) {
	_ = "STUB: not implemented"
	return
}

// calcGenericInfoAddr calculates the genericInfo from the adrp and add instructions. Example:
// ADRP X0, .+0xb3000
// ADD X0, X0, #0xec0
func (g *genericInfoInst) calcGenericInfoAddr() uintptr { _ = "STUB: not implemented"; return 0 }
