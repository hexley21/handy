// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/validator/validator.go
//
// Generated by this command:
//
//	mockgen -source=pkg/validator/validator.go -destination=pkg/validator/mock/mock_validator.go
//

// Package mock_validator is a generated GoMock package.
package mock_validator

import (
	reflect "reflect"

	rest "github.com/hexley21/fixup/pkg/http/rest"
	gomock "go.uber.org/mock/gomock"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidator) Validate(i any) *rest.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", i)
	ret0, _ := ret[0].(*rest.ErrorResponse)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockValidatorMockRecorder) Validate(i any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidator)(nil).Validate), i)
}
