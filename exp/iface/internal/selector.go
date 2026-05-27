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

type Selector interface {
	Match(info *funcInfo) bool
}

const (
	CTAnd CombineType = 0
	CTOr  CombineType = 1
)

type CombineType int

func NewCombinedSelector(combineType CombineType, selectors ...Selector) *CombinedSelector {
	_ = "STUB: not implemented"
	return nil
}

type CombinedSelector struct {
	combineType CombineType
	selectors   []Selector
	inverse     bool
}

func (s *CombinedSelector) Match(info *funcInfo) (res bool) {
	_ = "STUB: not implemented"
	return false
}

func (s *CombinedSelector) Add(selectors ...Selector) { _ = "STUB: not implemented"; return }

func (s *CombinedSelector) Not() { _ = "STUB: not implemented"; return }

const (
	MMExact   MatchMode = 0
	MMContain MatchMode = 1
)

type MatchMode int

func (m MatchMode) match(a, b string) bool { _ = "STUB: not implemented"; return false }

func NewPkgSelector(name string, mode MatchMode) PkgSelector {
	_ = "STUB: not implemented"
	return *new(PkgSelector)
}

type PkgSelector struct {
	mode MatchMode
	name string
}

func (s PkgSelector) Match(info *funcInfo) bool { _ = "STUB: not implemented"; return false }

func NewTypeSelector(name string, mode MatchMode) TypeSelector {
	_ = "STUB: not implemented"
	return *new(TypeSelector)
}

type TypeSelector struct {
	mode MatchMode
	name string
}

func (s TypeSelector) Match(info *funcInfo) bool { _ = "STUB: not implemented"; return false }
