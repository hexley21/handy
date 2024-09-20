// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/service/user.go
//
// Generated by this command:
//
//	mockgen -source=internal/user/service/user.go -destination=internal/user/service/mock/mock_user.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	io "io"
	reflect "reflect"

	dto "github.com/hexley21/fixup/internal/user/delivery/http/v1/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockUserService) ChangePassword(ctx context.Context, id int64, updateDto dto.UpdatePassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, id, updateDto)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserServiceMockRecorder) ChangePassword(ctx, id, updateDto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserService)(nil).ChangePassword), ctx, id, updateDto)
}

// DeleteUserById mocks base method.
func (m *MockUserService) DeleteUserById(ctx context.Context, userId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockUserServiceMockRecorder) DeleteUserById(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUserService)(nil).DeleteUserById), ctx, userId)
}

// FindUserById mocks base method.
func (m *MockUserService) FindUserById(ctx context.Context, userId int64) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", ctx, userId)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserServiceMockRecorder) FindUserById(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserService)(nil).FindUserById), ctx, userId)
}

// FindUserProfileById mocks base method.
func (m *MockUserService) FindUserProfileById(ctx context.Context, userId int64) (dto.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserProfileById", ctx, userId)
	ret0, _ := ret[0].(dto.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserProfileById indicates an expected call of FindUserProfileById.
func (mr *MockUserServiceMockRecorder) FindUserProfileById(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserProfileById", reflect.TypeOf((*MockUserService)(nil).FindUserProfileById), ctx, userId)
}

// SetProfilePicture mocks base method.
func (m *MockUserService) SetProfilePicture(ctx context.Context, userId int64, file io.Reader, fileName string, fileSize int64, fileType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProfilePicture", ctx, userId, file, fileName, fileSize, fileType)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProfilePicture indicates an expected call of SetProfilePicture.
func (mr *MockUserServiceMockRecorder) SetProfilePicture(ctx, userId, file, fileName, fileSize, fileType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProfilePicture", reflect.TypeOf((*MockUserService)(nil).SetProfilePicture), ctx, userId, file, fileName, fileSize, fileType)
}

// UpdateUserDataById mocks base method.
func (m *MockUserService) UpdateUserDataById(ctx context.Context, id int64, updateDto dto.UpdateUser) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDataById", ctx, id, updateDto)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserDataById indicates an expected call of UpdateUserDataById.
func (mr *MockUserServiceMockRecorder) UpdateUserDataById(ctx, id, updateDto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDataById", reflect.TypeOf((*MockUserService)(nil).UpdateUserDataById), ctx, id, updateDto)
}
