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

package mem

// WriteWithSTW copies data bytes to the target address and replaces the original bytes, during which it will stop the
// world (only the current goroutine's P is running).
func WriteWithSTW(target uintptr, data []byte) { _ = "STUB: not implemented"; return }

func suspendRuntime() (resume func()) { _ = "STUB: not implemented"; return nil }

// Suspend the system monitor thread to avoid SIGBUS errors during memory writes
// See https://github.com/bytedance/mockey/issues/68 for more details.
