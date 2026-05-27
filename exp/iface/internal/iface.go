//go:build go1.20 && !go1.26
// +build go1.20,!go1.26

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

package internal

import (
	"runtime"

	"github.com/bytedance/mockey/internal/fn"
	"github.com/bytedance/mockey/internal/monkey/linkname"
)

func FindImplementTargets(i interface{}, selector Selector) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Due to the lack of type information, our methods for finding targets are very limited.

// Exclude interface itself

// Exclude argument size not match

// Exclude function name not match

var funcInfoMap = make(map[string][]*funcInfo)

type funcInfo struct {
	Func     *runtime.Func
	Analyzer *fn.NameAnalyzer
}

func init() {
	for _, fun := range linkname.FuncList() {
		fullName := fun.Name()
		// Exclude internal functions that do not have a name
		if fullName == "" {
			continue
		}
		analyzer := fn.NewNameAnalyzer(fullName, false)
		// Exclude functions that do not have a receiver name
		if !analyzer.HasMiddleName() {
			continue
		}
		// Exclude methods that do not have pointer receivers. If a type implements an interface, all methods of that
		// interface must have pointer receivers.
		if !analyzer.IsPtrReceiver() {
			continue
		}
		name := analyzer.FuncName()
		funcInfoMap[name] = append(funcInfoMap[name], &funcInfo{Func: fun, Analyzer: analyzer})
	}
}

func totalArgSize(f *runtime.Func) int32 { _ = "STUB: not implemented"; return 0 }
