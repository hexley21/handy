// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/jwt/refresh_jwt/jwt.go
//
// Generated by this command:
//
//	mockgen -source=internal/user/jwt/refresh_jwt/jwt.go -destination=internal/user/jwt/refresh_jwt/mock/mock_jwt.go
//

// Package mock_refresh_jwt is a generated GoMock package.
package mock_refresh_jwt

import (
	reflect "reflect"

	refresh_jwt "github.com/hexley21/fixup/internal/user/jwt/refresh_jwt"
	rest "github.com/hexley21/fixup/pkg/http/rest"
	gomock "go.uber.org/mock/gomock"
)

// MockJWTManager is a mock of JWTManager interface.
type MockJWTManager struct {
	ctrl     *gomock.Controller
	recorder *MockJWTManagerMockRecorder
}

// MockJWTManagerMockRecorder is the mock recorder for MockJWTManager.
type MockJWTManagerMockRecorder struct {
	mock *MockJWTManager
}

// NewMockJWTManager creates a new mock instance.
func NewMockJWTManager(ctrl *gomock.Controller) *MockJWTManager {
	mock := &MockJWTManager{ctrl: ctrl}
	mock.recorder = &MockJWTManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTManager) EXPECT() *MockJWTManagerMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockJWTManager) Generate(id string) (string, *rest.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rest.ErrorResponse)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockJWTManagerMockRecorder) Generate(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockJWTManager)(nil).Generate), id)
}

// Verify mocks base method.
func (m *MockJWTManager) Verify(tokenString string) (refresh_jwt.RefreshClaims, *rest.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", tokenString)
	ret0, _ := ret[0].(refresh_jwt.RefreshClaims)
	ret1, _ := ret[1].(*rest.ErrorResponse)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockJWTManagerMockRecorder) Verify(tokenString any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockJWTManager)(nil).Verify), tokenString)
}

// MockGenerator is a mock of Generator interface.
type MockGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorMockRecorder
}

// MockGeneratorMockRecorder is the mock recorder for MockGenerator.
type MockGeneratorMockRecorder struct {
	mock *MockGenerator
}

// NewMockGenerator creates a new mock instance.
func NewMockGenerator(ctrl *gomock.Controller) *MockGenerator {
	mock := &MockGenerator{ctrl: ctrl}
	mock.recorder = &MockGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenerator) EXPECT() *MockGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockGenerator) Generate(id string) (string, *rest.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rest.ErrorResponse)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockGeneratorMockRecorder) Generate(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockGenerator)(nil).Generate), id)
}

// MockVerifier is a mock of Verifier interface.
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier.
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance.
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method.
func (m *MockVerifier) Verify(tokenString string) (refresh_jwt.RefreshClaims, *rest.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", tokenString)
	ret0, _ := ret[0].(refresh_jwt.RefreshClaims)
	ret1, _ := ret[1].(*rest.ErrorResponse)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockVerifierMockRecorder) Verify(tokenString any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifier)(nil).Verify), tokenString)
}
