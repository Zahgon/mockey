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

import (
	"sync"
)

type SequenceOpt interface {
	// private make sure it is mockey private interface
	private
	// GetNext is used by mockey, don't use it if you don't know what it does
	GetNext() []interface{}
}

type sequenceOpt interface {
	SequenceOpt
	Times(int) sequenceOpt
	Then(...interface{}) sequenceOpt
}

type sequence struct {
	Private  // make sure it does implements mockey SequenceOpt
	values   []*sequenceValue
	curV     int // current value
	curT     int // current value times
	readLock sync.Mutex
}

type sequenceValue struct {
	v []interface{}
	t int
}

func (s *sequence) GetNext() []interface{} { _ = "STUB: not implemented"; return nil }

func (s *sequence) Then(value ...interface{}) sequenceOpt {
	_ = "STUB: not implemented"
	return *new(sequenceOpt)
}

func (s *sequence) Times(t int) sequenceOpt { _ = "STUB: not implemented"; return *new(sequenceOpt) }

func Sequence(value ...interface{}) sequenceOpt {
	_ = "STUB: not implemented"
	return *new(sequenceOpt)
}
