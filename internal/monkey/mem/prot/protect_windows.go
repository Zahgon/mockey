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

package prot

import (
	"syscall"
)

const (
	protectRWX = 0x40
)

var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func MProtectRWX(addr uintptr) error { _ = "STUB: not implemented"; return nil }

func mProtectRX(b []byte) error { _ = "STUB: not implemented"; return nil }

func mProtectPage(page, prot uintptr) error { _ = "STUB: not implemented"; return nil }

func virtualProtect(lpAddress uintptr, dwSize int, flNewProtect uint32, lpflOldProtect uintptr) error {
	_ = "STUB: not implemented"
	return nil
}
