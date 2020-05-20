// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/typical-go/typical-rest-server/internal/server/service (interfaces: BookSvc)

// Package service_mock is a generated GoMock package.
package service_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repository "github.com/typical-go/typical-rest-server/internal/server/repository"
	reflect "reflect"
)

// MockBookSvc is a mock of BookSvc interface
type MockBookSvc struct {
	ctrl     *gomock.Controller
	recorder *MockBookSvcMockRecorder
}

// MockBookSvcMockRecorder is the mock recorder for MockBookSvc
type MockBookSvcMockRecorder struct {
	mock *MockBookSvc
}

// NewMockBookSvc creates a new mock instance
func NewMockBookSvc(ctrl *gomock.Controller) *MockBookSvc {
	mock := &MockBookSvc{ctrl: ctrl}
	mock.recorder = &MockBookSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookSvc) EXPECT() *MockBookSvcMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockBookSvc) Create(arg0 context.Context, arg1 *repository.Book) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBookSvcMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookSvc)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockBookSvc) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBookSvcMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookSvc)(nil).Delete), arg0, arg1)
}

// Find mocks base method
func (m *MockBookSvc) Find(arg0 context.Context) ([]*repository.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]*repository.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockBookSvcMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockBookSvc)(nil).Find), arg0)
}

// FindOne mocks base method
func (m *MockBookSvc) FindOne(arg0 context.Context, arg1 string) (*repository.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*repository.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockBookSvcMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockBookSvc)(nil).FindOne), arg0, arg1)
}

// Update mocks base method
func (m *MockBookSvc) Update(arg0 context.Context, arg1 string, arg2 *repository.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockBookSvcMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookSvc)(nil).Update), arg0, arg1, arg2)
}
