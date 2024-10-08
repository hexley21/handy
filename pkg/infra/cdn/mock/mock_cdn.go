// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/infra/cdn/cdn.go
//
// Generated by this command:
//
//	mockgen -source=pkg/infra/cdn/cdn.go -destination=pkg/infra/cdn/mock/mock_cdn.go
//

// Package mock_cdn is a generated GoMock package.
package mock_cdn

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFileInvalidator is a mock of FileInvalidator interface.
type MockFileInvalidator struct {
	ctrl     *gomock.Controller
	recorder *MockFileInvalidatorMockRecorder
}

// MockFileInvalidatorMockRecorder is the mock recorder for MockFileInvalidator.
type MockFileInvalidatorMockRecorder struct {
	mock *MockFileInvalidator
}

// NewMockFileInvalidator creates a new mock instance.
func NewMockFileInvalidator(ctrl *gomock.Controller) *MockFileInvalidator {
	mock := &MockFileInvalidator{ctrl: ctrl}
	mock.recorder = &MockFileInvalidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileInvalidator) EXPECT() *MockFileInvalidatorMockRecorder {
	return m.recorder
}

// InvalidateFile mocks base method.
func (m *MockFileInvalidator) InvalidateFile(ctx context.Context, fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateFile", ctx, fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateFile indicates an expected call of InvalidateFile.
func (mr *MockFileInvalidatorMockRecorder) InvalidateFile(ctx, fileName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateFile", reflect.TypeOf((*MockFileInvalidator)(nil).InvalidateFile), ctx, fileName)
}
