// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/interfaces.go

// Package mockRepository is a generated GoMock package.
package mockRepository

import (
	context "context"
	reflect "reflect"
	models "user-service/pkg/models"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// GetUserData mocks base method.
func (m *MockUserRepo) GetUserData(ctx context.Context, id int32) (models.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserData", ctx, id)
	ret0, _ := ret[0].(models.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserData indicates an expected call of GetUserData.
func (mr *MockUserRepoMockRecorder) GetUserData(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserData", reflect.TypeOf((*MockUserRepo)(nil).GetUserData), ctx, id)
}

// GetUserDataList mocks base method.
func (m *MockUserRepo) GetUserDataList(ctx context.Context, ids []int32) ([]models.UserData, models.NotFoundList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataList", ctx, ids)
	ret0, _ := ret[0].([]models.UserData)
	ret1, _ := ret[1].(models.NotFoundList)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserDataList indicates an expected call of GetUserDataList.
func (mr *MockUserRepoMockRecorder) GetUserDataList(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataList", reflect.TypeOf((*MockUserRepo)(nil).GetUserDataList), ctx, ids)
}
