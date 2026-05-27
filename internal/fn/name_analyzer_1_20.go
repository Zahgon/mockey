//go:build go1.20 && !go1.27
// +build go1.20,!go1.27

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

package fn

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/bytedance/mockey/internal/tool"
)

func NewNameAnalyzerByValue(fv reflect.Value) *NameAnalyzer { _ = "STUB: not implemented"; return nil }

func NewNameAnalyzer(fullname string, allowDebugLog bool) *NameAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

const (
	genericSubstr      = "[...]"
	globalSubstr       = "glob."
	ptrReceiverSubstr1 = "(*"
	ptrReceiverSubstr2 = ")"
)

var (
	anonymousNameReg = regexp.MustCompile(`func\d+(\.\d+)*$`)
)

type NameAnalyzer struct {
	// fullName name from runtime.FuncForPC
	// e.g. main.Foo[...], github.com/bytedance/mockey.Mock,
	fullName string

	// pkgName is the package name
	// e.g. main, github.com/bytedance/mockey
	pkgName string
	// middleName is the middle part separated by "." WITHOUT generic part
	// may be the type name of a method or an outer function name of an anonymous function
	// e.g. A, (*A), glob., foo
	middleName string
	// funcName is the function name WITHOUT generic part
	// e.g. Foo, Bar, func1.2.3
	funcName string
}

func (a *NameAnalyzer) init(allowDebugLog bool) *NameAnalyzer {
	if allowDebugLog {
		tool.DebugPrintf("[NameAnalyzer.init] fullName: %s\n", a.fullName)
	}
	tool.Assert(a.fullName != "", "function name is empty")

	restPart := strings.ReplaceAll(a.fullName, genericSubstr, "")

	// Extract function name
	if loc := anonymousNameReg.FindStringIndex(restPart); loc != nil {
		a.funcName = restPart[loc[0]:]
		restPart = restPart[:loc[0]-1]
	} else {
		lastDotIdx := strings.LastIndex(restPart, ".")
		a.funcName = restPart[lastDotIdx+1:]
		if lastDotIdx == -1 {
			restPart = ""
		} else {
			restPart = restPart[:lastDotIdx]
		}
	}

	// Extract package name
	firstSlashIdx := strings.LastIndex(restPart, "/")
	if firstSlashIdx > 0 {
		a.pkgName = restPart[:firstSlashIdx+1]
		restPart = restPart[firstSlashIdx+1:]
	}

	// Extract package name and middle name
	firstDotIdx := strings.Index(restPart, ".")
	if firstDotIdx > 0 {
		a.pkgName += restPart[:firstDotIdx]
		a.middleName = restPart[firstDotIdx+1:]
	} else {
		a.pkgName += restPart
	}
	if allowDebugLog {
		tool.DebugPrintf("[NameAnalyzer.init] pkgName: %s, middleName: %s, funcName: %s\n", a.pkgName, a.middleName, a.funcName)
	}
	return a
}

func (a *NameAnalyzer) FuncName() string { _ = "STUB: not implemented"; return "" }

func (a *NameAnalyzer) IsGeneric() bool { _ = "STUB: not implemented"; return false }

func (a *NameAnalyzer) HasMiddleName() bool { _ = "STUB: not implemented"; return false }

func (a *NameAnalyzer) IsExported() bool { _ = "STUB: not implemented"; return false }

func (a *NameAnalyzer) IsGlobal() bool { _ = "STUB: not implemented"; return false }

func (a *NameAnalyzer) IsPtrReceiver() bool { _ = "STUB: not implemented"; return false }

func (a *NameAnalyzer) IsAnonymousFormat() bool { _ = "STUB: not implemented"; return false }

// PkgName returns the package name
func (a *NameAnalyzer) PkgName() string {
	_ = "STUB: not implemented"

	// MiddleName returns the middle name (type name for methods)
	return ""
}

func (a *NameAnalyzer) MiddleName() string { _ = "STUB: not implemented"; return "" }
