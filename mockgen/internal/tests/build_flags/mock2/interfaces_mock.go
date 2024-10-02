// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/mock/mockgen/internal/tests/build_flags (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -destination=mock2/interfaces_mock.go -build_flags=-race -tags=tag2 . Interface
//

// Package mock_build_flags is a generated GoMock package.
package mock_build_flags

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Foo mocks base method.
func (m *MockInterface) Foo() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo")
}

// Foo indicates an expected call of Foo.
func (mr *MockInterfaceMockRecorder) Foo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockInterface)(nil).Foo))
}
