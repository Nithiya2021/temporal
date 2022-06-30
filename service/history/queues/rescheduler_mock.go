// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: rescheduler.go

// Package queues is a generated GoMock package.
package queues

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRescheduler is a mock of Rescheduler interface.
type MockRescheduler struct {
	ctrl     *gomock.Controller
	recorder *MockReschedulerMockRecorder
}

// MockReschedulerMockRecorder is the mock recorder for MockRescheduler.
type MockReschedulerMockRecorder struct {
	mock *MockRescheduler
}

// NewMockRescheduler creates a new mock instance.
func NewMockRescheduler(ctrl *gomock.Controller) *MockRescheduler {
	mock := &MockRescheduler{ctrl: ctrl}
	mock.recorder = &MockReschedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRescheduler) EXPECT() *MockReschedulerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRescheduler) Add(task Executable, rescheduleTime time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", task, rescheduleTime)
}

// Add indicates an expected call of Add.
func (mr *MockReschedulerMockRecorder) Add(task, rescheduleTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRescheduler)(nil).Add), task, rescheduleTime)
}

// Len mocks base method.
func (m *MockRescheduler) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockReschedulerMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockRescheduler)(nil).Len))
}

// Start mocks base method.
func (m *MockRescheduler) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockReschedulerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRescheduler)(nil).Start))
}

// Stop mocks base method.
func (m *MockRescheduler) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockReschedulerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRescheduler)(nil).Stop))
}
