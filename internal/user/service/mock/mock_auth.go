// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/service/auth.go
//
// Generated by this command:
//
//	mockgen -source=internal/user/service/auth.go -destination=internal/user/service/mock/mock_auth.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"
	time "time"

	dto "github.com/hexley21/fixup/internal/user/delivery/http/v1/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockAuthService) AuthenticateUser(ctx context.Context, loginDto dto.Login) (dto.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", ctx, loginDto)
	ret0, _ := ret[0].(dto.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockAuthServiceMockRecorder) AuthenticateUser(ctx, loginDto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockAuthService)(nil).AuthenticateUser), ctx, loginDto)
}

// GetUserConfirmationDetails mocks base method.
func (m *MockAuthService) GetUserConfirmationDetails(ctx context.Context, email string) (dto.UserConfirmationDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserConfirmationDetails", ctx, email)
	ret0, _ := ret[0].(dto.UserConfirmationDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserConfirmationDetails indicates an expected call of GetUserConfirmationDetails.
func (mr *MockAuthServiceMockRecorder) GetUserConfirmationDetails(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserConfirmationDetails", reflect.TypeOf((*MockAuthService)(nil).GetUserConfirmationDetails), ctx, email)
}

// GetUserRoleAndStatus mocks base method.
func (m *MockAuthService) GetUserRoleAndStatus(ctx context.Context, id int64) (dto.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRoleAndStatus", ctx, id)
	ret0, _ := ret[0].(dto.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRoleAndStatus indicates an expected call of GetUserRoleAndStatus.
func (mr *MockAuthServiceMockRecorder) GetUserRoleAndStatus(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRoleAndStatus", reflect.TypeOf((*MockAuthService)(nil).GetUserRoleAndStatus), ctx, id)
}

// RegisterCustomer mocks base method.
func (m *MockAuthService) RegisterCustomer(ctx context.Context, registerDto dto.RegisterUser) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCustomer", ctx, registerDto)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterCustomer indicates an expected call of RegisterCustomer.
func (mr *MockAuthServiceMockRecorder) RegisterCustomer(ctx, registerDto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCustomer", reflect.TypeOf((*MockAuthService)(nil).RegisterCustomer), ctx, registerDto)
}

// RegisterProvider mocks base method.
func (m *MockAuthService) RegisterProvider(ctx context.Context, registerDto dto.RegisterProvider) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterProvider", ctx, registerDto)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterProvider indicates an expected call of RegisterProvider.
func (mr *MockAuthServiceMockRecorder) RegisterProvider(ctx, registerDto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterProvider", reflect.TypeOf((*MockAuthService)(nil).RegisterProvider), ctx, registerDto)
}

// SendConfirmationLetter mocks base method.
func (m *MockAuthService) SendConfirmationLetter(ctx context.Context, token, email, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendConfirmationLetter", ctx, token, email, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendConfirmationLetter indicates an expected call of SendConfirmationLetter.
func (mr *MockAuthServiceMockRecorder) SendConfirmationLetter(ctx, token, email, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendConfirmationLetter", reflect.TypeOf((*MockAuthService)(nil).SendConfirmationLetter), ctx, token, email, name)
}

// SendVerifiedLetter mocks base method.
func (m *MockAuthService) SendVerifiedLetter(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerifiedLetter", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerifiedLetter indicates an expected call of SendVerifiedLetter.
func (mr *MockAuthServiceMockRecorder) SendVerifiedLetter(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerifiedLetter", reflect.TypeOf((*MockAuthService)(nil).SendVerifiedLetter), email)
}

// VerifyUser mocks base method.
func (m *MockAuthService) VerifyUser(ctx context.Context, token string, ttl time.Duration, id int64, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", ctx, token, ttl, id, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockAuthServiceMockRecorder) VerifyUser(ctx, token, ttl, id, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockAuthService)(nil).VerifyUser), ctx, token, ttl, id, email)
}
