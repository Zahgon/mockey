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

package iface

import (
	"github.com/bytedance/mockey"
)

type Mocker struct {
	builder *MockBuilder
	mockers []*mockey.Mocker
}

type MockBuilder struct {
	builders []*mockey.MockBuilder
}

// Mock mocks the given interface method. This will mock all the implemented methods of the interface. Note this is an
// experimental feature.
//
// Example:
// Mock(io.Reader.Read).Return(1, io.EOF).Build() // mock all implemented methods of io.Reader
//
// If you want to limit the mock scope to specific methods, you can use the select option.
//
// Example:
// Mock(io.Reader.Read, SelectType("Buffer"), SelectPkg("bytes")).Return(1, io.EOF).Build() // only mock bytes.Buffer
//
// For more details, please refer to https://github.com/bytedance/mockey/issues/3#issuecomment-3759010755.
func Mock(target interface{}, opt ...OptionFn) *MockBuilder { _ = "STUB: not implemented"; return nil }

func (builder *MockBuilder) When(when interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) To(hook interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) Return(results ...interface{}) *MockBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (builder *MockBuilder) Build() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) Patch() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) UnPatch() *Mocker { _ = "STUB: not implemented"; return nil }

func (mocker *Mocker) Times() int { _ = "STUB: not implemented"; return 0 }

func (mocker *Mocker) MockTimes() int { _ = "STUB: not implemented"; return 0 }
