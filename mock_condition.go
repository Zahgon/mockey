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

type mockCondition struct {
	when interface{} // condition
	hook interface{} // mock function

	builder *MockBuilder
}

func (m *mockCondition) Complete() bool { _ = "STUB: not implemented"; return false }

func (m *mockCondition) SetWhen(when interface{}) { _ = "STUB: not implemented"; return }

func (m *mockCondition) SetWhenForce(when interface{}) { _ = "STUB: not implemented"; return }

func (m *mockCondition) SetReturn(results ...interface{}) { _ = "STUB: not implemented"; return }

func (m *mockCondition) SetReturnForce(results ...interface{}) { _ = "STUB: not implemented"; return }

func (m *mockCondition) SetTo(to interface{}) { _ = "STUB: not implemented"; return }

func (m *mockCondition) SetToForce(to interface{}) { _ = "STUB: not implemented"; return }
