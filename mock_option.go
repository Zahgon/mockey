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

type mockOption struct {
	unsafe  bool
	generic *bool
	method  *bool
}

type mockOptionFn func(*mockOption)

func OptUnsafe(o *mockOption) { _ = "STUB: not implemented"; return }

func OptGeneric(o *mockOption) { _ = "STUB: not implemented"; return }

func OptMethod(o *mockOption) { _ = "STUB: not implemented"; return }

func resolveMockOpt(fn ...mockOptionFn) *mockOption { _ = "STUB: not implemented"; return nil }
