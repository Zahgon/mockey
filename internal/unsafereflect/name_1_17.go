//go:build go1.17 && !go1.27
// +build go1.17,!go1.27

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
 */

package unsafereflect

// name is an encoded type name with optional extra data.
type name struct {
	bytes *byte
}

func (n name) data(off int, whySafe string) *byte { _ = "STUB: not implemented"; return nil }

func (n name) readVarint(off int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (n name) name() (s string) { _ = "STUB: not implemented"; return "" }
