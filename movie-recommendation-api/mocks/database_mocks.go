// Code generated by MockGen. DO NOT EDIT.
// Source: movie-reccomendation-api/interfaces (interfaces: DatabaseService)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "movie-reccomendation-api/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDatabaseService is a mock of DatabaseService interface.
type MockDatabaseService struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseServiceMockRecorder
}

// MockDatabaseServiceMockRecorder is the mock recorder for MockDatabaseService.
type MockDatabaseServiceMockRecorder struct {
	mock *MockDatabaseService
}

// NewMockDatabaseService creates a new mock instance.
func NewMockDatabaseService(ctrl *gomock.Controller) *MockDatabaseService {
	mock := &MockDatabaseService{ctrl: ctrl}
	mock.recorder = &MockDatabaseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseService) EXPECT() *MockDatabaseServiceMockRecorder {
	return m.recorder
}

// AddFavorite mocks base method.
func (m *MockDatabaseService) AddFavorite(arg0, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavorite", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavorite indicates an expected call of AddFavorite.
func (mr *MockDatabaseServiceMockRecorder) AddFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavorite", reflect.TypeOf((*MockDatabaseService)(nil).AddFavorite), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockDatabaseService) CreateUser(arg0 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDatabaseServiceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabaseService)(nil).CreateUser), arg0)
}

// GetFavoriteByUserAndId mocks base method.
func (m *MockDatabaseService) GetFavoriteByUserAndId(arg0, arg1 int64) (*models.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteByUserAndId", arg0, arg1)
	ret0, _ := ret[0].(*models.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteByUserAndId indicates an expected call of GetFavoriteByUserAndId.
func (mr *MockDatabaseServiceMockRecorder) GetFavoriteByUserAndId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteByUserAndId", reflect.TypeOf((*MockDatabaseService)(nil).GetFavoriteByUserAndId), arg0, arg1)
}

// GetFavoritesByUser mocks base method.
func (m *MockDatabaseService) GetFavoritesByUser(arg0 int64) ([]models.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoritesByUser", arg0)
	ret0, _ := ret[0].([]models.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoritesByUser indicates an expected call of GetFavoritesByUser.
func (mr *MockDatabaseServiceMockRecorder) GetFavoritesByUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoritesByUser", reflect.TypeOf((*MockDatabaseService)(nil).GetFavoritesByUser), arg0)
}

// GetUserByID mocks base method.
func (m *MockDatabaseService) GetUserByID(arg0 int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockDatabaseServiceMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDatabaseService)(nil).GetUserByID), arg0)
}

// GetUserByUsernameOrEmail mocks base method.
func (m *MockDatabaseService) GetUserByUsernameOrEmail(arg0, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsernameOrEmail", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsernameOrEmail indicates an expected call of GetUserByUsernameOrEmail.
func (mr *MockDatabaseServiceMockRecorder) GetUserByUsernameOrEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsernameOrEmail", reflect.TypeOf((*MockDatabaseService)(nil).GetUserByUsernameOrEmail), arg0, arg1)
}

// LoginUser mocks base method.
func (m *MockDatabaseService) LoginUser(arg0, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockDatabaseServiceMockRecorder) LoginUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockDatabaseService)(nil).LoginUser), arg0, arg1)
}

// RemoveFavorite mocks base method.
func (m *MockDatabaseService) RemoveFavorite(arg0, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFavorite", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFavorite indicates an expected call of RemoveFavorite.
func (mr *MockDatabaseServiceMockRecorder) RemoveFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFavorite", reflect.TypeOf((*MockDatabaseService)(nil).RemoveFavorite), arg0, arg1)
}
