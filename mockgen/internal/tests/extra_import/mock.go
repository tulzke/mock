// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/mock/mockgen/internal/tests/extra_import (interfaces: Foo)
//
// Generated by this command:
//
//	mockgen -destination mock.go -package extra_import . Foo
//

// Package extra_import is a generated GoMock package.
package extra_import

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFoo is a mock of Foo interface.
type MockFoo struct {
	ctrl     *gomock.Controller
	recorder *MockFooMockRecorder
}

// MockFooMockRecorder is the mock recorder for MockFoo.
type MockFooMockRecorder struct {
	mock *MockFoo
}

// NewMockFoo creates a new mock instance.
func NewMockFoo(ctrl *gomock.Controller) *MockFoo {
	mock := &MockFoo{ctrl: ctrl}
	mock.recorder = &MockFooMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFoo) EXPECT() *MockFooMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockFoo) ISGOMOCK() struct{} {
	return struct{}{}
}

// Bar mocks base method.
func (m *MockFoo) Bar(channels []string, message chan<- Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Bar", channels, message)
}

// Bar indicates an expected call of Bar.
func (mr *MockFooMockRecorder) Bar(channels, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bar", reflect.TypeOf((*MockFoo)(nil).Bar), channels, message)
}
