/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logutils

import (
	"github.com/stretchr/testify/mock"
)

type MockLog struct {
	mock.Mock
}

func NewMockLog() *MockLog {
	return &MockLog{}
}

func (m *MockLog) Fatalf(format string, args ...any) {
	m.Called(append([]any{format}, args...)...)
}

func (m *MockLog) Panicf(format string, args ...any) {
	m.Called(append([]any{format}, args...)...)
}

func (m *MockLog) Errorf(format string, args ...any) {
	m.Called(append([]any{format}, args...)...)
}

func (m *MockLog) Warnf(format string, args ...any) {
	m.Called(append([]any{format}, args...)...)
}

func (m *MockLog) Infof(format string, args ...any) {
	m.Called(append([]any{format}, args...)...)
}

func (m *MockLog) Child(name string) Log {
	m.Called(name)
	return m
}

func (m *MockLog) SetLevel(level LogLevel) {
	m.Called(level)
}

func (m *MockLog) OnFatalf(format string, args ...any) *MockLog {
	arguments := append([]any{format}, args...)

	m.On("Fatalf", arguments...)

	return m
}

func (m *MockLog) OnPanicf(format string, args ...any) *MockLog {
	arguments := append([]any{format}, args...)

	m.On("Panicf", arguments...)

	return m
}

func (m *MockLog) OnErrorf(format string, args ...any) *MockLog {
	arguments := append([]any{format}, args...)

	m.On("Errorf", arguments...)

	return m
}

func (m *MockLog) OnWarnf(format string, args ...any) *MockLog {
	arguments := append([]any{format}, args...)

	m.On("Warnf", arguments...)

	return m
}

func (m *MockLog) OnInfof(format string, args ...any) *MockLog {
	arguments := append([]any{format}, args...)

	m.On("Infof", arguments...)

	return m
}
