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

package tool

import (
	"runtime"
)

type CallerInfo runtime.Frame

func (c CallerInfo) String() string { _ = "STUB: not implemented"; return "" }

// OuterCaller gets non-current package caller of a function
// For example, assume we have 3 files: a/b/foo.go, a/c/bar.go and a/c/innerBar.go,
// a/b/foo.Foo calls a/c/bar.Bar, and  a/c/bar.Bar calls a/c/innerBar.innerBar.
// Here is how innerBar looks like:
//
//	func innerBar() CallerInfo { /*do some thing*/ return Caller() }
//
// The return value of innerBar should represent the line in a/b/foo.go where a/b/foo.Foo calls a/c/bar.Bar
func OuterCaller() (info CallerInfo) { _ = "STUB: not implemented"; return *new(CallerInfo) }

func Caller() CallerInfo { _ = "STUB: not implemented"; return *new(CallerInfo) }

func getPackageAndFunction(pc uintptr) (string, string) { _ = "STUB: not implemented"; return "", "" }

// if mock run in an anonymous function of a global variable,
// the stack will looks like a.b.c.glob..func1(), so the
// second last part of the caller stack would not be guaranteed
// always to be non-empty.
